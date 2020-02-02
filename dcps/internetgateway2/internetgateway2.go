// Client for UPnP Device Control Protocol Internet Gateway Device v2.
//
// This DCP is documented in detail at: http://upnp.org/specs/gw/UPnP-gw-InternetGatewayDevice-v2-Device.pdf
//
// Typically, use one of the New* functions to create clients for services.
package internetgateway2

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
	URN_WANConnectionDevice_2 = "urn:schemas-upnp-org:device:WANConnectionDevice:2"
	URN_WANDevice_1           = "urn:schemas-upnp-org:device:WANDevice:1"
	URN_WANDevice_2           = "urn:schemas-upnp-org:device:WANDevice:2"
)

// Service URNs:
const (
	URN_DeviceProtection_1         = "urn:schemas-upnp-org:service:DeviceProtection:1"
	URN_LANHostConfigManagement_1  = "urn:schemas-upnp-org:service:LANHostConfigManagement:1"
	URN_Layer3Forwarding_1         = "urn:schemas-upnp-org:service:Layer3Forwarding:1"
	URN_WANCableLinkConfig_1       = "urn:schemas-upnp-org:service:WANCableLinkConfig:1"
	URN_WANCommonInterfaceConfig_1 = "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1"
	URN_WANDSLLinkConfig_1         = "urn:schemas-upnp-org:service:WANDSLLinkConfig:1"
	URN_WANEthernetLinkConfig_1    = "urn:schemas-upnp-org:service:WANEthernetLinkConfig:1"
	URN_WANIPConnection_1          = "urn:schemas-upnp-org:service:WANIPConnection:1"
	URN_WANIPConnection_2          = "urn:schemas-upnp-org:service:WANIPConnection:2"
	URN_WANIPv6FirewallControl_1   = "urn:schemas-upnp-org:service:WANIPv6FirewallControl:1"
	URN_WANPOTSLinkConfig_1        = "urn:schemas-upnp-org:service:WANPOTSLinkConfig:1"
	URN_WANPPPConnection_1         = "urn:schemas-upnp-org:service:WANPPPConnection:1"
)

// DeviceProtection1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:DeviceProtection:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type DeviceProtection1 struct {
	goupnp.ServiceClient
}

// NewDeviceProtection1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewDeviceProtection1Clients() (clients []*DeviceProtection1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_DeviceProtection_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newDeviceProtection1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewDeviceProtection1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewDeviceProtection1ClientsByURL(loc *url.URL) ([]*DeviceProtection1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_DeviceProtection_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newDeviceProtection1ClientsFromGenericClients(genericClients), nil
}

// NewDeviceProtection1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewDeviceProtection1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*DeviceProtection1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_DeviceProtection_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newDeviceProtection1ClientsFromGenericClients(genericClients), nil
}

func newDeviceProtection1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*DeviceProtection1 {
	clients := make([]*DeviceProtection1, len(genericClients))
	for i := range genericClients {
		clients[i] = &DeviceProtection1{genericClients[i]}
	}
	return clients
}

// DeviceProtection1SendSetupMessageRequest describes the request for DeviceProtection1.SendSetupMessage API
type DeviceProtection1SendSetupMessageRequest struct {
	ProtocolType soap.String
	InMessage    soap.BinBase64
}

// DeviceProtection1SendSetupMessageResponse describes the response for DeviceProtection1.SendSetupMessage API
type DeviceProtection1SendSetupMessageResponse struct {
	OutMessage soap.BinBase64
}

func (client *DeviceProtection1) SendSetupMessage(request DeviceProtection1SendSetupMessageRequest) (response *DeviceProtection1SendSetupMessageResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "SendSetupMessage", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// DeviceProtection1GetSupportedProtocolsRequest describes the request for DeviceProtection1.GetSupportedProtocols API
type DeviceProtection1GetSupportedProtocolsRequest struct {
}

// DeviceProtection1GetSupportedProtocolsResponse describes the response for DeviceProtection1.GetSupportedProtocols API
type DeviceProtection1GetSupportedProtocolsResponse struct {
	ProtocolList soap.String
}

func (client *DeviceProtection1) GetSupportedProtocols(request DeviceProtection1GetSupportedProtocolsRequest) (response *DeviceProtection1GetSupportedProtocolsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "GetSupportedProtocols", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// DeviceProtection1GetAssignedRolesRequest describes the request for DeviceProtection1.GetAssignedRoles API
type DeviceProtection1GetAssignedRolesRequest struct {
}

// DeviceProtection1GetAssignedRolesResponse describes the response for DeviceProtection1.GetAssignedRoles API
type DeviceProtection1GetAssignedRolesResponse struct {
	RoleList soap.String
}

func (client *DeviceProtection1) GetAssignedRoles(request DeviceProtection1GetAssignedRolesRequest) (response *DeviceProtection1GetAssignedRolesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "GetAssignedRoles", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// DeviceProtection1GetRolesForActionRequest describes the request for DeviceProtection1.GetRolesForAction API
type DeviceProtection1GetRolesForActionRequest struct {
	DeviceUDN  soap.String
	ServiceId  soap.String
	ActionName soap.String
}

// DeviceProtection1GetRolesForActionResponse describes the response for DeviceProtection1.GetRolesForAction API
type DeviceProtection1GetRolesForActionResponse struct {
	RoleList           soap.String
	RestrictedRoleList soap.String
}

func (client *DeviceProtection1) GetRolesForAction(request DeviceProtection1GetRolesForActionRequest) (response *DeviceProtection1GetRolesForActionResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "GetRolesForAction", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// DeviceProtection1GetUserLoginChallengeRequest describes the request for DeviceProtection1.GetUserLoginChallenge API
type DeviceProtection1GetUserLoginChallengeRequest struct {
	ProtocolType soap.String
	Name         soap.String
}

// DeviceProtection1GetUserLoginChallengeResponse describes the response for DeviceProtection1.GetUserLoginChallenge API
type DeviceProtection1GetUserLoginChallengeResponse struct {
	Salt      soap.BinBase64
	Challenge soap.BinBase64
}

func (client *DeviceProtection1) GetUserLoginChallenge(request DeviceProtection1GetUserLoginChallengeRequest) (response *DeviceProtection1GetUserLoginChallengeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "GetUserLoginChallenge", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// DeviceProtection1UserLoginRequest describes the request for DeviceProtection1.UserLogin API
type DeviceProtection1UserLoginRequest struct {
	ProtocolType  soap.String
	Challenge     soap.BinBase64
	Authenticator soap.BinBase64
}

// DeviceProtection1UserLoginResponse describes the response for DeviceProtection1.UserLogin API
type DeviceProtection1UserLoginResponse struct {
}

func (client *DeviceProtection1) UserLogin(request DeviceProtection1UserLoginRequest) (response *DeviceProtection1UserLoginResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "UserLogin", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// DeviceProtection1UserLogoutRequest describes the request for DeviceProtection1.UserLogout API
type DeviceProtection1UserLogoutRequest struct {
}

// DeviceProtection1UserLogoutResponse describes the response for DeviceProtection1.UserLogout API
type DeviceProtection1UserLogoutResponse struct {
}

func (client *DeviceProtection1) UserLogout(request DeviceProtection1UserLogoutRequest) (response *DeviceProtection1UserLogoutResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "UserLogout", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// DeviceProtection1GetACLDataRequest describes the request for DeviceProtection1.GetACLData API
type DeviceProtection1GetACLDataRequest struct {
}

// DeviceProtection1GetACLDataResponse describes the response for DeviceProtection1.GetACLData API
type DeviceProtection1GetACLDataResponse struct {
	ACL soap.String
}

func (client *DeviceProtection1) GetACLData(request DeviceProtection1GetACLDataRequest) (response *DeviceProtection1GetACLDataResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "GetACLData", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// DeviceProtection1AddIdentityListRequest describes the request for DeviceProtection1.AddIdentityList API
type DeviceProtection1AddIdentityListRequest struct {
	IdentityList soap.String
}

// DeviceProtection1AddIdentityListResponse describes the response for DeviceProtection1.AddIdentityList API
type DeviceProtection1AddIdentityListResponse struct {
	IdentityListResult soap.String
}

func (client *DeviceProtection1) AddIdentityList(request DeviceProtection1AddIdentityListRequest) (response *DeviceProtection1AddIdentityListResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "AddIdentityList", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// DeviceProtection1RemoveIdentityRequest describes the request for DeviceProtection1.RemoveIdentity API
type DeviceProtection1RemoveIdentityRequest struct {
	Identity soap.String
}

// DeviceProtection1RemoveIdentityResponse describes the response for DeviceProtection1.RemoveIdentity API
type DeviceProtection1RemoveIdentityResponse struct {
}

func (client *DeviceProtection1) RemoveIdentity(request DeviceProtection1RemoveIdentityRequest) (response *DeviceProtection1RemoveIdentityResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "RemoveIdentity", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// DeviceProtection1SetUserLoginPasswordRequest describes the request for DeviceProtection1.SetUserLoginPassword API
type DeviceProtection1SetUserLoginPasswordRequest struct {
	ProtocolType soap.String
	Name         soap.String
	Stored       soap.BinBase64
	Salt         soap.BinBase64
}

// DeviceProtection1SetUserLoginPasswordResponse describes the response for DeviceProtection1.SetUserLoginPassword API
type DeviceProtection1SetUserLoginPasswordResponse struct {
}

func (client *DeviceProtection1) SetUserLoginPassword(request DeviceProtection1SetUserLoginPasswordRequest) (response *DeviceProtection1SetUserLoginPasswordResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "SetUserLoginPassword", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// DeviceProtection1AddRolesForIdentityRequest describes the request for DeviceProtection1.AddRolesForIdentity API
type DeviceProtection1AddRolesForIdentityRequest struct {
	Identity soap.String
	RoleList soap.String
}

// DeviceProtection1AddRolesForIdentityResponse describes the response for DeviceProtection1.AddRolesForIdentity API
type DeviceProtection1AddRolesForIdentityResponse struct {
}

func (client *DeviceProtection1) AddRolesForIdentity(request DeviceProtection1AddRolesForIdentityRequest) (response *DeviceProtection1AddRolesForIdentityResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "AddRolesForIdentity", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// DeviceProtection1RemoveRolesForIdentityRequest describes the request for DeviceProtection1.RemoveRolesForIdentity API
type DeviceProtection1RemoveRolesForIdentityRequest struct {
	Identity soap.String
	RoleList soap.String
}

// DeviceProtection1RemoveRolesForIdentityResponse describes the response for DeviceProtection1.RemoveRolesForIdentity API
type DeviceProtection1RemoveRolesForIdentityResponse struct {
}

func (client *DeviceProtection1) RemoveRolesForIdentity(request DeviceProtection1RemoveRolesForIdentityRequest) (response *DeviceProtection1RemoveRolesForIdentityResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_DeviceProtection_1, "RemoveRolesForIdentity", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

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

func (client *LANHostConfigManagement1) SetDHCPServerConfigurable(request LANHostConfigManagement1SetDHCPServerConfigurableRequest) (response *LANHostConfigManagement1SetDHCPServerConfigurableResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDHCPServerConfigurable", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1GetDHCPServerConfigurableRequest describes the request for LANHostConfigManagement1.GetDHCPServerConfigurable API
type LANHostConfigManagement1GetDHCPServerConfigurableRequest struct {
}

// LANHostConfigManagement1GetDHCPServerConfigurableResponse describes the response for LANHostConfigManagement1.GetDHCPServerConfigurable API
type LANHostConfigManagement1GetDHCPServerConfigurableResponse struct {
	NewDHCPServerConfigurable soap.Bool
}

func (client *LANHostConfigManagement1) GetDHCPServerConfigurable(request LANHostConfigManagement1GetDHCPServerConfigurableRequest) (response *LANHostConfigManagement1GetDHCPServerConfigurableResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDHCPServerConfigurable", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1SetDHCPRelayRequest describes the request for LANHostConfigManagement1.SetDHCPRelay API
type LANHostConfigManagement1SetDHCPRelayRequest struct {
	NewDHCPRelay soap.Bool
}

// LANHostConfigManagement1SetDHCPRelayResponse describes the response for LANHostConfigManagement1.SetDHCPRelay API
type LANHostConfigManagement1SetDHCPRelayResponse struct {
}

func (client *LANHostConfigManagement1) SetDHCPRelay(request LANHostConfigManagement1SetDHCPRelayRequest) (response *LANHostConfigManagement1SetDHCPRelayResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDHCPRelay", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1GetDHCPRelayRequest describes the request for LANHostConfigManagement1.GetDHCPRelay API
type LANHostConfigManagement1GetDHCPRelayRequest struct {
}

// LANHostConfigManagement1GetDHCPRelayResponse describes the response for LANHostConfigManagement1.GetDHCPRelay API
type LANHostConfigManagement1GetDHCPRelayResponse struct {
	NewDHCPRelay soap.Bool
}

func (client *LANHostConfigManagement1) GetDHCPRelay(request LANHostConfigManagement1GetDHCPRelayRequest) (response *LANHostConfigManagement1GetDHCPRelayResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDHCPRelay", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1SetSubnetMaskRequest describes the request for LANHostConfigManagement1.SetSubnetMask API
type LANHostConfigManagement1SetSubnetMaskRequest struct {
	NewSubnetMask soap.String
}

// LANHostConfigManagement1SetSubnetMaskResponse describes the response for LANHostConfigManagement1.SetSubnetMask API
type LANHostConfigManagement1SetSubnetMaskResponse struct {
}

func (client *LANHostConfigManagement1) SetSubnetMask(request LANHostConfigManagement1SetSubnetMaskRequest) (response *LANHostConfigManagement1SetSubnetMaskResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetSubnetMask", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1GetSubnetMaskRequest describes the request for LANHostConfigManagement1.GetSubnetMask API
type LANHostConfigManagement1GetSubnetMaskRequest struct {
}

// LANHostConfigManagement1GetSubnetMaskResponse describes the response for LANHostConfigManagement1.GetSubnetMask API
type LANHostConfigManagement1GetSubnetMaskResponse struct {
	NewSubnetMask soap.String
}

func (client *LANHostConfigManagement1) GetSubnetMask(request LANHostConfigManagement1GetSubnetMaskRequest) (response *LANHostConfigManagement1GetSubnetMaskResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetSubnetMask", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1SetIPRouterRequest describes the request for LANHostConfigManagement1.SetIPRouter API
type LANHostConfigManagement1SetIPRouterRequest struct {
	NewIPRouters soap.String
}

// LANHostConfigManagement1SetIPRouterResponse describes the response for LANHostConfigManagement1.SetIPRouter API
type LANHostConfigManagement1SetIPRouterResponse struct {
}

func (client *LANHostConfigManagement1) SetIPRouter(request LANHostConfigManagement1SetIPRouterRequest) (response *LANHostConfigManagement1SetIPRouterResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetIPRouter", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1DeleteIPRouterRequest describes the request for LANHostConfigManagement1.DeleteIPRouter API
type LANHostConfigManagement1DeleteIPRouterRequest struct {
	NewIPRouters soap.String
}

// LANHostConfigManagement1DeleteIPRouterResponse describes the response for LANHostConfigManagement1.DeleteIPRouter API
type LANHostConfigManagement1DeleteIPRouterResponse struct {
}

func (client *LANHostConfigManagement1) DeleteIPRouter(request LANHostConfigManagement1DeleteIPRouterRequest) (response *LANHostConfigManagement1DeleteIPRouterResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteIPRouter", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1GetIPRoutersListRequest describes the request for LANHostConfigManagement1.GetIPRoutersList API
type LANHostConfigManagement1GetIPRoutersListRequest struct {
}

// LANHostConfigManagement1GetIPRoutersListResponse describes the response for LANHostConfigManagement1.GetIPRoutersList API
type LANHostConfigManagement1GetIPRoutersListResponse struct {
	NewIPRouters soap.String
}

func (client *LANHostConfigManagement1) GetIPRoutersList(request LANHostConfigManagement1GetIPRoutersListRequest) (response *LANHostConfigManagement1GetIPRoutersListResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetIPRoutersList", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1SetDomainNameRequest describes the request for LANHostConfigManagement1.SetDomainName API
type LANHostConfigManagement1SetDomainNameRequest struct {
	NewDomainName soap.String
}

// LANHostConfigManagement1SetDomainNameResponse describes the response for LANHostConfigManagement1.SetDomainName API
type LANHostConfigManagement1SetDomainNameResponse struct {
}

func (client *LANHostConfigManagement1) SetDomainName(request LANHostConfigManagement1SetDomainNameRequest) (response *LANHostConfigManagement1SetDomainNameResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDomainName", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1GetDomainNameRequest describes the request for LANHostConfigManagement1.GetDomainName API
type LANHostConfigManagement1GetDomainNameRequest struct {
}

// LANHostConfigManagement1GetDomainNameResponse describes the response for LANHostConfigManagement1.GetDomainName API
type LANHostConfigManagement1GetDomainNameResponse struct {
	NewDomainName soap.String
}

func (client *LANHostConfigManagement1) GetDomainName(request LANHostConfigManagement1GetDomainNameRequest) (response *LANHostConfigManagement1GetDomainNameResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDomainName", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1SetAddressRangeRequest describes the request for LANHostConfigManagement1.SetAddressRange API
type LANHostConfigManagement1SetAddressRangeRequest struct {
	NewMinAddress soap.String
	NewMaxAddress soap.String
}

// LANHostConfigManagement1SetAddressRangeResponse describes the response for LANHostConfigManagement1.SetAddressRange API
type LANHostConfigManagement1SetAddressRangeResponse struct {
}

func (client *LANHostConfigManagement1) SetAddressRange(request LANHostConfigManagement1SetAddressRangeRequest) (response *LANHostConfigManagement1SetAddressRangeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetAddressRange", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1GetAddressRangeRequest describes the request for LANHostConfigManagement1.GetAddressRange API
type LANHostConfigManagement1GetAddressRangeRequest struct {
}

// LANHostConfigManagement1GetAddressRangeResponse describes the response for LANHostConfigManagement1.GetAddressRange API
type LANHostConfigManagement1GetAddressRangeResponse struct {
	NewMinAddress soap.String
	NewMaxAddress soap.String
}

func (client *LANHostConfigManagement1) GetAddressRange(request LANHostConfigManagement1GetAddressRangeRequest) (response *LANHostConfigManagement1GetAddressRangeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetAddressRange", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1SetReservedAddressRequest describes the request for LANHostConfigManagement1.SetReservedAddress API
type LANHostConfigManagement1SetReservedAddressRequest struct {
	NewReservedAddresses soap.String
}

// LANHostConfigManagement1SetReservedAddressResponse describes the response for LANHostConfigManagement1.SetReservedAddress API
type LANHostConfigManagement1SetReservedAddressResponse struct {
}

func (client *LANHostConfigManagement1) SetReservedAddress(request LANHostConfigManagement1SetReservedAddressRequest) (response *LANHostConfigManagement1SetReservedAddressResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetReservedAddress", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1DeleteReservedAddressRequest describes the request for LANHostConfigManagement1.DeleteReservedAddress API
type LANHostConfigManagement1DeleteReservedAddressRequest struct {
	NewReservedAddresses soap.String
}

// LANHostConfigManagement1DeleteReservedAddressResponse describes the response for LANHostConfigManagement1.DeleteReservedAddress API
type LANHostConfigManagement1DeleteReservedAddressResponse struct {
}

func (client *LANHostConfigManagement1) DeleteReservedAddress(request LANHostConfigManagement1DeleteReservedAddressRequest) (response *LANHostConfigManagement1DeleteReservedAddressResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteReservedAddress", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1GetReservedAddressesRequest describes the request for LANHostConfigManagement1.GetReservedAddresses API
type LANHostConfigManagement1GetReservedAddressesRequest struct {
}

// LANHostConfigManagement1GetReservedAddressesResponse describes the response for LANHostConfigManagement1.GetReservedAddresses API
type LANHostConfigManagement1GetReservedAddressesResponse struct {
	NewReservedAddresses soap.String
}

func (client *LANHostConfigManagement1) GetReservedAddresses(request LANHostConfigManagement1GetReservedAddressesRequest) (response *LANHostConfigManagement1GetReservedAddressesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetReservedAddresses", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1SetDNSServerRequest describes the request for LANHostConfigManagement1.SetDNSServer API
type LANHostConfigManagement1SetDNSServerRequest struct {
	NewDNSServers soap.String
}

// LANHostConfigManagement1SetDNSServerResponse describes the response for LANHostConfigManagement1.SetDNSServer API
type LANHostConfigManagement1SetDNSServerResponse struct {
}

func (client *LANHostConfigManagement1) SetDNSServer(request LANHostConfigManagement1SetDNSServerRequest) (response *LANHostConfigManagement1SetDNSServerResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDNSServer", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1DeleteDNSServerRequest describes the request for LANHostConfigManagement1.DeleteDNSServer API
type LANHostConfigManagement1DeleteDNSServerRequest struct {
	NewDNSServers soap.String
}

// LANHostConfigManagement1DeleteDNSServerResponse describes the response for LANHostConfigManagement1.DeleteDNSServer API
type LANHostConfigManagement1DeleteDNSServerResponse struct {
}

func (client *LANHostConfigManagement1) DeleteDNSServer(request LANHostConfigManagement1DeleteDNSServerRequest) (response *LANHostConfigManagement1DeleteDNSServerResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteDNSServer", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// LANHostConfigManagement1GetDNSServersRequest describes the request for LANHostConfigManagement1.GetDNSServers API
type LANHostConfigManagement1GetDNSServersRequest struct {
}

// LANHostConfigManagement1GetDNSServersResponse describes the response for LANHostConfigManagement1.GetDNSServers API
type LANHostConfigManagement1GetDNSServersResponse struct {
	NewDNSServers soap.String
}

func (client *LANHostConfigManagement1) GetDNSServers(request LANHostConfigManagement1GetDNSServersRequest) (response *LANHostConfigManagement1GetDNSServersResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDNSServers", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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

func (client *Layer3Forwarding1) SetDefaultConnectionService(request Layer3Forwarding1SetDefaultConnectionServiceRequest) (response *Layer3Forwarding1SetDefaultConnectionServiceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_Layer3Forwarding_1, "SetDefaultConnectionService", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// Layer3Forwarding1GetDefaultConnectionServiceRequest describes the request for Layer3Forwarding1.GetDefaultConnectionService API
type Layer3Forwarding1GetDefaultConnectionServiceRequest struct {
}

// Layer3Forwarding1GetDefaultConnectionServiceResponse describes the response for Layer3Forwarding1.GetDefaultConnectionService API
type Layer3Forwarding1GetDefaultConnectionServiceResponse struct {
	NewDefaultConnectionService soap.String
}

func (client *Layer3Forwarding1) GetDefaultConnectionService(request Layer3Forwarding1GetDefaultConnectionServiceRequest) (response *Layer3Forwarding1GetDefaultConnectionServiceResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_Layer3Forwarding_1, "GetDefaultConnectionService", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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

// WANCableLinkConfig1GetCableLinkConfigInfoRequest describes the request for WANCableLinkConfig1.GetCableLinkConfigInfo API
type WANCableLinkConfig1GetCableLinkConfigInfoRequest struct {
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
func (client *WANCableLinkConfig1) GetCableLinkConfigInfo(request WANCableLinkConfig1GetCableLinkConfigInfoRequest) (response *WANCableLinkConfig1GetCableLinkConfigInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetCableLinkConfigInfo", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCableLinkConfig1GetDownstreamFrequencyRequest describes the request for WANCableLinkConfig1.GetDownstreamFrequency API
type WANCableLinkConfig1GetDownstreamFrequencyRequest struct {
}

// WANCableLinkConfig1GetDownstreamFrequencyResponse describes the response for WANCableLinkConfig1.GetDownstreamFrequency API
type WANCableLinkConfig1GetDownstreamFrequencyResponse struct {
	NewDownstreamFrequency soap.Ui4
}

func (client *WANCableLinkConfig1) GetDownstreamFrequency(request WANCableLinkConfig1GetDownstreamFrequencyRequest) (response *WANCableLinkConfig1GetDownstreamFrequencyResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetDownstreamFrequency", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCableLinkConfig1GetDownstreamModulationRequest describes the request for WANCableLinkConfig1.GetDownstreamModulation API
type WANCableLinkConfig1GetDownstreamModulationRequest struct {
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
func (client *WANCableLinkConfig1) GetDownstreamModulation(request WANCableLinkConfig1GetDownstreamModulationRequest) (response *WANCableLinkConfig1GetDownstreamModulationResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetDownstreamModulation", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCableLinkConfig1GetUpstreamFrequencyRequest describes the request for WANCableLinkConfig1.GetUpstreamFrequency API
type WANCableLinkConfig1GetUpstreamFrequencyRequest struct {
}

// WANCableLinkConfig1GetUpstreamFrequencyResponse describes the response for WANCableLinkConfig1.GetUpstreamFrequency API
type WANCableLinkConfig1GetUpstreamFrequencyResponse struct {
	NewUpstreamFrequency soap.Ui4
}

func (client *WANCableLinkConfig1) GetUpstreamFrequency(request WANCableLinkConfig1GetUpstreamFrequencyRequest) (response *WANCableLinkConfig1GetUpstreamFrequencyResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamFrequency", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCableLinkConfig1GetUpstreamModulationRequest describes the request for WANCableLinkConfig1.GetUpstreamModulation API
type WANCableLinkConfig1GetUpstreamModulationRequest struct {
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
func (client *WANCableLinkConfig1) GetUpstreamModulation(request WANCableLinkConfig1GetUpstreamModulationRequest) (response *WANCableLinkConfig1GetUpstreamModulationResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamModulation", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCableLinkConfig1GetUpstreamChannelIDRequest describes the request for WANCableLinkConfig1.GetUpstreamChannelID API
type WANCableLinkConfig1GetUpstreamChannelIDRequest struct {
}

// WANCableLinkConfig1GetUpstreamChannelIDResponse describes the response for WANCableLinkConfig1.GetUpstreamChannelID API
type WANCableLinkConfig1GetUpstreamChannelIDResponse struct {
	NewUpstreamChannelID soap.Ui4
}

func (client *WANCableLinkConfig1) GetUpstreamChannelID(request WANCableLinkConfig1GetUpstreamChannelIDRequest) (response *WANCableLinkConfig1GetUpstreamChannelIDResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamChannelID", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCableLinkConfig1GetUpstreamPowerLevelRequest describes the request for WANCableLinkConfig1.GetUpstreamPowerLevel API
type WANCableLinkConfig1GetUpstreamPowerLevelRequest struct {
}

// WANCableLinkConfig1GetUpstreamPowerLevelResponse describes the response for WANCableLinkConfig1.GetUpstreamPowerLevel API
type WANCableLinkConfig1GetUpstreamPowerLevelResponse struct {
	NewUpstreamPowerLevel soap.Ui4
}

func (client *WANCableLinkConfig1) GetUpstreamPowerLevel(request WANCableLinkConfig1GetUpstreamPowerLevelRequest) (response *WANCableLinkConfig1GetUpstreamPowerLevelResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamPowerLevel", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCableLinkConfig1GetBPIEncryptionEnabledRequest describes the request for WANCableLinkConfig1.GetBPIEncryptionEnabled API
type WANCableLinkConfig1GetBPIEncryptionEnabledRequest struct {
}

// WANCableLinkConfig1GetBPIEncryptionEnabledResponse describes the response for WANCableLinkConfig1.GetBPIEncryptionEnabled API
type WANCableLinkConfig1GetBPIEncryptionEnabledResponse struct {
	NewBPIEncryptionEnabled soap.Bool
}

func (client *WANCableLinkConfig1) GetBPIEncryptionEnabled(request WANCableLinkConfig1GetBPIEncryptionEnabledRequest) (response *WANCableLinkConfig1GetBPIEncryptionEnabledResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetBPIEncryptionEnabled", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCableLinkConfig1GetConfigFileRequest describes the request for WANCableLinkConfig1.GetConfigFile API
type WANCableLinkConfig1GetConfigFileRequest struct {
}

// WANCableLinkConfig1GetConfigFileResponse describes the response for WANCableLinkConfig1.GetConfigFile API
type WANCableLinkConfig1GetConfigFileResponse struct {
	NewConfigFile soap.String
}

func (client *WANCableLinkConfig1) GetConfigFile(request WANCableLinkConfig1GetConfigFileRequest) (response *WANCableLinkConfig1GetConfigFileResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetConfigFile", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCableLinkConfig1GetTFTPServerRequest describes the request for WANCableLinkConfig1.GetTFTPServer API
type WANCableLinkConfig1GetTFTPServerRequest struct {
}

// WANCableLinkConfig1GetTFTPServerResponse describes the response for WANCableLinkConfig1.GetTFTPServer API
type WANCableLinkConfig1GetTFTPServerResponse struct {
	NewTFTPServer soap.String
}

func (client *WANCableLinkConfig1) GetTFTPServer(request WANCableLinkConfig1GetTFTPServerRequest) (response *WANCableLinkConfig1GetTFTPServerResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetTFTPServer", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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

func (client *WANCommonInterfaceConfig1) SetEnabledForInternet(request WANCommonInterfaceConfig1SetEnabledForInternetRequest) (response *WANCommonInterfaceConfig1SetEnabledForInternetResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "SetEnabledForInternet", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCommonInterfaceConfig1GetEnabledForInternetRequest describes the request for WANCommonInterfaceConfig1.GetEnabledForInternet API
type WANCommonInterfaceConfig1GetEnabledForInternetRequest struct {
}

// WANCommonInterfaceConfig1GetEnabledForInternetResponse describes the response for WANCommonInterfaceConfig1.GetEnabledForInternet API
type WANCommonInterfaceConfig1GetEnabledForInternetResponse struct {
	NewEnabledForInternet soap.Bool
}

func (client *WANCommonInterfaceConfig1) GetEnabledForInternet(request WANCommonInterfaceConfig1GetEnabledForInternetRequest) (response *WANCommonInterfaceConfig1GetEnabledForInternetResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetEnabledForInternet", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCommonInterfaceConfig1GetCommonLinkPropertiesRequest describes the request for WANCommonInterfaceConfig1.GetCommonLinkProperties API
type WANCommonInterfaceConfig1GetCommonLinkPropertiesRequest struct {
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
func (client *WANCommonInterfaceConfig1) GetCommonLinkProperties(request WANCommonInterfaceConfig1GetCommonLinkPropertiesRequest) (response *WANCommonInterfaceConfig1GetCommonLinkPropertiesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetCommonLinkProperties", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCommonInterfaceConfig1GetWANAccessProviderRequest describes the request for WANCommonInterfaceConfig1.GetWANAccessProvider API
type WANCommonInterfaceConfig1GetWANAccessProviderRequest struct {
}

// WANCommonInterfaceConfig1GetWANAccessProviderResponse describes the response for WANCommonInterfaceConfig1.GetWANAccessProvider API
type WANCommonInterfaceConfig1GetWANAccessProviderResponse struct {
	NewWANAccessProvider soap.String
}

func (client *WANCommonInterfaceConfig1) GetWANAccessProvider(request WANCommonInterfaceConfig1GetWANAccessProviderRequest) (response *WANCommonInterfaceConfig1GetWANAccessProviderResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetWANAccessProvider", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCommonInterfaceConfig1GetMaximumActiveConnectionsRequest describes the request for WANCommonInterfaceConfig1.GetMaximumActiveConnections API
type WANCommonInterfaceConfig1GetMaximumActiveConnectionsRequest struct {
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
func (client *WANCommonInterfaceConfig1) GetMaximumActiveConnections(request WANCommonInterfaceConfig1GetMaximumActiveConnectionsRequest) (response *WANCommonInterfaceConfig1GetMaximumActiveConnectionsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetMaximumActiveConnections", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCommonInterfaceConfig1GetTotalBytesSentRequest describes the request for WANCommonInterfaceConfig1.GetTotalBytesSent API
type WANCommonInterfaceConfig1GetTotalBytesSentRequest struct {
}

// WANCommonInterfaceConfig1GetTotalBytesSentResponse describes the response for WANCommonInterfaceConfig1.GetTotalBytesSent API
type WANCommonInterfaceConfig1GetTotalBytesSentResponse struct {
	NewTotalBytesSent soap.Ui8
}

func (client *WANCommonInterfaceConfig1) GetTotalBytesSent(request WANCommonInterfaceConfig1GetTotalBytesSentRequest) (response *WANCommonInterfaceConfig1GetTotalBytesSentResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalBytesSent", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCommonInterfaceConfig1GetTotalBytesReceivedRequest describes the request for WANCommonInterfaceConfig1.GetTotalBytesReceived API
type WANCommonInterfaceConfig1GetTotalBytesReceivedRequest struct {
}

// WANCommonInterfaceConfig1GetTotalBytesReceivedResponse describes the response for WANCommonInterfaceConfig1.GetTotalBytesReceived API
type WANCommonInterfaceConfig1GetTotalBytesReceivedResponse struct {
	NewTotalBytesReceived soap.Ui8
}

func (client *WANCommonInterfaceConfig1) GetTotalBytesReceived(request WANCommonInterfaceConfig1GetTotalBytesReceivedRequest) (response *WANCommonInterfaceConfig1GetTotalBytesReceivedResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalBytesReceived", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCommonInterfaceConfig1GetTotalPacketsSentRequest describes the request for WANCommonInterfaceConfig1.GetTotalPacketsSent API
type WANCommonInterfaceConfig1GetTotalPacketsSentRequest struct {
}

// WANCommonInterfaceConfig1GetTotalPacketsSentResponse describes the response for WANCommonInterfaceConfig1.GetTotalPacketsSent API
type WANCommonInterfaceConfig1GetTotalPacketsSentResponse struct {
	NewTotalPacketsSent soap.Ui4
}

func (client *WANCommonInterfaceConfig1) GetTotalPacketsSent(request WANCommonInterfaceConfig1GetTotalPacketsSentRequest) (response *WANCommonInterfaceConfig1GetTotalPacketsSentResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalPacketsSent", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANCommonInterfaceConfig1GetTotalPacketsReceivedRequest describes the request for WANCommonInterfaceConfig1.GetTotalPacketsReceived API
type WANCommonInterfaceConfig1GetTotalPacketsReceivedRequest struct {
}

// WANCommonInterfaceConfig1GetTotalPacketsReceivedResponse describes the response for WANCommonInterfaceConfig1.GetTotalPacketsReceived API
type WANCommonInterfaceConfig1GetTotalPacketsReceivedResponse struct {
	NewTotalPacketsReceived soap.Ui4
}

func (client *WANCommonInterfaceConfig1) GetTotalPacketsReceived(request WANCommonInterfaceConfig1GetTotalPacketsReceivedRequest) (response *WANCommonInterfaceConfig1GetTotalPacketsReceivedResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalPacketsReceived", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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

func (client *WANCommonInterfaceConfig1) GetActiveConnection(request WANCommonInterfaceConfig1GetActiveConnectionRequest) (response *WANCommonInterfaceConfig1GetActiveConnectionResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetActiveConnection", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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

func (client *WANDSLLinkConfig1) SetDSLLinkType(request WANDSLLinkConfig1SetDSLLinkTypeRequest) (response *WANDSLLinkConfig1SetDSLLinkTypeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetDSLLinkType", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANDSLLinkConfig1GetDSLLinkInfoRequest describes the request for WANDSLLinkConfig1.GetDSLLinkInfo API
type WANDSLLinkConfig1GetDSLLinkInfoRequest struct {
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
func (client *WANDSLLinkConfig1) GetDSLLinkInfo(request WANDSLLinkConfig1GetDSLLinkInfoRequest) (response *WANDSLLinkConfig1GetDSLLinkInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetDSLLinkInfo", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANDSLLinkConfig1GetAutoConfigRequest describes the request for WANDSLLinkConfig1.GetAutoConfig API
type WANDSLLinkConfig1GetAutoConfigRequest struct {
}

// WANDSLLinkConfig1GetAutoConfigResponse describes the response for WANDSLLinkConfig1.GetAutoConfig API
type WANDSLLinkConfig1GetAutoConfigResponse struct {
	NewAutoConfig soap.Bool
}

func (client *WANDSLLinkConfig1) GetAutoConfig(request WANDSLLinkConfig1GetAutoConfigRequest) (response *WANDSLLinkConfig1GetAutoConfigResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetAutoConfig", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANDSLLinkConfig1GetModulationTypeRequest describes the request for WANDSLLinkConfig1.GetModulationType API
type WANDSLLinkConfig1GetModulationTypeRequest struct {
}

// WANDSLLinkConfig1GetModulationTypeResponse describes the response for WANDSLLinkConfig1.GetModulationType API
type WANDSLLinkConfig1GetModulationTypeResponse struct {
	NewModulationType soap.String
}

func (client *WANDSLLinkConfig1) GetModulationType(request WANDSLLinkConfig1GetModulationTypeRequest) (response *WANDSLLinkConfig1GetModulationTypeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetModulationType", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANDSLLinkConfig1SetDestinationAddressRequest describes the request for WANDSLLinkConfig1.SetDestinationAddress API
type WANDSLLinkConfig1SetDestinationAddressRequest struct {
	NewDestinationAddress soap.String
}

// WANDSLLinkConfig1SetDestinationAddressResponse describes the response for WANDSLLinkConfig1.SetDestinationAddress API
type WANDSLLinkConfig1SetDestinationAddressResponse struct {
}

func (client *WANDSLLinkConfig1) SetDestinationAddress(request WANDSLLinkConfig1SetDestinationAddressRequest) (response *WANDSLLinkConfig1SetDestinationAddressResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetDestinationAddress", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANDSLLinkConfig1GetDestinationAddressRequest describes the request for WANDSLLinkConfig1.GetDestinationAddress API
type WANDSLLinkConfig1GetDestinationAddressRequest struct {
}

// WANDSLLinkConfig1GetDestinationAddressResponse describes the response for WANDSLLinkConfig1.GetDestinationAddress API
type WANDSLLinkConfig1GetDestinationAddressResponse struct {
	NewDestinationAddress soap.String
}

func (client *WANDSLLinkConfig1) GetDestinationAddress(request WANDSLLinkConfig1GetDestinationAddressRequest) (response *WANDSLLinkConfig1GetDestinationAddressResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetDestinationAddress", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANDSLLinkConfig1SetATMEncapsulationRequest describes the request for WANDSLLinkConfig1.SetATMEncapsulation API
type WANDSLLinkConfig1SetATMEncapsulationRequest struct {
	NewATMEncapsulation soap.String
}

// WANDSLLinkConfig1SetATMEncapsulationResponse describes the response for WANDSLLinkConfig1.SetATMEncapsulation API
type WANDSLLinkConfig1SetATMEncapsulationResponse struct {
}

func (client *WANDSLLinkConfig1) SetATMEncapsulation(request WANDSLLinkConfig1SetATMEncapsulationRequest) (response *WANDSLLinkConfig1SetATMEncapsulationResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetATMEncapsulation", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANDSLLinkConfig1GetATMEncapsulationRequest describes the request for WANDSLLinkConfig1.GetATMEncapsulation API
type WANDSLLinkConfig1GetATMEncapsulationRequest struct {
}

// WANDSLLinkConfig1GetATMEncapsulationResponse describes the response for WANDSLLinkConfig1.GetATMEncapsulation API
type WANDSLLinkConfig1GetATMEncapsulationResponse struct {
	NewATMEncapsulation soap.String
}

func (client *WANDSLLinkConfig1) GetATMEncapsulation(request WANDSLLinkConfig1GetATMEncapsulationRequest) (response *WANDSLLinkConfig1GetATMEncapsulationResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetATMEncapsulation", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANDSLLinkConfig1SetFCSPreservedRequest describes the request for WANDSLLinkConfig1.SetFCSPreserved API
type WANDSLLinkConfig1SetFCSPreservedRequest struct {
	NewFCSPreserved soap.Bool
}

// WANDSLLinkConfig1SetFCSPreservedResponse describes the response for WANDSLLinkConfig1.SetFCSPreserved API
type WANDSLLinkConfig1SetFCSPreservedResponse struct {
}

func (client *WANDSLLinkConfig1) SetFCSPreserved(request WANDSLLinkConfig1SetFCSPreservedRequest) (response *WANDSLLinkConfig1SetFCSPreservedResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetFCSPreserved", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANDSLLinkConfig1GetFCSPreservedRequest describes the request for WANDSLLinkConfig1.GetFCSPreserved API
type WANDSLLinkConfig1GetFCSPreservedRequest struct {
}

// WANDSLLinkConfig1GetFCSPreservedResponse describes the response for WANDSLLinkConfig1.GetFCSPreserved API
type WANDSLLinkConfig1GetFCSPreservedResponse struct {
	NewFCSPreserved soap.Bool
}

func (client *WANDSLLinkConfig1) GetFCSPreserved(request WANDSLLinkConfig1GetFCSPreservedRequest) (response *WANDSLLinkConfig1GetFCSPreservedResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetFCSPreserved", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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

// WANEthernetLinkConfig1GetEthernetLinkStatusRequest describes the request for WANEthernetLinkConfig1.GetEthernetLinkStatus API
type WANEthernetLinkConfig1GetEthernetLinkStatusRequest struct {
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
func (client *WANEthernetLinkConfig1) GetEthernetLinkStatus(request WANEthernetLinkConfig1GetEthernetLinkStatusRequest) (response *WANEthernetLinkConfig1GetEthernetLinkStatusResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANEthernetLinkConfig_1, "GetEthernetLinkStatus", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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

func (client *WANIPConnection1) SetConnectionType(request WANIPConnection1SetConnectionTypeRequest) (response *WANIPConnection1SetConnectionTypeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetConnectionType", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1GetConnectionTypeInfoRequest describes the request for WANIPConnection1.GetConnectionTypeInfo API
type WANIPConnection1GetConnectionTypeInfoRequest struct {
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
func (client *WANIPConnection1) GetConnectionTypeInfo(request WANIPConnection1GetConnectionTypeInfoRequest) (response *WANIPConnection1GetConnectionTypeInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetConnectionTypeInfo", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1RequestConnectionRequest describes the request for WANIPConnection1.RequestConnection API
type WANIPConnection1RequestConnectionRequest struct {
}

// WANIPConnection1RequestConnectionResponse describes the response for WANIPConnection1.RequestConnection API
type WANIPConnection1RequestConnectionResponse struct {
}

func (client *WANIPConnection1) RequestConnection(request WANIPConnection1RequestConnectionRequest) (response *WANIPConnection1RequestConnectionResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "RequestConnection", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1RequestTerminationRequest describes the request for WANIPConnection1.RequestTermination API
type WANIPConnection1RequestTerminationRequest struct {
}

// WANIPConnection1RequestTerminationResponse describes the response for WANIPConnection1.RequestTermination API
type WANIPConnection1RequestTerminationResponse struct {
}

func (client *WANIPConnection1) RequestTermination(request WANIPConnection1RequestTerminationRequest) (response *WANIPConnection1RequestTerminationResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "RequestTermination", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1ForceTerminationRequest describes the request for WANIPConnection1.ForceTermination API
type WANIPConnection1ForceTerminationRequest struct {
}

// WANIPConnection1ForceTerminationResponse describes the response for WANIPConnection1.ForceTermination API
type WANIPConnection1ForceTerminationResponse struct {
}

func (client *WANIPConnection1) ForceTermination(request WANIPConnection1ForceTerminationRequest) (response *WANIPConnection1ForceTerminationResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "ForceTermination", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1SetAutoDisconnectTimeRequest describes the request for WANIPConnection1.SetAutoDisconnectTime API
type WANIPConnection1SetAutoDisconnectTimeRequest struct {
	NewAutoDisconnectTime soap.Ui4
}

// WANIPConnection1SetAutoDisconnectTimeResponse describes the response for WANIPConnection1.SetAutoDisconnectTime API
type WANIPConnection1SetAutoDisconnectTimeResponse struct {
}

func (client *WANIPConnection1) SetAutoDisconnectTime(request WANIPConnection1SetAutoDisconnectTimeRequest) (response *WANIPConnection1SetAutoDisconnectTimeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetAutoDisconnectTime", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1SetIdleDisconnectTimeRequest describes the request for WANIPConnection1.SetIdleDisconnectTime API
type WANIPConnection1SetIdleDisconnectTimeRequest struct {
	NewIdleDisconnectTime soap.Ui4
}

// WANIPConnection1SetIdleDisconnectTimeResponse describes the response for WANIPConnection1.SetIdleDisconnectTime API
type WANIPConnection1SetIdleDisconnectTimeResponse struct {
}

func (client *WANIPConnection1) SetIdleDisconnectTime(request WANIPConnection1SetIdleDisconnectTimeRequest) (response *WANIPConnection1SetIdleDisconnectTimeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetIdleDisconnectTime", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1SetWarnDisconnectDelayRequest describes the request for WANIPConnection1.SetWarnDisconnectDelay API
type WANIPConnection1SetWarnDisconnectDelayRequest struct {
	NewWarnDisconnectDelay soap.Ui4
}

// WANIPConnection1SetWarnDisconnectDelayResponse describes the response for WANIPConnection1.SetWarnDisconnectDelay API
type WANIPConnection1SetWarnDisconnectDelayResponse struct {
}

func (client *WANIPConnection1) SetWarnDisconnectDelay(request WANIPConnection1SetWarnDisconnectDelayRequest) (response *WANIPConnection1SetWarnDisconnectDelayResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetWarnDisconnectDelay", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1GetStatusInfoRequest describes the request for WANIPConnection1.GetStatusInfo API
type WANIPConnection1GetStatusInfoRequest struct {
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
func (client *WANIPConnection1) GetStatusInfo(request WANIPConnection1GetStatusInfoRequest) (response *WANIPConnection1GetStatusInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetStatusInfo", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1GetAutoDisconnectTimeRequest describes the request for WANIPConnection1.GetAutoDisconnectTime API
type WANIPConnection1GetAutoDisconnectTimeRequest struct {
}

// WANIPConnection1GetAutoDisconnectTimeResponse describes the response for WANIPConnection1.GetAutoDisconnectTime API
type WANIPConnection1GetAutoDisconnectTimeResponse struct {
	NewAutoDisconnectTime soap.Ui4
}

func (client *WANIPConnection1) GetAutoDisconnectTime(request WANIPConnection1GetAutoDisconnectTimeRequest) (response *WANIPConnection1GetAutoDisconnectTimeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetAutoDisconnectTime", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1GetIdleDisconnectTimeRequest describes the request for WANIPConnection1.GetIdleDisconnectTime API
type WANIPConnection1GetIdleDisconnectTimeRequest struct {
}

// WANIPConnection1GetIdleDisconnectTimeResponse describes the response for WANIPConnection1.GetIdleDisconnectTime API
type WANIPConnection1GetIdleDisconnectTimeResponse struct {
	NewIdleDisconnectTime soap.Ui4
}

func (client *WANIPConnection1) GetIdleDisconnectTime(request WANIPConnection1GetIdleDisconnectTimeRequest) (response *WANIPConnection1GetIdleDisconnectTimeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetIdleDisconnectTime", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1GetWarnDisconnectDelayRequest describes the request for WANIPConnection1.GetWarnDisconnectDelay API
type WANIPConnection1GetWarnDisconnectDelayRequest struct {
}

// WANIPConnection1GetWarnDisconnectDelayResponse describes the response for WANIPConnection1.GetWarnDisconnectDelay API
type WANIPConnection1GetWarnDisconnectDelayResponse struct {
	NewWarnDisconnectDelay soap.Ui4
}

func (client *WANIPConnection1) GetWarnDisconnectDelay(request WANIPConnection1GetWarnDisconnectDelayRequest) (response *WANIPConnection1GetWarnDisconnectDelayResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetWarnDisconnectDelay", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1GetNATRSIPStatusRequest describes the request for WANIPConnection1.GetNATRSIPStatus API
type WANIPConnection1GetNATRSIPStatusRequest struct {
}

// WANIPConnection1GetNATRSIPStatusResponse describes the response for WANIPConnection1.GetNATRSIPStatus API
type WANIPConnection1GetNATRSIPStatusResponse struct {
	NewRSIPAvailable soap.Bool
	NewNATEnabled    soap.Bool
}

func (client *WANIPConnection1) GetNATRSIPStatus(request WANIPConnection1GetNATRSIPStatusRequest) (response *WANIPConnection1GetNATRSIPStatusResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetNATRSIPStatus", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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
func (client *WANIPConnection1) GetGenericPortMappingEntry(request WANIPConnection1GetGenericPortMappingEntryRequest) (response *WANIPConnection1GetGenericPortMappingEntryResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetGenericPortMappingEntry", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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
func (client *WANIPConnection1) GetSpecificPortMappingEntry(request WANIPConnection1GetSpecificPortMappingEntryRequest) (response *WANIPConnection1GetSpecificPortMappingEntryResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetSpecificPortMappingEntry", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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
func (client *WANIPConnection1) AddPortMapping(request WANIPConnection1AddPortMappingRequest) (response *WANIPConnection1AddPortMappingResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "AddPortMapping", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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
func (client *WANIPConnection1) DeletePortMapping(request WANIPConnection1DeletePortMappingRequest) (response *WANIPConnection1DeletePortMappingResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "DeletePortMapping", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection1GetExternalIPAddressRequest describes the request for WANIPConnection1.GetExternalIPAddress API
type WANIPConnection1GetExternalIPAddressRequest struct {
}

// WANIPConnection1GetExternalIPAddressResponse describes the response for WANIPConnection1.GetExternalIPAddress API
type WANIPConnection1GetExternalIPAddressResponse struct {
	NewExternalIPAddress soap.String
}

func (client *WANIPConnection1) GetExternalIPAddress(request WANIPConnection1GetExternalIPAddressRequest) (response *WANIPConnection1GetExternalIPAddressResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetExternalIPAddress", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANIPConnection:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type WANIPConnection2 struct {
	goupnp.ServiceClient
}

// NewWANIPConnection2Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANIPConnection2Clients() (clients []*WANIPConnection2, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_WANIPConnection_2); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newWANIPConnection2ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewWANIPConnection2ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANIPConnection2ClientsByURL(loc *url.URL) ([]*WANIPConnection2, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_WANIPConnection_2)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANIPConnection2ClientsFromGenericClients(genericClients), nil
}

// NewWANIPConnection2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANIPConnection2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*WANIPConnection2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_WANIPConnection_2)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANIPConnection2ClientsFromGenericClients(genericClients), nil
}

func newWANIPConnection2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*WANIPConnection2 {
	clients := make([]*WANIPConnection2, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANIPConnection2{genericClients[i]}
	}
	return clients
}

// WANIPConnection2SetConnectionTypeRequest describes the request for WANIPConnection2.SetConnectionType API
type WANIPConnection2SetConnectionTypeRequest struct {
	NewConnectionType soap.String
}

// WANIPConnection2SetConnectionTypeResponse describes the response for WANIPConnection2.SetConnectionType API
type WANIPConnection2SetConnectionTypeResponse struct {
}

func (client *WANIPConnection2) SetConnectionType(request WANIPConnection2SetConnectionTypeRequest) (response *WANIPConnection2SetConnectionTypeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "SetConnectionType", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2GetConnectionTypeInfoRequest describes the request for WANIPConnection2.GetConnectionTypeInfo API
type WANIPConnection2GetConnectionTypeInfoRequest struct {
}

// WANIPConnection2GetConnectionTypeInfoResponse describes the response for WANIPConnection2.GetConnectionTypeInfo API
type WANIPConnection2GetConnectionTypeInfoResponse struct {
	NewConnectionType          soap.String
	NewPossibleConnectionTypes soap.String
}

func (client *WANIPConnection2) GetConnectionTypeInfo(request WANIPConnection2GetConnectionTypeInfoRequest) (response *WANIPConnection2GetConnectionTypeInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetConnectionTypeInfo", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2RequestConnectionRequest describes the request for WANIPConnection2.RequestConnection API
type WANIPConnection2RequestConnectionRequest struct {
}

// WANIPConnection2RequestConnectionResponse describes the response for WANIPConnection2.RequestConnection API
type WANIPConnection2RequestConnectionResponse struct {
}

func (client *WANIPConnection2) RequestConnection(request WANIPConnection2RequestConnectionRequest) (response *WANIPConnection2RequestConnectionResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "RequestConnection", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2RequestTerminationRequest describes the request for WANIPConnection2.RequestTermination API
type WANIPConnection2RequestTerminationRequest struct {
}

// WANIPConnection2RequestTerminationResponse describes the response for WANIPConnection2.RequestTermination API
type WANIPConnection2RequestTerminationResponse struct {
}

func (client *WANIPConnection2) RequestTermination(request WANIPConnection2RequestTerminationRequest) (response *WANIPConnection2RequestTerminationResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "RequestTermination", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2ForceTerminationRequest describes the request for WANIPConnection2.ForceTermination API
type WANIPConnection2ForceTerminationRequest struct {
}

// WANIPConnection2ForceTerminationResponse describes the response for WANIPConnection2.ForceTermination API
type WANIPConnection2ForceTerminationResponse struct {
}

func (client *WANIPConnection2) ForceTermination(request WANIPConnection2ForceTerminationRequest) (response *WANIPConnection2ForceTerminationResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "ForceTermination", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2SetAutoDisconnectTimeRequest describes the request for WANIPConnection2.SetAutoDisconnectTime API
type WANIPConnection2SetAutoDisconnectTimeRequest struct {
	NewAutoDisconnectTime soap.Ui4
}

// WANIPConnection2SetAutoDisconnectTimeResponse describes the response for WANIPConnection2.SetAutoDisconnectTime API
type WANIPConnection2SetAutoDisconnectTimeResponse struct {
}

func (client *WANIPConnection2) SetAutoDisconnectTime(request WANIPConnection2SetAutoDisconnectTimeRequest) (response *WANIPConnection2SetAutoDisconnectTimeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "SetAutoDisconnectTime", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2SetIdleDisconnectTimeRequest describes the request for WANIPConnection2.SetIdleDisconnectTime API
type WANIPConnection2SetIdleDisconnectTimeRequest struct {
	NewIdleDisconnectTime soap.Ui4
}

// WANIPConnection2SetIdleDisconnectTimeResponse describes the response for WANIPConnection2.SetIdleDisconnectTime API
type WANIPConnection2SetIdleDisconnectTimeResponse struct {
}

func (client *WANIPConnection2) SetIdleDisconnectTime(request WANIPConnection2SetIdleDisconnectTimeRequest) (response *WANIPConnection2SetIdleDisconnectTimeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "SetIdleDisconnectTime", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2SetWarnDisconnectDelayRequest describes the request for WANIPConnection2.SetWarnDisconnectDelay API
type WANIPConnection2SetWarnDisconnectDelayRequest struct {
	NewWarnDisconnectDelay soap.Ui4
}

// WANIPConnection2SetWarnDisconnectDelayResponse describes the response for WANIPConnection2.SetWarnDisconnectDelay API
type WANIPConnection2SetWarnDisconnectDelayResponse struct {
}

func (client *WANIPConnection2) SetWarnDisconnectDelay(request WANIPConnection2SetWarnDisconnectDelayRequest) (response *WANIPConnection2SetWarnDisconnectDelayResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "SetWarnDisconnectDelay", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2GetStatusInfoRequest describes the request for WANIPConnection2.GetStatusInfo API
type WANIPConnection2GetStatusInfoRequest struct {
}

// WANIPConnection2GetStatusInfoResponse describes the response for WANIPConnection2.GetStatusInfo API
type WANIPConnection2GetStatusInfoResponse struct {
	// NewConnectionStatus: allowed values: Unconfigured, Connecting, Connected, PendingDisconnect, Disconnecting, Disconnected
	NewConnectionStatus soap.String
	// NewLastConnectionError: allowed values: ERROR_NONE, ERROR_COMMAND_ABORTED, ERROR_NOT_ENABLED_FOR_INTERNET, ERROR_USER_DISCONNECT, ERROR_ISP_DISCONNECT, ERROR_IDLE_DISCONNECT, ERROR_FORCED_DISCONNECT, ERROR_NO_CARRIER, ERROR_IP_CONFIGURATION, ERROR_UNKNOWN
	NewLastConnectionError soap.String
	NewUptime              soap.Ui4
}

//
// Return value:
//
//  WANIPConnection2GetStatusInfoResponse
func (client *WANIPConnection2) GetStatusInfo(request WANIPConnection2GetStatusInfoRequest) (response *WANIPConnection2GetStatusInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetStatusInfo", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2GetAutoDisconnectTimeRequest describes the request for WANIPConnection2.GetAutoDisconnectTime API
type WANIPConnection2GetAutoDisconnectTimeRequest struct {
}

// WANIPConnection2GetAutoDisconnectTimeResponse describes the response for WANIPConnection2.GetAutoDisconnectTime API
type WANIPConnection2GetAutoDisconnectTimeResponse struct {
	NewAutoDisconnectTime soap.Ui4
}

func (client *WANIPConnection2) GetAutoDisconnectTime(request WANIPConnection2GetAutoDisconnectTimeRequest) (response *WANIPConnection2GetAutoDisconnectTimeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetAutoDisconnectTime", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2GetIdleDisconnectTimeRequest describes the request for WANIPConnection2.GetIdleDisconnectTime API
type WANIPConnection2GetIdleDisconnectTimeRequest struct {
}

// WANIPConnection2GetIdleDisconnectTimeResponse describes the response for WANIPConnection2.GetIdleDisconnectTime API
type WANIPConnection2GetIdleDisconnectTimeResponse struct {
	NewIdleDisconnectTime soap.Ui4
}

func (client *WANIPConnection2) GetIdleDisconnectTime(request WANIPConnection2GetIdleDisconnectTimeRequest) (response *WANIPConnection2GetIdleDisconnectTimeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetIdleDisconnectTime", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2GetWarnDisconnectDelayRequest describes the request for WANIPConnection2.GetWarnDisconnectDelay API
type WANIPConnection2GetWarnDisconnectDelayRequest struct {
}

// WANIPConnection2GetWarnDisconnectDelayResponse describes the response for WANIPConnection2.GetWarnDisconnectDelay API
type WANIPConnection2GetWarnDisconnectDelayResponse struct {
	NewWarnDisconnectDelay soap.Ui4
}

func (client *WANIPConnection2) GetWarnDisconnectDelay(request WANIPConnection2GetWarnDisconnectDelayRequest) (response *WANIPConnection2GetWarnDisconnectDelayResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetWarnDisconnectDelay", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2GetNATRSIPStatusRequest describes the request for WANIPConnection2.GetNATRSIPStatus API
type WANIPConnection2GetNATRSIPStatusRequest struct {
}

// WANIPConnection2GetNATRSIPStatusResponse describes the response for WANIPConnection2.GetNATRSIPStatus API
type WANIPConnection2GetNATRSIPStatusResponse struct {
	NewRSIPAvailable soap.Bool
	NewNATEnabled    soap.Bool
}

func (client *WANIPConnection2) GetNATRSIPStatus(request WANIPConnection2GetNATRSIPStatusRequest) (response *WANIPConnection2GetNATRSIPStatusResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetNATRSIPStatus", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2GetGenericPortMappingEntryRequest describes the request for WANIPConnection2.GetGenericPortMappingEntry API
type WANIPConnection2GetGenericPortMappingEntryRequest struct {
	NewPortMappingIndex soap.Ui2
}

// WANIPConnection2GetGenericPortMappingEntryResponse describes the response for WANIPConnection2.GetGenericPortMappingEntry API
type WANIPConnection2GetGenericPortMappingEntryResponse struct {
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
//  WANIPConnection2GetGenericPortMappingEntryResponse
func (client *WANIPConnection2) GetGenericPortMappingEntry(request WANIPConnection2GetGenericPortMappingEntryRequest) (response *WANIPConnection2GetGenericPortMappingEntryResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetGenericPortMappingEntry", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2GetSpecificPortMappingEntryRequest describes the request for WANIPConnection2.GetSpecificPortMappingEntry API
type WANIPConnection2GetSpecificPortMappingEntryRequest struct {
	NewRemoteHost   soap.String
	NewExternalPort soap.Ui2
	// NewProtocol: allowed values: TCP, UDP
	NewProtocol soap.String
}

// WANIPConnection2GetSpecificPortMappingEntryResponse describes the response for WANIPConnection2.GetSpecificPortMappingEntry API
type WANIPConnection2GetSpecificPortMappingEntryResponse struct {
	NewInternalPort           soap.Ui2
	NewInternalClient         soap.String
	NewEnabled                soap.Bool
	NewPortMappingDescription soap.String
	NewLeaseDuration          soap.Ui4
}

//
// Arguments:
//
//  WANIPConnection2GetSpecificPortMappingEntryRequest
func (client *WANIPConnection2) GetSpecificPortMappingEntry(request WANIPConnection2GetSpecificPortMappingEntryRequest) (response *WANIPConnection2GetSpecificPortMappingEntryResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetSpecificPortMappingEntry", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2AddPortMappingRequest describes the request for WANIPConnection2.AddPortMapping API
type WANIPConnection2AddPortMappingRequest struct {
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

// WANIPConnection2AddPortMappingResponse describes the response for WANIPConnection2.AddPortMapping API
type WANIPConnection2AddPortMappingResponse struct {
}

//
// Arguments:
//
//  WANIPConnection2AddPortMappingRequest
func (client *WANIPConnection2) AddPortMapping(request WANIPConnection2AddPortMappingRequest) (response *WANIPConnection2AddPortMappingResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "AddPortMapping", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2DeletePortMappingRequest describes the request for WANIPConnection2.DeletePortMapping API
type WANIPConnection2DeletePortMappingRequest struct {
	NewRemoteHost   soap.String
	NewExternalPort soap.Ui2
	// NewProtocol: allowed values: TCP, UDP
	NewProtocol soap.String
}

// WANIPConnection2DeletePortMappingResponse describes the response for WANIPConnection2.DeletePortMapping API
type WANIPConnection2DeletePortMappingResponse struct {
}

//
// Arguments:
//
//  WANIPConnection2DeletePortMappingRequest
func (client *WANIPConnection2) DeletePortMapping(request WANIPConnection2DeletePortMappingRequest) (response *WANIPConnection2DeletePortMappingResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "DeletePortMapping", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2DeletePortMappingRangeRequest describes the request for WANIPConnection2.DeletePortMappingRange API
type WANIPConnection2DeletePortMappingRangeRequest struct {
	NewStartPort soap.Ui2
	NewEndPort   soap.Ui2
	// NewProtocol: allowed values: TCP, UDP
	NewProtocol soap.String
	NewManage   soap.Bool
}

// WANIPConnection2DeletePortMappingRangeResponse describes the response for WANIPConnection2.DeletePortMappingRange API
type WANIPConnection2DeletePortMappingRangeResponse struct {
}

//
// Arguments:
//
//  WANIPConnection2DeletePortMappingRangeRequest
func (client *WANIPConnection2) DeletePortMappingRange(request WANIPConnection2DeletePortMappingRangeRequest) (response *WANIPConnection2DeletePortMappingRangeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "DeletePortMappingRange", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2GetExternalIPAddressRequest describes the request for WANIPConnection2.GetExternalIPAddress API
type WANIPConnection2GetExternalIPAddressRequest struct {
}

// WANIPConnection2GetExternalIPAddressResponse describes the response for WANIPConnection2.GetExternalIPAddress API
type WANIPConnection2GetExternalIPAddressResponse struct {
	NewExternalIPAddress soap.String
}

func (client *WANIPConnection2) GetExternalIPAddress(request WANIPConnection2GetExternalIPAddressRequest) (response *WANIPConnection2GetExternalIPAddressResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetExternalIPAddress", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2GetListOfPortMappingsRequest describes the request for WANIPConnection2.GetListOfPortMappings API
type WANIPConnection2GetListOfPortMappingsRequest struct {
	NewStartPort soap.Ui2
	NewEndPort   soap.Ui2
	// NewProtocol: allowed values: TCP, UDP
	NewProtocol      soap.String
	NewManage        soap.Bool
	NewNumberOfPorts soap.Ui2
}

// WANIPConnection2GetListOfPortMappingsResponse describes the response for WANIPConnection2.GetListOfPortMappings API
type WANIPConnection2GetListOfPortMappingsResponse struct {
	NewPortListing soap.String
}

//
// Arguments:
//
//  WANIPConnection2GetListOfPortMappingsRequest
func (client *WANIPConnection2) GetListOfPortMappings(request WANIPConnection2GetListOfPortMappingsRequest) (response *WANIPConnection2GetListOfPortMappingsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "GetListOfPortMappings", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPConnection2AddAnyPortMappingRequest describes the request for WANIPConnection2.AddAnyPortMapping API
type WANIPConnection2AddAnyPortMappingRequest struct {
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

// WANIPConnection2AddAnyPortMappingResponse describes the response for WANIPConnection2.AddAnyPortMapping API
type WANIPConnection2AddAnyPortMappingResponse struct {
	NewReservedPort soap.Ui2
}

//
// Arguments:
//
//  WANIPConnection2AddAnyPortMappingRequest
func (client *WANIPConnection2) AddAnyPortMapping(request WANIPConnection2AddAnyPortMappingRequest) (response *WANIPConnection2AddAnyPortMappingResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_2, "AddAnyPortMapping", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPv6FirewallControl1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANIPv6FirewallControl:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type WANIPv6FirewallControl1 struct {
	goupnp.ServiceClient
}

// NewWANIPv6FirewallControl1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANIPv6FirewallControl1Clients() (clients []*WANIPv6FirewallControl1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_WANIPv6FirewallControl_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newWANIPv6FirewallControl1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewWANIPv6FirewallControl1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANIPv6FirewallControl1ClientsByURL(loc *url.URL) ([]*WANIPv6FirewallControl1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_WANIPv6FirewallControl_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANIPv6FirewallControl1ClientsFromGenericClients(genericClients), nil
}

// NewWANIPv6FirewallControl1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANIPv6FirewallControl1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*WANIPv6FirewallControl1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_WANIPv6FirewallControl_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newWANIPv6FirewallControl1ClientsFromGenericClients(genericClients), nil
}

func newWANIPv6FirewallControl1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*WANIPv6FirewallControl1 {
	clients := make([]*WANIPv6FirewallControl1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANIPv6FirewallControl1{genericClients[i]}
	}
	return clients
}

// WANIPv6FirewallControl1GetFirewallStatusRequest describes the request for WANIPv6FirewallControl1.GetFirewallStatus API
type WANIPv6FirewallControl1GetFirewallStatusRequest struct {
}

// WANIPv6FirewallControl1GetFirewallStatusResponse describes the response for WANIPv6FirewallControl1.GetFirewallStatus API
type WANIPv6FirewallControl1GetFirewallStatusResponse struct {
	FirewallEnabled       soap.Bool
	InboundPinholeAllowed soap.Bool
}

func (client *WANIPv6FirewallControl1) GetFirewallStatus(request WANIPv6FirewallControl1GetFirewallStatusRequest) (response *WANIPv6FirewallControl1GetFirewallStatusResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "GetFirewallStatus", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPv6FirewallControl1GetOutboundPinholeTimeoutRequest describes the request for WANIPv6FirewallControl1.GetOutboundPinholeTimeout API
type WANIPv6FirewallControl1GetOutboundPinholeTimeoutRequest struct {
	RemoteHost     soap.String
	RemotePort     soap.Ui2
	InternalClient soap.String
	InternalPort   soap.Ui2
	Protocol       soap.Ui2
}

// WANIPv6FirewallControl1GetOutboundPinholeTimeoutResponse describes the response for WANIPv6FirewallControl1.GetOutboundPinholeTimeout API
type WANIPv6FirewallControl1GetOutboundPinholeTimeoutResponse struct {
	OutboundPinholeTimeout soap.Ui4
}

func (client *WANIPv6FirewallControl1) GetOutboundPinholeTimeout(request WANIPv6FirewallControl1GetOutboundPinholeTimeoutRequest) (response *WANIPv6FirewallControl1GetOutboundPinholeTimeoutResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "GetOutboundPinholeTimeout", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPv6FirewallControl1AddPinholeRequest describes the request for WANIPv6FirewallControl1.AddPinhole API
type WANIPv6FirewallControl1AddPinholeRequest struct {
	RemoteHost     soap.String
	RemotePort     soap.Ui2
	InternalClient soap.String
	InternalPort   soap.Ui2
	Protocol       soap.Ui2
	// LeaseTime: allowed value range: minimum=1, maximum=86400
	LeaseTime soap.Ui4
}

// WANIPv6FirewallControl1AddPinholeResponse describes the response for WANIPv6FirewallControl1.AddPinhole API
type WANIPv6FirewallControl1AddPinholeResponse struct {
	UniqueID soap.Ui2
}

//
// Arguments:
//
//  WANIPv6FirewallControl1AddPinholeRequest
func (client *WANIPv6FirewallControl1) AddPinhole(request WANIPv6FirewallControl1AddPinholeRequest) (response *WANIPv6FirewallControl1AddPinholeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "AddPinhole", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPv6FirewallControl1UpdatePinholeRequest describes the request for WANIPv6FirewallControl1.UpdatePinhole API
type WANIPv6FirewallControl1UpdatePinholeRequest struct {
	UniqueID soap.Ui2
	// NewLeaseTime: allowed value range: minimum=1, maximum=86400
	NewLeaseTime soap.Ui4
}

// WANIPv6FirewallControl1UpdatePinholeResponse describes the response for WANIPv6FirewallControl1.UpdatePinhole API
type WANIPv6FirewallControl1UpdatePinholeResponse struct {
}

//
// Arguments:
//
//  WANIPv6FirewallControl1UpdatePinholeRequest
func (client *WANIPv6FirewallControl1) UpdatePinhole(request WANIPv6FirewallControl1UpdatePinholeRequest) (response *WANIPv6FirewallControl1UpdatePinholeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "UpdatePinhole", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPv6FirewallControl1DeletePinholeRequest describes the request for WANIPv6FirewallControl1.DeletePinhole API
type WANIPv6FirewallControl1DeletePinholeRequest struct {
	UniqueID soap.Ui2
}

// WANIPv6FirewallControl1DeletePinholeResponse describes the response for WANIPv6FirewallControl1.DeletePinhole API
type WANIPv6FirewallControl1DeletePinholeResponse struct {
}

func (client *WANIPv6FirewallControl1) DeletePinhole(request WANIPv6FirewallControl1DeletePinholeRequest) (response *WANIPv6FirewallControl1DeletePinholeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "DeletePinhole", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPv6FirewallControl1GetPinholePacketsRequest describes the request for WANIPv6FirewallControl1.GetPinholePackets API
type WANIPv6FirewallControl1GetPinholePacketsRequest struct {
	UniqueID soap.Ui2
}

// WANIPv6FirewallControl1GetPinholePacketsResponse describes the response for WANIPv6FirewallControl1.GetPinholePackets API
type WANIPv6FirewallControl1GetPinholePacketsResponse struct {
	PinholePackets soap.Ui4
}

func (client *WANIPv6FirewallControl1) GetPinholePackets(request WANIPv6FirewallControl1GetPinholePacketsRequest) (response *WANIPv6FirewallControl1GetPinholePacketsResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "GetPinholePackets", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANIPv6FirewallControl1CheckPinholeWorkingRequest describes the request for WANIPv6FirewallControl1.CheckPinholeWorking API
type WANIPv6FirewallControl1CheckPinholeWorkingRequest struct {
	UniqueID soap.Ui2
}

// WANIPv6FirewallControl1CheckPinholeWorkingResponse describes the response for WANIPv6FirewallControl1.CheckPinholeWorking API
type WANIPv6FirewallControl1CheckPinholeWorkingResponse struct {
	IsWorking soap.Bool
}

func (client *WANIPv6FirewallControl1) CheckPinholeWorking(request WANIPv6FirewallControl1CheckPinholeWorkingRequest) (response *WANIPv6FirewallControl1CheckPinholeWorkingResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPv6FirewallControl_1, "CheckPinholeWorking", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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
func (client *WANPOTSLinkConfig1) SetISPInfo(request WANPOTSLinkConfig1SetISPInfoRequest) (response *WANPOTSLinkConfig1SetISPInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "SetISPInfo", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPOTSLinkConfig1SetCallRetryInfoRequest describes the request for WANPOTSLinkConfig1.SetCallRetryInfo API
type WANPOTSLinkConfig1SetCallRetryInfoRequest struct {
	NewNumberOfRetries     soap.Ui4
	NewDelayBetweenRetries soap.Ui4
}

// WANPOTSLinkConfig1SetCallRetryInfoResponse describes the response for WANPOTSLinkConfig1.SetCallRetryInfo API
type WANPOTSLinkConfig1SetCallRetryInfoResponse struct {
}

func (client *WANPOTSLinkConfig1) SetCallRetryInfo(request WANPOTSLinkConfig1SetCallRetryInfoRequest) (response *WANPOTSLinkConfig1SetCallRetryInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "SetCallRetryInfo", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPOTSLinkConfig1GetISPInfoRequest describes the request for WANPOTSLinkConfig1.GetISPInfo API
type WANPOTSLinkConfig1GetISPInfoRequest struct {
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
func (client *WANPOTSLinkConfig1) GetISPInfo(request WANPOTSLinkConfig1GetISPInfoRequest) (response *WANPOTSLinkConfig1GetISPInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetISPInfo", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPOTSLinkConfig1GetCallRetryInfoRequest describes the request for WANPOTSLinkConfig1.GetCallRetryInfo API
type WANPOTSLinkConfig1GetCallRetryInfoRequest struct {
}

// WANPOTSLinkConfig1GetCallRetryInfoResponse describes the response for WANPOTSLinkConfig1.GetCallRetryInfo API
type WANPOTSLinkConfig1GetCallRetryInfoResponse struct {
	NewNumberOfRetries     soap.Ui4
	NewDelayBetweenRetries soap.Ui4
}

func (client *WANPOTSLinkConfig1) GetCallRetryInfo(request WANPOTSLinkConfig1GetCallRetryInfoRequest) (response *WANPOTSLinkConfig1GetCallRetryInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetCallRetryInfo", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPOTSLinkConfig1GetFclassRequest describes the request for WANPOTSLinkConfig1.GetFclass API
type WANPOTSLinkConfig1GetFclassRequest struct {
}

// WANPOTSLinkConfig1GetFclassResponse describes the response for WANPOTSLinkConfig1.GetFclass API
type WANPOTSLinkConfig1GetFclassResponse struct {
	NewFclass soap.String
}

func (client *WANPOTSLinkConfig1) GetFclass(request WANPOTSLinkConfig1GetFclassRequest) (response *WANPOTSLinkConfig1GetFclassResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetFclass", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPOTSLinkConfig1GetDataModulationSupportedRequest describes the request for WANPOTSLinkConfig1.GetDataModulationSupported API
type WANPOTSLinkConfig1GetDataModulationSupportedRequest struct {
}

// WANPOTSLinkConfig1GetDataModulationSupportedResponse describes the response for WANPOTSLinkConfig1.GetDataModulationSupported API
type WANPOTSLinkConfig1GetDataModulationSupportedResponse struct {
	NewDataModulationSupported soap.String
}

func (client *WANPOTSLinkConfig1) GetDataModulationSupported(request WANPOTSLinkConfig1GetDataModulationSupportedRequest) (response *WANPOTSLinkConfig1GetDataModulationSupportedResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataModulationSupported", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPOTSLinkConfig1GetDataProtocolRequest describes the request for WANPOTSLinkConfig1.GetDataProtocol API
type WANPOTSLinkConfig1GetDataProtocolRequest struct {
}

// WANPOTSLinkConfig1GetDataProtocolResponse describes the response for WANPOTSLinkConfig1.GetDataProtocol API
type WANPOTSLinkConfig1GetDataProtocolResponse struct {
	NewDataProtocol soap.String
}

func (client *WANPOTSLinkConfig1) GetDataProtocol(request WANPOTSLinkConfig1GetDataProtocolRequest) (response *WANPOTSLinkConfig1GetDataProtocolResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataProtocol", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPOTSLinkConfig1GetDataCompressionRequest describes the request for WANPOTSLinkConfig1.GetDataCompression API
type WANPOTSLinkConfig1GetDataCompressionRequest struct {
}

// WANPOTSLinkConfig1GetDataCompressionResponse describes the response for WANPOTSLinkConfig1.GetDataCompression API
type WANPOTSLinkConfig1GetDataCompressionResponse struct {
	NewDataCompression soap.String
}

func (client *WANPOTSLinkConfig1) GetDataCompression(request WANPOTSLinkConfig1GetDataCompressionRequest) (response *WANPOTSLinkConfig1GetDataCompressionResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataCompression", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPOTSLinkConfig1GetPlusVTRCommandSupportedRequest describes the request for WANPOTSLinkConfig1.GetPlusVTRCommandSupported API
type WANPOTSLinkConfig1GetPlusVTRCommandSupportedRequest struct {
}

// WANPOTSLinkConfig1GetPlusVTRCommandSupportedResponse describes the response for WANPOTSLinkConfig1.GetPlusVTRCommandSupported API
type WANPOTSLinkConfig1GetPlusVTRCommandSupportedResponse struct {
	NewPlusVTRCommandSupported soap.Bool
}

func (client *WANPOTSLinkConfig1) GetPlusVTRCommandSupported(request WANPOTSLinkConfig1GetPlusVTRCommandSupportedRequest) (response *WANPOTSLinkConfig1GetPlusVTRCommandSupportedResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetPlusVTRCommandSupported", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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

func (client *WANPPPConnection1) SetConnectionType(request WANPPPConnection1SetConnectionTypeRequest) (response *WANPPPConnection1SetConnectionTypeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetConnectionType", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetConnectionTypeInfoRequest describes the request for WANPPPConnection1.GetConnectionTypeInfo API
type WANPPPConnection1GetConnectionTypeInfoRequest struct {
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
func (client *WANPPPConnection1) GetConnectionTypeInfo(request WANPPPConnection1GetConnectionTypeInfoRequest) (response *WANPPPConnection1GetConnectionTypeInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetConnectionTypeInfo", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1ConfigureConnectionRequest describes the request for WANPPPConnection1.ConfigureConnection API
type WANPPPConnection1ConfigureConnectionRequest struct {
	NewUserName soap.String
	NewPassword soap.String
}

// WANPPPConnection1ConfigureConnectionResponse describes the response for WANPPPConnection1.ConfigureConnection API
type WANPPPConnection1ConfigureConnectionResponse struct {
}

func (client *WANPPPConnection1) ConfigureConnection(request WANPPPConnection1ConfigureConnectionRequest) (response *WANPPPConnection1ConfigureConnectionResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "ConfigureConnection", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1RequestConnectionRequest describes the request for WANPPPConnection1.RequestConnection API
type WANPPPConnection1RequestConnectionRequest struct {
}

// WANPPPConnection1RequestConnectionResponse describes the response for WANPPPConnection1.RequestConnection API
type WANPPPConnection1RequestConnectionResponse struct {
}

func (client *WANPPPConnection1) RequestConnection(request WANPPPConnection1RequestConnectionRequest) (response *WANPPPConnection1RequestConnectionResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "RequestConnection", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1RequestTerminationRequest describes the request for WANPPPConnection1.RequestTermination API
type WANPPPConnection1RequestTerminationRequest struct {
}

// WANPPPConnection1RequestTerminationResponse describes the response for WANPPPConnection1.RequestTermination API
type WANPPPConnection1RequestTerminationResponse struct {
}

func (client *WANPPPConnection1) RequestTermination(request WANPPPConnection1RequestTerminationRequest) (response *WANPPPConnection1RequestTerminationResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "RequestTermination", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1ForceTerminationRequest describes the request for WANPPPConnection1.ForceTermination API
type WANPPPConnection1ForceTerminationRequest struct {
}

// WANPPPConnection1ForceTerminationResponse describes the response for WANPPPConnection1.ForceTermination API
type WANPPPConnection1ForceTerminationResponse struct {
}

func (client *WANPPPConnection1) ForceTermination(request WANPPPConnection1ForceTerminationRequest) (response *WANPPPConnection1ForceTerminationResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "ForceTermination", nil, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1SetAutoDisconnectTimeRequest describes the request for WANPPPConnection1.SetAutoDisconnectTime API
type WANPPPConnection1SetAutoDisconnectTimeRequest struct {
	NewAutoDisconnectTime soap.Ui4
}

// WANPPPConnection1SetAutoDisconnectTimeResponse describes the response for WANPPPConnection1.SetAutoDisconnectTime API
type WANPPPConnection1SetAutoDisconnectTimeResponse struct {
}

func (client *WANPPPConnection1) SetAutoDisconnectTime(request WANPPPConnection1SetAutoDisconnectTimeRequest) (response *WANPPPConnection1SetAutoDisconnectTimeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetAutoDisconnectTime", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1SetIdleDisconnectTimeRequest describes the request for WANPPPConnection1.SetIdleDisconnectTime API
type WANPPPConnection1SetIdleDisconnectTimeRequest struct {
	NewIdleDisconnectTime soap.Ui4
}

// WANPPPConnection1SetIdleDisconnectTimeResponse describes the response for WANPPPConnection1.SetIdleDisconnectTime API
type WANPPPConnection1SetIdleDisconnectTimeResponse struct {
}

func (client *WANPPPConnection1) SetIdleDisconnectTime(request WANPPPConnection1SetIdleDisconnectTimeRequest) (response *WANPPPConnection1SetIdleDisconnectTimeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetIdleDisconnectTime", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1SetWarnDisconnectDelayRequest describes the request for WANPPPConnection1.SetWarnDisconnectDelay API
type WANPPPConnection1SetWarnDisconnectDelayRequest struct {
	NewWarnDisconnectDelay soap.Ui4
}

// WANPPPConnection1SetWarnDisconnectDelayResponse describes the response for WANPPPConnection1.SetWarnDisconnectDelay API
type WANPPPConnection1SetWarnDisconnectDelayResponse struct {
}

func (client *WANPPPConnection1) SetWarnDisconnectDelay(request WANPPPConnection1SetWarnDisconnectDelayRequest) (response *WANPPPConnection1SetWarnDisconnectDelayResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetWarnDisconnectDelay", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetStatusInfoRequest describes the request for WANPPPConnection1.GetStatusInfo API
type WANPPPConnection1GetStatusInfoRequest struct {
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
func (client *WANPPPConnection1) GetStatusInfo(request WANPPPConnection1GetStatusInfoRequest) (response *WANPPPConnection1GetStatusInfoResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetStatusInfo", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetLinkLayerMaxBitRatesRequest describes the request for WANPPPConnection1.GetLinkLayerMaxBitRates API
type WANPPPConnection1GetLinkLayerMaxBitRatesRequest struct {
}

// WANPPPConnection1GetLinkLayerMaxBitRatesResponse describes the response for WANPPPConnection1.GetLinkLayerMaxBitRates API
type WANPPPConnection1GetLinkLayerMaxBitRatesResponse struct {
	NewUpstreamMaxBitRate   soap.Ui4
	NewDownstreamMaxBitRate soap.Ui4
}

func (client *WANPPPConnection1) GetLinkLayerMaxBitRates(request WANPPPConnection1GetLinkLayerMaxBitRatesRequest) (response *WANPPPConnection1GetLinkLayerMaxBitRatesResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetLinkLayerMaxBitRates", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetPPPEncryptionProtocolRequest describes the request for WANPPPConnection1.GetPPPEncryptionProtocol API
type WANPPPConnection1GetPPPEncryptionProtocolRequest struct {
}

// WANPPPConnection1GetPPPEncryptionProtocolResponse describes the response for WANPPPConnection1.GetPPPEncryptionProtocol API
type WANPPPConnection1GetPPPEncryptionProtocolResponse struct {
	NewPPPEncryptionProtocol soap.String
}

func (client *WANPPPConnection1) GetPPPEncryptionProtocol(request WANPPPConnection1GetPPPEncryptionProtocolRequest) (response *WANPPPConnection1GetPPPEncryptionProtocolResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPEncryptionProtocol", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetPPPCompressionProtocolRequest describes the request for WANPPPConnection1.GetPPPCompressionProtocol API
type WANPPPConnection1GetPPPCompressionProtocolRequest struct {
}

// WANPPPConnection1GetPPPCompressionProtocolResponse describes the response for WANPPPConnection1.GetPPPCompressionProtocol API
type WANPPPConnection1GetPPPCompressionProtocolResponse struct {
	NewPPPCompressionProtocol soap.String
}

func (client *WANPPPConnection1) GetPPPCompressionProtocol(request WANPPPConnection1GetPPPCompressionProtocolRequest) (response *WANPPPConnection1GetPPPCompressionProtocolResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPCompressionProtocol", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetPPPAuthenticationProtocolRequest describes the request for WANPPPConnection1.GetPPPAuthenticationProtocol API
type WANPPPConnection1GetPPPAuthenticationProtocolRequest struct {
}

// WANPPPConnection1GetPPPAuthenticationProtocolResponse describes the response for WANPPPConnection1.GetPPPAuthenticationProtocol API
type WANPPPConnection1GetPPPAuthenticationProtocolResponse struct {
	NewPPPAuthenticationProtocol soap.String
}

func (client *WANPPPConnection1) GetPPPAuthenticationProtocol(request WANPPPConnection1GetPPPAuthenticationProtocolRequest) (response *WANPPPConnection1GetPPPAuthenticationProtocolResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPAuthenticationProtocol", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetUserNameRequest describes the request for WANPPPConnection1.GetUserName API
type WANPPPConnection1GetUserNameRequest struct {
}

// WANPPPConnection1GetUserNameResponse describes the response for WANPPPConnection1.GetUserName API
type WANPPPConnection1GetUserNameResponse struct {
	NewUserName soap.String
}

func (client *WANPPPConnection1) GetUserName(request WANPPPConnection1GetUserNameRequest) (response *WANPPPConnection1GetUserNameResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetUserName", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetPasswordRequest describes the request for WANPPPConnection1.GetPassword API
type WANPPPConnection1GetPasswordRequest struct {
}

// WANPPPConnection1GetPasswordResponse describes the response for WANPPPConnection1.GetPassword API
type WANPPPConnection1GetPasswordResponse struct {
	NewPassword soap.String
}

func (client *WANPPPConnection1) GetPassword(request WANPPPConnection1GetPasswordRequest) (response *WANPPPConnection1GetPasswordResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPassword", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetAutoDisconnectTimeRequest describes the request for WANPPPConnection1.GetAutoDisconnectTime API
type WANPPPConnection1GetAutoDisconnectTimeRequest struct {
}

// WANPPPConnection1GetAutoDisconnectTimeResponse describes the response for WANPPPConnection1.GetAutoDisconnectTime API
type WANPPPConnection1GetAutoDisconnectTimeResponse struct {
	NewAutoDisconnectTime soap.Ui4
}

func (client *WANPPPConnection1) GetAutoDisconnectTime(request WANPPPConnection1GetAutoDisconnectTimeRequest) (response *WANPPPConnection1GetAutoDisconnectTimeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetAutoDisconnectTime", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetIdleDisconnectTimeRequest describes the request for WANPPPConnection1.GetIdleDisconnectTime API
type WANPPPConnection1GetIdleDisconnectTimeRequest struct {
}

// WANPPPConnection1GetIdleDisconnectTimeResponse describes the response for WANPPPConnection1.GetIdleDisconnectTime API
type WANPPPConnection1GetIdleDisconnectTimeResponse struct {
	NewIdleDisconnectTime soap.Ui4
}

func (client *WANPPPConnection1) GetIdleDisconnectTime(request WANPPPConnection1GetIdleDisconnectTimeRequest) (response *WANPPPConnection1GetIdleDisconnectTimeResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetIdleDisconnectTime", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetWarnDisconnectDelayRequest describes the request for WANPPPConnection1.GetWarnDisconnectDelay API
type WANPPPConnection1GetWarnDisconnectDelayRequest struct {
}

// WANPPPConnection1GetWarnDisconnectDelayResponse describes the response for WANPPPConnection1.GetWarnDisconnectDelay API
type WANPPPConnection1GetWarnDisconnectDelayResponse struct {
	NewWarnDisconnectDelay soap.Ui4
}

func (client *WANPPPConnection1) GetWarnDisconnectDelay(request WANPPPConnection1GetWarnDisconnectDelayRequest) (response *WANPPPConnection1GetWarnDisconnectDelayResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetWarnDisconnectDelay", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetNATRSIPStatusRequest describes the request for WANPPPConnection1.GetNATRSIPStatus API
type WANPPPConnection1GetNATRSIPStatusRequest struct {
}

// WANPPPConnection1GetNATRSIPStatusResponse describes the response for WANPPPConnection1.GetNATRSIPStatus API
type WANPPPConnection1GetNATRSIPStatusResponse struct {
	NewRSIPAvailable soap.Bool
	NewNATEnabled    soap.Bool
}

func (client *WANPPPConnection1) GetNATRSIPStatus(request WANPPPConnection1GetNATRSIPStatusRequest) (response *WANPPPConnection1GetNATRSIPStatusResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetNATRSIPStatus", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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
func (client *WANPPPConnection1) GetGenericPortMappingEntry(request WANPPPConnection1GetGenericPortMappingEntryRequest) (response *WANPPPConnection1GetGenericPortMappingEntryResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetGenericPortMappingEntry", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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
func (client *WANPPPConnection1) GetSpecificPortMappingEntry(request WANPPPConnection1GetSpecificPortMappingEntryRequest) (response *WANPPPConnection1GetSpecificPortMappingEntryResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetSpecificPortMappingEntry", &request, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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
func (client *WANPPPConnection1) AddPortMapping(request WANPPPConnection1AddPortMappingRequest) (response *WANPPPConnection1AddPortMappingResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "AddPortMapping", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
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
func (client *WANPPPConnection1) DeletePortMapping(request WANPPPConnection1DeletePortMappingRequest) (response *WANPPPConnection1DeletePortMappingResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "DeletePortMapping", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}

// WANPPPConnection1GetExternalIPAddressRequest describes the request for WANPPPConnection1.GetExternalIPAddress API
type WANPPPConnection1GetExternalIPAddressRequest struct {
}

// WANPPPConnection1GetExternalIPAddressResponse describes the response for WANPPPConnection1.GetExternalIPAddress API
type WANPPPConnection1GetExternalIPAddressResponse struct {
	NewExternalIPAddress soap.String
}

func (client *WANPPPConnection1) GetExternalIPAddress(request WANPPPConnection1GetExternalIPAddressRequest) (response *WANPPPConnection1GetExternalIPAddressResponse, err error) {
	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetExternalIPAddress", nil, response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}

	return response, nil
}
