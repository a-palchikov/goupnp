package main

import (
	"fmt"
	"log"

	"github.com/huin/goupnp/dcps/internetgateway1"
)

func main() {
	clients, errors, err := internetgateway1.NewWANPPPConnection1Clients()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got %d errors finding servers and %d successfully discovered.\n",
		len(errors), len(clients))
	for i, e := range errors {
		fmt.Printf("Error finding server #%d: %v\n", i+1, e)
	}

	for _, c := range clients {
		dev := &c.ServiceClient.RootDevice.Device
		srv := c.ServiceClient.Service
		fmt.Println(dev.FriendlyName, " :: ", srv.String())
		scpd, err := srv.RequestSCPD()
		if err != nil {
			fmt.Printf("  Error requesting service SCPD: %v\n", err)
		} else {
			fmt.Println("  Available actions:")
			for _, action := range scpd.Actions {
				fmt.Printf("  * %s\n", action.Name)
				for _, arg := range action.Arguments {
					var varDesc string
					if stateVar := scpd.GetStateVariable(arg.RelatedStateVariable); stateVar != nil {
						varDesc = fmt.Sprintf(" (%s)", stateVar.DataType.Name)
					}
					fmt.Printf("    * [%s] %s%s\n", arg.Direction, arg.Name, varDesc)
				}
			}
		}

		if scpd == nil || scpd.GetAction("GetExternalIPAddress") != nil {
			ip, err := c.GetExternalIPAddress()
			fmt.Println("GetExternalIPAddress: ", ip, err)
		}

		if scpd == nil || scpd.GetAction("GetStatusInfo") != nil {
			resp, err := c.GetStatusInfo()
			fmt.Println("GetStatusInfo: ", resp.NewConnectionStatus, resp.NewLastConnectionError, resp.NewUptime, err)
		}

		if scpd == nil || scpd.GetAction("GetIdleDisconnectTime") != nil {
			resp, err := c.GetIdleDisconnectTime()
			fmt.Println("GetIdleDisconnectTime: ", resp.NewIdleDisconnectTime, err)
		}

		if scpd == nil || scpd.GetAction("AddPortMapping") != nil {
			req := internetgateway1.WANPPPConnection1AddPortMappingRequest{
				NewExternalPort:           5000,
				NewProtocol:               "TCP",
				NewInternalPort:           5001,
				NewInternalClient:         "192.168.1.2",
				NewEnabled:                true,
				NewPortMappingDescription: "Test port mapping",
			}
			_, err := c.AddPortMapping(req)
			fmt.Println("AddPortMapping: ", err)
		}
		if scpd == nil || scpd.GetAction("DeletePortMapping") != nil {
			req := internetgateway1.WANPPPConnection1DeletePortMappingRequest{
				NewExternalPort: 5000,
				NewProtocol:     "TCP",
			}
			_, err := c.DeletePortMapping(req)
			fmt.Println("DeletePortMapping: ", err)
		}
	}
}
