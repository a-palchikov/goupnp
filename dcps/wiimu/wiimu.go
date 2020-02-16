// Client for UPnP Device Control Protocol Wiimu Link Player.
//
//
// Typically, use one of the New* functions to create clients for services.
package wiimu

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
	URN_MediaRenderer_1 = "urn:schemas-upnp-org:device:MediaRenderer:1"
)

// Service URNs:
const (
	URN_QPlay_1             = "urn:schemas-tencent-com:service:QPlay:1"
	URN_AVTransport_1       = "urn:schemas-upnp-org:service:AVTransport:1"
	URN_ConnectionManager_1 = "urn:schemas-upnp-org:service:ConnectionManager:1"
	URN_RenderingControl_1  = "urn:schemas-upnp-org:service:RenderingControl:1"
	URN_PlayQueue_1         = "urn:schemas-wiimu-com:service:PlayQueue:1"
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

// AVTransport1GetMuteRequest describes the request for AVTransport1.GetMute API
type AVTransport1GetMuteRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master, Single
	Channel soap.String
}

// AVTransport1GetMuteResponse describes the response for AVTransport1.GetMute API
type AVTransport1GetMuteResponse struct {
	CurrentMute soap.Bool
}

//
// Arguments:
//
//  AVTransport1GetMuteRequest
func (client *AVTransport1) GetMute(request AVTransport1GetMuteRequest) (*AVTransport1GetMuteResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1GetMuteResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "GetMute", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1GetVolumeRequest describes the request for AVTransport1.GetVolume API
type AVTransport1GetVolumeRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master, Single
	Channel soap.String
}

// AVTransport1GetVolumeResponse describes the response for AVTransport1.GetVolume API
type AVTransport1GetVolumeResponse struct {
	// CurrentVolume: allowed value range: minimum=0, maximum=100, step=1
	CurrentVolume soap.Ui2
}

//
// Arguments:
//
//  AVTransport1GetVolumeRequest
//
// Return value:
//
//  AVTransport1GetVolumeResponse
func (client *AVTransport1) GetVolume(request AVTransport1GetVolumeRequest) (*AVTransport1GetVolumeResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1GetVolumeResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "GetVolume", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1GetChannelRequest describes the request for AVTransport1.GetChannel API
type AVTransport1GetChannelRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master, Single
	Channel soap.String
}

// AVTransport1GetChannelResponse describes the response for AVTransport1.GetChannel API
type AVTransport1GetChannelResponse struct {
	CurrentChannel soap.Ui2
}

//
// Arguments:
//
//  AVTransport1GetChannelRequest
func (client *AVTransport1) GetChannel(request AVTransport1GetChannelRequest) (*AVTransport1GetChannelResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1GetChannelResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "GetChannel", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1GetEqualizerRequest describes the request for AVTransport1.GetEqualizer API
type AVTransport1GetEqualizerRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master, Single
	Channel soap.String
}

// AVTransport1GetEqualizerResponse describes the response for AVTransport1.GetEqualizer API
type AVTransport1GetEqualizerResponse struct {
	CurrentEqualizer soap.Ui2
}

//
// Arguments:
//
//  AVTransport1GetEqualizerRequest
func (client *AVTransport1) GetEqualizer(request AVTransport1GetEqualizerRequest) (*AVTransport1GetEqualizerResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1GetEqualizerResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "GetEqualizer", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1ListPresetsRequest describes the request for AVTransport1.ListPresets API
type AVTransport1ListPresetsRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1ListPresetsResponse describes the response for AVTransport1.ListPresets API
type AVTransport1ListPresetsResponse struct {
	CurrentPresetNameList soap.String
}

func (client *AVTransport1) ListPresets(request AVTransport1ListPresetsRequest) (*AVTransport1ListPresetsResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1ListPresetsResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "ListPresets", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1SelectPresetRequest describes the request for AVTransport1.SelectPreset API
type AVTransport1SelectPresetRequest struct {
	InstanceID soap.Ui4
	// PresetName: allowed values: FactoryDefaults, InstallationDefaults
	PresetName soap.String
}

// AVTransport1SelectPresetResponse describes the response for AVTransport1.SelectPreset API
type AVTransport1SelectPresetResponse struct {
}

//
// Arguments:
//
//  AVTransport1SelectPresetRequest
func (client *AVTransport1) SelectPreset(request AVTransport1SelectPresetRequest) (*AVTransport1SelectPresetResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1SelectPresetResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "SelectPreset", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1SetMuteRequest describes the request for AVTransport1.SetMute API
type AVTransport1SetMuteRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master, Single
	Channel     soap.String
	DesiredMute soap.Bool
}

// AVTransport1SetMuteResponse describes the response for AVTransport1.SetMute API
type AVTransport1SetMuteResponse struct {
}

//
// Arguments:
//
//  AVTransport1SetMuteRequest
func (client *AVTransport1) SetMute(request AVTransport1SetMuteRequest) (*AVTransport1SetMuteResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1SetMuteResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "SetMute", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1SetVolumeRequest describes the request for AVTransport1.SetVolume API
type AVTransport1SetVolumeRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master, Single
	Channel soap.String
	// DesiredVolume: allowed value range: minimum=0, maximum=100, step=1
	DesiredVolume soap.Ui2
}

// AVTransport1SetVolumeResponse describes the response for AVTransport1.SetVolume API
type AVTransport1SetVolumeResponse struct {
}

//
// Arguments:
//
//  AVTransport1SetVolumeRequest
func (client *AVTransport1) SetVolume(request AVTransport1SetVolumeRequest) (*AVTransport1SetVolumeResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1SetVolumeResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "SetVolume", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1SetChannelRequest describes the request for AVTransport1.SetChannel API
type AVTransport1SetChannelRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master, Single
	Channel        soap.String
	DesiredChannel soap.Ui2
}

// AVTransport1SetChannelResponse describes the response for AVTransport1.SetChannel API
type AVTransport1SetChannelResponse struct {
}

//
// Arguments:
//
//  AVTransport1SetChannelRequest
func (client *AVTransport1) SetChannel(request AVTransport1SetChannelRequest) (*AVTransport1SetChannelResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1SetChannelResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "SetChannel", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1SetEqualizerRequest describes the request for AVTransport1.SetEqualizer API
type AVTransport1SetEqualizerRequest struct {
	InstanceID soap.Ui4
	// Channel: allowed values: Master, Single
	Channel          soap.String
	DesiredEqualizer soap.Ui2
}

// AVTransport1SetEqualizerResponse describes the response for AVTransport1.SetEqualizer API
type AVTransport1SetEqualizerResponse struct {
}

//
// Arguments:
//
//  AVTransport1SetEqualizerRequest
func (client *AVTransport1) SetEqualizer(request AVTransport1SetEqualizerRequest) (*AVTransport1SetEqualizerResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1SetEqualizerResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "SetEqualizer", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1GetSimpleDeviceInfoRequest describes the request for AVTransport1.GetSimpleDeviceInfo API
type AVTransport1GetSimpleDeviceInfoRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1GetSimpleDeviceInfoResponse describes the response for AVTransport1.GetSimpleDeviceInfo API
type AVTransport1GetSimpleDeviceInfoResponse struct {
	MultiType soap.Ui2
	SlaveMask soap.Ui2
	Name      soap.String
	// CurrentVolume: allowed value range: minimum=0, maximum=100, step=1
	CurrentVolume  soap.Ui2
	CurrentChannel soap.Ui2
	SlaveList      soap.String
}

//
// Return value:
//
//  AVTransport1GetSimpleDeviceInfoResponse
func (client *AVTransport1) GetSimpleDeviceInfo(request AVTransport1GetSimpleDeviceInfoRequest) (*AVTransport1GetSimpleDeviceInfoResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1GetSimpleDeviceInfoResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "GetSimpleDeviceInfo", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1GetControlDeviceInfoRequest describes the request for AVTransport1.GetControlDeviceInfo API
type AVTransport1GetControlDeviceInfoRequest struct {
	InstanceID soap.Ui4
}

// AVTransport1GetControlDeviceInfoResponse describes the response for AVTransport1.GetControlDeviceInfo API
type AVTransport1GetControlDeviceInfoResponse struct {
	MultiType soap.Ui2
	Router    soap.String
	Ssid      soap.String
	SlaveMask soap.Ui2
	// CurrentVolume: allowed value range: minimum=0, maximum=100, step=1
	CurrentVolume  soap.Ui2
	CurrentMute    soap.Bool
	CurrentChannel soap.Ui2
	SlaveList      soap.String
	Status         soap.String
}

//
// Return value:
//
//  AVTransport1GetControlDeviceInfoResponse
func (client *AVTransport1) GetControlDeviceInfo(request AVTransport1GetControlDeviceInfoRequest) (*AVTransport1GetControlDeviceInfoResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1GetControlDeviceInfoResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "GetControlDeviceInfo", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1MultiPlaySlaveMaskRequest describes the request for AVTransport1.MultiPlaySlaveMask API
type AVTransport1MultiPlaySlaveMaskRequest struct {
	InstanceID soap.Ui4
	SlaveMask  soap.Ui2
}

// AVTransport1MultiPlaySlaveMaskResponse describes the response for AVTransport1.MultiPlaySlaveMask API
type AVTransport1MultiPlaySlaveMaskResponse struct {
}

func (client *AVTransport1) MultiPlaySlaveMask(request AVTransport1MultiPlaySlaveMaskRequest) (*AVTransport1MultiPlaySlaveMaskResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1MultiPlaySlaveMaskResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "MultiPlaySlaveMask", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1SetAlarmQueueRequest describes the request for AVTransport1.SetAlarmQueue API
type AVTransport1SetAlarmQueueRequest struct {
	AlarmContext soap.String
}

// AVTransport1SetAlarmQueueResponse describes the response for AVTransport1.SetAlarmQueue API
type AVTransport1SetAlarmQueueResponse struct {
}

func (client *AVTransport1) SetAlarmQueue(request AVTransport1SetAlarmQueueRequest) (*AVTransport1SetAlarmQueueResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1SetAlarmQueueResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "SetAlarmQueue", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1GetAlarmQueueRequest describes the request for AVTransport1.GetAlarmQueue API
type AVTransport1GetAlarmQueueRequest struct {
	AlarmName soap.String
}

// AVTransport1GetAlarmQueueResponse describes the response for AVTransport1.GetAlarmQueue API
type AVTransport1GetAlarmQueueResponse struct {
	AlarmContext soap.String
}

func (client *AVTransport1) GetAlarmQueue(request AVTransport1GetAlarmQueueRequest) (*AVTransport1GetAlarmQueueResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1GetAlarmQueueResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "GetAlarmQueue", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1DeleteAlarmQueueRequest describes the request for AVTransport1.DeleteAlarmQueue API
type AVTransport1DeleteAlarmQueueRequest struct {
	AlarmName soap.String
}

// AVTransport1DeleteAlarmQueueResponse describes the response for AVTransport1.DeleteAlarmQueue API
type AVTransport1DeleteAlarmQueueResponse struct {
}

func (client *AVTransport1) DeleteAlarmQueue(request AVTransport1DeleteAlarmQueueRequest) (*AVTransport1DeleteAlarmQueueResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1DeleteAlarmQueueResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "DeleteAlarmQueue", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// AVTransport1SetDeviceNameRequest describes the request for AVTransport1.SetDeviceName API
type AVTransport1SetDeviceNameRequest struct {
	Name soap.String
}

// AVTransport1SetDeviceNameResponse describes the response for AVTransport1.SetDeviceName API
type AVTransport1SetDeviceNameResponse struct {
}

func (client *AVTransport1) SetDeviceName(request AVTransport1SetDeviceNameRequest) (*AVTransport1SetDeviceNameResponse, error) {
	// Perform the SOAP call.
	var response AVTransport1SetDeviceNameResponse
	if err := client.SOAPClient.PerformAction(URN_AVTransport_1, "SetDeviceName", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
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

// ConnectionManager1GetCurrentConnectionIDsResponse describes the response for ConnectionManager1.GetCurrentConnectionIDs API
type ConnectionManager1GetCurrentConnectionIDsResponse struct {
	ConnectionIDs soap.String
}

func (client *ConnectionManager1) GetCurrentConnectionIDs() (*ConnectionManager1GetCurrentConnectionIDsResponse, error) {
	// Perform the SOAP call.
	var response ConnectionManager1GetCurrentConnectionIDsResponse
	if err := client.SOAPClient.PerformAction(URN_ConnectionManager_1, "GetCurrentConnectionIDs", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
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
func (client *ConnectionManager1) GetCurrentConnectionInfo(request ConnectionManager1GetCurrentConnectionInfoRequest) (*ConnectionManager1GetCurrentConnectionInfoResponse, error) {
	// Perform the SOAP call.
	var response ConnectionManager1GetCurrentConnectionInfoResponse
	if err := client.SOAPClient.PerformAction(URN_ConnectionManager_1, "GetCurrentConnectionInfo", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// ConnectionManager1GetProtocolInfoResponse describes the response for ConnectionManager1.GetProtocolInfo API
type ConnectionManager1GetProtocolInfoResponse struct {
	Source soap.String
	Sink   soap.String
}

func (client *ConnectionManager1) GetProtocolInfo() (*ConnectionManager1GetProtocolInfoResponse, error) {
	// Perform the SOAP call.
	var response ConnectionManager1GetProtocolInfoResponse
	if err := client.SOAPClient.PerformAction(URN_ConnectionManager_1, "GetProtocolInfo", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1 is a client for UPnP SOAP service with URN "urn:schemas-wiimu-com:service:PlayQueue:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type PlayQueue1 struct {
	goupnp.ServiceClient
}

// NewPlayQueue1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewPlayQueue1Clients() (clients []*PlayQueue1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_PlayQueue_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newPlayQueue1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewPlayQueue1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewPlayQueue1ClientsByURL(loc *url.URL) ([]*PlayQueue1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_PlayQueue_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newPlayQueue1ClientsFromGenericClients(genericClients), nil
}

// NewPlayQueue1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewPlayQueue1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*PlayQueue1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_PlayQueue_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newPlayQueue1ClientsFromGenericClients(genericClients), nil
}

func newPlayQueue1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*PlayQueue1 {
	clients := make([]*PlayQueue1, len(genericClients))
	for i := range genericClients {
		clients[i] = &PlayQueue1{genericClients[i]}
	}
	return clients
}

// PlayQueue1CreateQueueRequest describes the request for PlayQueue1.CreateQueue API
type PlayQueue1CreateQueueRequest struct {
	QueueContext soap.String
}

// PlayQueue1CreateQueueResponse describes the response for PlayQueue1.CreateQueue API
type PlayQueue1CreateQueueResponse struct {
}

func (client *PlayQueue1) CreateQueue(request PlayQueue1CreateQueueRequest) (*PlayQueue1CreateQueueResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1CreateQueueResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "CreateQueue", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1ReplaceQueueRequest describes the request for PlayQueue1.ReplaceQueue API
type PlayQueue1ReplaceQueueRequest struct {
	QueueContext soap.String
}

// PlayQueue1ReplaceQueueResponse describes the response for PlayQueue1.ReplaceQueue API
type PlayQueue1ReplaceQueueResponse struct {
}

func (client *PlayQueue1) ReplaceQueue(request PlayQueue1ReplaceQueueRequest) (*PlayQueue1ReplaceQueueResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1ReplaceQueueResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "ReplaceQueue", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1DeleteQueueRequest describes the request for PlayQueue1.DeleteQueue API
type PlayQueue1DeleteQueueRequest struct {
	QueueName soap.String
}

// PlayQueue1DeleteQueueResponse describes the response for PlayQueue1.DeleteQueue API
type PlayQueue1DeleteQueueResponse struct {
}

func (client *PlayQueue1) DeleteQueue(request PlayQueue1DeleteQueueRequest) (*PlayQueue1DeleteQueueResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1DeleteQueueResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "DeleteQueue", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1BackUpQueueRequest describes the request for PlayQueue1.BackUpQueue API
type PlayQueue1BackUpQueueRequest struct {
	QueueContext soap.String
}

// PlayQueue1BackUpQueueResponse describes the response for PlayQueue1.BackUpQueue API
type PlayQueue1BackUpQueueResponse struct {
}

func (client *PlayQueue1) BackUpQueue(request PlayQueue1BackUpQueueRequest) (*PlayQueue1BackUpQueueResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1BackUpQueueResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "BackUpQueue", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1AppendQueueRequest describes the request for PlayQueue1.AppendQueue API
type PlayQueue1AppendQueueRequest struct {
	QueueContext soap.String
}

// PlayQueue1AppendQueueResponse describes the response for PlayQueue1.AppendQueue API
type PlayQueue1AppendQueueResponse struct {
}

func (client *PlayQueue1) AppendQueue(request PlayQueue1AppendQueueRequest) (*PlayQueue1AppendQueueResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1AppendQueueResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "AppendQueue", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1BrowseQueueRequest describes the request for PlayQueue1.BrowseQueue API
type PlayQueue1BrowseQueueRequest struct {
	QueueName soap.String
}

// PlayQueue1BrowseQueueResponse describes the response for PlayQueue1.BrowseQueue API
type PlayQueue1BrowseQueueResponse struct {
	QueueContext soap.String
}

func (client *PlayQueue1) BrowseQueue(request PlayQueue1BrowseQueueRequest) (*PlayQueue1BrowseQueueResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1BrowseQueueResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "BrowseQueue", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1SetQueueLoopModeRequest describes the request for PlayQueue1.SetQueueLoopMode API
type PlayQueue1SetQueueLoopModeRequest struct {
	LoopMode soap.Ui4
}

// PlayQueue1SetQueueLoopModeResponse describes the response for PlayQueue1.SetQueueLoopMode API
type PlayQueue1SetQueueLoopModeResponse struct {
}

func (client *PlayQueue1) SetQueueLoopMode(request PlayQueue1SetQueueLoopModeRequest) (*PlayQueue1SetQueueLoopModeResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1SetQueueLoopModeResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "SetQueueLoopMode", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1GetQueueLoopModeResponse describes the response for PlayQueue1.GetQueueLoopMode API
type PlayQueue1GetQueueLoopModeResponse struct {
	LoopMode soap.Ui4
}

func (client *PlayQueue1) GetQueueLoopMode() (*PlayQueue1GetQueueLoopModeResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1GetQueueLoopModeResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "GetQueueLoopMode", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1SetQueuePolicyRequest describes the request for PlayQueue1.SetQueuePolicy API
type PlayQueue1SetQueuePolicyRequest struct {
	QueueName soap.String
}

// PlayQueue1SetQueuePolicyResponse describes the response for PlayQueue1.SetQueuePolicy API
type PlayQueue1SetQueuePolicyResponse struct {
}

func (client *PlayQueue1) SetQueuePolicy(request PlayQueue1SetQueuePolicyRequest) (*PlayQueue1SetQueuePolicyResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1SetQueuePolicyResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "SetQueuePolicy", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1PlayQueueWithIndexRequest describes the request for PlayQueue1.PlayQueueWithIndex API
type PlayQueue1PlayQueueWithIndexRequest struct {
	QueueName soap.String
	Index     soap.Ui4
}

// PlayQueue1PlayQueueWithIndexResponse describes the response for PlayQueue1.PlayQueueWithIndex API
type PlayQueue1PlayQueueWithIndexResponse struct {
}

func (client *PlayQueue1) PlayQueueWithIndex(request PlayQueue1PlayQueueWithIndexRequest) (*PlayQueue1PlayQueueWithIndexResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1PlayQueueWithIndexResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "PlayQueueWithIndex", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1AppendTracksInQueueRequest describes the request for PlayQueue1.AppendTracksInQueue API
type PlayQueue1AppendTracksInQueueRequest struct {
	QueueContext soap.String
}

// PlayQueue1AppendTracksInQueueResponse describes the response for PlayQueue1.AppendTracksInQueue API
type PlayQueue1AppendTracksInQueueResponse struct {
}

func (client *PlayQueue1) AppendTracksInQueue(request PlayQueue1AppendTracksInQueueRequest) (*PlayQueue1AppendTracksInQueueResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1AppendTracksInQueueResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "AppendTracksInQueue", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1RemoveTracksInQueueRequest describes the request for PlayQueue1.RemoveTracksInQueue API
type PlayQueue1RemoveTracksInQueueRequest struct {
	QueueName soap.String
	RangStart soap.Ui4
	RangEnd   soap.Ui4
}

// PlayQueue1RemoveTracksInQueueResponse describes the response for PlayQueue1.RemoveTracksInQueue API
type PlayQueue1RemoveTracksInQueueResponse struct {
}

func (client *PlayQueue1) RemoveTracksInQueue(request PlayQueue1RemoveTracksInQueueRequest) (*PlayQueue1RemoveTracksInQueueResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1RemoveTracksInQueueResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "RemoveTracksInQueue", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1AppendTracksInQueueExRequest describes the request for PlayQueue1.AppendTracksInQueueEx API
type PlayQueue1AppendTracksInQueueExRequest struct {
	QueueContext soap.String
	Direction    soap.Ui4
	StartIndex   soap.Ui4
}

// PlayQueue1AppendTracksInQueueExResponse describes the response for PlayQueue1.AppendTracksInQueueEx API
type PlayQueue1AppendTracksInQueueExResponse struct {
}

func (client *PlayQueue1) AppendTracksInQueueEx(request PlayQueue1AppendTracksInQueueExRequest) (*PlayQueue1AppendTracksInQueueExResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1AppendTracksInQueueExResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "AppendTracksInQueueEx", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1SetKeyMappingRequest describes the request for PlayQueue1.SetKeyMapping API
type PlayQueue1SetKeyMappingRequest struct {
	QueueContext soap.String
}

// PlayQueue1SetKeyMappingResponse describes the response for PlayQueue1.SetKeyMapping API
type PlayQueue1SetKeyMappingResponse struct {
}

func (client *PlayQueue1) SetKeyMapping(request PlayQueue1SetKeyMappingRequest) (*PlayQueue1SetKeyMappingResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1SetKeyMappingResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "SetKeyMapping", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1GetKeyMappingResponse describes the response for PlayQueue1.GetKeyMapping API
type PlayQueue1GetKeyMappingResponse struct {
	QueueContext soap.String
}

func (client *PlayQueue1) GetKeyMapping() (*PlayQueue1GetKeyMappingResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1GetKeyMappingResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "GetKeyMapping", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1GetQueueOnlineRequest describes the request for PlayQueue1.GetQueueOnline API
type PlayQueue1GetQueueOnlineRequest struct {
	QueueName       soap.String
	QueueID         soap.String
	QueueType       soap.String
	Queuelimit      soap.Ui4
	QueueAutoInsert soap.String
}

// PlayQueue1GetQueueOnlineResponse describes the response for PlayQueue1.GetQueueOnline API
type PlayQueue1GetQueueOnlineResponse struct {
	QueueContext soap.String
}

func (client *PlayQueue1) GetQueueOnline(request PlayQueue1GetQueueOnlineRequest) (*PlayQueue1GetQueueOnlineResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1GetQueueOnlineResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "GetQueueOnline", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1SearchQueueOnlineRequest describes the request for PlayQueue1.SearchQueueOnline API
type PlayQueue1SearchQueueOnlineRequest struct {
	QueueName  soap.String
	SearchKey  soap.String
	Queuelimit soap.Ui4
}

// PlayQueue1SearchQueueOnlineResponse describes the response for PlayQueue1.SearchQueueOnline API
type PlayQueue1SearchQueueOnlineResponse struct {
	QueueContext soap.String
}

func (client *PlayQueue1) SearchQueueOnline(request PlayQueue1SearchQueueOnlineRequest) (*PlayQueue1SearchQueueOnlineResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1SearchQueueOnlineResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "SearchQueueOnline", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1SetQueueRecordRequest describes the request for PlayQueue1.SetQueueRecord API
type PlayQueue1SetQueueRecordRequest struct {
	QueueName soap.String
	QueueID   soap.String
	Action    soap.String
}

// PlayQueue1SetQueueRecordResponse describes the response for PlayQueue1.SetQueueRecord API
type PlayQueue1SetQueueRecordResponse struct {
}

func (client *PlayQueue1) SetQueueRecord(request PlayQueue1SetQueueRecordRequest) (*PlayQueue1SetQueueRecordResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1SetQueueRecordResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "SetQueueRecord", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1SetSongsRecordRequest describes the request for PlayQueue1.SetSongsRecord API
type PlayQueue1SetSongsRecordRequest struct {
	QueueName soap.String
	SongID    soap.String
	Action    soap.String
}

// PlayQueue1SetSongsRecordResponse describes the response for PlayQueue1.SetSongsRecord API
type PlayQueue1SetSongsRecordResponse struct {
}

func (client *PlayQueue1) SetSongsRecord(request PlayQueue1SetSongsRecordRequest) (*PlayQueue1SetSongsRecordResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1SetSongsRecordResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "SetSongsRecord", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1UserRegisterRequest describes the request for PlayQueue1.UserRegister API
type PlayQueue1UserRegisterRequest struct {
	QueueName soap.String
	UserName  soap.String
	PassWord  soap.String
}

// PlayQueue1UserRegisterResponse describes the response for PlayQueue1.UserRegister API
type PlayQueue1UserRegisterResponse struct {
	Result soap.String
}

func (client *PlayQueue1) UserRegister(request PlayQueue1UserRegisterRequest) (*PlayQueue1UserRegisterResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1UserRegisterResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "UserRegister", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1UserLoginRequest describes the request for PlayQueue1.UserLogin API
type PlayQueue1UserLoginRequest struct {
	AccountSource soap.String
	UserName      soap.String
	PassWord      soap.String
	SavePass      soap.Ui4
	Code          soap.Ui4
	Proxy         soap.String
}

// PlayQueue1UserLoginResponse describes the response for PlayQueue1.UserLogin API
type PlayQueue1UserLoginResponse struct {
	Result soap.String
}

func (client *PlayQueue1) UserLogin(request PlayQueue1UserLoginRequest) (*PlayQueue1UserLoginResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1UserLoginResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "UserLogin", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1UserLogoutRequest describes the request for PlayQueue1.UserLogout API
type PlayQueue1UserLogoutRequest struct {
	AccountSource soap.String
}

// PlayQueue1UserLogoutResponse describes the response for PlayQueue1.UserLogout API
type PlayQueue1UserLogoutResponse struct {
	Result soap.String
}

func (client *PlayQueue1) UserLogout(request PlayQueue1UserLogoutRequest) (*PlayQueue1UserLogoutResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1UserLogoutResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "UserLogout", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1GetUserInfoRequest describes the request for PlayQueue1.GetUserInfo API
type PlayQueue1GetUserInfoRequest struct {
	AccountSource soap.String
}

// PlayQueue1GetUserInfoResponse describes the response for PlayQueue1.GetUserInfo API
type PlayQueue1GetUserInfoResponse struct {
	Result soap.String
}

func (client *PlayQueue1) GetUserInfo(request PlayQueue1GetUserInfoRequest) (*PlayQueue1GetUserInfoResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1GetUserInfoResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "GetUserInfo", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1GetUserAccountHistoryRequest describes the request for PlayQueue1.GetUserAccountHistory API
type PlayQueue1GetUserAccountHistoryRequest struct {
	AccountSource soap.String
	Number        soap.Ui4
}

// PlayQueue1GetUserAccountHistoryResponse describes the response for PlayQueue1.GetUserAccountHistory API
type PlayQueue1GetUserAccountHistoryResponse struct {
	Result soap.String
}

func (client *PlayQueue1) GetUserAccountHistory(request PlayQueue1GetUserAccountHistoryRequest) (*PlayQueue1GetUserAccountHistoryResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1GetUserAccountHistoryResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "GetUserAccountHistory", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1SetUserFavoritesRequest describes the request for PlayQueue1.SetUserFavorites API
type PlayQueue1SetUserFavoritesRequest struct {
	AccountSource soap.String
	Action        soap.String
	MediaType     soap.String
	MediaID       soap.String
	MediaContext  soap.String
}

// PlayQueue1SetUserFavoritesResponse describes the response for PlayQueue1.SetUserFavorites API
type PlayQueue1SetUserFavoritesResponse struct {
	Result soap.String
}

func (client *PlayQueue1) SetUserFavorites(request PlayQueue1SetUserFavoritesRequest) (*PlayQueue1SetUserFavoritesResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1SetUserFavoritesResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "SetUserFavorites", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1GetUserFavoritesRequest describes the request for PlayQueue1.GetUserFavorites API
type PlayQueue1GetUserFavoritesRequest struct {
	AccountSource soap.String
	MediaType     soap.String
	Filter        soap.String
}

// PlayQueue1GetUserFavoritesResponse describes the response for PlayQueue1.GetUserFavorites API
type PlayQueue1GetUserFavoritesResponse struct {
	Result soap.String
}

func (client *PlayQueue1) GetUserFavorites(request PlayQueue1GetUserFavoritesRequest) (*PlayQueue1GetUserFavoritesResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1GetUserFavoritesResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "GetUserFavorites", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1GetQueueIndexRequest describes the request for PlayQueue1.GetQueueIndex API
type PlayQueue1GetQueueIndexRequest struct {
	QueueName soap.String
}

// PlayQueue1GetQueueIndexResponse describes the response for PlayQueue1.GetQueueIndex API
type PlayQueue1GetQueueIndexResponse struct {
	CurrentIndex soap.Ui4
	CurrentPage  soap.Ui4
}

func (client *PlayQueue1) GetQueueIndex(request PlayQueue1GetQueueIndexRequest) (*PlayQueue1GetQueueIndexResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1GetQueueIndexResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "GetQueueIndex", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1SetSpotifyPresetRequest describes the request for PlayQueue1.SetSpotifyPreset API
type PlayQueue1SetSpotifyPresetRequest struct {
	KeyIndex soap.Ui4
}

// PlayQueue1SetSpotifyPresetResponse describes the response for PlayQueue1.SetSpotifyPreset API
type PlayQueue1SetSpotifyPresetResponse struct {
	Result soap.String
}

func (client *PlayQueue1) SetSpotifyPreset(request PlayQueue1SetSpotifyPresetRequest) (*PlayQueue1SetSpotifyPresetResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1SetSpotifyPresetResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "SetSpotifyPreset", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1BrowseSpotifyListResponse describes the response for PlayQueue1.BrowseSpotifyList API
type PlayQueue1BrowseSpotifyListResponse struct {
	Result soap.String
}

func (client *PlayQueue1) BrowseSpotifyList() (*PlayQueue1BrowseSpotifyListResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1BrowseSpotifyListResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "BrowseSpotifyList", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1AddSpotifyListResponse describes the response for PlayQueue1.AddSpotifyList API
type PlayQueue1AddSpotifyListResponse struct {
	Result soap.String
}

func (client *PlayQueue1) AddSpotifyList() (*PlayQueue1AddSpotifyListResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1AddSpotifyListResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "AddSpotifyList", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1DeleteSpotifyListRequest describes the request for PlayQueue1.DeleteSpotifyList API
type PlayQueue1DeleteSpotifyListRequest struct {
	QueueName soap.String
}

// PlayQueue1DeleteSpotifyListResponse describes the response for PlayQueue1.DeleteSpotifyList API
type PlayQueue1DeleteSpotifyListResponse struct {
	Result soap.String
}

func (client *PlayQueue1) DeleteSpotifyList(request PlayQueue1DeleteSpotifyListRequest) (*PlayQueue1DeleteSpotifyListResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1DeleteSpotifyListResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "DeleteSpotifyList", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1GetSpotifyPresetMaptoKeyResponse describes the response for PlayQueue1.GetSpotifyPresetMaptoKey API
type PlayQueue1GetSpotifyPresetMaptoKeyResponse struct {
	Result soap.String
}

func (client *PlayQueue1) GetSpotifyPresetMaptoKey() (*PlayQueue1GetSpotifyPresetMaptoKeyResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1GetSpotifyPresetMaptoKeyResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "GetSpotifyPresetMaptoKey", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// PlayQueue1DeleteActionQueueRequest describes the request for PlayQueue1.DeleteActionQueue API
type PlayQueue1DeleteActionQueueRequest struct {
	PressType soap.Ui4
}

// PlayQueue1DeleteActionQueueResponse describes the response for PlayQueue1.DeleteActionQueue API
type PlayQueue1DeleteActionQueueResponse struct {
}

func (client *PlayQueue1) DeleteActionQueue(request PlayQueue1DeleteActionQueueRequest) (*PlayQueue1DeleteActionQueueResponse, error) {
	// Perform the SOAP call.
	var response PlayQueue1DeleteActionQueueResponse
	if err := client.SOAPClient.PerformAction(URN_PlayQueue_1, "DeleteActionQueue", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// QPlay1 is a client for UPnP SOAP service with URN "urn:schemas-tencent-com:service:QPlay:1".
// See goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type QPlay1 struct {
	goupnp.ServiceClient
}

// NewQPlay1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewQPlay1Clients() (clients []*QPlay1, errs []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errs, err = goupnp.NewServiceClients(URN_QPlay_1); err != nil {
		return nil, errs, errors.Wrap(err, "creating service clients")
	}
	clients = newQPlay1ClientsFromGenericClients(genericClients)
	return clients, errs, nil
}

// NewQPlay1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewQPlay1ClientsByURL(loc *url.URL) ([]*QPlay1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_QPlay_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newQPlay1ClientsFromGenericClients(genericClients), nil
}

// NewQPlay1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewQPlay1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*QPlay1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_QPlay_1)
	if err != nil {
		return nil, errors.Wrap(err, "creating service clients")
	}
	return newQPlay1ClientsFromGenericClients(genericClients), nil
}

func newQPlay1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*QPlay1 {
	clients := make([]*QPlay1, len(genericClients))
	for i := range genericClients {
		clients[i] = &QPlay1{genericClients[i]}
	}
	return clients
}

// QPlay1InsertTracksRequest describes the request for QPlay1.InsertTracks API
type QPlay1InsertTracksRequest struct {
	QueueID        soap.String
	StartingIndex  soap.String
	TracksMetaData soap.String
}

// QPlay1InsertTracksResponse describes the response for QPlay1.InsertTracks API
type QPlay1InsertTracksResponse struct {
	NumberOfSuccess soap.String
}

func (client *QPlay1) InsertTracks(request QPlay1InsertTracksRequest) (*QPlay1InsertTracksResponse, error) {
	// Perform the SOAP call.
	var response QPlay1InsertTracksResponse
	if err := client.SOAPClient.PerformAction(URN_QPlay_1, "InsertTracks", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// QPlay1RemoveTracksRequest describes the request for QPlay1.RemoveTracks API
type QPlay1RemoveTracksRequest struct {
	QueueID        soap.String
	StartingIndex  soap.String
	NumberOfTracks soap.String
}

// QPlay1RemoveTracksResponse describes the response for QPlay1.RemoveTracks API
type QPlay1RemoveTracksResponse struct {
	NumberOfSuccess soap.String
}

func (client *QPlay1) RemoveTracks(request QPlay1RemoveTracksRequest) (*QPlay1RemoveTracksResponse, error) {
	// Perform the SOAP call.
	var response QPlay1RemoveTracksResponse
	if err := client.SOAPClient.PerformAction(URN_QPlay_1, "RemoveTracks", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// QPlay1RemoveAllTracksRequest describes the request for QPlay1.RemoveAllTracks API
type QPlay1RemoveAllTracksRequest struct {
	QueueID soap.String
}

// QPlay1RemoveAllTracksResponse describes the response for QPlay1.RemoveAllTracks API
type QPlay1RemoveAllTracksResponse struct {
	NumberOfSuccess soap.String
}

func (client *QPlay1) RemoveAllTracks(request QPlay1RemoveAllTracksRequest) (*QPlay1RemoveAllTracksResponse, error) {
	// Perform the SOAP call.
	var response QPlay1RemoveAllTracksResponse
	if err := client.SOAPClient.PerformAction(URN_QPlay_1, "RemoveAllTracks", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// QPlay1GetTracksInfoRequest describes the request for QPlay1.GetTracksInfo API
type QPlay1GetTracksInfoRequest struct {
	StartingIndex  soap.String
	NumberOfTracks soap.String
}

// QPlay1GetTracksInfoResponse describes the response for QPlay1.GetTracksInfo API
type QPlay1GetTracksInfoResponse struct {
	TracksMetaData soap.String
}

func (client *QPlay1) GetTracksInfo(request QPlay1GetTracksInfoRequest) (*QPlay1GetTracksInfoResponse, error) {
	// Perform the SOAP call.
	var response QPlay1GetTracksInfoResponse
	if err := client.SOAPClient.PerformAction(URN_QPlay_1, "GetTracksInfo", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// QPlay1SetTracksInfoRequest describes the request for QPlay1.SetTracksInfo API
type QPlay1SetTracksInfoRequest struct {
	QueueID        soap.String
	StartingIndex  soap.String
	NextIndex      soap.String
	TracksMetaData soap.String
}

// QPlay1SetTracksInfoResponse describes the response for QPlay1.SetTracksInfo API
type QPlay1SetTracksInfoResponse struct {
	NumberOfTracks soap.String
}

func (client *QPlay1) SetTracksInfo(request QPlay1SetTracksInfoRequest) (*QPlay1SetTracksInfoResponse, error) {
	// Perform the SOAP call.
	var response QPlay1SetTracksInfoResponse
	if err := client.SOAPClient.PerformAction(URN_QPlay_1, "SetTracksInfo", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// QPlay1GetTracksCountResponse describes the response for QPlay1.GetTracksCount API
type QPlay1GetTracksCountResponse struct {
	NrTracks soap.String
}

func (client *QPlay1) GetTracksCount() (*QPlay1GetTracksCountResponse, error) {
	// Perform the SOAP call.
	var response QPlay1GetTracksCountResponse
	if err := client.SOAPClient.PerformAction(URN_QPlay_1, "GetTracksCount", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// QPlay1GetMaxTracksResponse describes the response for QPlay1.GetMaxTracks API
type QPlay1GetMaxTracksResponse struct {
	MaxTracks soap.String
}

func (client *QPlay1) GetMaxTracks() (*QPlay1GetMaxTracksResponse, error) {
	// Perform the SOAP call.
	var response QPlay1GetMaxTracksResponse
	if err := client.SOAPClient.PerformAction(URN_QPlay_1, "GetMaxTracks", nil, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// QPlay1QPlayAuthRequest describes the request for QPlay1.QPlayAuth API
type QPlay1QPlayAuthRequest struct {
	Seed soap.String
}

// QPlay1QPlayAuthResponse describes the response for QPlay1.QPlayAuth API
type QPlay1QPlayAuthResponse struct {
	Code soap.String
	MID  soap.String
	DID  soap.String
}

func (client *QPlay1) QPlayAuth(request QPlay1QPlayAuthRequest) (*QPlay1QPlayAuthResponse, error) {
	// Perform the SOAP call.
	var response QPlay1QPlayAuthResponse
	if err := client.SOAPClient.PerformAction(URN_QPlay_1, "QPlayAuth", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// QPlay1SetNetworkRequest describes the request for QPlay1.SetNetwork API
type QPlay1SetNetworkRequest struct {
	SSID soap.String
	Key  soap.String
	// AuthAlgo: allowed values: open, shared, WPA, WPAPSK, WPA2, WPA2PSK
	AuthAlgo soap.String
	// CipherAlgo: allowed values: none, WEP, TKIP, AES
	CipherAlgo soap.String
}

// QPlay1SetNetworkResponse describes the response for QPlay1.SetNetwork API
type QPlay1SetNetworkResponse struct {
}

//
// Arguments:
//
//  QPlay1SetNetworkRequest
func (client *QPlay1) SetNetwork(request QPlay1SetNetworkRequest) (*QPlay1SetNetworkResponse, error) {
	// Perform the SOAP call.
	var response QPlay1SetNetworkResponse
	if err := client.SOAPClient.PerformAction(URN_QPlay_1, "SetNetwork", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
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

// RenderingControl1GetCurrentTransportActionsRequest describes the request for RenderingControl1.GetCurrentTransportActions API
type RenderingControl1GetCurrentTransportActionsRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetCurrentTransportActionsResponse describes the response for RenderingControl1.GetCurrentTransportActions API
type RenderingControl1GetCurrentTransportActionsResponse struct {
	Actions soap.String
}

func (client *RenderingControl1) GetCurrentTransportActions(request RenderingControl1GetCurrentTransportActionsRequest) (*RenderingControl1GetCurrentTransportActionsResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1GetCurrentTransportActionsResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetCurrentTransportActions", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1GetDeviceCapabilitiesRequest describes the request for RenderingControl1.GetDeviceCapabilities API
type RenderingControl1GetDeviceCapabilitiesRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetDeviceCapabilitiesResponse describes the response for RenderingControl1.GetDeviceCapabilities API
type RenderingControl1GetDeviceCapabilitiesResponse struct {
	// PlayMedia: allowed values: UNKNOWN, CD-DA, DVD-VIDEO, HDD, NETWORK, NONE
	PlayMedia soap.String
	// RecMedia: allowed values: NOT_IMPLEMENTED
	RecMedia soap.String
	// RecQualityModes: allowed values: NOT_IMPLEMENTED
	RecQualityModes soap.String
}

//
// Return value:
//
//  RenderingControl1GetDeviceCapabilitiesResponse
func (client *RenderingControl1) GetDeviceCapabilities(request RenderingControl1GetDeviceCapabilitiesRequest) (*RenderingControl1GetDeviceCapabilitiesResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1GetDeviceCapabilitiesResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetDeviceCapabilities", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1GetMediaInfoRequest describes the request for RenderingControl1.GetMediaInfo API
type RenderingControl1GetMediaInfoRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetMediaInfoResponse describes the response for RenderingControl1.GetMediaInfo API
type RenderingControl1GetMediaInfoResponse struct {
	// NrTracks: allowed value range: minimum=0, maximum=65535
	NrTracks           soap.Ui4
	MediaDuration      soap.String
	CurrentURI         soap.String
	CurrentURIMetaData soap.String
	NextURI            soap.String
	NextURIMetaData    soap.String
	TrackSource        soap.String
	// PlayMedium: allowed values: UNKNOWN, CD-DA, DVD-VIDEO, HDD, NETWORK, NONE
	PlayMedium soap.String
	// RecordMedium: allowed values: NOT_IMPLEMENTED
	RecordMedium soap.String
	// WriteStatus: allowed values: NOT_IMPLEMENTED
	WriteStatus soap.String
}

//
// Return value:
//
//  RenderingControl1GetMediaInfoResponse
func (client *RenderingControl1) GetMediaInfo(request RenderingControl1GetMediaInfoRequest) (*RenderingControl1GetMediaInfoResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1GetMediaInfoResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetMediaInfo", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1GetPositionInfoRequest describes the request for RenderingControl1.GetPositionInfo API
type RenderingControl1GetPositionInfoRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetPositionInfoResponse describes the response for RenderingControl1.GetPositionInfo API
type RenderingControl1GetPositionInfoResponse struct {
	// Track: allowed value range: minimum=0, maximum=65535, step=1
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
//  RenderingControl1GetPositionInfoResponse
func (client *RenderingControl1) GetPositionInfo(request RenderingControl1GetPositionInfoRequest) (*RenderingControl1GetPositionInfoResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1GetPositionInfoResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetPositionInfo", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1GetTransportInfoRequest describes the request for RenderingControl1.GetTransportInfo API
type RenderingControl1GetTransportInfoRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetTransportInfoResponse describes the response for RenderingControl1.GetTransportInfo API
type RenderingControl1GetTransportInfoResponse struct {
	// CurrentTransportState: allowed values: STOPPED, PAUSED_PLAYBACK, PLAYING, TRANSITIONING, NO_MEDIA_PRESENT
	CurrentTransportState soap.String
	// CurrentTransportStatus: allowed values: OK, ERROR_OCCURRED
	CurrentTransportStatus soap.String
	// CurrentSpeed: allowed values: 1
	CurrentSpeed soap.String
}

//
// Return value:
//
//  RenderingControl1GetTransportInfoResponse
func (client *RenderingControl1) GetTransportInfo(request RenderingControl1GetTransportInfoRequest) (*RenderingControl1GetTransportInfoResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1GetTransportInfoResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetTransportInfo", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1GetTransportSettingsRequest describes the request for RenderingControl1.GetTransportSettings API
type RenderingControl1GetTransportSettingsRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetTransportSettingsResponse describes the response for RenderingControl1.GetTransportSettings API
type RenderingControl1GetTransportSettingsResponse struct {
	// PlayMode: allowed values: NORMAL, REPEAT_TRACK, REPEAT_ALL
	PlayMode soap.String
	// RecQualityMode: allowed values: NOT_IMPLEMENTED
	RecQualityMode soap.String
}

//
// Return value:
//
//  RenderingControl1GetTransportSettingsResponse
func (client *RenderingControl1) GetTransportSettings(request RenderingControl1GetTransportSettingsRequest) (*RenderingControl1GetTransportSettingsResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1GetTransportSettingsResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetTransportSettings", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1NextRequest describes the request for RenderingControl1.Next API
type RenderingControl1NextRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1NextResponse describes the response for RenderingControl1.Next API
type RenderingControl1NextResponse struct {
}

func (client *RenderingControl1) Next(request RenderingControl1NextRequest) (*RenderingControl1NextResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1NextResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "Next", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1PauseRequest describes the request for RenderingControl1.Pause API
type RenderingControl1PauseRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1PauseResponse describes the response for RenderingControl1.Pause API
type RenderingControl1PauseResponse struct {
}

func (client *RenderingControl1) Pause(request RenderingControl1PauseRequest) (*RenderingControl1PauseResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1PauseResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "Pause", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1PlayRequest describes the request for RenderingControl1.Play API
type RenderingControl1PlayRequest struct {
	InstanceID soap.Ui4
	// Speed: allowed values: 1
	Speed soap.String
}

// RenderingControl1PlayResponse describes the response for RenderingControl1.Play API
type RenderingControl1PlayResponse struct {
}

//
// Arguments:
//
//  RenderingControl1PlayRequest
func (client *RenderingControl1) Play(request RenderingControl1PlayRequest) (*RenderingControl1PlayResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1PlayResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "Play", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1PreviousRequest describes the request for RenderingControl1.Previous API
type RenderingControl1PreviousRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1PreviousResponse describes the response for RenderingControl1.Previous API
type RenderingControl1PreviousResponse struct {
}

func (client *RenderingControl1) Previous(request RenderingControl1PreviousRequest) (*RenderingControl1PreviousResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1PreviousResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "Previous", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1SeekRequest describes the request for RenderingControl1.Seek API
type RenderingControl1SeekRequest struct {
	InstanceID soap.Ui4
	// Unit: allowed values: REL_TIME, TRACK_NR
	Unit   soap.String
	Target soap.String
}

// RenderingControl1SeekResponse describes the response for RenderingControl1.Seek API
type RenderingControl1SeekResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SeekRequest
func (client *RenderingControl1) Seek(request RenderingControl1SeekRequest) (*RenderingControl1SeekResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1SeekResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "Seek", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1SetAVTransportURIRequest describes the request for RenderingControl1.SetAVTransportURI API
type RenderingControl1SetAVTransportURIRequest struct {
	InstanceID         soap.Ui4
	CurrentURI         soap.String
	CurrentURIMetaData soap.String
}

// RenderingControl1SetAVTransportURIResponse describes the response for RenderingControl1.SetAVTransportURI API
type RenderingControl1SetAVTransportURIResponse struct {
}

func (client *RenderingControl1) SetAVTransportURI(request RenderingControl1SetAVTransportURIRequest) (*RenderingControl1SetAVTransportURIResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1SetAVTransportURIResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetAVTransportURI", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1StopRequest describes the request for RenderingControl1.Stop API
type RenderingControl1StopRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1StopResponse describes the response for RenderingControl1.Stop API
type RenderingControl1StopResponse struct {
}

func (client *RenderingControl1) Stop(request RenderingControl1StopRequest) (*RenderingControl1StopResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1StopResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "Stop", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1GetPlayTypeRequest describes the request for RenderingControl1.GetPlayType API
type RenderingControl1GetPlayTypeRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetPlayTypeResponse describes the response for RenderingControl1.GetPlayType API
type RenderingControl1GetPlayTypeResponse struct {
	PlayType soap.String
}

func (client *RenderingControl1) GetPlayType(request RenderingControl1GetPlayTypeRequest) (*RenderingControl1GetPlayTypeResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1GetPlayTypeResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetPlayType", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1GetInfoExRequest describes the request for RenderingControl1.GetInfoEx API
type RenderingControl1GetInfoExRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1GetInfoExResponse describes the response for RenderingControl1.GetInfoEx API
type RenderingControl1GetInfoExResponse struct {
	// CurrentTransportState: allowed values: STOPPED, PAUSED_PLAYBACK, PLAYING, TRANSITIONING, NO_MEDIA_PRESENT
	CurrentTransportState soap.String
	// CurrentTransportStatus: allowed values: OK, ERROR_OCCURRED
	CurrentTransportStatus soap.String
	// CurrentSpeed: allowed values: 1
	CurrentSpeed soap.String
	// Track: allowed value range: minimum=0, maximum=65535, step=1
	Track          soap.Ui4
	TrackDuration  soap.String
	TrackMetaData  soap.String
	TrackURI       soap.String
	RelTime        soap.String
	AbsTime        soap.String
	RelCount       soap.I4
	AbsCount       soap.I4
	LoopMode       soap.Ui4
	CurrentVolume  soap.Ui4
	CurrentChannel soap.Ui4
	SlaveList      soap.String
	// PlayMedium: allowed values: UNKNOWN, CD-DA, DVD-VIDEO, HDD, NETWORK, NONE
	PlayMedium      soap.String
	TrackSource     soap.String
	InternetAccess  soap.Ui4
	VerUpdateFlag   soap.Ui4
	VerUpdateStatus soap.Ui4
	BatteryFlag     soap.Ui4
	BatteryPercent  soap.Ui4
	AlarmFlag       soap.Ui4
	TimeStamp       soap.Ui4
	SubNum          soap.Ui4
	SpotifyActive   soap.Ui4
}

//
// Return value:
//
//  RenderingControl1GetInfoExResponse
func (client *RenderingControl1) GetInfoEx(request RenderingControl1GetInfoExRequest) (*RenderingControl1GetInfoExResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1GetInfoExResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "GetInfoEx", &request, &response); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1SetPlayModeRequest describes the request for RenderingControl1.SetPlayMode API
type RenderingControl1SetPlayModeRequest struct {
	InstanceID soap.Ui4
	// NewPlayMode: allowed values: NORMAL, REPEAT_TRACK, REPEAT_ALL
	NewPlayMode soap.String
}

// RenderingControl1SetPlayModeResponse describes the response for RenderingControl1.SetPlayMode API
type RenderingControl1SetPlayModeResponse struct {
}

//
// Arguments:
//
//  RenderingControl1SetPlayModeRequest
func (client *RenderingControl1) SetPlayMode(request RenderingControl1SetPlayModeRequest) (*RenderingControl1SetPlayModeResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1SetPlayModeResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "SetPlayMode", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1SeekForwardRequest describes the request for RenderingControl1.SeekForward API
type RenderingControl1SeekForwardRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1SeekForwardResponse describes the response for RenderingControl1.SeekForward API
type RenderingControl1SeekForwardResponse struct {
}

func (client *RenderingControl1) SeekForward(request RenderingControl1SeekForwardRequest) (*RenderingControl1SeekForwardResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1SeekForwardResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "SeekForward", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}

// RenderingControl1SeekBackwardRequest describes the request for RenderingControl1.SeekBackward API
type RenderingControl1SeekBackwardRequest struct {
	InstanceID soap.Ui4
}

// RenderingControl1SeekBackwardResponse describes the response for RenderingControl1.SeekBackward API
type RenderingControl1SeekBackwardResponse struct {
}

func (client *RenderingControl1) SeekBackward(request RenderingControl1SeekBackwardRequest) (*RenderingControl1SeekBackwardResponse, error) {
	// Perform the SOAP call.
	var response RenderingControl1SeekBackwardResponse
	if err := client.SOAPClient.PerformAction(URN_RenderingControl_1, "SeekBackward", &request, nil); err != nil {
		return nil, errors.Wrap(err, "performing SOAP request")
	}
	return &response, nil
}
