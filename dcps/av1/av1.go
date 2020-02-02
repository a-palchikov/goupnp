// Client for UPnP Device Control Protocol MediaServer v1 and MediaRenderer v1.
//
// This DCP is documented in detail at: http://upnp.org/specs/av/av1/
//
// Typically, use one of the New* functions to create clients for services.
package av1

// ***********************************************************
// GENERATED FILE - DO NOT EDIT BY HAND. See README.md
// ***********************************************************

import (
	"net/url"
	"time"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/soap"
	"github.com/pkg/errors"
)

// Hack to avoid Go complaining if time isn't used.
var _ time.Time

// Device URNs:
const ()

// Service URNs:
const (
	URN_AVTransport_1        = "urn:schemas-upnp-org:service:AVTransport:1"
	URN_AVTransport_2        = "urn:schemas-upnp-org:service:AVTransport:2"
	URN_ConnectionManager_1  = "urn:schemas-upnp-org:service:ConnectionManager:1"
	URN_ConnectionManager_2  = "urn:schemas-upnp-org:service:ConnectionManager:2"
	URN_ContentDirectory_1   = "urn:schemas-upnp-org:service:ContentDirectory:1"
	URN_ContentDirectory_2   = "urn:schemas-upnp-org:service:ContentDirectory:2"
	URN_ContentDirectory_3   = "urn:schemas-upnp-org:service:ContentDirectory:3"
	URN_RenderingControl_1   = "urn:schemas-upnp-org:service:RenderingControl:1"
	URN_RenderingControl_2   = "urn:schemas-upnp-org:service:RenderingControl:2"
	URN_ScheduledRecording_1 = "urn:schemas-upnp-org:service:ScheduledRecording:1"
	URN_ScheduledRecording_2 = "urn:schemas-upnp-org:service:ScheduledRecording:2"
)

// AVTransport1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:AVTransport:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type AVTransport1 struct {
	goupnp.ServiceClient
}

// NewAVTransport1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewAVTransport1Clients() (clients []*AVTransport1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_AVTransport_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newAVTransport1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewAVTransport1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewAVTransport1ClientsByURL(loc *url.URL) ([]*AVTransport1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_AVTransport_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newAVTransport1ClientsFromGenericClients(genericClients), nil
}

// NewAVTransport1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewAVTransport1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*AVTransport1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_AVTransport_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newAVTransport1ClientsFromGenericClients(genericClients), nil
}

func newAVTransport1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*AVTransport1 {
	clients := make([]*AVTransport1, len(genericClients))
	for i := range genericClients {
		clients[i] = &AVTransport1{genericClients[i]}
	}
	return clients
}

// AVTransport1SetAVTransportURIRequest describes the request for AVTransport1.SetAVTransportURI API
type AVTransport1SetAVTransportURIRequest struct {
	InstanceID         soap.Ui4
	CurrentURI         soap.String
	CurrentURIMetaData soap.String
}

// AVTransport1SetAVTransportURIResponse describes the response for AVTransport1.SetAVTransportURI API
type AVTransport1SetAVTransportURIResponse struct {
}

func (client *AVTransport1) SetAVTransportURI(request AVTransport1SetAVTransportURIRequest) (response *AVTransport1SetAVTransportURIResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "SetAVTransportURI", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1SetNextAVTransportURIRequest describes the request for AVTransport1.SetNextAVTransportURI API
type AVTransport1SetNextAVTransportURIRequest struct {
	InstanceID      soap.Ui4
	NextURI         soap.String
	NextURIMetaData soap.String
}

// AVTransport1SetNextAVTransportURIResponse describes the response for AVTransport1.SetNextAVTransportURI API
type AVTransport1SetNextAVTransportURIResponse struct {
}

func (client *AVTransport1) SetNextAVTransportURI(request AVTransport1SetNextAVTransportURIRequest) (response *AVTransport1SetNextAVTransportURIResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "SetNextAVTransportURI", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1GetMediaInfoRequest describes the request for AVTransport1.GetMediaInfo API
type AVTransport1GetMediaInfoRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1GetMediaInfoResponse describes the response for AVTransport1.GetMediaInfo API
type AVTransport1GetMediaInfoResponse struct {
	// NrTracks: allowed value range: minimum=0
	NrTracks           soap.Ui4
	MediaDuration      soap.String
	CurrentURI         soap.String
	CurrentURIMetaData soap.String
	NextURI            soap.String
	NextURIMetaData    soap.String
	PlayMedium         soap.String
	RecordMedium       soap.String
	WriteStatus        soap.String
}

//
// Return value:
//
//  AVTransport1GetMediaInfoResponse
func (client *AVTransport1) GetMediaInfo(request AVTransport1GetMediaInfoRequest) (response *AVTransport1GetMediaInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "GetMediaInfo", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1GetTransportInfoRequest describes the request for AVTransport1.GetTransportInfo API
type AVTransport1GetTransportInfoRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1GetTransportInfoResponse describes the response for AVTransport1.GetTransportInfo API
type AVTransport1GetTransportInfoResponse struct {
	// CurrentTransportState: allowed values: STOPPED, PLAYING
	CurrentTransportState soap.String
	// CurrentTransportStatus: allowed values: OK, ERROR_OCCURRED
	CurrentTransportStatus soap.String
	// CurrentSpeed: allowed values: 1
	CurrentSpeed soap.String
}

//
// Return value:
//
//  AVTransport1GetTransportInfoResponse
func (client *AVTransport1) GetTransportInfo(request AVTransport1GetTransportInfoRequest) (response *AVTransport1GetTransportInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "GetTransportInfo", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1GetPositionInfoRequest describes the request for AVTransport1.GetPositionInfo API
type AVTransport1GetPositionInfoRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1GetPositionInfoResponse describes the response for AVTransport1.GetPositionInfo API
type AVTransport1GetPositionInfoResponse struct {
	// Track: allowed value range: minimum=0, step=1
	Track         soap.Ui4
	TrackDuration soap.String
	TrackMetaData soap.String
	TrackURI      soap.String
	RelTime       soap.String
	AbsTime       soap.String
	RelCount      soap.I4
	AbsCount      soap.I4
}

//
// Return value:
//
//  AVTransport1GetPositionInfoResponse
func (client *AVTransport1) GetPositionInfo(request AVTransport1GetPositionInfoRequest) (response *AVTransport1GetPositionInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "GetPositionInfo", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1GetDeviceCapabilitiesRequest describes the request for AVTransport1.GetDeviceCapabilities API
type AVTransport1GetDeviceCapabilitiesRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1GetDeviceCapabilitiesResponse describes the response for AVTransport1.GetDeviceCapabilities API
type AVTransport1GetDeviceCapabilitiesResponse struct {
	PlayMedia       soap.String
	RecMedia        soap.String
	RecQualityModes soap.String
}

func (client *AVTransport1) GetDeviceCapabilities(request AVTransport1GetDeviceCapabilitiesRequest) (response *AVTransport1GetDeviceCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "GetDeviceCapabilities", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1GetTransportSettingsRequest describes the request for AVTransport1.GetTransportSettings API
type AVTransport1GetTransportSettingsRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1GetTransportSettingsResponse describes the response for AVTransport1.GetTransportSettings API
type AVTransport1GetTransportSettingsResponse struct {
	// PlayMode: allowed values: NORMAL
	PlayMode       soap.String
	RecQualityMode soap.String
}

//
// Return value:
//
//  AVTransport1GetTransportSettingsResponse
func (client *AVTransport1) GetTransportSettings(request AVTransport1GetTransportSettingsRequest) (response *AVTransport1GetTransportSettingsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "GetTransportSettings", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1StopRequest describes the request for AVTransport1.Stop API
type AVTransport1StopRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1StopResponse describes the response for AVTransport1.Stop API
type AVTransport1StopResponse struct {
}

func (client *AVTransport1) Stop(request AVTransport1StopRequest) (response *AVTransport1StopResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "Stop", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1PlayRequest describes the request for AVTransport1.Play API
type AVTransport1PlayRequest struct {
	InstanceID soap.Ui4
	// Speed: allowed values: 1
	Speed soap.String
}

// AVTransport1PlayResponse describes the response for AVTransport1.Play API
type AVTransport1PlayResponse struct {
}

//
// Arguments:
//
//  AVTransport1PlayRequest
func (client *AVTransport1) Play(request AVTransport1PlayRequest) (response *AVTransport1PlayResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "Play", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1PauseRequest describes the request for AVTransport1.Pause API
type AVTransport1PauseRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1PauseResponse describes the response for AVTransport1.Pause API
type AVTransport1PauseResponse struct {
}

func (client *AVTransport1) Pause(request AVTransport1PauseRequest) (response *AVTransport1PauseResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "Pause", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1RecordRequest describes the request for AVTransport1.Record API
type AVTransport1RecordRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1RecordResponse describes the response for AVTransport1.Record API
type AVTransport1RecordResponse struct {
}

func (client *AVTransport1) Record(request AVTransport1RecordRequest) (response *AVTransport1RecordResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "Record", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1SeekRequest describes the request for AVTransport1.Seek API
type AVTransport1SeekRequest struct {
	InstanceID soap.Ui4
	// Unit: allowed values: TRACK_NR
	Unit   soap.String
	Target soap.String
}

// AVTransport1SeekResponse describes the response for AVTransport1.Seek API
type AVTransport1SeekResponse struct {
}

//
// Arguments:
//
//  AVTransport1SeekRequest
func (client *AVTransport1) Seek(request AVTransport1SeekRequest) (response *AVTransport1SeekResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "Seek", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1NextRequest describes the request for AVTransport1.Next API
type AVTransport1NextRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1NextResponse describes the response for AVTransport1.Next API
type AVTransport1NextResponse struct {
}

func (client *AVTransport1) Next(request AVTransport1NextRequest) (response *AVTransport1NextResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "Next", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1PreviousRequest describes the request for AVTransport1.Previous API
type AVTransport1PreviousRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1PreviousResponse describes the response for AVTransport1.Previous API
type AVTransport1PreviousResponse struct {
}

func (client *AVTransport1) Previous(request AVTransport1PreviousRequest) (response *AVTransport1PreviousResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "Previous", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1SetPlayModeRequest describes the request for AVTransport1.SetPlayMode API
type AVTransport1SetPlayModeRequest struct {
	InstanceID soap.Ui4
	// NewPlayMode: allowed values: NORMAL
	NewPlayMode soap.String
}

// AVTransport1SetPlayModeResponse describes the response for AVTransport1.SetPlayMode API
type AVTransport1SetPlayModeResponse struct {
}

//
// Arguments:
//
//  AVTransport1SetPlayModeRequest
func (client *AVTransport1) SetPlayMode(request AVTransport1SetPlayModeRequest) (response *AVTransport1SetPlayModeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "SetPlayMode", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1SetRecordQualityModeRequest describes the request for AVTransport1.SetRecordQualityMode API
type AVTransport1SetRecordQualityModeRequest struct {
	InstanceID           soap.Ui4
	NewRecordQualityMode soap.String
}

// AVTransport1SetRecordQualityModeResponse describes the response for AVTransport1.SetRecordQualityMode API
type AVTransport1SetRecordQualityModeResponse struct {
}

func (client *AVTransport1) SetRecordQualityMode(request AVTransport1SetRecordQualityModeRequest) (response *AVTransport1SetRecordQualityModeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "SetRecordQualityMode", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport1GetCurrentTransportActionsRequest describes the request for AVTransport1.GetCurrentTransportActions API
type AVTransport1GetCurrentTransportActionsRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1GetCurrentTransportActionsResponse describes the response for AVTransport1.GetCurrentTransportActions API
type AVTransport1GetCurrentTransportActionsResponse struct {
	Actions soap.String
}

func (client *AVTransport1) GetCurrentTransportActions(request AVTransport1GetCurrentTransportActionsRequest) (response *AVTransport1GetCurrentTransportActionsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_1, "GetCurrentTransportActions", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:AVTransport:2".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type AVTransport2 struct {
	goupnp.ServiceClient
}

// NewAVTransport2Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewAVTransport2Clients() (clients []*AVTransport2, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_AVTransport_2); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newAVTransport2ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewAVTransport2ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewAVTransport2ClientsByURL(loc *url.URL) ([]*AVTransport2, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_AVTransport_2)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newAVTransport2ClientsFromGenericClients(genericClients), nil
}

// NewAVTransport2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewAVTransport2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*AVTransport2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_AVTransport_2)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newAVTransport2ClientsFromGenericClients(genericClients), nil
}

func newAVTransport2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*AVTransport2 {
	clients := make([]*AVTransport2, len(genericClients))
	for i := range genericClients {
		clients[i] = &AVTransport2{genericClients[i]}
	}
	return clients
}

// AVTransport2SetAVTransportURIRequest describes the request for AVTransport2.SetAVTransportURI API
type AVTransport2SetAVTransportURIRequest struct {
	InstanceID         soap.Ui4
	CurrentURI         soap.String
	CurrentURIMetaData soap.String
}

// AVTransport2SetAVTransportURIResponse describes the response for AVTransport2.SetAVTransportURI API
type AVTransport2SetAVTransportURIResponse struct {
}

func (client *AVTransport2) SetAVTransportURI(request AVTransport2SetAVTransportURIRequest) (response *AVTransport2SetAVTransportURIResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "SetAVTransportURI", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2SetNextAVTransportURIRequest describes the request for AVTransport2.SetNextAVTransportURI API
type AVTransport2SetNextAVTransportURIRequest struct {
	InstanceID      soap.Ui4
	NextURI         soap.String
	NextURIMetaData soap.String
}

// AVTransport2SetNextAVTransportURIResponse describes the response for AVTransport2.SetNextAVTransportURI API
type AVTransport2SetNextAVTransportURIResponse struct {
}

func (client *AVTransport2) SetNextAVTransportURI(request AVTransport2SetNextAVTransportURIRequest) (response *AVTransport2SetNextAVTransportURIResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "SetNextAVTransportURI", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2GetMediaInfoRequest describes the request for AVTransport2.GetMediaInfo API
type AVTransport2GetMediaInfoRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2GetMediaInfoResponse describes the response for AVTransport2.GetMediaInfo API
type AVTransport2GetMediaInfoResponse struct {
	// NrTracks: allowed value range: minimum=0
	NrTracks           soap.Ui4
	MediaDuration      soap.String
	CurrentURI         soap.String
	CurrentURIMetaData soap.String
	NextURI            soap.String
	NextURIMetaData    soap.String
	PlayMedium         soap.String
	RecordMedium       soap.String
	WriteStatus        soap.String
}

//
// Return value:
//
//  AVTransport2GetMediaInfoResponse
func (client *AVTransport2) GetMediaInfo(request AVTransport2GetMediaInfoRequest) (response *AVTransport2GetMediaInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "GetMediaInfo", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2GetMediaInfo_ExtRequest describes the request for AVTransport2.GetMediaInfo_Ext API
type AVTransport2GetMediaInfo_ExtRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2GetMediaInfo_ExtResponse describes the response for AVTransport2.GetMediaInfo_Ext API
type AVTransport2GetMediaInfo_ExtResponse struct {
	// CurrentType: allowed values: NO_MEDIA, TRACK_AWARE, TRACK_UNAWARE
	CurrentType soap.String
	// NrTracks: allowed value range: minimum=0
	NrTracks           soap.Ui4
	MediaDuration      soap.String
	CurrentURI         soap.String
	CurrentURIMetaData soap.String
	NextURI            soap.String
	NextURIMetaData    soap.String
	PlayMedium         soap.String
	RecordMedium       soap.String
	WriteStatus        soap.String
}

//
// Return value:
//
//  AVTransport2GetMediaInfo_ExtResponse
func (client *AVTransport2) GetMediaInfo_Ext(request AVTransport2GetMediaInfo_ExtRequest) (response *AVTransport2GetMediaInfo_ExtResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "GetMediaInfo_Ext", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2GetTransportInfoRequest describes the request for AVTransport2.GetTransportInfo API
type AVTransport2GetTransportInfoRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2GetTransportInfoResponse describes the response for AVTransport2.GetTransportInfo API
type AVTransport2GetTransportInfoResponse struct {
	// CurrentTransportState: allowed values: STOPPED, PLAYING
	CurrentTransportState soap.String
	// CurrentTransportStatus: allowed values: OK, ERROR_OCCURRED
	CurrentTransportStatus soap.String
	// CurrentSpeed: allowed values: 1
	CurrentSpeed soap.String
}

//
// Return value:
//
//  AVTransport2GetTransportInfoResponse
func (client *AVTransport2) GetTransportInfo(request AVTransport2GetTransportInfoRequest) (response *AVTransport2GetTransportInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "GetTransportInfo", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2GetPositionInfoRequest describes the request for AVTransport2.GetPositionInfo API
type AVTransport2GetPositionInfoRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2GetPositionInfoResponse describes the response for AVTransport2.GetPositionInfo API
type AVTransport2GetPositionInfoResponse struct {
	// Track: allowed value range: minimum=0, step=1
	Track         soap.Ui4
	TrackDuration soap.String
	TrackMetaData soap.String
	TrackURI      soap.String
	RelTime       soap.String
	AbsTime       soap.String
	RelCount      soap.I4
	AbsCount      soap.I4
}

//
// Return value:
//
//  AVTransport2GetPositionInfoResponse
func (client *AVTransport2) GetPositionInfo(request AVTransport2GetPositionInfoRequest) (response *AVTransport2GetPositionInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "GetPositionInfo", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2GetDeviceCapabilitiesRequest describes the request for AVTransport2.GetDeviceCapabilities API
type AVTransport2GetDeviceCapabilitiesRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2GetDeviceCapabilitiesResponse describes the response for AVTransport2.GetDeviceCapabilities API
type AVTransport2GetDeviceCapabilitiesResponse struct {
	PlayMedia       soap.String
	RecMedia        soap.String
	RecQualityModes soap.String
}

func (client *AVTransport2) GetDeviceCapabilities(request AVTransport2GetDeviceCapabilitiesRequest) (response *AVTransport2GetDeviceCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "GetDeviceCapabilities", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2GetTransportSettingsRequest describes the request for AVTransport2.GetTransportSettings API
type AVTransport2GetTransportSettingsRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2GetTransportSettingsResponse describes the response for AVTransport2.GetTransportSettings API
type AVTransport2GetTransportSettingsResponse struct {
	// PlayMode: allowed values: NORMAL
	PlayMode       soap.String
	RecQualityMode soap.String
}

//
// Return value:
//
//  AVTransport2GetTransportSettingsResponse
func (client *AVTransport2) GetTransportSettings(request AVTransport2GetTransportSettingsRequest) (response *AVTransport2GetTransportSettingsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "GetTransportSettings", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2StopRequest describes the request for AVTransport2.Stop API
type AVTransport2StopRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2StopResponse describes the response for AVTransport2.Stop API
type AVTransport2StopResponse struct {
}

func (client *AVTransport2) Stop(request AVTransport2StopRequest) (response *AVTransport2StopResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "Stop", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2PlayRequest describes the request for AVTransport2.Play API
type AVTransport2PlayRequest struct {
	InstanceID soap.Ui4
	// Speed: allowed values: 1
	Speed soap.String
}

// AVTransport2PlayResponse describes the response for AVTransport2.Play API
type AVTransport2PlayResponse struct {
}

//
// Arguments:
//
//  AVTransport2PlayRequest
func (client *AVTransport2) Play(request AVTransport2PlayRequest) (response *AVTransport2PlayResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "Play", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2PauseRequest describes the request for AVTransport2.Pause API
type AVTransport2PauseRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2PauseResponse describes the response for AVTransport2.Pause API
type AVTransport2PauseResponse struct {
}

func (client *AVTransport2) Pause(request AVTransport2PauseRequest) (response *AVTransport2PauseResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "Pause", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2RecordRequest describes the request for AVTransport2.Record API
type AVTransport2RecordRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2RecordResponse describes the response for AVTransport2.Record API
type AVTransport2RecordResponse struct {
}

func (client *AVTransport2) Record(request AVTransport2RecordRequest) (response *AVTransport2RecordResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "Record", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2SeekRequest describes the request for AVTransport2.Seek API
type AVTransport2SeekRequest struct {
	InstanceID soap.Ui4
	// Unit: allowed values: TRACK_NR
	Unit   soap.String
	Target soap.String
}

// AVTransport2SeekResponse describes the response for AVTransport2.Seek API
type AVTransport2SeekResponse struct {
}

//
// Arguments:
//
//  AVTransport2SeekRequest
func (client *AVTransport2) Seek(request AVTransport2SeekRequest) (response *AVTransport2SeekResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "Seek", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2NextRequest describes the request for AVTransport2.Next API
type AVTransport2NextRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2NextResponse describes the response for AVTransport2.Next API
type AVTransport2NextResponse struct {
}

func (client *AVTransport2) Next(request AVTransport2NextRequest) (response *AVTransport2NextResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "Next", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2PreviousRequest describes the request for AVTransport2.Previous API
type AVTransport2PreviousRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2PreviousResponse describes the response for AVTransport2.Previous API
type AVTransport2PreviousResponse struct {
}

func (client *AVTransport2) Previous(request AVTransport2PreviousRequest) (response *AVTransport2PreviousResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "Previous", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2SetPlayModeRequest describes the request for AVTransport2.SetPlayMode API
type AVTransport2SetPlayModeRequest struct {
	InstanceID soap.Ui4
	// NewPlayMode: allowed values: NORMAL
	NewPlayMode soap.String
}

// AVTransport2SetPlayModeResponse describes the response for AVTransport2.SetPlayMode API
type AVTransport2SetPlayModeResponse struct {
}

//
// Arguments:
//
//  AVTransport2SetPlayModeRequest
func (client *AVTransport2) SetPlayMode(request AVTransport2SetPlayModeRequest) (response *AVTransport2SetPlayModeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "SetPlayMode", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2SetRecordQualityModeRequest describes the request for AVTransport2.SetRecordQualityMode API
type AVTransport2SetRecordQualityModeRequest struct {
	InstanceID           soap.Ui4
	NewRecordQualityMode soap.String
}

// AVTransport2SetRecordQualityModeResponse describes the response for AVTransport2.SetRecordQualityMode API
type AVTransport2SetRecordQualityModeResponse struct {
}

func (client *AVTransport2) SetRecordQualityMode(request AVTransport2SetRecordQualityModeRequest) (response *AVTransport2SetRecordQualityModeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "SetRecordQualityMode", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2GetCurrentTransportActionsRequest describes the request for AVTransport2.GetCurrentTransportActions API
type AVTransport2GetCurrentTransportActionsRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2GetCurrentTransportActionsResponse describes the response for AVTransport2.GetCurrentTransportActions API
type AVTransport2GetCurrentTransportActionsResponse struct {
	Actions soap.String
}

func (client *AVTransport2) GetCurrentTransportActions(request AVTransport2GetCurrentTransportActionsRequest) (response *AVTransport2GetCurrentTransportActionsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "GetCurrentTransportActions", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2GetDRMStateRequest describes the request for AVTransport2.GetDRMState API
type AVTransport2GetDRMStateRequest struct {
	InstanceID soap.Ui4
}

// AVTransport2GetDRMStateResponse describes the response for AVTransport2.GetDRMState API
type AVTransport2GetDRMStateResponse struct {
	// CurrentDRMState: allowed values: OK
	CurrentDRMState soap.String
}

//
// Return value:
//
//  AVTransport2GetDRMStateResponse
func (client *AVTransport2) GetDRMState(request AVTransport2GetDRMStateRequest) (response *AVTransport2GetDRMStateResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "GetDRMState", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2GetStateVariablesRequest describes the request for AVTransport2.GetStateVariables API
type AVTransport2GetStateVariablesRequest struct {
	InstanceID        soap.Ui4
	StateVariableList soap.String
}

// AVTransport2GetStateVariablesResponse describes the response for AVTransport2.GetStateVariables API
type AVTransport2GetStateVariablesResponse struct {
	StateVariableValuePairs soap.String
}

func (client *AVTransport2) GetStateVariables(request AVTransport2GetStateVariablesRequest) (response *AVTransport2GetStateVariablesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "GetStateVariables", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// AVTransport2SetStateVariablesRequest describes the request for AVTransport2.SetStateVariables API
type AVTransport2SetStateVariablesRequest struct {
	InstanceID              soap.Ui4
	AVTransportUDN          soap.String
	ServiceType             soap.String
	ServiceId               soap.String
	StateVariableValuePairs soap.String
}

// AVTransport2SetStateVariablesResponse describes the response for AVTransport2.SetStateVariables API
type AVTransport2SetStateVariablesResponse struct {
	StateVariableList soap.String
}

func (client *AVTransport2) SetStateVariables(request AVTransport2SetStateVariablesRequest) (response *AVTransport2SetStateVariablesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_AVTransport_2, "SetStateVariables", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ConnectionManager1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ConnectionManager:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ConnectionManager1 struct {
	goupnp.ServiceClient
}

// NewConnectionManager1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewConnectionManager1Clients() (clients []*ConnectionManager1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_ConnectionManager_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newConnectionManager1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewConnectionManager1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewConnectionManager1ClientsByURL(loc *url.URL) ([]*ConnectionManager1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_ConnectionManager_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newConnectionManager1ClientsFromGenericClients(genericClients), nil
}

// NewConnectionManager1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewConnectionManager1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ConnectionManager1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ConnectionManager_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newConnectionManager1ClientsFromGenericClients(genericClients), nil
}

func newConnectionManager1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ConnectionManager1 {
	clients := make([]*ConnectionManager1, len(genericClients))
	for i := range genericClients {
		clients[i] = &ConnectionManager1{genericClients[i]}
	}
	return clients
}

// ConnectionManager1GetProtocolInfoRequest describes the request for ConnectionManager1.GetProtocolInfo API
type ConnectionManager1GetProtocolInfoRequest struct {
}

// ConnectionManager1GetProtocolInfoResponse describes the response for ConnectionManager1.GetProtocolInfo API
type ConnectionManager1GetProtocolInfoResponse struct {
	Source soap.String
	Sink   soap.String
}

func (client *ConnectionManager1) GetProtocolInfo(request ConnectionManager1GetProtocolInfoRequest) (response *ConnectionManager1GetProtocolInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ConnectionManager_1, "GetProtocolInfo", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ConnectionManager1PrepareForConnectionRequest describes the request for ConnectionManager1.PrepareForConnection API
type ConnectionManager1PrepareForConnectionRequest struct {
	RemoteProtocolInfo    soap.String
	PeerConnectionManager soap.String
	PeerConnectionID      soap.I4
	// Direction: allowed values: Input, Output
	Direction soap.String
}

// ConnectionManager1PrepareForConnectionResponse describes the response for ConnectionManager1.PrepareForConnection API
type ConnectionManager1PrepareForConnectionResponse struct {
	ConnectionID  soap.I4
	AVTransportID soap.I4
	RcsID         soap.I4
}

//
// Arguments:
//
//  ConnectionManager1PrepareForConnectionRequest
func (client *ConnectionManager1) PrepareForConnection(request ConnectionManager1PrepareForConnectionRequest) (response *ConnectionManager1PrepareForConnectionResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ConnectionManager_1, "PrepareForConnection", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ConnectionManager1ConnectionCompleteRequest describes the request for ConnectionManager1.ConnectionComplete API
type ConnectionManager1ConnectionCompleteRequest struct {
	ConnectionID soap.I4
}

// ConnectionManager1ConnectionCompleteResponse describes the response for ConnectionManager1.ConnectionComplete API
type ConnectionManager1ConnectionCompleteResponse struct {
}

func (client *ConnectionManager1) ConnectionComplete(request ConnectionManager1ConnectionCompleteRequest) (response *ConnectionManager1ConnectionCompleteResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ConnectionManager_1, "ConnectionComplete", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ConnectionManager1GetCurrentConnectionIDsRequest describes the request for ConnectionManager1.GetCurrentConnectionIDs API
type ConnectionManager1GetCurrentConnectionIDsRequest struct {
}

// ConnectionManager1GetCurrentConnectionIDsResponse describes the response for ConnectionManager1.GetCurrentConnectionIDs API
type ConnectionManager1GetCurrentConnectionIDsResponse struct {
	ConnectionIDs soap.String
}

func (client *ConnectionManager1) GetCurrentConnectionIDs(request ConnectionManager1GetCurrentConnectionIDsRequest) (response *ConnectionManager1GetCurrentConnectionIDsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ConnectionManager_1, "GetCurrentConnectionIDs", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ConnectionManager1GetCurrentConnectionInfoRequest describes the request for ConnectionManager1.GetCurrentConnectionInfo API
type ConnectionManager1GetCurrentConnectionInfoRequest struct {
	ConnectionID soap.I4
}

// ConnectionManager1GetCurrentConnectionInfoResponse describes the response for ConnectionManager1.GetCurrentConnectionInfo API
type ConnectionManager1GetCurrentConnectionInfoResponse struct {
	RcsID                 soap.I4
	AVTransportID         soap.I4
	ProtocolInfo          soap.String
	PeerConnectionManager soap.String
	PeerConnectionID      soap.I4
	// Direction: allowed values: Input, Output
	Direction soap.String
	// Status: allowed values: OK, ContentFormatMismatch, InsufficientBandwidth, UnreliableChannel, Unknown
	Status soap.String
}

//
// Return value:
//
//  ConnectionManager1GetCurrentConnectionInfoResponse
func (client *ConnectionManager1) GetCurrentConnectionInfo(request ConnectionManager1GetCurrentConnectionInfoRequest) (response *ConnectionManager1GetCurrentConnectionInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ConnectionManager_1, "GetCurrentConnectionInfo", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ConnectionManager2 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ConnectionManager:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ConnectionManager2 struct {
	goupnp.ServiceClient
}

// NewConnectionManager2Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewConnectionManager2Clients() (clients []*ConnectionManager2, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_ConnectionManager_2); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newConnectionManager2ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewConnectionManager2ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewConnectionManager2ClientsByURL(loc *url.URL) ([]*ConnectionManager2, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_ConnectionManager_2)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newConnectionManager2ClientsFromGenericClients(genericClients), nil
}

// NewConnectionManager2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewConnectionManager2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ConnectionManager2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ConnectionManager_2)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newConnectionManager2ClientsFromGenericClients(genericClients), nil
}

func newConnectionManager2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ConnectionManager2 {
	clients := make([]*ConnectionManager2, len(genericClients))
	for i := range genericClients {
		clients[i] = &ConnectionManager2{genericClients[i]}
	}
	return clients
}

// ConnectionManager2GetProtocolInfoRequest describes the request for ConnectionManager2.GetProtocolInfo API
type ConnectionManager2GetProtocolInfoRequest struct {
}

// ConnectionManager2GetProtocolInfoResponse describes the response for ConnectionManager2.GetProtocolInfo API
type ConnectionManager2GetProtocolInfoResponse struct {
	Source soap.String
	Sink   soap.String
}

func (client *ConnectionManager2) GetProtocolInfo(request ConnectionManager2GetProtocolInfoRequest) (response *ConnectionManager2GetProtocolInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ConnectionManager_2, "GetProtocolInfo", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ConnectionManager2PrepareForConnectionRequest describes the request for ConnectionManager2.PrepareForConnection API
type ConnectionManager2PrepareForConnectionRequest struct {
	RemoteProtocolInfo    soap.String
	PeerConnectionManager soap.String
	PeerConnectionID      soap.I4
	// Direction: allowed values: Input, Output
	Direction soap.String
}

// ConnectionManager2PrepareForConnectionResponse describes the response for ConnectionManager2.PrepareForConnection API
type ConnectionManager2PrepareForConnectionResponse struct {
	ConnectionID  soap.I4
	AVTransportID soap.I4
	RcsID         soap.I4
}

//
// Arguments:
//
//  ConnectionManager2PrepareForConnectionRequest
func (client *ConnectionManager2) PrepareForConnection(request ConnectionManager2PrepareForConnectionRequest) (response *ConnectionManager2PrepareForConnectionResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ConnectionManager_2, "PrepareForConnection", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ConnectionManager2ConnectionCompleteRequest describes the request for ConnectionManager2.ConnectionComplete API
type ConnectionManager2ConnectionCompleteRequest struct {
	ConnectionID soap.I4
}

// ConnectionManager2ConnectionCompleteResponse describes the response for ConnectionManager2.ConnectionComplete API
type ConnectionManager2ConnectionCompleteResponse struct {
}

func (client *ConnectionManager2) ConnectionComplete(request ConnectionManager2ConnectionCompleteRequest) (response *ConnectionManager2ConnectionCompleteResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ConnectionManager_2, "ConnectionComplete", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ConnectionManager2GetCurrentConnectionIDsRequest describes the request for ConnectionManager2.GetCurrentConnectionIDs API
type ConnectionManager2GetCurrentConnectionIDsRequest struct {
}

// ConnectionManager2GetCurrentConnectionIDsResponse describes the response for ConnectionManager2.GetCurrentConnectionIDs API
type ConnectionManager2GetCurrentConnectionIDsResponse struct {
	ConnectionIDs soap.String
}

func (client *ConnectionManager2) GetCurrentConnectionIDs(request ConnectionManager2GetCurrentConnectionIDsRequest) (response *ConnectionManager2GetCurrentConnectionIDsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ConnectionManager_2, "GetCurrentConnectionIDs", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ConnectionManager2GetCurrentConnectionInfoRequest describes the request for ConnectionManager2.GetCurrentConnectionInfo API
type ConnectionManager2GetCurrentConnectionInfoRequest struct {
	ConnectionID soap.I4
}

// ConnectionManager2GetCurrentConnectionInfoResponse describes the response for ConnectionManager2.GetCurrentConnectionInfo API
type ConnectionManager2GetCurrentConnectionInfoResponse struct {
	RcsID                 soap.I4
	AVTransportID         soap.I4
	ProtocolInfo          soap.String
	PeerConnectionManager soap.String
	PeerConnectionID      soap.I4
	// Direction: allowed values: Input, Output
	Direction soap.String
	// Status: allowed values: OK, ContentFormatMismatch, InsufficientBandwidth, UnreliableChannel, Unknown
	Status soap.String
}

//
// Return value:
//
//  ConnectionManager2GetCurrentConnectionInfoResponse
func (client *ConnectionManager2) GetCurrentConnectionInfo(request ConnectionManager2GetCurrentConnectionInfoRequest) (response *ConnectionManager2GetCurrentConnectionInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ConnectionManager_2, "GetCurrentConnectionInfo", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ContentDirectory:3".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ContentDirectory1 struct {
	goupnp.ServiceClient
}

// NewContentDirectory1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewContentDirectory1Clients() (clients []*ContentDirectory1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_ContentDirectory_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newContentDirectory1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewContentDirectory1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewContentDirectory1ClientsByURL(loc *url.URL) ([]*ContentDirectory1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_ContentDirectory_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newContentDirectory1ClientsFromGenericClients(genericClients), nil
}

// NewContentDirectory1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewContentDirectory1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ContentDirectory1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ContentDirectory_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newContentDirectory1ClientsFromGenericClients(genericClients), nil
}

func newContentDirectory1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ContentDirectory1 {
	clients := make([]*ContentDirectory1, len(genericClients))
	for i := range genericClients {
		clients[i] = &ContentDirectory1{genericClients[i]}
	}
	return clients
}

// ContentDirectory1GetSearchCapabilitiesRequest describes the request for ContentDirectory1.GetSearchCapabilities API
type ContentDirectory1GetSearchCapabilitiesRequest struct {
}

// ContentDirectory1GetSearchCapabilitiesResponse describes the response for ContentDirectory1.GetSearchCapabilities API
type ContentDirectory1GetSearchCapabilitiesResponse struct {
	SearchCaps soap.String
}

func (client *ContentDirectory1) GetSearchCapabilities(request ContentDirectory1GetSearchCapabilitiesRequest) (response *ContentDirectory1GetSearchCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "GetSearchCapabilities", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1GetSortCapabilitiesRequest describes the request for ContentDirectory1.GetSortCapabilities API
type ContentDirectory1GetSortCapabilitiesRequest struct {
}

// ContentDirectory1GetSortCapabilitiesResponse describes the response for ContentDirectory1.GetSortCapabilities API
type ContentDirectory1GetSortCapabilitiesResponse struct {
	SortCaps soap.String
}

func (client *ContentDirectory1) GetSortCapabilities(request ContentDirectory1GetSortCapabilitiesRequest) (response *ContentDirectory1GetSortCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "GetSortCapabilities", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1GetSystemUpdateIDRequest describes the request for ContentDirectory1.GetSystemUpdateID API
type ContentDirectory1GetSystemUpdateIDRequest struct {
}

// ContentDirectory1GetSystemUpdateIDResponse describes the response for ContentDirectory1.GetSystemUpdateID API
type ContentDirectory1GetSystemUpdateIDResponse struct {
	Id soap.Ui4
}

func (client *ContentDirectory1) GetSystemUpdateID(request ContentDirectory1GetSystemUpdateIDRequest) (response *ContentDirectory1GetSystemUpdateIDResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "GetSystemUpdateID", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1BrowseRequest describes the request for ContentDirectory1.Browse API
type ContentDirectory1BrowseRequest struct {
	ObjectID soap.String
	// BrowseFlag: allowed values: BrowseMetadata, BrowseDirectChildren
	BrowseFlag     soap.String
	Filter         soap.String
	StartingIndex  soap.Ui4
	RequestedCount soap.Ui4
	SortCriteria   soap.String
}

// ContentDirectory1BrowseResponse describes the response for ContentDirectory1.Browse API
type ContentDirectory1BrowseResponse struct {
	Result         soap.String
	NumberReturned soap.Ui4
	TotalMatches   soap.Ui4
	UpdateID       soap.Ui4
}

//
// Arguments:
//
//  ContentDirectory1BrowseRequest
func (client *ContentDirectory1) Browse(request ContentDirectory1BrowseRequest) (response *ContentDirectory1BrowseResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "Browse", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1SearchRequest describes the request for ContentDirectory1.Search API
type ContentDirectory1SearchRequest struct {
	ContainerID    soap.String
	SearchCriteria soap.String
	Filter         soap.String
	StartingIndex  soap.Ui4
	RequestedCount soap.Ui4
	SortCriteria   soap.String
}

// ContentDirectory1SearchResponse describes the response for ContentDirectory1.Search API
type ContentDirectory1SearchResponse struct {
	Result         soap.String
	NumberReturned soap.Ui4
	TotalMatches   soap.Ui4
	UpdateID       soap.Ui4
}

func (client *ContentDirectory1) Search(request ContentDirectory1SearchRequest) (response *ContentDirectory1SearchResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "Search", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1CreateObjectRequest describes the request for ContentDirectory1.CreateObject API
type ContentDirectory1CreateObjectRequest struct {
	ContainerID soap.String
	Elements    soap.String
}

// ContentDirectory1CreateObjectResponse describes the response for ContentDirectory1.CreateObject API
type ContentDirectory1CreateObjectResponse struct {
	ObjectID soap.String
	Result   soap.String
}

func (client *ContentDirectory1) CreateObject(request ContentDirectory1CreateObjectRequest) (response *ContentDirectory1CreateObjectResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "CreateObject", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1DestroyObjectRequest describes the request for ContentDirectory1.DestroyObject API
type ContentDirectory1DestroyObjectRequest struct {
	ObjectID soap.String
}

// ContentDirectory1DestroyObjectResponse describes the response for ContentDirectory1.DestroyObject API
type ContentDirectory1DestroyObjectResponse struct {
}

func (client *ContentDirectory1) DestroyObject(request ContentDirectory1DestroyObjectRequest) (response *ContentDirectory1DestroyObjectResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "DestroyObject", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1UpdateObjectRequest describes the request for ContentDirectory1.UpdateObject API
type ContentDirectory1UpdateObjectRequest struct {
	ObjectID        soap.String
	CurrentTagValue soap.String
	NewTagValue     soap.String
}

// ContentDirectory1UpdateObjectResponse describes the response for ContentDirectory1.UpdateObject API
type ContentDirectory1UpdateObjectResponse struct {
}

func (client *ContentDirectory1) UpdateObject(request ContentDirectory1UpdateObjectRequest) (response *ContentDirectory1UpdateObjectResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "UpdateObject", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1ImportResourceRequest describes the request for ContentDirectory1.ImportResource API
type ContentDirectory1ImportResourceRequest struct {
	SourceURI      soap.URI
	DestinationURI soap.URI
}

// ContentDirectory1ImportResourceResponse describes the response for ContentDirectory1.ImportResource API
type ContentDirectory1ImportResourceResponse struct {
	TransferID soap.Ui4
}

func (client *ContentDirectory1) ImportResource(request ContentDirectory1ImportResourceRequest) (response *ContentDirectory1ImportResourceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "ImportResource", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1ExportResourceRequest describes the request for ContentDirectory1.ExportResource API
type ContentDirectory1ExportResourceRequest struct {
	SourceURI      soap.URI
	DestinationURI soap.URI
}

// ContentDirectory1ExportResourceResponse describes the response for ContentDirectory1.ExportResource API
type ContentDirectory1ExportResourceResponse struct {
	TransferID soap.Ui4
}

func (client *ContentDirectory1) ExportResource(request ContentDirectory1ExportResourceRequest) (response *ContentDirectory1ExportResourceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "ExportResource", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1StopTransferResourceRequest describes the request for ContentDirectory1.StopTransferResource API
type ContentDirectory1StopTransferResourceRequest struct {
	TransferID soap.Ui4
}

// ContentDirectory1StopTransferResourceResponse describes the response for ContentDirectory1.StopTransferResource API
type ContentDirectory1StopTransferResourceResponse struct {
}

func (client *ContentDirectory1) StopTransferResource(request ContentDirectory1StopTransferResourceRequest) (response *ContentDirectory1StopTransferResourceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "StopTransferResource", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1GetTransferProgressRequest describes the request for ContentDirectory1.GetTransferProgress API
type ContentDirectory1GetTransferProgressRequest struct {
	TransferID soap.Ui4
}

// ContentDirectory1GetTransferProgressResponse describes the response for ContentDirectory1.GetTransferProgress API
type ContentDirectory1GetTransferProgressResponse struct {
	// TransferStatus: allowed values: COMPLETED, ERROR, IN_PROGRESS, STOPPED
	TransferStatus soap.String
	TransferLength soap.String
	TransferTotal  soap.String
}

//
// Return value:
//
//  ContentDirectory1GetTransferProgressResponse
func (client *ContentDirectory1) GetTransferProgress(request ContentDirectory1GetTransferProgressRequest) (response *ContentDirectory1GetTransferProgressResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "GetTransferProgress", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1DeleteResourceRequest describes the request for ContentDirectory1.DeleteResource API
type ContentDirectory1DeleteResourceRequest struct {
	ResourceURI soap.URI
}

// ContentDirectory1DeleteResourceResponse describes the response for ContentDirectory1.DeleteResource API
type ContentDirectory1DeleteResourceResponse struct {
}

func (client *ContentDirectory1) DeleteResource(request ContentDirectory1DeleteResourceRequest) (response *ContentDirectory1DeleteResourceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "DeleteResource", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory1CreateReferenceRequest describes the request for ContentDirectory1.CreateReference API
type ContentDirectory1CreateReferenceRequest struct {
	ContainerID soap.String
	ObjectID    soap.String
}

// ContentDirectory1CreateReferenceResponse describes the response for ContentDirectory1.CreateReference API
type ContentDirectory1CreateReferenceResponse struct {
	NewID soap.String
}

func (client *ContentDirectory1) CreateReference(request ContentDirectory1CreateReferenceRequest) (response *ContentDirectory1CreateReferenceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_1, "CreateReference", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ContentDirectory:2".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ContentDirectory2 struct {
	goupnp.ServiceClient
}

// NewContentDirectory2Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewContentDirectory2Clients() (clients []*ContentDirectory2, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_ContentDirectory_2); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newContentDirectory2ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewContentDirectory2ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewContentDirectory2ClientsByURL(loc *url.URL) ([]*ContentDirectory2, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_ContentDirectory_2)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newContentDirectory2ClientsFromGenericClients(genericClients), nil
}

// NewContentDirectory2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewContentDirectory2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ContentDirectory2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ContentDirectory_2)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newContentDirectory2ClientsFromGenericClients(genericClients), nil
}

func newContentDirectory2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ContentDirectory2 {
	clients := make([]*ContentDirectory2, len(genericClients))
	for i := range genericClients {
		clients[i] = &ContentDirectory2{genericClients[i]}
	}
	return clients
}

// ContentDirectory2GetSearchCapabilitiesRequest describes the request for ContentDirectory2.GetSearchCapabilities API
type ContentDirectory2GetSearchCapabilitiesRequest struct {
}

// ContentDirectory2GetSearchCapabilitiesResponse describes the response for ContentDirectory2.GetSearchCapabilities API
type ContentDirectory2GetSearchCapabilitiesResponse struct {
	SearchCaps soap.String
}

func (client *ContentDirectory2) GetSearchCapabilities(request ContentDirectory2GetSearchCapabilitiesRequest) (response *ContentDirectory2GetSearchCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "GetSearchCapabilities", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2GetSortCapabilitiesRequest describes the request for ContentDirectory2.GetSortCapabilities API
type ContentDirectory2GetSortCapabilitiesRequest struct {
}

// ContentDirectory2GetSortCapabilitiesResponse describes the response for ContentDirectory2.GetSortCapabilities API
type ContentDirectory2GetSortCapabilitiesResponse struct {
	SortCaps soap.String
}

func (client *ContentDirectory2) GetSortCapabilities(request ContentDirectory2GetSortCapabilitiesRequest) (response *ContentDirectory2GetSortCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "GetSortCapabilities", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2GetSortExtensionCapabilitiesRequest describes the request for ContentDirectory2.GetSortExtensionCapabilities API
type ContentDirectory2GetSortExtensionCapabilitiesRequest struct {
}

// ContentDirectory2GetSortExtensionCapabilitiesResponse describes the response for ContentDirectory2.GetSortExtensionCapabilities API
type ContentDirectory2GetSortExtensionCapabilitiesResponse struct {
	SortExtensionCaps soap.String
}

func (client *ContentDirectory2) GetSortExtensionCapabilities(request ContentDirectory2GetSortExtensionCapabilitiesRequest) (response *ContentDirectory2GetSortExtensionCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "GetSortExtensionCapabilities", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2GetFeatureListRequest describes the request for ContentDirectory2.GetFeatureList API
type ContentDirectory2GetFeatureListRequest struct {
}

// ContentDirectory2GetFeatureListResponse describes the response for ContentDirectory2.GetFeatureList API
type ContentDirectory2GetFeatureListResponse struct {
	FeatureList soap.String
}

func (client *ContentDirectory2) GetFeatureList(request ContentDirectory2GetFeatureListRequest) (response *ContentDirectory2GetFeatureListResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "GetFeatureList", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2GetSystemUpdateIDRequest describes the request for ContentDirectory2.GetSystemUpdateID API
type ContentDirectory2GetSystemUpdateIDRequest struct {
}

// ContentDirectory2GetSystemUpdateIDResponse describes the response for ContentDirectory2.GetSystemUpdateID API
type ContentDirectory2GetSystemUpdateIDResponse struct {
	Id soap.Ui4
}

func (client *ContentDirectory2) GetSystemUpdateID(request ContentDirectory2GetSystemUpdateIDRequest) (response *ContentDirectory2GetSystemUpdateIDResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "GetSystemUpdateID", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2BrowseRequest describes the request for ContentDirectory2.Browse API
type ContentDirectory2BrowseRequest struct {
	ObjectID soap.String
	// BrowseFlag: allowed values: BrowseMetadata, BrowseDirectChildren
	BrowseFlag     soap.String
	Filter         soap.String
	StartingIndex  soap.Ui4
	RequestedCount soap.Ui4
	SortCriteria   soap.String
}

// ContentDirectory2BrowseResponse describes the response for ContentDirectory2.Browse API
type ContentDirectory2BrowseResponse struct {
	Result         soap.String
	NumberReturned soap.Ui4
	TotalMatches   soap.Ui4
	UpdateID       soap.Ui4
}

//
// Arguments:
//
//  ContentDirectory2BrowseRequest
func (client *ContentDirectory2) Browse(request ContentDirectory2BrowseRequest) (response *ContentDirectory2BrowseResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "Browse", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2SearchRequest describes the request for ContentDirectory2.Search API
type ContentDirectory2SearchRequest struct {
	ContainerID    soap.String
	SearchCriteria soap.String
	Filter         soap.String
	StartingIndex  soap.Ui4
	RequestedCount soap.Ui4
	SortCriteria   soap.String
}

// ContentDirectory2SearchResponse describes the response for ContentDirectory2.Search API
type ContentDirectory2SearchResponse struct {
	Result         soap.String
	NumberReturned soap.Ui4
	TotalMatches   soap.Ui4
	UpdateID       soap.Ui4
}

func (client *ContentDirectory2) Search(request ContentDirectory2SearchRequest) (response *ContentDirectory2SearchResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "Search", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2CreateObjectRequest describes the request for ContentDirectory2.CreateObject API
type ContentDirectory2CreateObjectRequest struct {
	ContainerID soap.String
	Elements    soap.String
}

// ContentDirectory2CreateObjectResponse describes the response for ContentDirectory2.CreateObject API
type ContentDirectory2CreateObjectResponse struct {
	ObjectID soap.String
	Result   soap.String
}

func (client *ContentDirectory2) CreateObject(request ContentDirectory2CreateObjectRequest) (response *ContentDirectory2CreateObjectResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "CreateObject", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2DestroyObjectRequest describes the request for ContentDirectory2.DestroyObject API
type ContentDirectory2DestroyObjectRequest struct {
	ObjectID soap.String
}

// ContentDirectory2DestroyObjectResponse describes the response for ContentDirectory2.DestroyObject API
type ContentDirectory2DestroyObjectResponse struct {
}

func (client *ContentDirectory2) DestroyObject(request ContentDirectory2DestroyObjectRequest) (response *ContentDirectory2DestroyObjectResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "DestroyObject", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2UpdateObjectRequest describes the request for ContentDirectory2.UpdateObject API
type ContentDirectory2UpdateObjectRequest struct {
	ObjectID        soap.String
	CurrentTagValue soap.String
	NewTagValue     soap.String
}

// ContentDirectory2UpdateObjectResponse describes the response for ContentDirectory2.UpdateObject API
type ContentDirectory2UpdateObjectResponse struct {
}

func (client *ContentDirectory2) UpdateObject(request ContentDirectory2UpdateObjectRequest) (response *ContentDirectory2UpdateObjectResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "UpdateObject", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2MoveObjectRequest describes the request for ContentDirectory2.MoveObject API
type ContentDirectory2MoveObjectRequest struct {
	ObjectID    soap.String
	NewParentID soap.String
}

// ContentDirectory2MoveObjectResponse describes the response for ContentDirectory2.MoveObject API
type ContentDirectory2MoveObjectResponse struct {
	NewObjectID soap.String
}

func (client *ContentDirectory2) MoveObject(request ContentDirectory2MoveObjectRequest) (response *ContentDirectory2MoveObjectResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "MoveObject", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2ImportResourceRequest describes the request for ContentDirectory2.ImportResource API
type ContentDirectory2ImportResourceRequest struct {
	SourceURI      soap.URI
	DestinationURI soap.URI
}

// ContentDirectory2ImportResourceResponse describes the response for ContentDirectory2.ImportResource API
type ContentDirectory2ImportResourceResponse struct {
	TransferID soap.Ui4
}

func (client *ContentDirectory2) ImportResource(request ContentDirectory2ImportResourceRequest) (response *ContentDirectory2ImportResourceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "ImportResource", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2ExportResourceRequest describes the request for ContentDirectory2.ExportResource API
type ContentDirectory2ExportResourceRequest struct {
	SourceURI      soap.URI
	DestinationURI soap.URI
}

// ContentDirectory2ExportResourceResponse describes the response for ContentDirectory2.ExportResource API
type ContentDirectory2ExportResourceResponse struct {
	TransferID soap.Ui4
}

func (client *ContentDirectory2) ExportResource(request ContentDirectory2ExportResourceRequest) (response *ContentDirectory2ExportResourceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "ExportResource", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2DeleteResourceRequest describes the request for ContentDirectory2.DeleteResource API
type ContentDirectory2DeleteResourceRequest struct {
	ResourceURI soap.URI
}

// ContentDirectory2DeleteResourceResponse describes the response for ContentDirectory2.DeleteResource API
type ContentDirectory2DeleteResourceResponse struct {
}

func (client *ContentDirectory2) DeleteResource(request ContentDirectory2DeleteResourceRequest) (response *ContentDirectory2DeleteResourceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "DeleteResource", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2StopTransferResourceRequest describes the request for ContentDirectory2.StopTransferResource API
type ContentDirectory2StopTransferResourceRequest struct {
	TransferID soap.Ui4
}

// ContentDirectory2StopTransferResourceResponse describes the response for ContentDirectory2.StopTransferResource API
type ContentDirectory2StopTransferResourceResponse struct {
}

func (client *ContentDirectory2) StopTransferResource(request ContentDirectory2StopTransferResourceRequest) (response *ContentDirectory2StopTransferResourceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "StopTransferResource", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2GetTransferProgressRequest describes the request for ContentDirectory2.GetTransferProgress API
type ContentDirectory2GetTransferProgressRequest struct {
	TransferID soap.Ui4
}

// ContentDirectory2GetTransferProgressResponse describes the response for ContentDirectory2.GetTransferProgress API
type ContentDirectory2GetTransferProgressResponse struct {
	// TransferStatus: allowed values: COMPLETED, ERROR, IN_PROGRESS, STOPPED
	TransferStatus soap.String
	TransferLength soap.String
	TransferTotal  soap.String
}

//
// Return value:
//
//  ContentDirectory2GetTransferProgressResponse
func (client *ContentDirectory2) GetTransferProgress(request ContentDirectory2GetTransferProgressRequest) (response *ContentDirectory2GetTransferProgressResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "GetTransferProgress", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory2CreateReferenceRequest describes the request for ContentDirectory2.CreateReference API
type ContentDirectory2CreateReferenceRequest struct {
	ContainerID soap.String
	ObjectID    soap.String
}

// ContentDirectory2CreateReferenceResponse describes the response for ContentDirectory2.CreateReference API
type ContentDirectory2CreateReferenceResponse struct {
	NewID soap.String
}

func (client *ContentDirectory2) CreateReference(request ContentDirectory2CreateReferenceRequest) (response *ContentDirectory2CreateReferenceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_2, "CreateReference", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ContentDirectory:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ContentDirectory3 struct {
	goupnp.ServiceClient
}

// NewContentDirectory3Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewContentDirectory3Clients() (clients []*ContentDirectory3, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_ContentDirectory_3); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newContentDirectory3ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewContentDirectory3ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewContentDirectory3ClientsByURL(loc *url.URL) ([]*ContentDirectory3, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_ContentDirectory_3)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newContentDirectory3ClientsFromGenericClients(genericClients), nil
}

// NewContentDirectory3ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewContentDirectory3ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ContentDirectory3, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ContentDirectory_3)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newContentDirectory3ClientsFromGenericClients(genericClients), nil
}

func newContentDirectory3ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ContentDirectory3 {
	clients := make([]*ContentDirectory3, len(genericClients))
	for i := range genericClients {
		clients[i] = &ContentDirectory3{genericClients[i]}
	}
	return clients
}

// ContentDirectory3GetSearchCapabilitiesRequest describes the request for ContentDirectory3.GetSearchCapabilities API
type ContentDirectory3GetSearchCapabilitiesRequest struct {
}

// ContentDirectory3GetSearchCapabilitiesResponse describes the response for ContentDirectory3.GetSearchCapabilities API
type ContentDirectory3GetSearchCapabilitiesResponse struct {
	SearchCaps soap.String
}

func (client *ContentDirectory3) GetSearchCapabilities(request ContentDirectory3GetSearchCapabilitiesRequest) (response *ContentDirectory3GetSearchCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "GetSearchCapabilities", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3GetSortCapabilitiesRequest describes the request for ContentDirectory3.GetSortCapabilities API
type ContentDirectory3GetSortCapabilitiesRequest struct {
}

// ContentDirectory3GetSortCapabilitiesResponse describes the response for ContentDirectory3.GetSortCapabilities API
type ContentDirectory3GetSortCapabilitiesResponse struct {
	SortCaps soap.String
}

func (client *ContentDirectory3) GetSortCapabilities(request ContentDirectory3GetSortCapabilitiesRequest) (response *ContentDirectory3GetSortCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "GetSortCapabilities", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3GetSortExtensionCapabilitiesRequest describes the request for ContentDirectory3.GetSortExtensionCapabilities API
type ContentDirectory3GetSortExtensionCapabilitiesRequest struct {
}

// ContentDirectory3GetSortExtensionCapabilitiesResponse describes the response for ContentDirectory3.GetSortExtensionCapabilities API
type ContentDirectory3GetSortExtensionCapabilitiesResponse struct {
	SortExtensionCaps soap.String
}

func (client *ContentDirectory3) GetSortExtensionCapabilities(request ContentDirectory3GetSortExtensionCapabilitiesRequest) (response *ContentDirectory3GetSortExtensionCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "GetSortExtensionCapabilities", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3GetFeatureListRequest describes the request for ContentDirectory3.GetFeatureList API
type ContentDirectory3GetFeatureListRequest struct {
}

// ContentDirectory3GetFeatureListResponse describes the response for ContentDirectory3.GetFeatureList API
type ContentDirectory3GetFeatureListResponse struct {
	FeatureList soap.String
}

func (client *ContentDirectory3) GetFeatureList(request ContentDirectory3GetFeatureListRequest) (response *ContentDirectory3GetFeatureListResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "GetFeatureList", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3GetSystemUpdateIDRequest describes the request for ContentDirectory3.GetSystemUpdateID API
type ContentDirectory3GetSystemUpdateIDRequest struct {
}

// ContentDirectory3GetSystemUpdateIDResponse describes the response for ContentDirectory3.GetSystemUpdateID API
type ContentDirectory3GetSystemUpdateIDResponse struct {
	Id soap.Ui4
}

func (client *ContentDirectory3) GetSystemUpdateID(request ContentDirectory3GetSystemUpdateIDRequest) (response *ContentDirectory3GetSystemUpdateIDResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "GetSystemUpdateID", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3GetServiceResetTokenRequest describes the request for ContentDirectory3.GetServiceResetToken API
type ContentDirectory3GetServiceResetTokenRequest struct {
}

// ContentDirectory3GetServiceResetTokenResponse describes the response for ContentDirectory3.GetServiceResetToken API
type ContentDirectory3GetServiceResetTokenResponse struct {
	ResetToken soap.String
}

func (client *ContentDirectory3) GetServiceResetToken(request ContentDirectory3GetServiceResetTokenRequest) (response *ContentDirectory3GetServiceResetTokenResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "GetServiceResetToken", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3BrowseRequest describes the request for ContentDirectory3.Browse API
type ContentDirectory3BrowseRequest struct {
	ObjectID soap.String
	// BrowseFlag: allowed values: BrowseMetadata, BrowseDirectChildren
	BrowseFlag     soap.String
	Filter         soap.String
	StartingIndex  soap.Ui4
	RequestedCount soap.Ui4
	SortCriteria   soap.String
}

// ContentDirectory3BrowseResponse describes the response for ContentDirectory3.Browse API
type ContentDirectory3BrowseResponse struct {
	Result         soap.String
	NumberReturned soap.Ui4
	TotalMatches   soap.Ui4
	UpdateID       soap.Ui4
}

//
// Arguments:
//
//  ContentDirectory3BrowseRequest
func (client *ContentDirectory3) Browse(request ContentDirectory3BrowseRequest) (response *ContentDirectory3BrowseResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "Browse", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3SearchRequest describes the request for ContentDirectory3.Search API
type ContentDirectory3SearchRequest struct {
	ContainerID    soap.String
	SearchCriteria soap.String
	Filter         soap.String
	StartingIndex  soap.Ui4
	RequestedCount soap.Ui4
	SortCriteria   soap.String
}

// ContentDirectory3SearchResponse describes the response for ContentDirectory3.Search API
type ContentDirectory3SearchResponse struct {
	Result         soap.String
	NumberReturned soap.Ui4
	TotalMatches   soap.Ui4
	UpdateID       soap.Ui4
}

func (client *ContentDirectory3) Search(request ContentDirectory3SearchRequest) (response *ContentDirectory3SearchResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "Search", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3CreateObjectRequest describes the request for ContentDirectory3.CreateObject API
type ContentDirectory3CreateObjectRequest struct {
	ContainerID soap.String
	Elements    soap.String
}

// ContentDirectory3CreateObjectResponse describes the response for ContentDirectory3.CreateObject API
type ContentDirectory3CreateObjectResponse struct {
	ObjectID soap.String
	Result   soap.String
}

func (client *ContentDirectory3) CreateObject(request ContentDirectory3CreateObjectRequest) (response *ContentDirectory3CreateObjectResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "CreateObject", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3DestroyObjectRequest describes the request for ContentDirectory3.DestroyObject API
type ContentDirectory3DestroyObjectRequest struct {
	ObjectID soap.String
}

// ContentDirectory3DestroyObjectResponse describes the response for ContentDirectory3.DestroyObject API
type ContentDirectory3DestroyObjectResponse struct {
}

func (client *ContentDirectory3) DestroyObject(request ContentDirectory3DestroyObjectRequest) (response *ContentDirectory3DestroyObjectResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "DestroyObject", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3UpdateObjectRequest describes the request for ContentDirectory3.UpdateObject API
type ContentDirectory3UpdateObjectRequest struct {
	ObjectID        soap.String
	CurrentTagValue soap.String
	NewTagValue     soap.String
}

// ContentDirectory3UpdateObjectResponse describes the response for ContentDirectory3.UpdateObject API
type ContentDirectory3UpdateObjectResponse struct {
}

func (client *ContentDirectory3) UpdateObject(request ContentDirectory3UpdateObjectRequest) (response *ContentDirectory3UpdateObjectResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "UpdateObject", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3MoveObjectRequest describes the request for ContentDirectory3.MoveObject API
type ContentDirectory3MoveObjectRequest struct {
	ObjectID    soap.String
	NewParentID soap.String
}

// ContentDirectory3MoveObjectResponse describes the response for ContentDirectory3.MoveObject API
type ContentDirectory3MoveObjectResponse struct {
	NewObjectID soap.String
}

func (client *ContentDirectory3) MoveObject(request ContentDirectory3MoveObjectRequest) (response *ContentDirectory3MoveObjectResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "MoveObject", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3ImportResourceRequest describes the request for ContentDirectory3.ImportResource API
type ContentDirectory3ImportResourceRequest struct {
	SourceURI      soap.URI
	DestinationURI soap.URI
}

// ContentDirectory3ImportResourceResponse describes the response for ContentDirectory3.ImportResource API
type ContentDirectory3ImportResourceResponse struct {
	TransferID soap.Ui4
}

func (client *ContentDirectory3) ImportResource(request ContentDirectory3ImportResourceRequest) (response *ContentDirectory3ImportResourceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "ImportResource", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3ExportResourceRequest describes the request for ContentDirectory3.ExportResource API
type ContentDirectory3ExportResourceRequest struct {
	SourceURI      soap.URI
	DestinationURI soap.URI
}

// ContentDirectory3ExportResourceResponse describes the response for ContentDirectory3.ExportResource API
type ContentDirectory3ExportResourceResponse struct {
	TransferID soap.Ui4
}

func (client *ContentDirectory3) ExportResource(request ContentDirectory3ExportResourceRequest) (response *ContentDirectory3ExportResourceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "ExportResource", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3DeleteResourceRequest describes the request for ContentDirectory3.DeleteResource API
type ContentDirectory3DeleteResourceRequest struct {
	ResourceURI soap.URI
}

// ContentDirectory3DeleteResourceResponse describes the response for ContentDirectory3.DeleteResource API
type ContentDirectory3DeleteResourceResponse struct {
}

func (client *ContentDirectory3) DeleteResource(request ContentDirectory3DeleteResourceRequest) (response *ContentDirectory3DeleteResourceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "DeleteResource", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3StopTransferResourceRequest describes the request for ContentDirectory3.StopTransferResource API
type ContentDirectory3StopTransferResourceRequest struct {
	TransferID soap.Ui4
}

// ContentDirectory3StopTransferResourceResponse describes the response for ContentDirectory3.StopTransferResource API
type ContentDirectory3StopTransferResourceResponse struct {
}

func (client *ContentDirectory3) StopTransferResource(request ContentDirectory3StopTransferResourceRequest) (response *ContentDirectory3StopTransferResourceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "StopTransferResource", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3GetTransferProgressRequest describes the request for ContentDirectory3.GetTransferProgress API
type ContentDirectory3GetTransferProgressRequest struct {
	TransferID soap.Ui4
}

// ContentDirectory3GetTransferProgressResponse describes the response for ContentDirectory3.GetTransferProgress API
type ContentDirectory3GetTransferProgressResponse struct {
	// TransferStatus: allowed values: COMPLETED, ERROR, IN_PROGRESS, STOPPED
	TransferStatus soap.String
	TransferLength soap.String
	TransferTotal  soap.String
}

//
// Return value:
//
//  ContentDirectory3GetTransferProgressResponse
func (client *ContentDirectory3) GetTransferProgress(request ContentDirectory3GetTransferProgressRequest) (response *ContentDirectory3GetTransferProgressResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "GetTransferProgress", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3CreateReferenceRequest describes the request for ContentDirectory3.CreateReference API
type ContentDirectory3CreateReferenceRequest struct {
	ContainerID soap.String
	ObjectID    soap.String
}

// ContentDirectory3CreateReferenceResponse describes the response for ContentDirectory3.CreateReference API
type ContentDirectory3CreateReferenceResponse struct {
	NewID soap.String
}

func (client *ContentDirectory3) CreateReference(request ContentDirectory3CreateReferenceRequest) (response *ContentDirectory3CreateReferenceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "CreateReference", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3FreeFormQueryRequest describes the request for ContentDirectory3.FreeFormQuery API
type ContentDirectory3FreeFormQueryRequest struct {
	ContainerID  soap.String
	CDSView      soap.Ui4
	QueryRequest soap.String
}

// ContentDirectory3FreeFormQueryResponse describes the response for ContentDirectory3.FreeFormQuery API
type ContentDirectory3FreeFormQueryResponse struct {
	QueryResult soap.String
	UpdateID    soap.Ui4
}

func (client *ContentDirectory3) FreeFormQuery(request ContentDirectory3FreeFormQueryRequest) (response *ContentDirectory3FreeFormQueryResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "FreeFormQuery", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ContentDirectory3GetFreeFormQueryCapabilitiesRequest describes the request for ContentDirectory3.GetFreeFormQueryCapabilities API
type ContentDirectory3GetFreeFormQueryCapabilitiesRequest struct {
}

// ContentDirectory3GetFreeFormQueryCapabilitiesResponse describes the response for ContentDirectory3.GetFreeFormQueryCapabilities API
type ContentDirectory3GetFreeFormQueryCapabilitiesResponse struct {
	FFQCapabilities soap.String
}

func (client *ContentDirectory3) GetFreeFormQueryCapabilities(request ContentDirectory3GetFreeFormQueryCapabilitiesRequest) (response *ContentDirectory3GetFreeFormQueryCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ContentDirectory_3, "GetFreeFormQueryCapabilities", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:RenderingControl:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type RenderingControl1 struct {
	goupnp.ServiceClient
}

// NewRenderingControl1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewRenderingControl1Clients() (clients []*RenderingControl1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_RenderingControl_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newRenderingControl1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewRenderingControl1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewRenderingControl1ClientsByURL(loc *url.URL) ([]*RenderingControl1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_RenderingControl_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newRenderingControl1ClientsFromGenericClients(genericClients), nil
}

// NewRenderingControl1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewRenderingControl1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*RenderingControl1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_RenderingControl_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newRenderingControl1ClientsFromGenericClients(genericClients), nil
}

func newRenderingControl1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*RenderingControl1 {
	clients := make([]*RenderingControl1, len(genericClients))
	for i := range genericClients {
		clients[i] = &RenderingControl1{genericClients[i]}
	}
	return clients
}

// RenderingControl1ListPresetsRequest describes the request for RenderingControl1.ListPresets API
type RenderingControl1ListPresetsRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1ListPresetsResponse describes the response for RenderingControl1.ListPresets API
type RenderingControl1ListPresetsResponse struct {
	CurrentPresetNameList soap.String
}

func (client *RenderingControl1) ListPresets(request RenderingControl1ListPresetsRequest) (response *RenderingControl1ListPresetsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "ListPresets", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SelectPresetRequest describes the request for RenderingControl1.SelectPreset API
type RenderingControl1SelectPresetRequest struct {
	InstanceID soap.Ui4
	// PresetName: allowed values: FactoryDefaults
	PresetName soap.String
}

// RenderingControl1SelectPresetResponse describes the response for RenderingControl1.SelectPreset API
type RenderingControl1SelectPresetResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SelectPresetRequest
func (client *RenderingControl1) SelectPreset(request RenderingControl1SelectPresetRequest) (response *RenderingControl1SelectPresetResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SelectPreset", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetBrightnessRequest describes the request for RenderingControl1.GetBrightness API
type RenderingControl1GetBrightnessRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetBrightnessResponse describes the response for RenderingControl1.GetBrightness API
type RenderingControl1GetBrightnessResponse struct {
	// CurrentBrightness: allowed value range: minimum=0, step=1
	CurrentBrightness soap.Ui2
}

//
// Return value:
//
//  RenderingControl1GetBrightnessResponse
func (client *RenderingControl1) GetBrightness(request RenderingControl1GetBrightnessRequest) (response *RenderingControl1GetBrightnessResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetBrightness", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetBrightnessRequest describes the request for RenderingControl1.SetBrightness API
type RenderingControl1SetBrightnessRequest struct {
	InstanceID soap.Ui4
	// DesiredBrightness: allowed value range: minimum=0, step=1
	DesiredBrightness soap.Ui2
}

// RenderingControl1SetBrightnessResponse describes the response for RenderingControl1.SetBrightness API
type RenderingControl1SetBrightnessResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetBrightnessRequest
func (client *RenderingControl1) SetBrightness(request RenderingControl1SetBrightnessRequest) (response *RenderingControl1SetBrightnessResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetBrightness", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetContrastRequest describes the request for RenderingControl1.GetContrast API
type RenderingControl1GetContrastRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetContrastResponse describes the response for RenderingControl1.GetContrast API
type RenderingControl1GetContrastResponse struct {
	// CurrentContrast: allowed value range: minimum=0, step=1
	CurrentContrast soap.Ui2
}

//
// Return value:
//
//  RenderingControl1GetContrastResponse
func (client *RenderingControl1) GetContrast(request RenderingControl1GetContrastRequest) (response *RenderingControl1GetContrastResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetContrast", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetContrastRequest describes the request for RenderingControl1.SetContrast API
type RenderingControl1SetContrastRequest struct {
	InstanceID soap.Ui4
	// DesiredContrast: allowed value range: minimum=0, step=1
	DesiredContrast soap.Ui2
}

// RenderingControl1SetContrastResponse describes the response for RenderingControl1.SetContrast API
type RenderingControl1SetContrastResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetContrastRequest
func (client *RenderingControl1) SetContrast(request RenderingControl1SetContrastRequest) (response *RenderingControl1SetContrastResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetContrast", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetSharpnessRequest describes the request for RenderingControl1.GetSharpness API
type RenderingControl1GetSharpnessRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetSharpnessResponse describes the response for RenderingControl1.GetSharpness API
type RenderingControl1GetSharpnessResponse struct {
	// CurrentSharpness: allowed value range: minimum=0, step=1
	CurrentSharpness soap.Ui2
}

//
// Return value:
//
//  RenderingControl1GetSharpnessResponse
func (client *RenderingControl1) GetSharpness(request RenderingControl1GetSharpnessRequest) (response *RenderingControl1GetSharpnessResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetSharpness", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetSharpnessRequest describes the request for RenderingControl1.SetSharpness API
type RenderingControl1SetSharpnessRequest struct {
	InstanceID soap.Ui4
	// DesiredSharpness: allowed value range: minimum=0, step=1
	DesiredSharpness soap.Ui2
}

// RenderingControl1SetSharpnessResponse describes the response for RenderingControl1.SetSharpness API
type RenderingControl1SetSharpnessResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetSharpnessRequest
func (client *RenderingControl1) SetSharpness(request RenderingControl1SetSharpnessRequest) (response *RenderingControl1SetSharpnessResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetSharpness", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetRedVideoGainRequest describes the request for RenderingControl1.GetRedVideoGain API
type RenderingControl1GetRedVideoGainRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetRedVideoGainResponse describes the response for RenderingControl1.GetRedVideoGain API
type RenderingControl1GetRedVideoGainResponse struct {
	CurrentRedVideoGain soap.Ui2
}

func (client *RenderingControl1) GetRedVideoGain(request RenderingControl1GetRedVideoGainRequest) (response *RenderingControl1GetRedVideoGainResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetRedVideoGain", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetRedVideoGainRequest describes the request for RenderingControl1.SetRedVideoGain API
type RenderingControl1SetRedVideoGainRequest struct {
	InstanceID          soap.Ui4
	DesiredRedVideoGain soap.Ui2
}

// RenderingControl1SetRedVideoGainResponse describes the response for RenderingControl1.SetRedVideoGain API
type RenderingControl1SetRedVideoGainResponse struct {
}

func (client *RenderingControl1) SetRedVideoGain(request RenderingControl1SetRedVideoGainRequest) (response *RenderingControl1SetRedVideoGainResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetRedVideoGain", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetGreenVideoGainRequest describes the request for RenderingControl1.GetGreenVideoGain API
type RenderingControl1GetGreenVideoGainRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetGreenVideoGainResponse describes the response for RenderingControl1.GetGreenVideoGain API
type RenderingControl1GetGreenVideoGainResponse struct {
	// CurrentGreenVideoGain: allowed value range: minimum=0, step=1
	CurrentGreenVideoGain soap.Ui2
}

//
// Return value:
//
//  RenderingControl1GetGreenVideoGainResponse
func (client *RenderingControl1) GetGreenVideoGain(request RenderingControl1GetGreenVideoGainRequest) (response *RenderingControl1GetGreenVideoGainResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetGreenVideoGain", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetGreenVideoGainRequest describes the request for RenderingControl1.SetGreenVideoGain API
type RenderingControl1SetGreenVideoGainRequest struct {
	InstanceID soap.Ui4
	// DesiredGreenVideoGain: allowed value range: minimum=0, step=1
	DesiredGreenVideoGain soap.Ui2
}

// RenderingControl1SetGreenVideoGainResponse describes the response for RenderingControl1.SetGreenVideoGain API
type RenderingControl1SetGreenVideoGainResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetGreenVideoGainRequest
func (client *RenderingControl1) SetGreenVideoGain(request RenderingControl1SetGreenVideoGainRequest) (response *RenderingControl1SetGreenVideoGainResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetGreenVideoGain", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetBlueVideoGainRequest describes the request for RenderingControl1.GetBlueVideoGain API
type RenderingControl1GetBlueVideoGainRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetBlueVideoGainResponse describes the response for RenderingControl1.GetBlueVideoGain API
type RenderingControl1GetBlueVideoGainResponse struct {
	// CurrentBlueVideoGain: allowed value range: minimum=0, step=1
	CurrentBlueVideoGain soap.Ui2
}

//
// Return value:
//
//  RenderingControl1GetBlueVideoGainResponse
func (client *RenderingControl1) GetBlueVideoGain(request RenderingControl1GetBlueVideoGainRequest) (response *RenderingControl1GetBlueVideoGainResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetBlueVideoGain", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetBlueVideoGainRequest describes the request for RenderingControl1.SetBlueVideoGain API
type RenderingControl1SetBlueVideoGainRequest struct {
	InstanceID soap.Ui4
	// DesiredBlueVideoGain: allowed value range: minimum=0, step=1
	DesiredBlueVideoGain soap.Ui2
}

// RenderingControl1SetBlueVideoGainResponse describes the response for RenderingControl1.SetBlueVideoGain API
type RenderingControl1SetBlueVideoGainResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetBlueVideoGainRequest
func (client *RenderingControl1) SetBlueVideoGain(request RenderingControl1SetBlueVideoGainRequest) (response *RenderingControl1SetBlueVideoGainResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetBlueVideoGain", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetRedVideoBlackLevelRequest describes the request for RenderingControl1.GetRedVideoBlackLevel API
type RenderingControl1GetRedVideoBlackLevelRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetRedVideoBlackLevelResponse describes the response for RenderingControl1.GetRedVideoBlackLevel API
type RenderingControl1GetRedVideoBlackLevelResponse struct {
	// CurrentRedVideoBlackLevel: allowed value range: minimum=0, step=1
	CurrentRedVideoBlackLevel soap.Ui2
}

//
// Return value:
//
//  RenderingControl1GetRedVideoBlackLevelResponse
func (client *RenderingControl1) GetRedVideoBlackLevel(request RenderingControl1GetRedVideoBlackLevelRequest) (response *RenderingControl1GetRedVideoBlackLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetRedVideoBlackLevel", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetRedVideoBlackLevelRequest describes the request for RenderingControl1.SetRedVideoBlackLevel API
type RenderingControl1SetRedVideoBlackLevelRequest struct {
	InstanceID soap.Ui4
	// DesiredRedVideoBlackLevel: allowed value range: minimum=0, step=1
	DesiredRedVideoBlackLevel soap.Ui2
}

// RenderingControl1SetRedVideoBlackLevelResponse describes the response for RenderingControl1.SetRedVideoBlackLevel API
type RenderingControl1SetRedVideoBlackLevelResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetRedVideoBlackLevelRequest
func (client *RenderingControl1) SetRedVideoBlackLevel(request RenderingControl1SetRedVideoBlackLevelRequest) (response *RenderingControl1SetRedVideoBlackLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetRedVideoBlackLevel", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetGreenVideoBlackLevelRequest describes the request for RenderingControl1.GetGreenVideoBlackLevel API
type RenderingControl1GetGreenVideoBlackLevelRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetGreenVideoBlackLevelResponse describes the response for RenderingControl1.GetGreenVideoBlackLevel API
type RenderingControl1GetGreenVideoBlackLevelResponse struct {
	// CurrentGreenVideoBlackLevel: allowed value range: minimum=0, step=1
	CurrentGreenVideoBlackLevel soap.Ui2
}

//
// Return value:
//
//  RenderingControl1GetGreenVideoBlackLevelResponse
func (client *RenderingControl1) GetGreenVideoBlackLevel(request RenderingControl1GetGreenVideoBlackLevelRequest) (response *RenderingControl1GetGreenVideoBlackLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetGreenVideoBlackLevel", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetGreenVideoBlackLevelRequest describes the request for RenderingControl1.SetGreenVideoBlackLevel API
type RenderingControl1SetGreenVideoBlackLevelRequest struct {
	InstanceID soap.Ui4
	// DesiredGreenVideoBlackLevel: allowed value range: minimum=0, step=1
	DesiredGreenVideoBlackLevel soap.Ui2
}

// RenderingControl1SetGreenVideoBlackLevelResponse describes the response for RenderingControl1.SetGreenVideoBlackLevel API
type RenderingControl1SetGreenVideoBlackLevelResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetGreenVideoBlackLevelRequest
func (client *RenderingControl1) SetGreenVideoBlackLevel(request RenderingControl1SetGreenVideoBlackLevelRequest) (response *RenderingControl1SetGreenVideoBlackLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetGreenVideoBlackLevel", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetBlueVideoBlackLevelRequest describes the request for RenderingControl1.GetBlueVideoBlackLevel API
type RenderingControl1GetBlueVideoBlackLevelRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetBlueVideoBlackLevelResponse describes the response for RenderingControl1.GetBlueVideoBlackLevel API
type RenderingControl1GetBlueVideoBlackLevelResponse struct {
	// CurrentBlueVideoBlackLevel: allowed value range: minimum=0, step=1
	CurrentBlueVideoBlackLevel soap.Ui2
}

//
// Return value:
//
//  RenderingControl1GetBlueVideoBlackLevelResponse
func (client *RenderingControl1) GetBlueVideoBlackLevel(request RenderingControl1GetBlueVideoBlackLevelRequest) (response *RenderingControl1GetBlueVideoBlackLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetBlueVideoBlackLevel", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetBlueVideoBlackLevelRequest describes the request for RenderingControl1.SetBlueVideoBlackLevel API
type RenderingControl1SetBlueVideoBlackLevelRequest struct {
	InstanceID soap.Ui4
	// DesiredBlueVideoBlackLevel: allowed value range: minimum=0, step=1
	DesiredBlueVideoBlackLevel soap.Ui2
}

// RenderingControl1SetBlueVideoBlackLevelResponse describes the response for RenderingControl1.SetBlueVideoBlackLevel API
type RenderingControl1SetBlueVideoBlackLevelResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetBlueVideoBlackLevelRequest
func (client *RenderingControl1) SetBlueVideoBlackLevel(request RenderingControl1SetBlueVideoBlackLevelRequest) (response *RenderingControl1SetBlueVideoBlackLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetBlueVideoBlackLevel", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetColorTemperatureRequest describes the request for RenderingControl1.GetColorTemperature API
type RenderingControl1GetColorTemperatureRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetColorTemperatureResponse describes the response for RenderingControl1.GetColorTemperature API
type RenderingControl1GetColorTemperatureResponse struct {
	// CurrentColorTemperature: allowed value range: minimum=0, step=1
	CurrentColorTemperature soap.Ui2
}

//
// Return value:
//
//  RenderingControl1GetColorTemperatureResponse
func (client *RenderingControl1) GetColorTemperature(request RenderingControl1GetColorTemperatureRequest) (response *RenderingControl1GetColorTemperatureResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetColorTemperature", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetColorTemperatureRequest describes the request for RenderingControl1.SetColorTemperature API
type RenderingControl1SetColorTemperatureRequest struct {
	InstanceID soap.Ui4
	// DesiredColorTemperature: allowed value range: minimum=0, step=1
	DesiredColorTemperature soap.Ui2
}

// RenderingControl1SetColorTemperatureResponse describes the response for RenderingControl1.SetColorTemperature API
type RenderingControl1SetColorTemperatureResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetColorTemperatureRequest
func (client *RenderingControl1) SetColorTemperature(request RenderingControl1SetColorTemperatureRequest) (response *RenderingControl1SetColorTemperatureResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetColorTemperature", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetHorizontalKeystoneRequest describes the request for RenderingControl1.GetHorizontalKeystone API
type RenderingControl1GetHorizontalKeystoneRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetHorizontalKeystoneResponse describes the response for RenderingControl1.GetHorizontalKeystone API
type RenderingControl1GetHorizontalKeystoneResponse struct {
	// CurrentHorizontalKeystone: allowed value range: step=1
	CurrentHorizontalKeystone soap.I2
}

//
// Return value:
//
//  RenderingControl1GetHorizontalKeystoneResponse
func (client *RenderingControl1) GetHorizontalKeystone(request RenderingControl1GetHorizontalKeystoneRequest) (response *RenderingControl1GetHorizontalKeystoneResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetHorizontalKeystone", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetHorizontalKeystoneRequest describes the request for RenderingControl1.SetHorizontalKeystone API
type RenderingControl1SetHorizontalKeystoneRequest struct {
	InstanceID soap.Ui4
	// DesiredHorizontalKeystone: allowed value range: step=1
	DesiredHorizontalKeystone soap.I2
}

// RenderingControl1SetHorizontalKeystoneResponse describes the response for RenderingControl1.SetHorizontalKeystone API
type RenderingControl1SetHorizontalKeystoneResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetHorizontalKeystoneRequest
func (client *RenderingControl1) SetHorizontalKeystone(request RenderingControl1SetHorizontalKeystoneRequest) (response *RenderingControl1SetHorizontalKeystoneResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetHorizontalKeystone", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetVerticalKeystoneRequest describes the request for RenderingControl1.GetVerticalKeystone API
type RenderingControl1GetVerticalKeystoneRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetVerticalKeystoneResponse describes the response for RenderingControl1.GetVerticalKeystone API
type RenderingControl1GetVerticalKeystoneResponse struct {
	// CurrentVerticalKeystone: allowed value range: step=1
	CurrentVerticalKeystone soap.I2
}

//
// Return value:
//
//  RenderingControl1GetVerticalKeystoneResponse
func (client *RenderingControl1) GetVerticalKeystone(request RenderingControl1GetVerticalKeystoneRequest) (response *RenderingControl1GetVerticalKeystoneResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetVerticalKeystone", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetVerticalKeystoneRequest describes the request for RenderingControl1.SetVerticalKeystone API
type RenderingControl1SetVerticalKeystoneRequest struct {
	InstanceID soap.Ui4
	// DesiredVerticalKeystone: allowed value range: step=1
	DesiredVerticalKeystone soap.I2
}

// RenderingControl1SetVerticalKeystoneResponse describes the response for RenderingControl1.SetVerticalKeystone API
type RenderingControl1SetVerticalKeystoneResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetVerticalKeystoneRequest
func (client *RenderingControl1) SetVerticalKeystone(request RenderingControl1SetVerticalKeystoneRequest) (response *RenderingControl1SetVerticalKeystoneResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetVerticalKeystone", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetMuteRequest describes the request for RenderingControl1.GetMute API
type RenderingControl1GetMuteRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel soap.String
}

// RenderingControl1GetMuteResponse describes the response for RenderingControl1.GetMute API
type RenderingControl1GetMuteResponse struct {
	CurrentMute soap.Bool
}

//
// Arguments:
//
//  RenderingControl1GetMuteRequest
func (client *RenderingControl1) GetMute(request RenderingControl1GetMuteRequest) (response *RenderingControl1GetMuteResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetMute", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetMuteRequest describes the request for RenderingControl1.SetMute API
type RenderingControl1SetMuteRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel     soap.String
	DesiredMute soap.Bool
}

// RenderingControl1SetMuteResponse describes the response for RenderingControl1.SetMute API
type RenderingControl1SetMuteResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetMuteRequest
func (client *RenderingControl1) SetMute(request RenderingControl1SetMuteRequest) (response *RenderingControl1SetMuteResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetMute", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetVolumeRequest describes the request for RenderingControl1.GetVolume API
type RenderingControl1GetVolumeRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel soap.String
}

// RenderingControl1GetVolumeResponse describes the response for RenderingControl1.GetVolume API
type RenderingControl1GetVolumeResponse struct {
	// CurrentVolume: allowed value range: minimum=0, step=1
	CurrentVolume soap.Ui2
}

//
// Arguments:
//
//  RenderingControl1GetVolumeRequest
//
// Return value:
//
//  RenderingControl1GetVolumeResponse
func (client *RenderingControl1) GetVolume(request RenderingControl1GetVolumeRequest) (response *RenderingControl1GetVolumeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetVolume", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetVolumeRequest describes the request for RenderingControl1.SetVolume API
type RenderingControl1SetVolumeRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel soap.String
	// DesiredVolume: allowed value range: minimum=0, step=1
	DesiredVolume soap.Ui2
}

// RenderingControl1SetVolumeResponse describes the response for RenderingControl1.SetVolume API
type RenderingControl1SetVolumeResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetVolumeRequest
func (client *RenderingControl1) SetVolume(request RenderingControl1SetVolumeRequest) (response *RenderingControl1SetVolumeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetVolume", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetVolumeDBRequest describes the request for RenderingControl1.GetVolumeDB API
type RenderingControl1GetVolumeDBRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel soap.String
}

// RenderingControl1GetVolumeDBResponse describes the response for RenderingControl1.GetVolumeDB API
type RenderingControl1GetVolumeDBResponse struct {
	CurrentVolume soap.I2
}

//
// Arguments:
//
//  RenderingControl1GetVolumeDBRequest
func (client *RenderingControl1) GetVolumeDB(request RenderingControl1GetVolumeDBRequest) (response *RenderingControl1GetVolumeDBResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetVolumeDB", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetVolumeDBRequest describes the request for RenderingControl1.SetVolumeDB API
type RenderingControl1SetVolumeDBRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel       soap.String
	DesiredVolume soap.I2
}

// RenderingControl1SetVolumeDBResponse describes the response for RenderingControl1.SetVolumeDB API
type RenderingControl1SetVolumeDBResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetVolumeDBRequest
func (client *RenderingControl1) SetVolumeDB(request RenderingControl1SetVolumeDBRequest) (response *RenderingControl1SetVolumeDBResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetVolumeDB", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetVolumeDBRangeRequest describes the request for RenderingControl1.GetVolumeDBRange API
type RenderingControl1GetVolumeDBRangeRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel soap.String
}

// RenderingControl1GetVolumeDBRangeResponse describes the response for RenderingControl1.GetVolumeDBRange API
type RenderingControl1GetVolumeDBRangeResponse struct {
	MinValue soap.I2
	MaxValue soap.I2
}

//
// Arguments:
//
//  RenderingControl1GetVolumeDBRangeRequest
func (client *RenderingControl1) GetVolumeDBRange(request RenderingControl1GetVolumeDBRangeRequest) (response *RenderingControl1GetVolumeDBRangeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetVolumeDBRange", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1GetLoudnessRequest describes the request for RenderingControl1.GetLoudness API
type RenderingControl1GetLoudnessRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel soap.String
}

// RenderingControl1GetLoudnessResponse describes the response for RenderingControl1.GetLoudness API
type RenderingControl1GetLoudnessResponse struct {
	CurrentLoudness soap.Bool
}

//
// Arguments:
//
//  RenderingControl1GetLoudnessRequest
func (client *RenderingControl1) GetLoudness(request RenderingControl1GetLoudnessRequest) (response *RenderingControl1GetLoudnessResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetLoudness", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl1SetLoudnessRequest describes the request for RenderingControl1.SetLoudness API
type RenderingControl1SetLoudnessRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel         soap.String
	DesiredLoudness soap.Bool
}

// RenderingControl1SetLoudnessResponse describes the response for RenderingControl1.SetLoudness API
type RenderingControl1SetLoudnessResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetLoudnessRequest
func (client *RenderingControl1) SetLoudness(request RenderingControl1SetLoudnessRequest) (response *RenderingControl1SetLoudnessResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetLoudness", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:RenderingControl:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type RenderingControl2 struct {
	goupnp.ServiceClient
}

// NewRenderingControl2Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewRenderingControl2Clients() (clients []*RenderingControl2, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_RenderingControl_2); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newRenderingControl2ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewRenderingControl2ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewRenderingControl2ClientsByURL(loc *url.URL) ([]*RenderingControl2, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_RenderingControl_2)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newRenderingControl2ClientsFromGenericClients(genericClients), nil
}

// NewRenderingControl2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewRenderingControl2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*RenderingControl2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_RenderingControl_2)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newRenderingControl2ClientsFromGenericClients(genericClients), nil
}

func newRenderingControl2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*RenderingControl2 {
	clients := make([]*RenderingControl2, len(genericClients))
	for i := range genericClients {
		clients[i] = &RenderingControl2{genericClients[i]}
	}
	return clients
}

// RenderingControl2ListPresetsRequest describes the request for RenderingControl2.ListPresets API
type RenderingControl2ListPresetsRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2ListPresetsResponse describes the response for RenderingControl2.ListPresets API
type RenderingControl2ListPresetsResponse struct {
	CurrentPresetNameList soap.String
}

func (client *RenderingControl2) ListPresets(request RenderingControl2ListPresetsRequest) (response *RenderingControl2ListPresetsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "ListPresets", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SelectPresetRequest describes the request for RenderingControl2.SelectPreset API
type RenderingControl2SelectPresetRequest struct {
	InstanceID soap.Ui4
	// PresetName: allowed values: FactoryDefaults
	PresetName soap.String
}

// RenderingControl2SelectPresetResponse describes the response for RenderingControl2.SelectPreset API
type RenderingControl2SelectPresetResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SelectPresetRequest
func (client *RenderingControl2) SelectPreset(request RenderingControl2SelectPresetRequest) (response *RenderingControl2SelectPresetResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SelectPreset", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetBrightnessRequest describes the request for RenderingControl2.GetBrightness API
type RenderingControl2GetBrightnessRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2GetBrightnessResponse describes the response for RenderingControl2.GetBrightness API
type RenderingControl2GetBrightnessResponse struct {
	// CurrentBrightness: allowed value range: minimum=0, step=1
	CurrentBrightness soap.Ui2
}

//
// Return value:
//
//  RenderingControl2GetBrightnessResponse
func (client *RenderingControl2) GetBrightness(request RenderingControl2GetBrightnessRequest) (response *RenderingControl2GetBrightnessResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetBrightness", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetBrightnessRequest describes the request for RenderingControl2.SetBrightness API
type RenderingControl2SetBrightnessRequest struct {
	InstanceID soap.Ui4
	// DesiredBrightness: allowed value range: minimum=0, step=1
	DesiredBrightness soap.Ui2
}

// RenderingControl2SetBrightnessResponse describes the response for RenderingControl2.SetBrightness API
type RenderingControl2SetBrightnessResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetBrightnessRequest
func (client *RenderingControl2) SetBrightness(request RenderingControl2SetBrightnessRequest) (response *RenderingControl2SetBrightnessResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetBrightness", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetContrastRequest describes the request for RenderingControl2.GetContrast API
type RenderingControl2GetContrastRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2GetContrastResponse describes the response for RenderingControl2.GetContrast API
type RenderingControl2GetContrastResponse struct {
	// CurrentContrast: allowed value range: minimum=0, step=1
	CurrentContrast soap.Ui2
}

//
// Return value:
//
//  RenderingControl2GetContrastResponse
func (client *RenderingControl2) GetContrast(request RenderingControl2GetContrastRequest) (response *RenderingControl2GetContrastResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetContrast", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetContrastRequest describes the request for RenderingControl2.SetContrast API
type RenderingControl2SetContrastRequest struct {
	InstanceID soap.Ui4
	// DesiredContrast: allowed value range: minimum=0, step=1
	DesiredContrast soap.Ui2
}

// RenderingControl2SetContrastResponse describes the response for RenderingControl2.SetContrast API
type RenderingControl2SetContrastResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetContrastRequest
func (client *RenderingControl2) SetContrast(request RenderingControl2SetContrastRequest) (response *RenderingControl2SetContrastResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetContrast", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetSharpnessRequest describes the request for RenderingControl2.GetSharpness API
type RenderingControl2GetSharpnessRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2GetSharpnessResponse describes the response for RenderingControl2.GetSharpness API
type RenderingControl2GetSharpnessResponse struct {
	// CurrentSharpness: allowed value range: minimum=0, step=1
	CurrentSharpness soap.Ui2
}

//
// Return value:
//
//  RenderingControl2GetSharpnessResponse
func (client *RenderingControl2) GetSharpness(request RenderingControl2GetSharpnessRequest) (response *RenderingControl2GetSharpnessResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetSharpness", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetSharpnessRequest describes the request for RenderingControl2.SetSharpness API
type RenderingControl2SetSharpnessRequest struct {
	InstanceID soap.Ui4
	// DesiredSharpness: allowed value range: minimum=0, step=1
	DesiredSharpness soap.Ui2
}

// RenderingControl2SetSharpnessResponse describes the response for RenderingControl2.SetSharpness API
type RenderingControl2SetSharpnessResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetSharpnessRequest
func (client *RenderingControl2) SetSharpness(request RenderingControl2SetSharpnessRequest) (response *RenderingControl2SetSharpnessResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetSharpness", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetRedVideoGainRequest describes the request for RenderingControl2.GetRedVideoGain API
type RenderingControl2GetRedVideoGainRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2GetRedVideoGainResponse describes the response for RenderingControl2.GetRedVideoGain API
type RenderingControl2GetRedVideoGainResponse struct {
	// CurrentRedVideoGain: allowed value range: minimum=0, step=1
	CurrentRedVideoGain soap.Ui2
}

//
// Return value:
//
//  RenderingControl2GetRedVideoGainResponse
func (client *RenderingControl2) GetRedVideoGain(request RenderingControl2GetRedVideoGainRequest) (response *RenderingControl2GetRedVideoGainResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetRedVideoGain", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetRedVideoGainRequest describes the request for RenderingControl2.SetRedVideoGain API
type RenderingControl2SetRedVideoGainRequest struct {
	InstanceID soap.Ui4
	// DesiredRedVideoGain: allowed value range: minimum=0, step=1
	DesiredRedVideoGain soap.Ui2
}

// RenderingControl2SetRedVideoGainResponse describes the response for RenderingControl2.SetRedVideoGain API
type RenderingControl2SetRedVideoGainResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetRedVideoGainRequest
func (client *RenderingControl2) SetRedVideoGain(request RenderingControl2SetRedVideoGainRequest) (response *RenderingControl2SetRedVideoGainResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetRedVideoGain", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetGreenVideoGainRequest describes the request for RenderingControl2.GetGreenVideoGain API
type RenderingControl2GetGreenVideoGainRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2GetGreenVideoGainResponse describes the response for RenderingControl2.GetGreenVideoGain API
type RenderingControl2GetGreenVideoGainResponse struct {
	// CurrentGreenVideoGain: allowed value range: minimum=0, step=1
	CurrentGreenVideoGain soap.Ui2
}

//
// Return value:
//
//  RenderingControl2GetGreenVideoGainResponse
func (client *RenderingControl2) GetGreenVideoGain(request RenderingControl2GetGreenVideoGainRequest) (response *RenderingControl2GetGreenVideoGainResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetGreenVideoGain", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetGreenVideoGainRequest describes the request for RenderingControl2.SetGreenVideoGain API
type RenderingControl2SetGreenVideoGainRequest struct {
	InstanceID soap.Ui4
	// DesiredGreenVideoGain: allowed value range: minimum=0, step=1
	DesiredGreenVideoGain soap.Ui2
}

// RenderingControl2SetGreenVideoGainResponse describes the response for RenderingControl2.SetGreenVideoGain API
type RenderingControl2SetGreenVideoGainResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetGreenVideoGainRequest
func (client *RenderingControl2) SetGreenVideoGain(request RenderingControl2SetGreenVideoGainRequest) (response *RenderingControl2SetGreenVideoGainResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetGreenVideoGain", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetBlueVideoGainRequest describes the request for RenderingControl2.GetBlueVideoGain API
type RenderingControl2GetBlueVideoGainRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2GetBlueVideoGainResponse describes the response for RenderingControl2.GetBlueVideoGain API
type RenderingControl2GetBlueVideoGainResponse struct {
	// CurrentBlueVideoGain: allowed value range: minimum=0, step=1
	CurrentBlueVideoGain soap.Ui2
}

//
// Return value:
//
//  RenderingControl2GetBlueVideoGainResponse
func (client *RenderingControl2) GetBlueVideoGain(request RenderingControl2GetBlueVideoGainRequest) (response *RenderingControl2GetBlueVideoGainResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetBlueVideoGain", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetBlueVideoGainRequest describes the request for RenderingControl2.SetBlueVideoGain API
type RenderingControl2SetBlueVideoGainRequest struct {
	InstanceID soap.Ui4
	// DesiredBlueVideoGain: allowed value range: minimum=0, step=1
	DesiredBlueVideoGain soap.Ui2
}

// RenderingControl2SetBlueVideoGainResponse describes the response for RenderingControl2.SetBlueVideoGain API
type RenderingControl2SetBlueVideoGainResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetBlueVideoGainRequest
func (client *RenderingControl2) SetBlueVideoGain(request RenderingControl2SetBlueVideoGainRequest) (response *RenderingControl2SetBlueVideoGainResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetBlueVideoGain", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetRedVideoBlackLevelRequest describes the request for RenderingControl2.GetRedVideoBlackLevel API
type RenderingControl2GetRedVideoBlackLevelRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2GetRedVideoBlackLevelResponse describes the response for RenderingControl2.GetRedVideoBlackLevel API
type RenderingControl2GetRedVideoBlackLevelResponse struct {
	// CurrentRedVideoBlackLevel: allowed value range: minimum=0, step=1
	CurrentRedVideoBlackLevel soap.Ui2
}

//
// Return value:
//
//  RenderingControl2GetRedVideoBlackLevelResponse
func (client *RenderingControl2) GetRedVideoBlackLevel(request RenderingControl2GetRedVideoBlackLevelRequest) (response *RenderingControl2GetRedVideoBlackLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetRedVideoBlackLevel", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetRedVideoBlackLevelRequest describes the request for RenderingControl2.SetRedVideoBlackLevel API
type RenderingControl2SetRedVideoBlackLevelRequest struct {
	InstanceID soap.Ui4
	// DesiredRedVideoBlackLevel: allowed value range: minimum=0, step=1
	DesiredRedVideoBlackLevel soap.Ui2
}

// RenderingControl2SetRedVideoBlackLevelResponse describes the response for RenderingControl2.SetRedVideoBlackLevel API
type RenderingControl2SetRedVideoBlackLevelResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetRedVideoBlackLevelRequest
func (client *RenderingControl2) SetRedVideoBlackLevel(request RenderingControl2SetRedVideoBlackLevelRequest) (response *RenderingControl2SetRedVideoBlackLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetRedVideoBlackLevel", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetGreenVideoBlackLevelRequest describes the request for RenderingControl2.GetGreenVideoBlackLevel API
type RenderingControl2GetGreenVideoBlackLevelRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2GetGreenVideoBlackLevelResponse describes the response for RenderingControl2.GetGreenVideoBlackLevel API
type RenderingControl2GetGreenVideoBlackLevelResponse struct {
	// CurrentGreenVideoBlackLevel: allowed value range: minimum=0, step=1
	CurrentGreenVideoBlackLevel soap.Ui2
}

//
// Return value:
//
//  RenderingControl2GetGreenVideoBlackLevelResponse
func (client *RenderingControl2) GetGreenVideoBlackLevel(request RenderingControl2GetGreenVideoBlackLevelRequest) (response *RenderingControl2GetGreenVideoBlackLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetGreenVideoBlackLevel", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetGreenVideoBlackLevelRequest describes the request for RenderingControl2.SetGreenVideoBlackLevel API
type RenderingControl2SetGreenVideoBlackLevelRequest struct {
	InstanceID soap.Ui4
	// DesiredGreenVideoBlackLevel: allowed value range: minimum=0, step=1
	DesiredGreenVideoBlackLevel soap.Ui2
}

// RenderingControl2SetGreenVideoBlackLevelResponse describes the response for RenderingControl2.SetGreenVideoBlackLevel API
type RenderingControl2SetGreenVideoBlackLevelResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetGreenVideoBlackLevelRequest
func (client *RenderingControl2) SetGreenVideoBlackLevel(request RenderingControl2SetGreenVideoBlackLevelRequest) (response *RenderingControl2SetGreenVideoBlackLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetGreenVideoBlackLevel", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetBlueVideoBlackLevelRequest describes the request for RenderingControl2.GetBlueVideoBlackLevel API
type RenderingControl2GetBlueVideoBlackLevelRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2GetBlueVideoBlackLevelResponse describes the response for RenderingControl2.GetBlueVideoBlackLevel API
type RenderingControl2GetBlueVideoBlackLevelResponse struct {
	// CurrentBlueVideoBlackLevel: allowed value range: minimum=0, step=1
	CurrentBlueVideoBlackLevel soap.Ui2
}

//
// Return value:
//
//  RenderingControl2GetBlueVideoBlackLevelResponse
func (client *RenderingControl2) GetBlueVideoBlackLevel(request RenderingControl2GetBlueVideoBlackLevelRequest) (response *RenderingControl2GetBlueVideoBlackLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetBlueVideoBlackLevel", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetBlueVideoBlackLevelRequest describes the request for RenderingControl2.SetBlueVideoBlackLevel API
type RenderingControl2SetBlueVideoBlackLevelRequest struct {
	InstanceID soap.Ui4
	// DesiredBlueVideoBlackLevel: allowed value range: minimum=0, step=1
	DesiredBlueVideoBlackLevel soap.Ui2
}

// RenderingControl2SetBlueVideoBlackLevelResponse describes the response for RenderingControl2.SetBlueVideoBlackLevel API
type RenderingControl2SetBlueVideoBlackLevelResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetBlueVideoBlackLevelRequest
func (client *RenderingControl2) SetBlueVideoBlackLevel(request RenderingControl2SetBlueVideoBlackLevelRequest) (response *RenderingControl2SetBlueVideoBlackLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetBlueVideoBlackLevel", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetColorTemperatureRequest describes the request for RenderingControl2.GetColorTemperature API
type RenderingControl2GetColorTemperatureRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2GetColorTemperatureResponse describes the response for RenderingControl2.GetColorTemperature API
type RenderingControl2GetColorTemperatureResponse struct {
	// CurrentColorTemperature: allowed value range: minimum=0, step=1
	CurrentColorTemperature soap.Ui2
}

//
// Return value:
//
//  RenderingControl2GetColorTemperatureResponse
func (client *RenderingControl2) GetColorTemperature(request RenderingControl2GetColorTemperatureRequest) (response *RenderingControl2GetColorTemperatureResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetColorTemperature", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetColorTemperatureRequest describes the request for RenderingControl2.SetColorTemperature API
type RenderingControl2SetColorTemperatureRequest struct {
	InstanceID soap.Ui4
	// DesiredColorTemperature: allowed value range: minimum=0, step=1
	DesiredColorTemperature soap.Ui2
}

// RenderingControl2SetColorTemperatureResponse describes the response for RenderingControl2.SetColorTemperature API
type RenderingControl2SetColorTemperatureResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetColorTemperatureRequest
func (client *RenderingControl2) SetColorTemperature(request RenderingControl2SetColorTemperatureRequest) (response *RenderingControl2SetColorTemperatureResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetColorTemperature", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetHorizontalKeystoneRequest describes the request for RenderingControl2.GetHorizontalKeystone API
type RenderingControl2GetHorizontalKeystoneRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2GetHorizontalKeystoneResponse describes the response for RenderingControl2.GetHorizontalKeystone API
type RenderingControl2GetHorizontalKeystoneResponse struct {
	// CurrentHorizontalKeystone: allowed value range: step=1
	CurrentHorizontalKeystone soap.I2
}

//
// Return value:
//
//  RenderingControl2GetHorizontalKeystoneResponse
func (client *RenderingControl2) GetHorizontalKeystone(request RenderingControl2GetHorizontalKeystoneRequest) (response *RenderingControl2GetHorizontalKeystoneResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetHorizontalKeystone", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetHorizontalKeystoneRequest describes the request for RenderingControl2.SetHorizontalKeystone API
type RenderingControl2SetHorizontalKeystoneRequest struct {
	InstanceID soap.Ui4
	// DesiredHorizontalKeystone: allowed value range: step=1
	DesiredHorizontalKeystone soap.I2
}

// RenderingControl2SetHorizontalKeystoneResponse describes the response for RenderingControl2.SetHorizontalKeystone API
type RenderingControl2SetHorizontalKeystoneResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetHorizontalKeystoneRequest
func (client *RenderingControl2) SetHorizontalKeystone(request RenderingControl2SetHorizontalKeystoneRequest) (response *RenderingControl2SetHorizontalKeystoneResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetHorizontalKeystone", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetVerticalKeystoneRequest describes the request for RenderingControl2.GetVerticalKeystone API
type RenderingControl2GetVerticalKeystoneRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl2GetVerticalKeystoneResponse describes the response for RenderingControl2.GetVerticalKeystone API
type RenderingControl2GetVerticalKeystoneResponse struct {
	// CurrentVerticalKeystone: allowed value range: step=1
	CurrentVerticalKeystone soap.I2
}

//
// Return value:
//
//  RenderingControl2GetVerticalKeystoneResponse
func (client *RenderingControl2) GetVerticalKeystone(request RenderingControl2GetVerticalKeystoneRequest) (response *RenderingControl2GetVerticalKeystoneResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetVerticalKeystone", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetVerticalKeystoneRequest describes the request for RenderingControl2.SetVerticalKeystone API
type RenderingControl2SetVerticalKeystoneRequest struct {
	InstanceID soap.Ui4
	// DesiredVerticalKeystone: allowed value range: step=1
	DesiredVerticalKeystone soap.I2
}

// RenderingControl2SetVerticalKeystoneResponse describes the response for RenderingControl2.SetVerticalKeystone API
type RenderingControl2SetVerticalKeystoneResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetVerticalKeystoneRequest
func (client *RenderingControl2) SetVerticalKeystone(request RenderingControl2SetVerticalKeystoneRequest) (response *RenderingControl2SetVerticalKeystoneResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetVerticalKeystone", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetMuteRequest describes the request for RenderingControl2.GetMute API
type RenderingControl2GetMuteRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel soap.String
}

// RenderingControl2GetMuteResponse describes the response for RenderingControl2.GetMute API
type RenderingControl2GetMuteResponse struct {
	CurrentMute soap.Bool
}

//
// Arguments:
//
//  RenderingControl2GetMuteRequest
func (client *RenderingControl2) GetMute(request RenderingControl2GetMuteRequest) (response *RenderingControl2GetMuteResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetMute", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetMuteRequest describes the request for RenderingControl2.SetMute API
type RenderingControl2SetMuteRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel     soap.String
	DesiredMute soap.Bool
}

// RenderingControl2SetMuteResponse describes the response for RenderingControl2.SetMute API
type RenderingControl2SetMuteResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetMuteRequest
func (client *RenderingControl2) SetMute(request RenderingControl2SetMuteRequest) (response *RenderingControl2SetMuteResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetMute", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetVolumeRequest describes the request for RenderingControl2.GetVolume API
type RenderingControl2GetVolumeRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel soap.String
}

// RenderingControl2GetVolumeResponse describes the response for RenderingControl2.GetVolume API
type RenderingControl2GetVolumeResponse struct {
	// CurrentVolume: allowed value range: minimum=0, step=1
	CurrentVolume soap.Ui2
}

//
// Arguments:
//
//  RenderingControl2GetVolumeRequest
//
// Return value:
//
//  RenderingControl2GetVolumeResponse
func (client *RenderingControl2) GetVolume(request RenderingControl2GetVolumeRequest) (response *RenderingControl2GetVolumeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetVolume", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetVolumeRequest describes the request for RenderingControl2.SetVolume API
type RenderingControl2SetVolumeRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel soap.String
	// DesiredVolume: allowed value range: minimum=0, step=1
	DesiredVolume soap.Ui2
}

// RenderingControl2SetVolumeResponse describes the response for RenderingControl2.SetVolume API
type RenderingControl2SetVolumeResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetVolumeRequest
func (client *RenderingControl2) SetVolume(request RenderingControl2SetVolumeRequest) (response *RenderingControl2SetVolumeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetVolume", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetVolumeDBRequest describes the request for RenderingControl2.GetVolumeDB API
type RenderingControl2GetVolumeDBRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel soap.String
}

// RenderingControl2GetVolumeDBResponse describes the response for RenderingControl2.GetVolumeDB API
type RenderingControl2GetVolumeDBResponse struct {
	CurrentVolume soap.I2
}

//
// Arguments:
//
//  RenderingControl2GetVolumeDBRequest
func (client *RenderingControl2) GetVolumeDB(request RenderingControl2GetVolumeDBRequest) (response *RenderingControl2GetVolumeDBResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetVolumeDB", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetVolumeDBRequest describes the request for RenderingControl2.SetVolumeDB API
type RenderingControl2SetVolumeDBRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel       soap.String
	DesiredVolume soap.I2
}

// RenderingControl2SetVolumeDBResponse describes the response for RenderingControl2.SetVolumeDB API
type RenderingControl2SetVolumeDBResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetVolumeDBRequest
func (client *RenderingControl2) SetVolumeDB(request RenderingControl2SetVolumeDBRequest) (response *RenderingControl2SetVolumeDBResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetVolumeDB", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetVolumeDBRangeRequest describes the request for RenderingControl2.GetVolumeDBRange API
type RenderingControl2GetVolumeDBRangeRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel soap.String
}

// RenderingControl2GetVolumeDBRangeResponse describes the response for RenderingControl2.GetVolumeDBRange API
type RenderingControl2GetVolumeDBRangeResponse struct {
	MinValue soap.I2
	MaxValue soap.I2
}

//
// Arguments:
//
//  RenderingControl2GetVolumeDBRangeRequest
func (client *RenderingControl2) GetVolumeDBRange(request RenderingControl2GetVolumeDBRangeRequest) (response *RenderingControl2GetVolumeDBRangeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetVolumeDBRange", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetLoudnessRequest describes the request for RenderingControl2.GetLoudness API
type RenderingControl2GetLoudnessRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel soap.String
}

// RenderingControl2GetLoudnessResponse describes the response for RenderingControl2.GetLoudness API
type RenderingControl2GetLoudnessResponse struct {
	CurrentLoudness soap.Bool
}

//
// Arguments:
//
//  RenderingControl2GetLoudnessRequest
func (client *RenderingControl2) GetLoudness(request RenderingControl2GetLoudnessRequest) (response *RenderingControl2GetLoudnessResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetLoudness", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetLoudnessRequest describes the request for RenderingControl2.SetLoudness API
type RenderingControl2SetLoudnessRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master
	Channel         soap.String
	DesiredLoudness soap.Bool
}

// RenderingControl2SetLoudnessResponse describes the response for RenderingControl2.SetLoudness API
type RenderingControl2SetLoudnessResponse struct {
}

//
// Arguments:
//
//  RenderingControl2SetLoudnessRequest
func (client *RenderingControl2) SetLoudness(request RenderingControl2SetLoudnessRequest) (response *RenderingControl2SetLoudnessResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetLoudness", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2GetStateVariablesRequest describes the request for RenderingControl2.GetStateVariables API
type RenderingControl2GetStateVariablesRequest struct {
	InstanceID        soap.Ui4
	StateVariableList soap.String
}

// RenderingControl2GetStateVariablesResponse describes the response for RenderingControl2.GetStateVariables API
type RenderingControl2GetStateVariablesResponse struct {
	StateVariableValuePairs soap.String
}

func (client *RenderingControl2) GetStateVariables(request RenderingControl2GetStateVariablesRequest) (response *RenderingControl2GetStateVariablesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "GetStateVariables", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// RenderingControl2SetStateVariablesRequest describes the request for RenderingControl2.SetStateVariables API
type RenderingControl2SetStateVariablesRequest struct {
	InstanceID              soap.Ui4
	RenderingControlUDN     soap.String
	ServiceType             soap.String
	ServiceId               soap.String
	StateVariableValuePairs soap.String
}

// RenderingControl2SetStateVariablesResponse describes the response for RenderingControl2.SetStateVariables API
type RenderingControl2SetStateVariablesResponse struct {
	StateVariableList soap.String
}

func (client *RenderingControl2) SetStateVariables(request RenderingControl2SetStateVariablesRequest) (response *RenderingControl2SetStateVariablesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_RenderingControl_2, "SetStateVariables", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ScheduledRecording:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ScheduledRecording1 struct {
	goupnp.ServiceClient
}

// NewScheduledRecording1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewScheduledRecording1Clients() (clients []*ScheduledRecording1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_ScheduledRecording_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newScheduledRecording1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewScheduledRecording1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewScheduledRecording1ClientsByURL(loc *url.URL) ([]*ScheduledRecording1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_ScheduledRecording_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newScheduledRecording1ClientsFromGenericClients(genericClients), nil
}

// NewScheduledRecording1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewScheduledRecording1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ScheduledRecording1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ScheduledRecording_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newScheduledRecording1ClientsFromGenericClients(genericClients), nil
}

func newScheduledRecording1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ScheduledRecording1 {
	clients := make([]*ScheduledRecording1, len(genericClients))
	for i := range genericClients {
		clients[i] = &ScheduledRecording1{genericClients[i]}
	}
	return clients
}

// ScheduledRecording1GetSortCapabilitiesRequest describes the request for ScheduledRecording1.GetSortCapabilities API
type ScheduledRecording1GetSortCapabilitiesRequest struct {
}

// ScheduledRecording1GetSortCapabilitiesResponse describes the response for ScheduledRecording1.GetSortCapabilities API
type ScheduledRecording1GetSortCapabilitiesResponse struct {
	SortCaps     soap.String
	SortLevelCap soap.Ui4
}

func (client *ScheduledRecording1) GetSortCapabilities(request ScheduledRecording1GetSortCapabilitiesRequest) (response *ScheduledRecording1GetSortCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "GetSortCapabilities", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1GetPropertyListRequest describes the request for ScheduledRecording1.GetPropertyList API
type ScheduledRecording1GetPropertyListRequest struct {
	// DataTypeID: allowed values: A_ARG_TYPE_RecordSchedule, A_ARG_TYPE_RecordTask, A_ARG_TYPE_RecordScheduleParts
	DataTypeID soap.String
}

// ScheduledRecording1GetPropertyListResponse describes the response for ScheduledRecording1.GetPropertyList API
type ScheduledRecording1GetPropertyListResponse struct {
	PropertyList soap.String
}

//
// Arguments:
//
//  ScheduledRecording1GetPropertyListRequest
func (client *ScheduledRecording1) GetPropertyList(request ScheduledRecording1GetPropertyListRequest) (response *ScheduledRecording1GetPropertyListResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "GetPropertyList", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1GetAllowedValuesRequest describes the request for ScheduledRecording1.GetAllowedValues API
type ScheduledRecording1GetAllowedValuesRequest struct {
	// DataTypeID: allowed values: A_ARG_TYPE_RecordSchedule, A_ARG_TYPE_RecordTask, A_ARG_TYPE_RecordScheduleParts
	DataTypeID soap.String
	Filter     soap.String
}

// ScheduledRecording1GetAllowedValuesResponse describes the response for ScheduledRecording1.GetAllowedValues API
type ScheduledRecording1GetAllowedValuesResponse struct {
	PropertyInfo soap.String
}

//
// Arguments:
//
//  ScheduledRecording1GetAllowedValuesRequest
func (client *ScheduledRecording1) GetAllowedValues(request ScheduledRecording1GetAllowedValuesRequest) (response *ScheduledRecording1GetAllowedValuesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "GetAllowedValues", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1GetStateUpdateIDRequest describes the request for ScheduledRecording1.GetStateUpdateID API
type ScheduledRecording1GetStateUpdateIDRequest struct {
}

// ScheduledRecording1GetStateUpdateIDResponse describes the response for ScheduledRecording1.GetStateUpdateID API
type ScheduledRecording1GetStateUpdateIDResponse struct {
	Id soap.Ui4
}

func (client *ScheduledRecording1) GetStateUpdateID(request ScheduledRecording1GetStateUpdateIDRequest) (response *ScheduledRecording1GetStateUpdateIDResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "GetStateUpdateID", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1BrowseRecordSchedulesRequest describes the request for ScheduledRecording1.BrowseRecordSchedules API
type ScheduledRecording1BrowseRecordSchedulesRequest struct {
	Filter         soap.String
	StartingIndex  soap.Ui4
	RequestedCount soap.Ui4
	SortCriteria   soap.String
}

// ScheduledRecording1BrowseRecordSchedulesResponse describes the response for ScheduledRecording1.BrowseRecordSchedules API
type ScheduledRecording1BrowseRecordSchedulesResponse struct {
	Result         soap.String
	NumberReturned soap.Ui4
	TotalMatches   soap.Ui4
	UpdateID       soap.Ui4
}

func (client *ScheduledRecording1) BrowseRecordSchedules(request ScheduledRecording1BrowseRecordSchedulesRequest) (response *ScheduledRecording1BrowseRecordSchedulesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "BrowseRecordSchedules", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1BrowseRecordTasksRequest describes the request for ScheduledRecording1.BrowseRecordTasks API
type ScheduledRecording1BrowseRecordTasksRequest struct {
	RecordScheduleID soap.String
	Filter           soap.String
	StartingIndex    soap.Ui4
	RequestedCount   soap.Ui4
	SortCriteria     soap.String
}

// ScheduledRecording1BrowseRecordTasksResponse describes the response for ScheduledRecording1.BrowseRecordTasks API
type ScheduledRecording1BrowseRecordTasksResponse struct {
	Result         soap.String
	NumberReturned soap.Ui4
	TotalMatches   soap.Ui4
	UpdateID       soap.Ui4
}

func (client *ScheduledRecording1) BrowseRecordTasks(request ScheduledRecording1BrowseRecordTasksRequest) (response *ScheduledRecording1BrowseRecordTasksResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "BrowseRecordTasks", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1CreateRecordScheduleRequest describes the request for ScheduledRecording1.CreateRecordSchedule API
type ScheduledRecording1CreateRecordScheduleRequest struct {
	Elements soap.String
}

// ScheduledRecording1CreateRecordScheduleResponse describes the response for ScheduledRecording1.CreateRecordSchedule API
type ScheduledRecording1CreateRecordScheduleResponse struct {
	RecordScheduleID soap.String
	Result           soap.String
	UpdateID         soap.Ui4
}

func (client *ScheduledRecording1) CreateRecordSchedule(request ScheduledRecording1CreateRecordScheduleRequest) (response *ScheduledRecording1CreateRecordScheduleResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "CreateRecordSchedule", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1DeleteRecordScheduleRequest describes the request for ScheduledRecording1.DeleteRecordSchedule API
type ScheduledRecording1DeleteRecordScheduleRequest struct {
	RecordScheduleID soap.String
}

// ScheduledRecording1DeleteRecordScheduleResponse describes the response for ScheduledRecording1.DeleteRecordSchedule API
type ScheduledRecording1DeleteRecordScheduleResponse struct {
}

func (client *ScheduledRecording1) DeleteRecordSchedule(request ScheduledRecording1DeleteRecordScheduleRequest) (response *ScheduledRecording1DeleteRecordScheduleResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "DeleteRecordSchedule", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1GetRecordScheduleRequest describes the request for ScheduledRecording1.GetRecordSchedule API
type ScheduledRecording1GetRecordScheduleRequest struct {
	RecordScheduleID soap.String
	Filter           soap.String
}

// ScheduledRecording1GetRecordScheduleResponse describes the response for ScheduledRecording1.GetRecordSchedule API
type ScheduledRecording1GetRecordScheduleResponse struct {
	Result   soap.String
	UpdateID soap.Ui4
}

func (client *ScheduledRecording1) GetRecordSchedule(request ScheduledRecording1GetRecordScheduleRequest) (response *ScheduledRecording1GetRecordScheduleResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "GetRecordSchedule", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1EnableRecordScheduleRequest describes the request for ScheduledRecording1.EnableRecordSchedule API
type ScheduledRecording1EnableRecordScheduleRequest struct {
	RecordScheduleID soap.String
}

// ScheduledRecording1EnableRecordScheduleResponse describes the response for ScheduledRecording1.EnableRecordSchedule API
type ScheduledRecording1EnableRecordScheduleResponse struct {
}

func (client *ScheduledRecording1) EnableRecordSchedule(request ScheduledRecording1EnableRecordScheduleRequest) (response *ScheduledRecording1EnableRecordScheduleResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "EnableRecordSchedule", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1DisableRecordScheduleRequest describes the request for ScheduledRecording1.DisableRecordSchedule API
type ScheduledRecording1DisableRecordScheduleRequest struct {
	RecordScheduleID soap.String
}

// ScheduledRecording1DisableRecordScheduleResponse describes the response for ScheduledRecording1.DisableRecordSchedule API
type ScheduledRecording1DisableRecordScheduleResponse struct {
}

func (client *ScheduledRecording1) DisableRecordSchedule(request ScheduledRecording1DisableRecordScheduleRequest) (response *ScheduledRecording1DisableRecordScheduleResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "DisableRecordSchedule", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1DeleteRecordTaskRequest describes the request for ScheduledRecording1.DeleteRecordTask API
type ScheduledRecording1DeleteRecordTaskRequest struct {
	RecordTaskID soap.String
}

// ScheduledRecording1DeleteRecordTaskResponse describes the response for ScheduledRecording1.DeleteRecordTask API
type ScheduledRecording1DeleteRecordTaskResponse struct {
}

func (client *ScheduledRecording1) DeleteRecordTask(request ScheduledRecording1DeleteRecordTaskRequest) (response *ScheduledRecording1DeleteRecordTaskResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "DeleteRecordTask", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1GetRecordTaskRequest describes the request for ScheduledRecording1.GetRecordTask API
type ScheduledRecording1GetRecordTaskRequest struct {
	RecordTaskID soap.String
	Filter       soap.String
}

// ScheduledRecording1GetRecordTaskResponse describes the response for ScheduledRecording1.GetRecordTask API
type ScheduledRecording1GetRecordTaskResponse struct {
	Result   soap.String
	UpdateID soap.Ui4
}

func (client *ScheduledRecording1) GetRecordTask(request ScheduledRecording1GetRecordTaskRequest) (response *ScheduledRecording1GetRecordTaskResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "GetRecordTask", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1EnableRecordTaskRequest describes the request for ScheduledRecording1.EnableRecordTask API
type ScheduledRecording1EnableRecordTaskRequest struct {
	RecordTaskID soap.String
}

// ScheduledRecording1EnableRecordTaskResponse describes the response for ScheduledRecording1.EnableRecordTask API
type ScheduledRecording1EnableRecordTaskResponse struct {
}

func (client *ScheduledRecording1) EnableRecordTask(request ScheduledRecording1EnableRecordTaskRequest) (response *ScheduledRecording1EnableRecordTaskResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "EnableRecordTask", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1DisableRecordTaskRequest describes the request for ScheduledRecording1.DisableRecordTask API
type ScheduledRecording1DisableRecordTaskRequest struct {
	RecordTaskID soap.String
}

// ScheduledRecording1DisableRecordTaskResponse describes the response for ScheduledRecording1.DisableRecordTask API
type ScheduledRecording1DisableRecordTaskResponse struct {
}

func (client *ScheduledRecording1) DisableRecordTask(request ScheduledRecording1DisableRecordTaskRequest) (response *ScheduledRecording1DisableRecordTaskResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "DisableRecordTask", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1ResetRecordTaskRequest describes the request for ScheduledRecording1.ResetRecordTask API
type ScheduledRecording1ResetRecordTaskRequest struct {
	RecordTaskID soap.String
}

// ScheduledRecording1ResetRecordTaskResponse describes the response for ScheduledRecording1.ResetRecordTask API
type ScheduledRecording1ResetRecordTaskResponse struct {
}

func (client *ScheduledRecording1) ResetRecordTask(request ScheduledRecording1ResetRecordTaskRequest) (response *ScheduledRecording1ResetRecordTaskResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "ResetRecordTask", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1GetRecordScheduleConflictsRequest describes the request for ScheduledRecording1.GetRecordScheduleConflicts API
type ScheduledRecording1GetRecordScheduleConflictsRequest struct {
	RecordScheduleID soap.String
}

// ScheduledRecording1GetRecordScheduleConflictsResponse describes the response for ScheduledRecording1.GetRecordScheduleConflicts API
type ScheduledRecording1GetRecordScheduleConflictsResponse struct {
	RecordScheduleConflictIDList soap.String
	UpdateID                     soap.Ui4
}

func (client *ScheduledRecording1) GetRecordScheduleConflicts(request ScheduledRecording1GetRecordScheduleConflictsRequest) (response *ScheduledRecording1GetRecordScheduleConflictsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "GetRecordScheduleConflicts", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording1GetRecordTaskConflictsRequest describes the request for ScheduledRecording1.GetRecordTaskConflicts API
type ScheduledRecording1GetRecordTaskConflictsRequest struct {
	RecordTaskID soap.String
}

// ScheduledRecording1GetRecordTaskConflictsResponse describes the response for ScheduledRecording1.GetRecordTaskConflicts API
type ScheduledRecording1GetRecordTaskConflictsResponse struct {
	RecordTaskConflictIDList soap.String
	UpdateID                 soap.Ui4
}

func (client *ScheduledRecording1) GetRecordTaskConflicts(request ScheduledRecording1GetRecordTaskConflictsRequest) (response *ScheduledRecording1GetRecordTaskConflictsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_1, "GetRecordTaskConflicts", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ScheduledRecording:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ScheduledRecording2 struct {
	goupnp.ServiceClient
}

// NewScheduledRecording2Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewScheduledRecording2Clients() (clients []*ScheduledRecording2, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_ScheduledRecording_2); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newScheduledRecording2ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewScheduledRecording2ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewScheduledRecording2ClientsByURL(loc *url.URL) ([]*ScheduledRecording2, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_ScheduledRecording_2)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newScheduledRecording2ClientsFromGenericClients(genericClients), nil
}

// NewScheduledRecording2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewScheduledRecording2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ScheduledRecording2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ScheduledRecording_2)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newScheduledRecording2ClientsFromGenericClients(genericClients), nil
}

func newScheduledRecording2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ScheduledRecording2 {
	clients := make([]*ScheduledRecording2, len(genericClients))
	for i := range genericClients {
		clients[i] = &ScheduledRecording2{genericClients[i]}
	}
	return clients
}

// ScheduledRecording2GetSortCapabilitiesRequest describes the request for ScheduledRecording2.GetSortCapabilities API
type ScheduledRecording2GetSortCapabilitiesRequest struct {
}

// ScheduledRecording2GetSortCapabilitiesResponse describes the response for ScheduledRecording2.GetSortCapabilities API
type ScheduledRecording2GetSortCapabilitiesResponse struct {
	SortCaps     soap.String
	SortLevelCap soap.Ui4
}

func (client *ScheduledRecording2) GetSortCapabilities(request ScheduledRecording2GetSortCapabilitiesRequest) (response *ScheduledRecording2GetSortCapabilitiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "GetSortCapabilities", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2GetPropertyListRequest describes the request for ScheduledRecording2.GetPropertyList API
type ScheduledRecording2GetPropertyListRequest struct {
	// DataTypeID: allowed values: A_ARG_TYPE_RecordSchedule, A_ARG_TYPE_RecordTask, A_ARG_TYPE_RecordScheduleParts
	DataTypeID soap.String
}

// ScheduledRecording2GetPropertyListResponse describes the response for ScheduledRecording2.GetPropertyList API
type ScheduledRecording2GetPropertyListResponse struct {
	PropertyList soap.String
}

//
// Arguments:
//
//  ScheduledRecording2GetPropertyListRequest
func (client *ScheduledRecording2) GetPropertyList(request ScheduledRecording2GetPropertyListRequest) (response *ScheduledRecording2GetPropertyListResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "GetPropertyList", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2GetAllowedValuesRequest describes the request for ScheduledRecording2.GetAllowedValues API
type ScheduledRecording2GetAllowedValuesRequest struct {
	// DataTypeID: allowed values: A_ARG_TYPE_RecordSchedule, A_ARG_TYPE_RecordTask, A_ARG_TYPE_RecordScheduleParts
	DataTypeID soap.String
	Filter     soap.String
}

// ScheduledRecording2GetAllowedValuesResponse describes the response for ScheduledRecording2.GetAllowedValues API
type ScheduledRecording2GetAllowedValuesResponse struct {
	PropertyInfo soap.String
}

//
// Arguments:
//
//  ScheduledRecording2GetAllowedValuesRequest
func (client *ScheduledRecording2) GetAllowedValues(request ScheduledRecording2GetAllowedValuesRequest) (response *ScheduledRecording2GetAllowedValuesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "GetAllowedValues", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2GetStateUpdateIDRequest describes the request for ScheduledRecording2.GetStateUpdateID API
type ScheduledRecording2GetStateUpdateIDRequest struct {
}

// ScheduledRecording2GetStateUpdateIDResponse describes the response for ScheduledRecording2.GetStateUpdateID API
type ScheduledRecording2GetStateUpdateIDResponse struct {
	Id soap.Ui4
}

func (client *ScheduledRecording2) GetStateUpdateID(request ScheduledRecording2GetStateUpdateIDRequest) (response *ScheduledRecording2GetStateUpdateIDResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "GetStateUpdateID", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2BrowseRecordSchedulesRequest describes the request for ScheduledRecording2.BrowseRecordSchedules API
type ScheduledRecording2BrowseRecordSchedulesRequest struct {
	Filter         soap.String
	StartingIndex  soap.Ui4
	RequestedCount soap.Ui4
	SortCriteria   soap.String
}

// ScheduledRecording2BrowseRecordSchedulesResponse describes the response for ScheduledRecording2.BrowseRecordSchedules API
type ScheduledRecording2BrowseRecordSchedulesResponse struct {
	Result         soap.String
	NumberReturned soap.Ui4
	TotalMatches   soap.Ui4
	UpdateID       soap.Ui4
}

func (client *ScheduledRecording2) BrowseRecordSchedules(request ScheduledRecording2BrowseRecordSchedulesRequest) (response *ScheduledRecording2BrowseRecordSchedulesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "BrowseRecordSchedules", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2BrowseRecordTasksRequest describes the request for ScheduledRecording2.BrowseRecordTasks API
type ScheduledRecording2BrowseRecordTasksRequest struct {
	RecordScheduleID soap.String
	Filter           soap.String
	StartingIndex    soap.Ui4
	RequestedCount   soap.Ui4
	SortCriteria     soap.String
}

// ScheduledRecording2BrowseRecordTasksResponse describes the response for ScheduledRecording2.BrowseRecordTasks API
type ScheduledRecording2BrowseRecordTasksResponse struct {
	Result         soap.String
	NumberReturned soap.Ui4
	TotalMatches   soap.Ui4
	UpdateID       soap.Ui4
}

func (client *ScheduledRecording2) BrowseRecordTasks(request ScheduledRecording2BrowseRecordTasksRequest) (response *ScheduledRecording2BrowseRecordTasksResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "BrowseRecordTasks", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2CreateRecordScheduleRequest describes the request for ScheduledRecording2.CreateRecordSchedule API
type ScheduledRecording2CreateRecordScheduleRequest struct {
	Elements soap.String
}

// ScheduledRecording2CreateRecordScheduleResponse describes the response for ScheduledRecording2.CreateRecordSchedule API
type ScheduledRecording2CreateRecordScheduleResponse struct {
	RecordScheduleID soap.String
	Result           soap.String
	UpdateID         soap.Ui4
}

func (client *ScheduledRecording2) CreateRecordSchedule(request ScheduledRecording2CreateRecordScheduleRequest) (response *ScheduledRecording2CreateRecordScheduleResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "CreateRecordSchedule", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2DeleteRecordScheduleRequest describes the request for ScheduledRecording2.DeleteRecordSchedule API
type ScheduledRecording2DeleteRecordScheduleRequest struct {
	RecordScheduleID soap.String
}

// ScheduledRecording2DeleteRecordScheduleResponse describes the response for ScheduledRecording2.DeleteRecordSchedule API
type ScheduledRecording2DeleteRecordScheduleResponse struct {
}

func (client *ScheduledRecording2) DeleteRecordSchedule(request ScheduledRecording2DeleteRecordScheduleRequest) (response *ScheduledRecording2DeleteRecordScheduleResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "DeleteRecordSchedule", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2GetRecordScheduleRequest describes the request for ScheduledRecording2.GetRecordSchedule API
type ScheduledRecording2GetRecordScheduleRequest struct {
	RecordScheduleID soap.String
	Filter           soap.String
}

// ScheduledRecording2GetRecordScheduleResponse describes the response for ScheduledRecording2.GetRecordSchedule API
type ScheduledRecording2GetRecordScheduleResponse struct {
	Result   soap.String
	UpdateID soap.Ui4
}

func (client *ScheduledRecording2) GetRecordSchedule(request ScheduledRecording2GetRecordScheduleRequest) (response *ScheduledRecording2GetRecordScheduleResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "GetRecordSchedule", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2EnableRecordScheduleRequest describes the request for ScheduledRecording2.EnableRecordSchedule API
type ScheduledRecording2EnableRecordScheduleRequest struct {
	RecordScheduleID soap.String
}

// ScheduledRecording2EnableRecordScheduleResponse describes the response for ScheduledRecording2.EnableRecordSchedule API
type ScheduledRecording2EnableRecordScheduleResponse struct {
}

func (client *ScheduledRecording2) EnableRecordSchedule(request ScheduledRecording2EnableRecordScheduleRequest) (response *ScheduledRecording2EnableRecordScheduleResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "EnableRecordSchedule", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2DisableRecordScheduleRequest describes the request for ScheduledRecording2.DisableRecordSchedule API
type ScheduledRecording2DisableRecordScheduleRequest struct {
	RecordScheduleID soap.String
}

// ScheduledRecording2DisableRecordScheduleResponse describes the response for ScheduledRecording2.DisableRecordSchedule API
type ScheduledRecording2DisableRecordScheduleResponse struct {
}

func (client *ScheduledRecording2) DisableRecordSchedule(request ScheduledRecording2DisableRecordScheduleRequest) (response *ScheduledRecording2DisableRecordScheduleResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "DisableRecordSchedule", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2DeleteRecordTaskRequest describes the request for ScheduledRecording2.DeleteRecordTask API
type ScheduledRecording2DeleteRecordTaskRequest struct {
	RecordTaskID soap.String
}

// ScheduledRecording2DeleteRecordTaskResponse describes the response for ScheduledRecording2.DeleteRecordTask API
type ScheduledRecording2DeleteRecordTaskResponse struct {
}

func (client *ScheduledRecording2) DeleteRecordTask(request ScheduledRecording2DeleteRecordTaskRequest) (response *ScheduledRecording2DeleteRecordTaskResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "DeleteRecordTask", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2GetRecordTaskRequest describes the request for ScheduledRecording2.GetRecordTask API
type ScheduledRecording2GetRecordTaskRequest struct {
	RecordTaskID soap.String
	Filter       soap.String
}

// ScheduledRecording2GetRecordTaskResponse describes the response for ScheduledRecording2.GetRecordTask API
type ScheduledRecording2GetRecordTaskResponse struct {
	Result   soap.String
	UpdateID soap.Ui4
}

func (client *ScheduledRecording2) GetRecordTask(request ScheduledRecording2GetRecordTaskRequest) (response *ScheduledRecording2GetRecordTaskResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "GetRecordTask", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2EnableRecordTaskRequest describes the request for ScheduledRecording2.EnableRecordTask API
type ScheduledRecording2EnableRecordTaskRequest struct {
	RecordTaskID soap.String
}

// ScheduledRecording2EnableRecordTaskResponse describes the response for ScheduledRecording2.EnableRecordTask API
type ScheduledRecording2EnableRecordTaskResponse struct {
}

func (client *ScheduledRecording2) EnableRecordTask(request ScheduledRecording2EnableRecordTaskRequest) (response *ScheduledRecording2EnableRecordTaskResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "EnableRecordTask", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2DisableRecordTaskRequest describes the request for ScheduledRecording2.DisableRecordTask API
type ScheduledRecording2DisableRecordTaskRequest struct {
	RecordTaskID soap.String
}

// ScheduledRecording2DisableRecordTaskResponse describes the response for ScheduledRecording2.DisableRecordTask API
type ScheduledRecording2DisableRecordTaskResponse struct {
}

func (client *ScheduledRecording2) DisableRecordTask(request ScheduledRecording2DisableRecordTaskRequest) (response *ScheduledRecording2DisableRecordTaskResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "DisableRecordTask", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2ResetRecordTaskRequest describes the request for ScheduledRecording2.ResetRecordTask API
type ScheduledRecording2ResetRecordTaskRequest struct {
	RecordTaskID soap.String
}

// ScheduledRecording2ResetRecordTaskResponse describes the response for ScheduledRecording2.ResetRecordTask API
type ScheduledRecording2ResetRecordTaskResponse struct {
}

func (client *ScheduledRecording2) ResetRecordTask(request ScheduledRecording2ResetRecordTaskRequest) (response *ScheduledRecording2ResetRecordTaskResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "ResetRecordTask", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2GetRecordScheduleConflictsRequest describes the request for ScheduledRecording2.GetRecordScheduleConflicts API
type ScheduledRecording2GetRecordScheduleConflictsRequest struct {
	RecordScheduleID soap.String
}

// ScheduledRecording2GetRecordScheduleConflictsResponse describes the response for ScheduledRecording2.GetRecordScheduleConflicts API
type ScheduledRecording2GetRecordScheduleConflictsResponse struct {
	RecordScheduleConflictIDList soap.String
	UpdateID                     soap.Ui4
}

func (client *ScheduledRecording2) GetRecordScheduleConflicts(request ScheduledRecording2GetRecordScheduleConflictsRequest) (response *ScheduledRecording2GetRecordScheduleConflictsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "GetRecordScheduleConflicts", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// ScheduledRecording2GetRecordTaskConflictsRequest describes the request for ScheduledRecording2.GetRecordTaskConflicts API
type ScheduledRecording2GetRecordTaskConflictsRequest struct {
	RecordTaskID soap.String
}

// ScheduledRecording2GetRecordTaskConflictsResponse describes the response for ScheduledRecording2.GetRecordTaskConflicts API
type ScheduledRecording2GetRecordTaskConflictsResponse struct {
	RecordTaskConflictIDList soap.String
	UpdateID                 soap.Ui4
}

func (client *ScheduledRecording2) GetRecordTaskConflicts(request ScheduledRecording2GetRecordTaskConflictsRequest) (response *ScheduledRecording2GetRecordTaskConflictsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_ScheduledRecording_2, "GetRecordTaskConflicts", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}
