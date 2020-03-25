package main

import (
	"fmt"
	"log"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
)

// Client is Webex Teams API client
var Client *webexteams.Client

func main() {
	Client = webexteams.NewClient()

	/*

		DEVICES

	*/

	// GET devices
	queryParams := &webexteams.ListDevicesQueryParams{
		DisplayName: "", // Change to the device you want to look for
		Max:         10,
	}

	devices, resp, err := Client.Devices.ListDevices(queryParams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode())
	deviceID := ""
	for id, device := range devices.Items {
		fmt.Println("GET:", id, device.ID, device.DisplayName, device.Product)
		deviceID = device.ID
	}

	// GET devices/<id>
	device, _, err := Client.Devices.GetDevice(deviceID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", device.ID, device.DisplayName)

}
