// Definition for the SOAP structure required for UPnP's SOAP usage.

package soap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	soapEncodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"
	soapPrefix        = xml.Header + `<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><s:Body>`
	soapSuffix        = `</s:Body></s:Envelope>`
)

type SOAPClient struct {
	EndpointURL url.URL
	HTTPClient  http.Client
}

func NewSOAPClient(endpointURL url.URL) *SOAPClient {
	return &SOAPClient{
		EndpointURL: endpointURL,
	}
}

// PerformAction makes a SOAP request, with the given action.
// inAction and outAction must both be pointers to structs with string fields
// only.
func (client *SOAPClient) PerformAction(actionNamespace, actionName string, inAction, outAction interface{}) error {
	requestBytes, err := encodeRequestAction(actionNamespace, actionName, inAction)
	if err != nil {
		return err
	}

	request := &http.Request{
		Method: "POST",
		URL:    &client.EndpointURL,
		Header: http.Header{
			"SOAPACTION":   []string{`"` + actionNamespace + "#" + actionName + `"`},
			"CONTENT-TYPE": []string{"text/xml; charset=\"utf-8\""},
		},
		Body: ioutil.NopCloser(bytes.NewBuffer(requestBytes)),
		// Set ContentLength to avoid chunked encoding - some servers might not support it.
		ContentLength: int64(len(requestBytes)),
	}
	requestDumpBytes, err := httputil.DumpRequest(request, true)
	if err != nil {
		return err
	}
	log.WithField("req", string(requestDumpBytes)).Info("Encode request.")
	response, err := client.HTTPClient.Do(request)
	if err != nil {
		return errors.Wrapf(err, "goupnp: error performing SOAP HTTP request")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		responseBytes, err := ioutil.ReadAll(response.Body)
		log.WithFields(log.Fields{
			log.ErrorKey: err,
			"resp":       string(responseBytes),
		}).Warn("Response with error.")
		return errors.Errorf("goupnp: SOAP request got HTTP %s", response.Status)
	}

	responseEnv := newSOAPEnvelope()
	decoder := xml.NewDecoder(response.Body)
	if err := decoder.Decode(responseEnv); err != nil {
		return errors.Wrapf(err, "goupnp: error decoding response body: %v")
	}

	log.WithField("resp", responseEnv).Info("Recv response.")

	if responseEnv.Body.Fault != nil {
		return responseEnv.Body.Fault
	}

	if outAction != nil {
		if err := xml.Unmarshal(responseEnv.Body.RawAction, outAction); err != nil {
			return errors.Wrapf(err, "goupnp: error unmarshalling out action: %v", responseEnv.Body.RawAction)
		}
	}

	return nil
}

// newSOAPAction creates a soapEnvelope with the given action and arguments.
func newSOAPEnvelope() *soapEnvelope {
	return &soapEnvelope{
		EncodingStyle: soapEncodingStyle,
	}
}

// encodeRequestAction is a hacky way to create an encoded SOAP envelope
// containing the given action. Experiments with one router have shown that it
// 500s for requests where the outer default xmlns is set to the SOAP
// namespace, and then reassigning the default namespace within that to the
// service namespace. Hand-coding the outer XML to work-around this.
func encodeRequestAction(actionNamespace, actionName string, inAction interface{}) ([]byte, error) {
	requestBuf := new(bytes.Buffer)
	requestBuf.WriteString(soapPrefix)
	requestBuf.WriteString(`<u:`)
	//nolint:errcheck
	xml.EscapeText(requestBuf, []byte(actionName))
	requestBuf.WriteString(` xmlns:u="`)
	//nolint:errcheck
	xml.EscapeText(requestBuf, []byte(actionNamespace))
	requestBuf.WriteString(`">`)
	if inAction != nil {
		if err := encodeRequestArgs(requestBuf, inAction); err != nil {
			return nil, errors.Wrap(err, "encoding request arguments")
		}
	}
	requestBuf.WriteString(`</u:`)
	//nolint:errcheck
	xml.EscapeText(requestBuf, []byte(actionName))
	requestBuf.WriteString(`>`)
	requestBuf.WriteString(soapSuffix)
	return requestBuf.Bytes(), nil
}

func encodeRequestArgs(w *bytes.Buffer, inAction interface{}) error {
	in := reflect.Indirect(reflect.ValueOf(inAction))
	if in.Kind() != reflect.Struct {
		return errors.Errorf("goupnp: SOAP inAction is not a struct but of type %v", in.Type())
	}
	enc := xml.NewEncoder(w)
	nFields := in.NumField()
	inType := in.Type()
	for i := 0; i < nFields; i++ {
		field := inType.Field(i)
		argName := field.Name
		if nameOverride := field.Tag.Get("soap"); nameOverride != "" {
			argName = nameOverride
		}
		value := in.Field(i)
		if value.Kind() != reflect.String {
			return errors.Errorf("goupnp: SOAP arg %q is not of type string, but of type %v", argName, value.Type())
		}
		elem := xml.StartElement{Name: xml.Name{Local: argName}}
		if err := enc.EncodeToken(elem); err != nil {
			return errors.Wrapf(err, "goupnp: error encoding start element for SOAP arg %q", argName)
		}
		if err := enc.Flush(); err != nil {
			return errors.Wrapf(err, "goupnp: error flushing start element for SOAP arg %q", argName)
		}
		text := value.Interface().(string)
		if !strings.HasPrefix(text, "<![CDATA") {
			text = escapeXMLText(text)
		}
		if _, err := w.Write([]byte(text)); err != nil {
			return errors.Wrapf(err, "goupnp: error writing value for SOAP arg %q: %v", argName)
		}
		if err := enc.EncodeToken(elem.End()); err != nil {
			return errors.Wrapf(err, "goupnp: error encoding end element for SOAP arg %q: %v", argName)
		}
	}
	enc.Flush()
	return nil
}

var xmlCharRx = regexp.MustCompile("[<>&]")

// escapeXMLText is used by generated code to escape text in XML, but only
// escaping the characters `<`, `>`, and `&`.
//
// This is provided in order to work around SOAP server implementations that
// fail to decode XML correctly, specifically failing to decode `"`, `'`. Note
// that this can only be safely used for injecting into XML text, but not into
// attributes or other contexts.
func escapeXMLText(s string) string {
	return xmlCharRx.ReplaceAllStringFunc(s, replaceEntity)
}

func replaceEntity(s string) string {
	switch s {
	case "<":
		return "&lt;"
	case ">":
		return "&gt;"
	case "&":
		return "&amp;"
	}
	return s
}

func (r *soapEnvelope) String() string {
	return r.Body.String()
}

type soapEnvelope struct {
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	EncodingStyle string   `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
	Body          soapBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

func (r *soapBody) String() string {
	var buf bytes.Buffer
	if r.Fault != nil {
		fmt.Print(&buf, "fault:", *r.Fault, ",")
	}
	fmt.Fprint(&buf, "body:", string(r.RawAction))
	return buf.String()
}

type soapBody struct {
	Fault     *SOAPFaultError `xml:"Fault"`
	RawAction []byte          `xml:",innerxml"`
}

// SOAPFaultError implements error, and contains SOAP fault information.
type SOAPFaultError struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	Detail      string `xml:"detail"`
}

func (err *SOAPFaultError) Error() string {
	return fmt.Sprintf("SOAP fault: %s", err.FaultString)
}
