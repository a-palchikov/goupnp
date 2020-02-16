// Client for UPnP Device Control Protocol Internet Gateway Device v1.
//
// This DCP is documented in detail at: http://upnp.org/specs/gw/UPnP-gw-InternetGatewayDevice-v1-Device.pdf
//
// Typically, use one of the New* functions to create clients for services.
package internetgateway1

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
const (
	URN_LANDevice_1           = "urn:schemas-upnp-org:device:LANDevice:1"
	URN_WANConnectionDevice_1 = "urn:schemas-upnp-org:device:WANConnectionDevice:1"
	URN_WANDevice_1           = "urn:schemas-upnp-org:device:WANDevice:1"
)

// Service URNs:
const (
	URN_LANHostConfigManagement_1  = "urn:schemas-upnp-org:service:LANHostConfigManagement:1"
	URN_Layer3Forwarding_1         = "urn:schemas-upnp-org:service:Layer3Forwarding:1"
	URN_WANCableLinkConfig_1       = "urn:schemas-upnp-org:service:WANCableLinkConfig:1"
	URN_WANCommonInterfaceConfig_1 = "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1"
	URN_WANDSLLinkConfig_1         = "urn:schemas-upnp-org:service:WANDSLLinkConfig:1"
	URN_WANEthernetLinkConfig_1    = "urn:schemas-upnp-org:service:WANEthernetLinkConfig:1"
	URN_WANIPConnection_1          = "urn:schemas-upnp-org:service:WANIPConnection:1"
	URN_WANPOTSLinkConfig_1        = "urn:schemas-upnp-org:service:WANPOTSLinkConfig:1"
	URN_WANPPPConnection_1         = "urn:schemas-upnp-org:service:WANPPPConnection:1"
)

// LANHostConfigManagement1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:LANHostConfigManagement:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type LANHostConfigManagement1 struct {
	goupnp.ServiceClient
}

// NewLANHostConfigManagement1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewLANHostConfigManagement1Clients() (clients []*LANHostConfigManagement1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_LANHostConfigManagement_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newLANHostConfigManagement1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewLANHostConfigManagement1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewLANHostConfigManagement1ClientsByURL(loc *url.URL) ([]*LANHostConfigManagement1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_LANHostConfigManagement_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newLANHostConfigManagement1ClientsFromGenericClients(genericClients), nil
}

// NewLANHostConfigManagement1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewLANHostConfigManagement1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*LANHostConfigManagement1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_LANHostConfigManagement_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newLANHostConfigManagement1ClientsFromGenericClients(genericClients), nil
}

func newLANHostConfigManagement1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*LANHostConfigManagement1 {
	clients := make([]*LANHostConfigManagement1, len(genericClients))
	for i := range genericClients {
		clients[i] = &LANHostConfigManagement1{genericClients[i]}
	}
	return clients
}

// LANHostConfigManagement1SetDHCPServerConfigurableRequest describes the request for LANHostConfigManagement1.SetDHCPServerConfigurable API
type LANHostConfigManagement1SetDHCPServerConfigurableRequest struct {
	NewDHCPServerConfigurable soap.Bool
}

// LANHostConfigManagement1SetDHCPServerConfigurableResponse describes the response for LANHostConfigManagement1.SetDHCPServerConfigurable API
type LANHostConfigManagement1SetDHCPServerConfigurableResponse struct {
}

func (client *LANHostConfigManagement1) SetDHCPServerConfigurable(request LANHostConfigManagement1SetDHCPServerConfigurableRequest) (*LANHostConfigManagement1SetDHCPServerConfigurableResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1SetDHCPServerConfigurableResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDHCPServerConfigurable", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1GetDHCPServerConfigurableResponse describes the response for LANHostConfigManagement1.GetDHCPServerConfigurable API
type LANHostConfigManagement1GetDHCPServerConfigurableResponse struct {
	NewDHCPServerConfigurable soap.Bool
}

func (client *LANHostConfigManagement1) GetDHCPServerConfigurable() (*LANHostConfigManagement1GetDHCPServerConfigurableResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1GetDHCPServerConfigurableResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDHCPServerConfigurable", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1SetDHCPRelayRequest describes the request for LANHostConfigManagement1.SetDHCPRelay API
type LANHostConfigManagement1SetDHCPRelayRequest struct {
	NewDHCPRelay soap.Bool
}

// LANHostConfigManagement1SetDHCPRelayResponse describes the response for LANHostConfigManagement1.SetDHCPRelay API
type LANHostConfigManagement1SetDHCPRelayResponse struct {
}

func (client *LANHostConfigManagement1) SetDHCPRelay(request LANHostConfigManagement1SetDHCPRelayRequest) (*LANHostConfigManagement1SetDHCPRelayResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1SetDHCPRelayResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDHCPRelay", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1GetDHCPRelayResponse describes the response for LANHostConfigManagement1.GetDHCPRelay API
type LANHostConfigManagement1GetDHCPRelayResponse struct {
	NewDHCPRelay soap.Bool
}

func (client *LANHostConfigManagement1) GetDHCPRelay() (*LANHostConfigManagement1GetDHCPRelayResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1GetDHCPRelayResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDHCPRelay", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1SetSubnetMaskRequest describes the request for LANHostConfigManagement1.SetSubnetMask API
type LANHostConfigManagement1SetSubnetMaskRequest struct {
	NewSubnetMask soap.String
}

// LANHostConfigManagement1SetSubnetMaskResponse describes the response for LANHostConfigManagement1.SetSubnetMask API
type LANHostConfigManagement1SetSubnetMaskResponse struct {
}

func (client *LANHostConfigManagement1) SetSubnetMask(request LANHostConfigManagement1SetSubnetMaskRequest) (*LANHostConfigManagement1SetSubnetMaskResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1SetSubnetMaskResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetSubnetMask", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1GetSubnetMaskResponse describes the response for LANHostConfigManagement1.GetSubnetMask API
type LANHostConfigManagement1GetSubnetMaskResponse struct {
	NewSubnetMask soap.String
}

func (client *LANHostConfigManagement1) GetSubnetMask() (*LANHostConfigManagement1GetSubnetMaskResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1GetSubnetMaskResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetSubnetMask", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1SetIPRouterRequest describes the request for LANHostConfigManagement1.SetIPRouter API
type LANHostConfigManagement1SetIPRouterRequest struct {
	NewIPRouters soap.String
}

// LANHostConfigManagement1SetIPRouterResponse describes the response for LANHostConfigManagement1.SetIPRouter API
type LANHostConfigManagement1SetIPRouterResponse struct {
}

func (client *LANHostConfigManagement1) SetIPRouter(request LANHostConfigManagement1SetIPRouterRequest) (*LANHostConfigManagement1SetIPRouterResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1SetIPRouterResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetIPRouter", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1DeleteIPRouterRequest describes the request for LANHostConfigManagement1.DeleteIPRouter API
type LANHostConfigManagement1DeleteIPRouterRequest struct {
	NewIPRouters soap.String
}

// LANHostConfigManagement1DeleteIPRouterResponse describes the response for LANHostConfigManagement1.DeleteIPRouter API
type LANHostConfigManagement1DeleteIPRouterResponse struct {
}

func (client *LANHostConfigManagement1) DeleteIPRouter(request LANHostConfigManagement1DeleteIPRouterRequest) (*LANHostConfigManagement1DeleteIPRouterResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1DeleteIPRouterResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteIPRouter", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1GetIPRoutersListResponse describes the response for LANHostConfigManagement1.GetIPRoutersList API
type LANHostConfigManagement1GetIPRoutersListResponse struct {
	NewIPRouters soap.String
}

func (client *LANHostConfigManagement1) GetIPRoutersList() (*LANHostConfigManagement1GetIPRoutersListResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1GetIPRoutersListResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetIPRoutersList", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1SetDomainNameRequest describes the request for LANHostConfigManagement1.SetDomainName API
type LANHostConfigManagement1SetDomainNameRequest struct {
	NewDomainName soap.String
}

// LANHostConfigManagement1SetDomainNameResponse describes the response for LANHostConfigManagement1.SetDomainName API
type LANHostConfigManagement1SetDomainNameResponse struct {
}

func (client *LANHostConfigManagement1) SetDomainName(request LANHostConfigManagement1SetDomainNameRequest) (*LANHostConfigManagement1SetDomainNameResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1SetDomainNameResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDomainName", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1GetDomainNameResponse describes the response for LANHostConfigManagement1.GetDomainName API
type LANHostConfigManagement1GetDomainNameResponse struct {
	NewDomainName soap.String
}

func (client *LANHostConfigManagement1) GetDomainName() (*LANHostConfigManagement1GetDomainNameResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1GetDomainNameResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDomainName", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1SetAddressRangeRequest describes the request for LANHostConfigManagement1.SetAddressRange API
type LANHostConfigManagement1SetAddressRangeRequest struct {
	NewMinAddress soap.String
	NewMaxAddress soap.String
}

// LANHostConfigManagement1SetAddressRangeResponse describes the response for LANHostConfigManagement1.SetAddressRange API
type LANHostConfigManagement1SetAddressRangeResponse struct {
}

func (client *LANHostConfigManagement1) SetAddressRange(request LANHostConfigManagement1SetAddressRangeRequest) (*LANHostConfigManagement1SetAddressRangeResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1SetAddressRangeResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetAddressRange", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1GetAddressRangeResponse describes the response for LANHostConfigManagement1.GetAddressRange API
type LANHostConfigManagement1GetAddressRangeResponse struct {
	NewMinAddress soap.String
	NewMaxAddress soap.String
}

func (client *LANHostConfigManagement1) GetAddressRange() (*LANHostConfigManagement1GetAddressRangeResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1GetAddressRangeResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetAddressRange", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1SetReservedAddressRequest describes the request for LANHostConfigManagement1.SetReservedAddress API
type LANHostConfigManagement1SetReservedAddressRequest struct {
	NewReservedAddresses soap.String
}

// LANHostConfigManagement1SetReservedAddressResponse describes the response for LANHostConfigManagement1.SetReservedAddress API
type LANHostConfigManagement1SetReservedAddressResponse struct {
}

func (client *LANHostConfigManagement1) SetReservedAddress(request LANHostConfigManagement1SetReservedAddressRequest) (*LANHostConfigManagement1SetReservedAddressResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1SetReservedAddressResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetReservedAddress", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1DeleteReservedAddressRequest describes the request for LANHostConfigManagement1.DeleteReservedAddress API
type LANHostConfigManagement1DeleteReservedAddressRequest struct {
	NewReservedAddresses soap.String
}

// LANHostConfigManagement1DeleteReservedAddressResponse describes the response for LANHostConfigManagement1.DeleteReservedAddress API
type LANHostConfigManagement1DeleteReservedAddressResponse struct {
}

func (client *LANHostConfigManagement1) DeleteReservedAddress(request LANHostConfigManagement1DeleteReservedAddressRequest) (*LANHostConfigManagement1DeleteReservedAddressResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1DeleteReservedAddressResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteReservedAddress", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1GetReservedAddressesResponse describes the response for LANHostConfigManagement1.GetReservedAddresses API
type LANHostConfigManagement1GetReservedAddressesResponse struct {
	NewReservedAddresses soap.String
}

func (client *LANHostConfigManagement1) GetReservedAddresses() (*LANHostConfigManagement1GetReservedAddressesResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1GetReservedAddressesResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetReservedAddresses", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1SetDNSServerRequest describes the request for LANHostConfigManagement1.SetDNSServer API
type LANHostConfigManagement1SetDNSServerRequest struct {
	NewDNSServers soap.String
}

// LANHostConfigManagement1SetDNSServerResponse describes the response for LANHostConfigManagement1.SetDNSServer API
type LANHostConfigManagement1SetDNSServerResponse struct {
}

func (client *LANHostConfigManagement1) SetDNSServer(request LANHostConfigManagement1SetDNSServerRequest) (*LANHostConfigManagement1SetDNSServerResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1SetDNSServerResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDNSServer", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1DeleteDNSServerRequest describes the request for LANHostConfigManagement1.DeleteDNSServer API
type LANHostConfigManagement1DeleteDNSServerRequest struct {
	NewDNSServers soap.String
}

// LANHostConfigManagement1DeleteDNSServerResponse describes the response for LANHostConfigManagement1.DeleteDNSServer API
type LANHostConfigManagement1DeleteDNSServerResponse struct {
}

func (client *LANHostConfigManagement1) DeleteDNSServer(request LANHostConfigManagement1DeleteDNSServerRequest) (*LANHostConfigManagement1DeleteDNSServerResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1DeleteDNSServerResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteDNSServer", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// LANHostConfigManagement1GetDNSServersResponse describes the response for LANHostConfigManagement1.GetDNSServers API
type LANHostConfigManagement1GetDNSServersResponse struct {
	NewDNSServers soap.String
}

func (client *LANHostConfigManagement1) GetDNSServers() (*LANHostConfigManagement1GetDNSServersResponse, error) {
	// Perform the SOAP call.
	var response LANHostConfigManagement1GetDNSServersResponse
	if err := client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDNSServers", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// Layer3Forwarding1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:Layer3Forwarding:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Layer3Forwarding1 struct {
	goupnp.ServiceClient
}

// NewLayer3Forwarding1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewLayer3Forwarding1Clients() (clients []*Layer3Forwarding1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_Layer3Forwarding_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newLayer3Forwarding1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewLayer3Forwarding1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewLayer3Forwarding1ClientsByURL(loc *url.URL) ([]*Layer3Forwarding1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_Layer3Forwarding_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newLayer3Forwarding1ClientsFromGenericClients(genericClients), nil
}

// NewLayer3Forwarding1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewLayer3Forwarding1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Layer3Forwarding1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Layer3Forwarding_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newLayer3Forwarding1ClientsFromGenericClients(genericClients), nil
}

func newLayer3Forwarding1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Layer3Forwarding1 {
	clients := make([]*Layer3Forwarding1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Layer3Forwarding1{genericClients[i]}
	}
	return clients
}

// Layer3Forwarding1SetDefaultConnectionServiceRequest describes the request for Layer3Forwarding1.SetDefaultConnectionService API
type Layer3Forwarding1SetDefaultConnectionServiceRequest struct {
	NewDefaultConnectionService soap.String
}

// Layer3Forwarding1SetDefaultConnectionServiceResponse describes the response for Layer3Forwarding1.SetDefaultConnectionService API
type Layer3Forwarding1SetDefaultConnectionServiceResponse struct {
}

func (client *Layer3Forwarding1) SetDefaultConnectionService(request Layer3Forwarding1SetDefaultConnectionServiceRequest) (*Layer3Forwarding1SetDefaultConnectionServiceResponse, error) {
	// Perform the SOAP call.
	var response Layer3Forwarding1SetDefaultConnectionServiceResponse
	if err := client.SOAPClient.PerformAction(URN_Layer3Forwarding_1, "SetDefaultConnectionService", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// Layer3Forwarding1GetDefaultConnectionServiceResponse describes the response for Layer3Forwarding1.GetDefaultConnectionService API
type Layer3Forwarding1GetDefaultConnectionServiceResponse struct {
	NewDefaultConnectionService soap.String
}

func (client *Layer3Forwarding1) GetDefaultConnectionService() (*Layer3Forwarding1GetDefaultConnectionServiceResponse, error) {
	// Perform the SOAP call.
	var response Layer3Forwarding1GetDefaultConnectionServiceResponse
	if err := client.SOAPClient.PerformAction(URN_Layer3Forwarding_1, "GetDefaultConnectionService", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCableLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANCableLinkConfig:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type WANCableLinkConfig1 struct {
	goupnp.ServiceClient
}

// NewWANCableLinkConfig1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANCableLinkConfig1Clients() (clients []*WANCableLinkConfig1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_WANCableLinkConfig_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newWANCableLinkConfig1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewWANCableLinkConfig1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANCableLinkConfig1ClientsByURL(loc *url.URL) ([]*WANCableLinkConfig1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_WANCableLinkConfig_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANCableLinkConfig1ClientsFromGenericClients(genericClients), nil
}

// NewWANCableLinkConfig1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANCableLinkConfig1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*WANCableLinkConfig1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_WANCableLinkConfig_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANCableLinkConfig1ClientsFromGenericClients(genericClients), nil
}

func newWANCableLinkConfig1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*WANCableLinkConfig1 {
	clients := make([]*WANCableLinkConfig1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANCableLinkConfig1{genericClients[i]}
	}
	return clients
}

// WANCableLinkConfig1GetCableLinkConfigInfoResponse describes the response for WANCableLinkConfig1.GetCableLinkConfigInfo API
type WANCableLinkConfig1GetCableLinkConfigInfoResponse struct {
	// NewCableLinkConfigState: allowed values: notReady, dsSyncComplete, usParamAcquired, rangingComplete, ipComplete, todEstablished, paramTransferComplete, registrationComplete, operational, accessDenied
	NewCableLinkConfigState soap.String
	// NewLinkType: allowed values: Ethernet
	NewLinkType soap.String
}

//
// Return value:
//
//  WANCableLinkConfig1GetCableLinkConfigInfoResponse
func (client *WANCableLinkConfig1) GetCableLinkConfigInfo() (*WANCableLinkConfig1GetCableLinkConfigInfoResponse, error) {
	// Perform the SOAP call.
	var response WANCableLinkConfig1GetCableLinkConfigInfoResponse
	if err := client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetCableLinkConfigInfo", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCableLinkConfig1GetDownstreamFrequencyResponse describes the response for WANCableLinkConfig1.GetDownstreamFrequency API
type WANCableLinkConfig1GetDownstreamFrequencyResponse struct {
	NewDownstreamFrequency soap.Ui4
}

func (client *WANCableLinkConfig1) GetDownstreamFrequency() (*WANCableLinkConfig1GetDownstreamFrequencyResponse, error) {
	// Perform the SOAP call.
	var response WANCableLinkConfig1GetDownstreamFrequencyResponse
	if err := client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetDownstreamFrequency", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCableLinkConfig1GetDownstreamModulationResponse describes the response for WANCableLinkConfig1.GetDownstreamModulation API
type WANCableLinkConfig1GetDownstreamModulationResponse struct {
	// NewDownstreamModulation: allowed values: 64QAM, 256QAM
	NewDownstreamModulation soap.String
}

//
// Return value:
//
//  WANCableLinkConfig1GetDownstreamModulationResponse
func (client *WANCableLinkConfig1) GetDownstreamModulation() (*WANCableLinkConfig1GetDownstreamModulationResponse, error) {
	// Perform the SOAP call.
	var response WANCableLinkConfig1GetDownstreamModulationResponse
	if err := client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetDownstreamModulation", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCableLinkConfig1GetUpstreamFrequencyResponse describes the response for WANCableLinkConfig1.GetUpstreamFrequency API
type WANCableLinkConfig1GetUpstreamFrequencyResponse struct {
	NewUpstreamFrequency soap.Ui4
}

func (client *WANCableLinkConfig1) GetUpstreamFrequency() (*WANCableLinkConfig1GetUpstreamFrequencyResponse, error) {
	// Perform the SOAP call.
	var response WANCableLinkConfig1GetUpstreamFrequencyResponse
	if err := client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamFrequency", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCableLinkConfig1GetUpstreamModulationResponse describes the response for WANCableLinkConfig1.GetUpstreamModulation API
type WANCableLinkConfig1GetUpstreamModulationResponse struct {
	// NewUpstreamModulation: allowed values: QPSK, 16QAM
	NewUpstreamModulation soap.String
}

//
// Return value:
//
//  WANCableLinkConfig1GetUpstreamModulationResponse
func (client *WANCableLinkConfig1) GetUpstreamModulation() (*WANCableLinkConfig1GetUpstreamModulationResponse, error) {
	// Perform the SOAP call.
	var response WANCableLinkConfig1GetUpstreamModulationResponse
	if err := client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamModulation", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCableLinkConfig1GetUpstreamChannelIDResponse describes the response for WANCableLinkConfig1.GetUpstreamChannelID API
type WANCableLinkConfig1GetUpstreamChannelIDResponse struct {
	NewUpstreamChannelID soap.Ui4
}

func (client *WANCableLinkConfig1) GetUpstreamChannelID() (*WANCableLinkConfig1GetUpstreamChannelIDResponse, error) {
	// Perform the SOAP call.
	var response WANCableLinkConfig1GetUpstreamChannelIDResponse
	if err := client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamChannelID", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCableLinkConfig1GetUpstreamPowerLevelResponse describes the response for WANCableLinkConfig1.GetUpstreamPowerLevel API
type WANCableLinkConfig1GetUpstreamPowerLevelResponse struct {
	NewUpstreamPowerLevel soap.Ui4
}

func (client *WANCableLinkConfig1) GetUpstreamPowerLevel() (*WANCableLinkConfig1GetUpstreamPowerLevelResponse, error) {
	// Perform the SOAP call.
	var response WANCableLinkConfig1GetUpstreamPowerLevelResponse
	if err := client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamPowerLevel", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCableLinkConfig1GetBPIEncryptionEnabledResponse describes the response for WANCableLinkConfig1.GetBPIEncryptionEnabled API
type WANCableLinkConfig1GetBPIEncryptionEnabledResponse struct {
	NewBPIEncryptionEnabled soap.Bool
}

func (client *WANCableLinkConfig1) GetBPIEncryptionEnabled() (*WANCableLinkConfig1GetBPIEncryptionEnabledResponse, error) {
	// Perform the SOAP call.
	var response WANCableLinkConfig1GetBPIEncryptionEnabledResponse
	if err := client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetBPIEncryptionEnabled", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCableLinkConfig1GetConfigFileResponse describes the response for WANCableLinkConfig1.GetConfigFile API
type WANCableLinkConfig1GetConfigFileResponse struct {
	NewConfigFile soap.String
}

func (client *WANCableLinkConfig1) GetConfigFile() (*WANCableLinkConfig1GetConfigFileResponse, error) {
	// Perform the SOAP call.
	var response WANCableLinkConfig1GetConfigFileResponse
	if err := client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetConfigFile", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCableLinkConfig1GetTFTPServerResponse describes the response for WANCableLinkConfig1.GetTFTPServer API
type WANCableLinkConfig1GetTFTPServerResponse struct {
	NewTFTPServer soap.String
}

func (client *WANCableLinkConfig1) GetTFTPServer() (*WANCableLinkConfig1GetTFTPServerResponse, error) {
	// Perform the SOAP call.
	var response WANCableLinkConfig1GetTFTPServerResponse
	if err := client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetTFTPServer", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCommonInterfaceConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type WANCommonInterfaceConfig1 struct {
	goupnp.ServiceClient
}

// NewWANCommonInterfaceConfig1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANCommonInterfaceConfig1Clients() (clients []*WANCommonInterfaceConfig1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_WANCommonInterfaceConfig_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newWANCommonInterfaceConfig1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewWANCommonInterfaceConfig1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANCommonInterfaceConfig1ClientsByURL(loc *url.URL) ([]*WANCommonInterfaceConfig1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_WANCommonInterfaceConfig_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANCommonInterfaceConfig1ClientsFromGenericClients(genericClients), nil
}

// NewWANCommonInterfaceConfig1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANCommonInterfaceConfig1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*WANCommonInterfaceConfig1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_WANCommonInterfaceConfig_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANCommonInterfaceConfig1ClientsFromGenericClients(genericClients), nil
}

func newWANCommonInterfaceConfig1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*WANCommonInterfaceConfig1 {
	clients := make([]*WANCommonInterfaceConfig1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANCommonInterfaceConfig1{genericClients[i]}
	}
	return clients
}

// WANCommonInterfaceConfig1SetEnabledForInternetRequest describes the request for WANCommonInterfaceConfig1.SetEnabledForInternet API
type WANCommonInterfaceConfig1SetEnabledForInternetRequest struct {
	NewEnabledForInternet soap.Bool
}

// WANCommonInterfaceConfig1SetEnabledForInternetResponse describes the response for WANCommonInterfaceConfig1.SetEnabledForInternet API
type WANCommonInterfaceConfig1SetEnabledForInternetResponse struct {
}

func (client *WANCommonInterfaceConfig1) SetEnabledForInternet(request WANCommonInterfaceConfig1SetEnabledForInternetRequest) (*WANCommonInterfaceConfig1SetEnabledForInternetResponse, error) {
	// Perform the SOAP call.
	var response WANCommonInterfaceConfig1SetEnabledForInternetResponse
	if err := client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "SetEnabledForInternet", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCommonInterfaceConfig1GetEnabledForInternetResponse describes the response for WANCommonInterfaceConfig1.GetEnabledForInternet API
type WANCommonInterfaceConfig1GetEnabledForInternetResponse struct {
	NewEnabledForInternet soap.Bool
}

func (client *WANCommonInterfaceConfig1) GetEnabledForInternet() (*WANCommonInterfaceConfig1GetEnabledForInternetResponse, error) {
	// Perform the SOAP call.
	var response WANCommonInterfaceConfig1GetEnabledForInternetResponse
	if err := client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetEnabledForInternet", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCommonInterfaceConfig1GetCommonLinkPropertiesResponse describes the response for WANCommonInterfaceConfig1.GetCommonLinkProperties API
type WANCommonInterfaceConfig1GetCommonLinkPropertiesResponse struct {
	// NewWANAccessType: allowed values: DSL, POTS, Cable, Ethernet
	NewWANAccessType              soap.String
	NewLayer1UpstreamMaxBitRate   soap.Ui4
	NewLayer1DownstreamMaxBitRate soap.Ui4
	// NewPhysicalLinkStatus: allowed values: Up, Down
	NewPhysicalLinkStatus soap.String
}

//
// Return value:
//
//  WANCommonInterfaceConfig1GetCommonLinkPropertiesResponse
func (client *WANCommonInterfaceConfig1) GetCommonLinkProperties() (*WANCommonInterfaceConfig1GetCommonLinkPropertiesResponse, error) {
	// Perform the SOAP call.
	var response WANCommonInterfaceConfig1GetCommonLinkPropertiesResponse
	if err := client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetCommonLinkProperties", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCommonInterfaceConfig1GetWANAccessProviderResponse describes the response for WANCommonInterfaceConfig1.GetWANAccessProvider API
type WANCommonInterfaceConfig1GetWANAccessProviderResponse struct {
	NewWANAccessProvider soap.String
}

func (client *WANCommonInterfaceConfig1) GetWANAccessProvider() (*WANCommonInterfaceConfig1GetWANAccessProviderResponse, error) {
	// Perform the SOAP call.
	var response WANCommonInterfaceConfig1GetWANAccessProviderResponse
	if err := client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetWANAccessProvider", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCommonInterfaceConfig1GetMaximumActiveConnectionsResponse describes the response for WANCommonInterfaceConfig1.GetMaximumActiveConnections API
type WANCommonInterfaceConfig1GetMaximumActiveConnectionsResponse struct {
	// NewMaximumActiveConnections: allowed value range: minimum=1, step=1
	NewMaximumActiveConnections soap.Ui2
}

//
// Return value:
//
//  WANCommonInterfaceConfig1GetMaximumActiveConnectionsResponse
func (client *WANCommonInterfaceConfig1) GetMaximumActiveConnections() (*WANCommonInterfaceConfig1GetMaximumActiveConnectionsResponse, error) {
	// Perform the SOAP call.
	var response WANCommonInterfaceConfig1GetMaximumActiveConnectionsResponse
	if err := client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetMaximumActiveConnections", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCommonInterfaceConfig1GetTotalBytesSentResponse describes the response for WANCommonInterfaceConfig1.GetTotalBytesSent API
type WANCommonInterfaceConfig1GetTotalBytesSentResponse struct {
	NewTotalBytesSent soap.Ui8
}

func (client *WANCommonInterfaceConfig1) GetTotalBytesSent() (*WANCommonInterfaceConfig1GetTotalBytesSentResponse, error) {
	// Perform the SOAP call.
	var response WANCommonInterfaceConfig1GetTotalBytesSentResponse
	if err := client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalBytesSent", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCommonInterfaceConfig1GetTotalBytesReceivedResponse describes the response for WANCommonInterfaceConfig1.GetTotalBytesReceived API
type WANCommonInterfaceConfig1GetTotalBytesReceivedResponse struct {
	NewTotalBytesReceived soap.Ui8
}

func (client *WANCommonInterfaceConfig1) GetTotalBytesReceived() (*WANCommonInterfaceConfig1GetTotalBytesReceivedResponse, error) {
	// Perform the SOAP call.
	var response WANCommonInterfaceConfig1GetTotalBytesReceivedResponse
	if err := client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalBytesReceived", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCommonInterfaceConfig1GetTotalPacketsSentResponse describes the response for WANCommonInterfaceConfig1.GetTotalPacketsSent API
type WANCommonInterfaceConfig1GetTotalPacketsSentResponse struct {
	NewTotalPacketsSent soap.Ui4
}

func (client *WANCommonInterfaceConfig1) GetTotalPacketsSent() (*WANCommonInterfaceConfig1GetTotalPacketsSentResponse, error) {
	// Perform the SOAP call.
	var response WANCommonInterfaceConfig1GetTotalPacketsSentResponse
	if err := client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalPacketsSent", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCommonInterfaceConfig1GetTotalPacketsReceivedResponse describes the response for WANCommonInterfaceConfig1.GetTotalPacketsReceived API
type WANCommonInterfaceConfig1GetTotalPacketsReceivedResponse struct {
	NewTotalPacketsReceived soap.Ui4
}

func (client *WANCommonInterfaceConfig1) GetTotalPacketsReceived() (*WANCommonInterfaceConfig1GetTotalPacketsReceivedResponse, error) {
	// Perform the SOAP call.
	var response WANCommonInterfaceConfig1GetTotalPacketsReceivedResponse
	if err := client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalPacketsReceived", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANCommonInterfaceConfig1GetActiveConnectionRequest describes the request for WANCommonInterfaceConfig1.GetActiveConnection API
type WANCommonInterfaceConfig1GetActiveConnectionRequest struct {
	NewActiveConnectionIndex soap.Ui2
}

// WANCommonInterfaceConfig1GetActiveConnectionResponse describes the response for WANCommonInterfaceConfig1.GetActiveConnection API
type WANCommonInterfaceConfig1GetActiveConnectionResponse struct {
	NewActiveConnDeviceContainer soap.String
	NewActiveConnectionServiceID soap.String
}

func (client *WANCommonInterfaceConfig1) GetActiveConnection(request WANCommonInterfaceConfig1GetActiveConnectionRequest) (*WANCommonInterfaceConfig1GetActiveConnectionResponse, error) {
	// Perform the SOAP call.
	var response WANCommonInterfaceConfig1GetActiveConnectionResponse
	if err := client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetActiveConnection", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANDSLLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANDSLLinkConfig:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type WANDSLLinkConfig1 struct {
	goupnp.ServiceClient
}

// NewWANDSLLinkConfig1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANDSLLinkConfig1Clients() (clients []*WANDSLLinkConfig1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_WANDSLLinkConfig_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newWANDSLLinkConfig1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewWANDSLLinkConfig1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANDSLLinkConfig1ClientsByURL(loc *url.URL) ([]*WANDSLLinkConfig1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_WANDSLLinkConfig_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANDSLLinkConfig1ClientsFromGenericClients(genericClients), nil
}

// NewWANDSLLinkConfig1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANDSLLinkConfig1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*WANDSLLinkConfig1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_WANDSLLinkConfig_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANDSLLinkConfig1ClientsFromGenericClients(genericClients), nil
}

func newWANDSLLinkConfig1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*WANDSLLinkConfig1 {
	clients := make([]*WANDSLLinkConfig1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANDSLLinkConfig1{genericClients[i]}
	}
	return clients
}

// WANDSLLinkConfig1SetDSLLinkTypeRequest describes the request for WANDSLLinkConfig1.SetDSLLinkType API
type WANDSLLinkConfig1SetDSLLinkTypeRequest struct {
	NewLinkType soap.String
}

// WANDSLLinkConfig1SetDSLLinkTypeResponse describes the response for WANDSLLinkConfig1.SetDSLLinkType API
type WANDSLLinkConfig1SetDSLLinkTypeResponse struct {
}

func (client *WANDSLLinkConfig1) SetDSLLinkType(request WANDSLLinkConfig1SetDSLLinkTypeRequest) (*WANDSLLinkConfig1SetDSLLinkTypeResponse, error) {
	// Perform the SOAP call.
	var response WANDSLLinkConfig1SetDSLLinkTypeResponse
	if err := client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetDSLLinkType", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANDSLLinkConfig1GetDSLLinkInfoResponse describes the response for WANDSLLinkConfig1.GetDSLLinkInfo API
type WANDSLLinkConfig1GetDSLLinkInfoResponse struct {
	NewLinkType soap.String
	// NewLinkStatus: allowed values: Up, Down
	NewLinkStatus soap.String
}

//
// Return value:
//
//  WANDSLLinkConfig1GetDSLLinkInfoResponse
func (client *WANDSLLinkConfig1) GetDSLLinkInfo() (*WANDSLLinkConfig1GetDSLLinkInfoResponse, error) {
	// Perform the SOAP call.
	var response WANDSLLinkConfig1GetDSLLinkInfoResponse
	if err := client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetDSLLinkInfo", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANDSLLinkConfig1GetAutoConfigResponse describes the response for WANDSLLinkConfig1.GetAutoConfig API
type WANDSLLinkConfig1GetAutoConfigResponse struct {
	NewAutoConfig soap.Bool
}

func (client *WANDSLLinkConfig1) GetAutoConfig() (*WANDSLLinkConfig1GetAutoConfigResponse, error) {
	// Perform the SOAP call.
	var response WANDSLLinkConfig1GetAutoConfigResponse
	if err := client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetAutoConfig", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANDSLLinkConfig1GetModulationTypeResponse describes the response for WANDSLLinkConfig1.GetModulationType API
type WANDSLLinkConfig1GetModulationTypeResponse struct {
	NewModulationType soap.String
}

func (client *WANDSLLinkConfig1) GetModulationType() (*WANDSLLinkConfig1GetModulationTypeResponse, error) {
	// Perform the SOAP call.
	var response WANDSLLinkConfig1GetModulationTypeResponse
	if err := client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetModulationType", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANDSLLinkConfig1SetDestinationAddressRequest describes the request for WANDSLLinkConfig1.SetDestinationAddress API
type WANDSLLinkConfig1SetDestinationAddressRequest struct {
	NewDestinationAddress soap.String
}

// WANDSLLinkConfig1SetDestinationAddressResponse describes the response for WANDSLLinkConfig1.SetDestinationAddress API
type WANDSLLinkConfig1SetDestinationAddressResponse struct {
}

func (client *WANDSLLinkConfig1) SetDestinationAddress(request WANDSLLinkConfig1SetDestinationAddressRequest) (*WANDSLLinkConfig1SetDestinationAddressResponse, error) {
	// Perform the SOAP call.
	var response WANDSLLinkConfig1SetDestinationAddressResponse
	if err := client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetDestinationAddress", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANDSLLinkConfig1GetDestinationAddressResponse describes the response for WANDSLLinkConfig1.GetDestinationAddress API
type WANDSLLinkConfig1GetDestinationAddressResponse struct {
	NewDestinationAddress soap.String
}

func (client *WANDSLLinkConfig1) GetDestinationAddress() (*WANDSLLinkConfig1GetDestinationAddressResponse, error) {
	// Perform the SOAP call.
	var response WANDSLLinkConfig1GetDestinationAddressResponse
	if err := client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetDestinationAddress", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANDSLLinkConfig1SetATMEncapsulationRequest describes the request for WANDSLLinkConfig1.SetATMEncapsulation API
type WANDSLLinkConfig1SetATMEncapsulationRequest struct {
	NewATMEncapsulation soap.String
}

// WANDSLLinkConfig1SetATMEncapsulationResponse describes the response for WANDSLLinkConfig1.SetATMEncapsulation API
type WANDSLLinkConfig1SetATMEncapsulationResponse struct {
}

func (client *WANDSLLinkConfig1) SetATMEncapsulation(request WANDSLLinkConfig1SetATMEncapsulationRequest) (*WANDSLLinkConfig1SetATMEncapsulationResponse, error) {
	// Perform the SOAP call.
	var response WANDSLLinkConfig1SetATMEncapsulationResponse
	if err := client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetATMEncapsulation", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANDSLLinkConfig1GetATMEncapsulationResponse describes the response for WANDSLLinkConfig1.GetATMEncapsulation API
type WANDSLLinkConfig1GetATMEncapsulationResponse struct {
	NewATMEncapsulation soap.String
}

func (client *WANDSLLinkConfig1) GetATMEncapsulation() (*WANDSLLinkConfig1GetATMEncapsulationResponse, error) {
	// Perform the SOAP call.
	var response WANDSLLinkConfig1GetATMEncapsulationResponse
	if err := client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetATMEncapsulation", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANDSLLinkConfig1SetFCSPreservedRequest describes the request for WANDSLLinkConfig1.SetFCSPreserved API
type WANDSLLinkConfig1SetFCSPreservedRequest struct {
	NewFCSPreserved soap.Bool
}

// WANDSLLinkConfig1SetFCSPreservedResponse describes the response for WANDSLLinkConfig1.SetFCSPreserved API
type WANDSLLinkConfig1SetFCSPreservedResponse struct {
}

func (client *WANDSLLinkConfig1) SetFCSPreserved(request WANDSLLinkConfig1SetFCSPreservedRequest) (*WANDSLLinkConfig1SetFCSPreservedResponse, error) {
	// Perform the SOAP call.
	var response WANDSLLinkConfig1SetFCSPreservedResponse
	if err := client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetFCSPreserved", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANDSLLinkConfig1GetFCSPreservedResponse describes the response for WANDSLLinkConfig1.GetFCSPreserved API
type WANDSLLinkConfig1GetFCSPreservedResponse struct {
	NewFCSPreserved soap.Bool
}

func (client *WANDSLLinkConfig1) GetFCSPreserved() (*WANDSLLinkConfig1GetFCSPreservedResponse, error) {
	// Perform the SOAP call.
	var response WANDSLLinkConfig1GetFCSPreservedResponse
	if err := client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetFCSPreserved", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANEthernetLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANEthernetLinkConfig:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type WANEthernetLinkConfig1 struct {
	goupnp.ServiceClient
}

// NewWANEthernetLinkConfig1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANEthernetLinkConfig1Clients() (clients []*WANEthernetLinkConfig1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_WANEthernetLinkConfig_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newWANEthernetLinkConfig1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewWANEthernetLinkConfig1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANEthernetLinkConfig1ClientsByURL(loc *url.URL) ([]*WANEthernetLinkConfig1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_WANEthernetLinkConfig_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANEthernetLinkConfig1ClientsFromGenericClients(genericClients), nil
}

// NewWANEthernetLinkConfig1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANEthernetLinkConfig1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*WANEthernetLinkConfig1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_WANEthernetLinkConfig_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANEthernetLinkConfig1ClientsFromGenericClients(genericClients), nil
}

func newWANEthernetLinkConfig1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*WANEthernetLinkConfig1 {
	clients := make([]*WANEthernetLinkConfig1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANEthernetLinkConfig1{genericClients[i]}
	}
	return clients
}

// WANEthernetLinkConfig1GetEthernetLinkStatusResponse describes the response for WANEthernetLinkConfig1.GetEthernetLinkStatus API
type WANEthernetLinkConfig1GetEthernetLinkStatusResponse struct {
	// NewEthernetLinkStatus: allowed values: Up, Down
	NewEthernetLinkStatus soap.String
}

//
// Return value:
//
//  WANEthernetLinkConfig1GetEthernetLinkStatusResponse
func (client *WANEthernetLinkConfig1) GetEthernetLinkStatus() (*WANEthernetLinkConfig1GetEthernetLinkStatusResponse, error) {
	// Perform the SOAP call.
	var response WANEthernetLinkConfig1GetEthernetLinkStatusResponse
	if err := client.SOAPClient.PerformAction(URN_WANEthernetLinkConfig_1, "GetEthernetLinkStatus", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANIPConnection:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type WANIPConnection1 struct {
	goupnp.ServiceClient
}

// NewWANIPConnection1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANIPConnection1Clients() (clients []*WANIPConnection1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_WANIPConnection_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newWANIPConnection1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewWANIPConnection1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANIPConnection1ClientsByURL(loc *url.URL) ([]*WANIPConnection1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_WANIPConnection_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANIPConnection1ClientsFromGenericClients(genericClients), nil
}

// NewWANIPConnection1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANIPConnection1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*WANIPConnection1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_WANIPConnection_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANIPConnection1ClientsFromGenericClients(genericClients), nil
}

func newWANIPConnection1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*WANIPConnection1 {
	clients := make([]*WANIPConnection1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANIPConnection1{genericClients[i]}
	}
	return clients
}

// WANIPConnection1SetConnectionTypeRequest describes the request for WANIPConnection1.SetConnectionType API
type WANIPConnection1SetConnectionTypeRequest struct {
	NewConnectionType soap.String
}

// WANIPConnection1SetConnectionTypeResponse describes the response for WANIPConnection1.SetConnectionType API
type WANIPConnection1SetConnectionTypeResponse struct {
}

func (client *WANIPConnection1) SetConnectionType(request WANIPConnection1SetConnectionTypeRequest) (*WANIPConnection1SetConnectionTypeResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1SetConnectionTypeResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetConnectionType", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1GetConnectionTypeInfoResponse describes the response for WANIPConnection1.GetConnectionTypeInfo API
type WANIPConnection1GetConnectionTypeInfoResponse struct {
	NewConnectionType soap.String
	// NewPossibleConnectionTypes: allowed values: Unconfigured, IP_Routed, IP_Bridged
	NewPossibleConnectionTypes soap.String
}

//
// Return value:
//
//  WANIPConnection1GetConnectionTypeInfoResponse
func (client *WANIPConnection1) GetConnectionTypeInfo() (*WANIPConnection1GetConnectionTypeInfoResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1GetConnectionTypeInfoResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetConnectionTypeInfo", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1RequestConnectionResponse describes the response for WANIPConnection1.RequestConnection API
type WANIPConnection1RequestConnectionResponse struct {
}

func (client *WANIPConnection1) RequestConnection() (*WANIPConnection1RequestConnectionResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1RequestConnectionResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "RequestConnection", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1RequestTerminationResponse describes the response for WANIPConnection1.RequestTermination API
type WANIPConnection1RequestTerminationResponse struct {
}

func (client *WANIPConnection1) RequestTermination() (*WANIPConnection1RequestTerminationResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1RequestTerminationResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "RequestTermination", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1ForceTerminationResponse describes the response for WANIPConnection1.ForceTermination API
type WANIPConnection1ForceTerminationResponse struct {
}

func (client *WANIPConnection1) ForceTermination() (*WANIPConnection1ForceTerminationResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1ForceTerminationResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "ForceTermination", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1SetAutoDisconnectTimeRequest describes the request for WANIPConnection1.SetAutoDisconnectTime API
type WANIPConnection1SetAutoDisconnectTimeRequest struct {
	NewAutoDisconnectTime soap.Ui4
}

// WANIPConnection1SetAutoDisconnectTimeResponse describes the response for WANIPConnection1.SetAutoDisconnectTime API
type WANIPConnection1SetAutoDisconnectTimeResponse struct {
}

func (client *WANIPConnection1) SetAutoDisconnectTime(request WANIPConnection1SetAutoDisconnectTimeRequest) (*WANIPConnection1SetAutoDisconnectTimeResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1SetAutoDisconnectTimeResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetAutoDisconnectTime", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1SetIdleDisconnectTimeRequest describes the request for WANIPConnection1.SetIdleDisconnectTime API
type WANIPConnection1SetIdleDisconnectTimeRequest struct {
	NewIdleDisconnectTime soap.Ui4
}

// WANIPConnection1SetIdleDisconnectTimeResponse describes the response for WANIPConnection1.SetIdleDisconnectTime API
type WANIPConnection1SetIdleDisconnectTimeResponse struct {
}

func (client *WANIPConnection1) SetIdleDisconnectTime(request WANIPConnection1SetIdleDisconnectTimeRequest) (*WANIPConnection1SetIdleDisconnectTimeResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1SetIdleDisconnectTimeResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetIdleDisconnectTime", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1SetWarnDisconnectDelayRequest describes the request for WANIPConnection1.SetWarnDisconnectDelay API
type WANIPConnection1SetWarnDisconnectDelayRequest struct {
	NewWarnDisconnectDelay soap.Ui4
}

// WANIPConnection1SetWarnDisconnectDelayResponse describes the response for WANIPConnection1.SetWarnDisconnectDelay API
type WANIPConnection1SetWarnDisconnectDelayResponse struct {
}

func (client *WANIPConnection1) SetWarnDisconnectDelay(request WANIPConnection1SetWarnDisconnectDelayRequest) (*WANIPConnection1SetWarnDisconnectDelayResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1SetWarnDisconnectDelayResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetWarnDisconnectDelay", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1GetStatusInfoResponse describes the response for WANIPConnection1.GetStatusInfo API
type WANIPConnection1GetStatusInfoResponse struct {
	// NewConnectionStatus: allowed values: Unconfigured, Connected, Disconnected
	NewConnectionStatus soap.String
	// NewLastConnectionError: allowed values: ERROR_NONE
	NewLastConnectionError soap.String
	NewUptime              soap.Ui4
}

//
// Return value:
//
//  WANIPConnection1GetStatusInfoResponse
func (client *WANIPConnection1) GetStatusInfo() (*WANIPConnection1GetStatusInfoResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1GetStatusInfoResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetStatusInfo", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1GetAutoDisconnectTimeResponse describes the response for WANIPConnection1.GetAutoDisconnectTime API
type WANIPConnection1GetAutoDisconnectTimeResponse struct {
	NewAutoDisconnectTime soap.Ui4
}

func (client *WANIPConnection1) GetAutoDisconnectTime() (*WANIPConnection1GetAutoDisconnectTimeResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1GetAutoDisconnectTimeResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetAutoDisconnectTime", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1GetIdleDisconnectTimeResponse describes the response for WANIPConnection1.GetIdleDisconnectTime API
type WANIPConnection1GetIdleDisconnectTimeResponse struct {
	NewIdleDisconnectTime soap.Ui4
}

func (client *WANIPConnection1) GetIdleDisconnectTime() (*WANIPConnection1GetIdleDisconnectTimeResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1GetIdleDisconnectTimeResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetIdleDisconnectTime", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1GetWarnDisconnectDelayResponse describes the response for WANIPConnection1.GetWarnDisconnectDelay API
type WANIPConnection1GetWarnDisconnectDelayResponse struct {
	NewWarnDisconnectDelay soap.Ui4
}

func (client *WANIPConnection1) GetWarnDisconnectDelay() (*WANIPConnection1GetWarnDisconnectDelayResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1GetWarnDisconnectDelayResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetWarnDisconnectDelay", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1GetNATRSIPStatusResponse describes the response for WANIPConnection1.GetNATRSIPStatus API
type WANIPConnection1GetNATRSIPStatusResponse struct {
	NewRSIPAvailable soap.Bool
	NewNATEnabled    soap.Bool
}

func (client *WANIPConnection1) GetNATRSIPStatus() (*WANIPConnection1GetNATRSIPStatusResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1GetNATRSIPStatusResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetNATRSIPStatus", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1GetGenericPortMappingEntryRequest describes the request for WANIPConnection1.GetGenericPortMappingEntry API
type WANIPConnection1GetGenericPortMappingEntryRequest struct {
	NewPortMappingIndex soap.Ui2
}

// WANIPConnection1GetGenericPortMappingEntryResponse describes the response for WANIPConnection1.GetGenericPortMappingEntry API
type WANIPConnection1GetGenericPortMappingEntryResponse struct {
	NewRemoteHost   soap.String
	NewExternalPort soap.Ui2
	// NewProtocol: allowed values: TCP, UDP
	NewProtocol               soap.String
	NewInternalPort           soap.Ui2
	NewInternalClient         soap.String
	NewEnabled                soap.Bool
	NewPortMappingDescription soap.String
	NewLeaseDuration          soap.Ui4
}

//
// Return value:
//
//  WANIPConnection1GetGenericPortMappingEntryResponse
func (client *WANIPConnection1) GetGenericPortMappingEntry(request WANIPConnection1GetGenericPortMappingEntryRequest) (*WANIPConnection1GetGenericPortMappingEntryResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1GetGenericPortMappingEntryResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetGenericPortMappingEntry", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1GetSpecificPortMappingEntryRequest describes the request for WANIPConnection1.GetSpecificPortMappingEntry API
type WANIPConnection1GetSpecificPortMappingEntryRequest struct {
	NewRemoteHost   soap.String
	NewExternalPort soap.Ui2
	// NewProtocol: allowed values: TCP, UDP
	NewProtocol soap.String
}

// WANIPConnection1GetSpecificPortMappingEntryResponse describes the response for WANIPConnection1.GetSpecificPortMappingEntry API
type WANIPConnection1GetSpecificPortMappingEntryResponse struct {
	NewInternalPort           soap.Ui2
	NewInternalClient         soap.String
	NewEnabled                soap.Bool
	NewPortMappingDescription soap.String
	NewLeaseDuration          soap.Ui4
}

//
// Arguments:
//
//  WANIPConnection1GetSpecificPortMappingEntryRequest
func (client *WANIPConnection1) GetSpecificPortMappingEntry(request WANIPConnection1GetSpecificPortMappingEntryRequest) (*WANIPConnection1GetSpecificPortMappingEntryResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1GetSpecificPortMappingEntryResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetSpecificPortMappingEntry", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1AddPortMappingRequest describes the request for WANIPConnection1.AddPortMapping API
type WANIPConnection1AddPortMappingRequest struct {
	NewRemoteHost   soap.String
	NewExternalPort soap.Ui2
	// NewProtocol: allowed values: TCP, UDP
	NewProtocol               soap.String
	NewInternalPort           soap.Ui2
	NewInternalClient         soap.String
	NewEnabled                soap.Bool
	NewPortMappingDescription soap.String
	NewLeaseDuration          soap.Ui4
}

// WANIPConnection1AddPortMappingResponse describes the response for WANIPConnection1.AddPortMapping API
type WANIPConnection1AddPortMappingResponse struct {
}

//
// Arguments:
//
//  WANIPConnection1AddPortMappingRequest
func (client *WANIPConnection1) AddPortMapping(request WANIPConnection1AddPortMappingRequest) (*WANIPConnection1AddPortMappingResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1AddPortMappingResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "AddPortMapping", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1DeletePortMappingRequest describes the request for WANIPConnection1.DeletePortMapping API
type WANIPConnection1DeletePortMappingRequest struct {
	NewRemoteHost   soap.String
	NewExternalPort soap.Ui2
	// NewProtocol: allowed values: TCP, UDP
	NewProtocol soap.String
}

// WANIPConnection1DeletePortMappingResponse describes the response for WANIPConnection1.DeletePortMapping API
type WANIPConnection1DeletePortMappingResponse struct {
}

//
// Arguments:
//
//  WANIPConnection1DeletePortMappingRequest
func (client *WANIPConnection1) DeletePortMapping(request WANIPConnection1DeletePortMappingRequest) (*WANIPConnection1DeletePortMappingResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1DeletePortMappingResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "DeletePortMapping", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANIPConnection1GetExternalIPAddressResponse describes the response for WANIPConnection1.GetExternalIPAddress API
type WANIPConnection1GetExternalIPAddressResponse struct {
	NewExternalIPAddress soap.String
}

func (client *WANIPConnection1) GetExternalIPAddress() (*WANIPConnection1GetExternalIPAddressResponse, error) {
	// Perform the SOAP call.
	var response WANIPConnection1GetExternalIPAddressResponse
	if err := client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetExternalIPAddress", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPOTSLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANPOTSLinkConfig:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type WANPOTSLinkConfig1 struct {
	goupnp.ServiceClient
}

// NewWANPOTSLinkConfig1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANPOTSLinkConfig1Clients() (clients []*WANPOTSLinkConfig1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_WANPOTSLinkConfig_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newWANPOTSLinkConfig1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewWANPOTSLinkConfig1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANPOTSLinkConfig1ClientsByURL(loc *url.URL) ([]*WANPOTSLinkConfig1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_WANPOTSLinkConfig_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANPOTSLinkConfig1ClientsFromGenericClients(genericClients), nil
}

// NewWANPOTSLinkConfig1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANPOTSLinkConfig1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*WANPOTSLinkConfig1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_WANPOTSLinkConfig_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANPOTSLinkConfig1ClientsFromGenericClients(genericClients), nil
}

func newWANPOTSLinkConfig1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*WANPOTSLinkConfig1 {
	clients := make([]*WANPOTSLinkConfig1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANPOTSLinkConfig1{genericClients[i]}
	}
	return clients
}

// WANPOTSLinkConfig1SetISPInfoRequest describes the request for WANPOTSLinkConfig1.SetISPInfo API
type WANPOTSLinkConfig1SetISPInfoRequest struct {
	NewISPPhoneNumber soap.String
	NewISPInfo        soap.String
	// NewLinkType: allowed values: PPP_Dialup
	NewLinkType soap.String
}

// WANPOTSLinkConfig1SetISPInfoResponse describes the response for WANPOTSLinkConfig1.SetISPInfo API
type WANPOTSLinkConfig1SetISPInfoResponse struct {
}

//
// Arguments:
//
//  WANPOTSLinkConfig1SetISPInfoRequest
func (client *WANPOTSLinkConfig1) SetISPInfo(request WANPOTSLinkConfig1SetISPInfoRequest) (*WANPOTSLinkConfig1SetISPInfoResponse, error) {
	// Perform the SOAP call.
	var response WANPOTSLinkConfig1SetISPInfoResponse
	if err := client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "SetISPInfo", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPOTSLinkConfig1SetCallRetryInfoRequest describes the request for WANPOTSLinkConfig1.SetCallRetryInfo API
type WANPOTSLinkConfig1SetCallRetryInfoRequest struct {
	NewNumberOfRetries     soap.Ui4
	NewDelayBetweenRetries soap.Ui4
}

// WANPOTSLinkConfig1SetCallRetryInfoResponse describes the response for WANPOTSLinkConfig1.SetCallRetryInfo API
type WANPOTSLinkConfig1SetCallRetryInfoResponse struct {
}

func (client *WANPOTSLinkConfig1) SetCallRetryInfo(request WANPOTSLinkConfig1SetCallRetryInfoRequest) (*WANPOTSLinkConfig1SetCallRetryInfoResponse, error) {
	// Perform the SOAP call.
	var response WANPOTSLinkConfig1SetCallRetryInfoResponse
	if err := client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "SetCallRetryInfo", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPOTSLinkConfig1GetISPInfoResponse describes the response for WANPOTSLinkConfig1.GetISPInfo API
type WANPOTSLinkConfig1GetISPInfoResponse struct {
	NewISPPhoneNumber soap.String
	NewISPInfo        soap.String
	// NewLinkType: allowed values: PPP_Dialup
	NewLinkType soap.String
}

//
// Return value:
//
//  WANPOTSLinkConfig1GetISPInfoResponse
func (client *WANPOTSLinkConfig1) GetISPInfo() (*WANPOTSLinkConfig1GetISPInfoResponse, error) {
	// Perform the SOAP call.
	var response WANPOTSLinkConfig1GetISPInfoResponse
	if err := client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetISPInfo", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPOTSLinkConfig1GetCallRetryInfoResponse describes the response for WANPOTSLinkConfig1.GetCallRetryInfo API
type WANPOTSLinkConfig1GetCallRetryInfoResponse struct {
	NewNumberOfRetries     soap.Ui4
	NewDelayBetweenRetries soap.Ui4
}

func (client *WANPOTSLinkConfig1) GetCallRetryInfo() (*WANPOTSLinkConfig1GetCallRetryInfoResponse, error) {
	// Perform the SOAP call.
	var response WANPOTSLinkConfig1GetCallRetryInfoResponse
	if err := client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetCallRetryInfo", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPOTSLinkConfig1GetFclassResponse describes the response for WANPOTSLinkConfig1.GetFclass API
type WANPOTSLinkConfig1GetFclassResponse struct {
	NewFclass soap.String
}

func (client *WANPOTSLinkConfig1) GetFclass() (*WANPOTSLinkConfig1GetFclassResponse, error) {
	// Perform the SOAP call.
	var response WANPOTSLinkConfig1GetFclassResponse
	if err := client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetFclass", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPOTSLinkConfig1GetDataModulationSupportedResponse describes the response for WANPOTSLinkConfig1.GetDataModulationSupported API
type WANPOTSLinkConfig1GetDataModulationSupportedResponse struct {
	NewDataModulationSupported soap.String
}

func (client *WANPOTSLinkConfig1) GetDataModulationSupported() (*WANPOTSLinkConfig1GetDataModulationSupportedResponse, error) {
	// Perform the SOAP call.
	var response WANPOTSLinkConfig1GetDataModulationSupportedResponse
	if err := client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataModulationSupported", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPOTSLinkConfig1GetDataProtocolResponse describes the response for WANPOTSLinkConfig1.GetDataProtocol API
type WANPOTSLinkConfig1GetDataProtocolResponse struct {
	NewDataProtocol soap.String
}

func (client *WANPOTSLinkConfig1) GetDataProtocol() (*WANPOTSLinkConfig1GetDataProtocolResponse, error) {
	// Perform the SOAP call.
	var response WANPOTSLinkConfig1GetDataProtocolResponse
	if err := client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataProtocol", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPOTSLinkConfig1GetDataCompressionResponse describes the response for WANPOTSLinkConfig1.GetDataCompression API
type WANPOTSLinkConfig1GetDataCompressionResponse struct {
	NewDataCompression soap.String
}

func (client *WANPOTSLinkConfig1) GetDataCompression() (*WANPOTSLinkConfig1GetDataCompressionResponse, error) {
	// Perform the SOAP call.
	var response WANPOTSLinkConfig1GetDataCompressionResponse
	if err := client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataCompression", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPOTSLinkConfig1GetPlusVTRCommandSupportedResponse describes the response for WANPOTSLinkConfig1.GetPlusVTRCommandSupported API
type WANPOTSLinkConfig1GetPlusVTRCommandSupportedResponse struct {
	NewPlusVTRCommandSupported soap.Bool
}

func (client *WANPOTSLinkConfig1) GetPlusVTRCommandSupported() (*WANPOTSLinkConfig1GetPlusVTRCommandSupportedResponse, error) {
	// Perform the SOAP call.
	var response WANPOTSLinkConfig1GetPlusVTRCommandSupportedResponse
	if err := client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetPlusVTRCommandSupported", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANPPPConnection:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type WANPPPConnection1 struct {
	goupnp.ServiceClient
}

// NewWANPPPConnection1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANPPPConnection1Clients() (clients []*WANPPPConnection1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_WANPPPConnection_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newWANPPPConnection1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewWANPPPConnection1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANPPPConnection1ClientsByURL(loc *url.URL) ([]*WANPPPConnection1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_WANPPPConnection_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANPPPConnection1ClientsFromGenericClients(genericClients), nil
}

// NewWANPPPConnection1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANPPPConnection1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*WANPPPConnection1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_WANPPPConnection_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANPPPConnection1ClientsFromGenericClients(genericClients), nil
}

func newWANPPPConnection1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*WANPPPConnection1 {
	clients := make([]*WANPPPConnection1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANPPPConnection1{genericClients[i]}
	}
	return clients
}

// WANPPPConnection1SetConnectionTypeRequest describes the request for WANPPPConnection1.SetConnectionType API
type WANPPPConnection1SetConnectionTypeRequest struct {
	NewConnectionType soap.String
}

// WANPPPConnection1SetConnectionTypeResponse describes the response for WANPPPConnection1.SetConnectionType API
type WANPPPConnection1SetConnectionTypeResponse struct {
}

func (client *WANPPPConnection1) SetConnectionType(request WANPPPConnection1SetConnectionTypeRequest) (*WANPPPConnection1SetConnectionTypeResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1SetConnectionTypeResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetConnectionType", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetConnectionTypeInfoResponse describes the response for WANPPPConnection1.GetConnectionTypeInfo API
type WANPPPConnection1GetConnectionTypeInfoResponse struct {
	NewConnectionType soap.String
	// NewPossibleConnectionTypes: allowed values: Unconfigured, IP_Routed, DHCP_Spoofed, PPPoE_Bridged, PPTP_Relay, L2TP_Relay, PPPoE_Relay
	NewPossibleConnectionTypes soap.String
}

//
// Return value:
//
//  WANPPPConnection1GetConnectionTypeInfoResponse
func (client *WANPPPConnection1) GetConnectionTypeInfo() (*WANPPPConnection1GetConnectionTypeInfoResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetConnectionTypeInfoResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetConnectionTypeInfo", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1ConfigureConnectionRequest describes the request for WANPPPConnection1.ConfigureConnection API
type WANPPPConnection1ConfigureConnectionRequest struct {
	NewUserName soap.String
	NewPassword soap.String
}

// WANPPPConnection1ConfigureConnectionResponse describes the response for WANPPPConnection1.ConfigureConnection API
type WANPPPConnection1ConfigureConnectionResponse struct {
}

func (client *WANPPPConnection1) ConfigureConnection(request WANPPPConnection1ConfigureConnectionRequest) (*WANPPPConnection1ConfigureConnectionResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1ConfigureConnectionResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "ConfigureConnection", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1RequestConnectionResponse describes the response for WANPPPConnection1.RequestConnection API
type WANPPPConnection1RequestConnectionResponse struct {
}

func (client *WANPPPConnection1) RequestConnection() (*WANPPPConnection1RequestConnectionResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1RequestConnectionResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "RequestConnection", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1RequestTerminationResponse describes the response for WANPPPConnection1.RequestTermination API
type WANPPPConnection1RequestTerminationResponse struct {
}

func (client *WANPPPConnection1) RequestTermination() (*WANPPPConnection1RequestTerminationResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1RequestTerminationResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "RequestTermination", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1ForceTerminationResponse describes the response for WANPPPConnection1.ForceTermination API
type WANPPPConnection1ForceTerminationResponse struct {
}

func (client *WANPPPConnection1) ForceTermination() (*WANPPPConnection1ForceTerminationResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1ForceTerminationResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "ForceTermination", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1SetAutoDisconnectTimeRequest describes the request for WANPPPConnection1.SetAutoDisconnectTime API
type WANPPPConnection1SetAutoDisconnectTimeRequest struct {
	NewAutoDisconnectTime soap.Ui4
}

// WANPPPConnection1SetAutoDisconnectTimeResponse describes the response for WANPPPConnection1.SetAutoDisconnectTime API
type WANPPPConnection1SetAutoDisconnectTimeResponse struct {
}

func (client *WANPPPConnection1) SetAutoDisconnectTime(request WANPPPConnection1SetAutoDisconnectTimeRequest) (*WANPPPConnection1SetAutoDisconnectTimeResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1SetAutoDisconnectTimeResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetAutoDisconnectTime", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1SetIdleDisconnectTimeRequest describes the request for WANPPPConnection1.SetIdleDisconnectTime API
type WANPPPConnection1SetIdleDisconnectTimeRequest struct {
	NewIdleDisconnectTime soap.Ui4
}

// WANPPPConnection1SetIdleDisconnectTimeResponse describes the response for WANPPPConnection1.SetIdleDisconnectTime API
type WANPPPConnection1SetIdleDisconnectTimeResponse struct {
}

func (client *WANPPPConnection1) SetIdleDisconnectTime(request WANPPPConnection1SetIdleDisconnectTimeRequest) (*WANPPPConnection1SetIdleDisconnectTimeResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1SetIdleDisconnectTimeResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetIdleDisconnectTime", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1SetWarnDisconnectDelayRequest describes the request for WANPPPConnection1.SetWarnDisconnectDelay API
type WANPPPConnection1SetWarnDisconnectDelayRequest struct {
	NewWarnDisconnectDelay soap.Ui4
}

// WANPPPConnection1SetWarnDisconnectDelayResponse describes the response for WANPPPConnection1.SetWarnDisconnectDelay API
type WANPPPConnection1SetWarnDisconnectDelayResponse struct {
}

func (client *WANPPPConnection1) SetWarnDisconnectDelay(request WANPPPConnection1SetWarnDisconnectDelayRequest) (*WANPPPConnection1SetWarnDisconnectDelayResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1SetWarnDisconnectDelayResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetWarnDisconnectDelay", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetStatusInfoResponse describes the response for WANPPPConnection1.GetStatusInfo API
type WANPPPConnection1GetStatusInfoResponse struct {
	// NewConnectionStatus: allowed values: Unconfigured, Connected, Disconnected
	NewConnectionStatus soap.String
	// NewLastConnectionError: allowed values: ERROR_NONE
	NewLastConnectionError soap.String
	NewUptime              soap.Ui4
}

//
// Return value:
//
//  WANPPPConnection1GetStatusInfoResponse
func (client *WANPPPConnection1) GetStatusInfo() (*WANPPPConnection1GetStatusInfoResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetStatusInfoResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetStatusInfo", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetLinkLayerMaxBitRatesResponse describes the response for WANPPPConnection1.GetLinkLayerMaxBitRates API
type WANPPPConnection1GetLinkLayerMaxBitRatesResponse struct {
	NewUpstreamMaxBitRate   soap.Ui4
	NewDownstreamMaxBitRate soap.Ui4
}

func (client *WANPPPConnection1) GetLinkLayerMaxBitRates() (*WANPPPConnection1GetLinkLayerMaxBitRatesResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetLinkLayerMaxBitRatesResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetLinkLayerMaxBitRates", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetPPPEncryptionProtocolResponse describes the response for WANPPPConnection1.GetPPPEncryptionProtocol API
type WANPPPConnection1GetPPPEncryptionProtocolResponse struct {
	NewPPPEncryptionProtocol soap.String
}

func (client *WANPPPConnection1) GetPPPEncryptionProtocol() (*WANPPPConnection1GetPPPEncryptionProtocolResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetPPPEncryptionProtocolResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPEncryptionProtocol", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetPPPCompressionProtocolResponse describes the response for WANPPPConnection1.GetPPPCompressionProtocol API
type WANPPPConnection1GetPPPCompressionProtocolResponse struct {
	NewPPPCompressionProtocol soap.String
}

func (client *WANPPPConnection1) GetPPPCompressionProtocol() (*WANPPPConnection1GetPPPCompressionProtocolResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetPPPCompressionProtocolResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPCompressionProtocol", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetPPPAuthenticationProtocolResponse describes the response for WANPPPConnection1.GetPPPAuthenticationProtocol API
type WANPPPConnection1GetPPPAuthenticationProtocolResponse struct {
	NewPPPAuthenticationProtocol soap.String
}

func (client *WANPPPConnection1) GetPPPAuthenticationProtocol() (*WANPPPConnection1GetPPPAuthenticationProtocolResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetPPPAuthenticationProtocolResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPAuthenticationProtocol", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetUserNameResponse describes the response for WANPPPConnection1.GetUserName API
type WANPPPConnection1GetUserNameResponse struct {
	NewUserName soap.String
}

func (client *WANPPPConnection1) GetUserName() (*WANPPPConnection1GetUserNameResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetUserNameResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetUserName", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetPasswordResponse describes the response for WANPPPConnection1.GetPassword API
type WANPPPConnection1GetPasswordResponse struct {
	NewPassword soap.String
}

func (client *WANPPPConnection1) GetPassword() (*WANPPPConnection1GetPasswordResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetPasswordResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPassword", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetAutoDisconnectTimeResponse describes the response for WANPPPConnection1.GetAutoDisconnectTime API
type WANPPPConnection1GetAutoDisconnectTimeResponse struct {
	NewAutoDisconnectTime soap.Ui4
}

func (client *WANPPPConnection1) GetAutoDisconnectTime() (*WANPPPConnection1GetAutoDisconnectTimeResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetAutoDisconnectTimeResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetAutoDisconnectTime", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetIdleDisconnectTimeResponse describes the response for WANPPPConnection1.GetIdleDisconnectTime API
type WANPPPConnection1GetIdleDisconnectTimeResponse struct {
	NewIdleDisconnectTime soap.Ui4
}

func (client *WANPPPConnection1) GetIdleDisconnectTime() (*WANPPPConnection1GetIdleDisconnectTimeResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetIdleDisconnectTimeResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetIdleDisconnectTime", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetWarnDisconnectDelayResponse describes the response for WANPPPConnection1.GetWarnDisconnectDelay API
type WANPPPConnection1GetWarnDisconnectDelayResponse struct {
	NewWarnDisconnectDelay soap.Ui4
}

func (client *WANPPPConnection1) GetWarnDisconnectDelay() (*WANPPPConnection1GetWarnDisconnectDelayResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetWarnDisconnectDelayResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetWarnDisconnectDelay", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetNATRSIPStatusResponse describes the response for WANPPPConnection1.GetNATRSIPStatus API
type WANPPPConnection1GetNATRSIPStatusResponse struct {
	NewRSIPAvailable soap.Bool
	NewNATEnabled    soap.Bool
}

func (client *WANPPPConnection1) GetNATRSIPStatus() (*WANPPPConnection1GetNATRSIPStatusResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetNATRSIPStatusResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetNATRSIPStatus", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetGenericPortMappingEntryRequest describes the request for WANPPPConnection1.GetGenericPortMappingEntry API
type WANPPPConnection1GetGenericPortMappingEntryRequest struct {
	NewPortMappingIndex soap.Ui2
}

// WANPPPConnection1GetGenericPortMappingEntryResponse describes the response for WANPPPConnection1.GetGenericPortMappingEntry API
type WANPPPConnection1GetGenericPortMappingEntryResponse struct {
	NewRemoteHost   soap.String
	NewExternalPort soap.Ui2
	// NewProtocol: allowed values: TCP, UDP
	NewProtocol               soap.String
	NewInternalPort           soap.Ui2
	NewInternalClient         soap.String
	NewEnabled                soap.Bool
	NewPortMappingDescription soap.String
	NewLeaseDuration          soap.Ui4
}

//
// Return value:
//
//  WANPPPConnection1GetGenericPortMappingEntryResponse
func (client *WANPPPConnection1) GetGenericPortMappingEntry(request WANPPPConnection1GetGenericPortMappingEntryRequest) (*WANPPPConnection1GetGenericPortMappingEntryResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetGenericPortMappingEntryResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetGenericPortMappingEntry", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetSpecificPortMappingEntryRequest describes the request for WANPPPConnection1.GetSpecificPortMappingEntry API
type WANPPPConnection1GetSpecificPortMappingEntryRequest struct {
	NewRemoteHost   soap.String
	NewExternalPort soap.Ui2
	// NewProtocol: allowed values: TCP, UDP
	NewProtocol soap.String
}

// WANPPPConnection1GetSpecificPortMappingEntryResponse describes the response for WANPPPConnection1.GetSpecificPortMappingEntry API
type WANPPPConnection1GetSpecificPortMappingEntryResponse struct {
	NewInternalPort           soap.Ui2
	NewInternalClient         soap.String
	NewEnabled                soap.Bool
	NewPortMappingDescription soap.String
	NewLeaseDuration          soap.Ui4
}

//
// Arguments:
//
//  WANPPPConnection1GetSpecificPortMappingEntryRequest
func (client *WANPPPConnection1) GetSpecificPortMappingEntry(request WANPPPConnection1GetSpecificPortMappingEntryRequest) (*WANPPPConnection1GetSpecificPortMappingEntryResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetSpecificPortMappingEntryResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetSpecificPortMappingEntry", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1AddPortMappingRequest describes the request for WANPPPConnection1.AddPortMapping API
type WANPPPConnection1AddPortMappingRequest struct {
	NewRemoteHost   soap.String
	NewExternalPort soap.Ui2
	// NewProtocol: allowed values: TCP, UDP
	NewProtocol               soap.String
	NewInternalPort           soap.Ui2
	NewInternalClient         soap.String
	NewEnabled                soap.Bool
	NewPortMappingDescription soap.String
	NewLeaseDuration          soap.Ui4
}

// WANPPPConnection1AddPortMappingResponse describes the response for WANPPPConnection1.AddPortMapping API
type WANPPPConnection1AddPortMappingResponse struct {
}

//
// Arguments:
//
//  WANPPPConnection1AddPortMappingRequest
func (client *WANPPPConnection1) AddPortMapping(request WANPPPConnection1AddPortMappingRequest) (*WANPPPConnection1AddPortMappingResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1AddPortMappingResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "AddPortMapping", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1DeletePortMappingRequest describes the request for WANPPPConnection1.DeletePortMapping API
type WANPPPConnection1DeletePortMappingRequest struct {
	NewRemoteHost   soap.String
	NewExternalPort soap.Ui2
	// NewProtocol: allowed values: TCP, UDP
	NewProtocol soap.String
}

// WANPPPConnection1DeletePortMappingResponse describes the response for WANPPPConnection1.DeletePortMapping API
type WANPPPConnection1DeletePortMappingResponse struct {
}

//
// Arguments:
//
//  WANPPPConnection1DeletePortMappingRequest
func (client *WANPPPConnection1) DeletePortMapping(request WANPPPConnection1DeletePortMappingRequest) (*WANPPPConnection1DeletePortMappingResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1DeletePortMappingResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "DeletePortMapping", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// WANPPPConnection1GetExternalIPAddressResponse describes the response for WANPPPConnection1.GetExternalIPAddress API
type WANPPPConnection1GetExternalIPAddressResponse struct {
	NewExternalIPAddress soap.String
}

func (client *WANPPPConnection1) GetExternalIPAddress() (*WANPPPConnection1GetExternalIPAddressResponse, error) {
	// Perform the SOAP call.
	var response WANPPPConnection1GetExternalIPAddressResponse
	if err := client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetExternalIPAddress", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}
