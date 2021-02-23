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

		Events

	*/

	// GET events
	eventQueryParams := &webexteams.ListEventsQueryParams{
		Resource: "messages",
		From:     "2021-02-23T00:00:00.000Z",
		To:       "2021-02-23T12:00:00.000Z",
	}

	events, _, err := Client.Events.ListEvents(eventQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, event := range events.Items {
		fmt.Println("GET:", id, event.ID, event.Created, event.Data.ID)
	}

	// GET events/<ID>

	eventGet, _, err := Client.Events.GetEvent("<EVENT ID>")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", eventGet.ID, eventGet.Data.Text, eventGet.Created, eventGet.Data.PersonEmail, eventGet.Data.RoomType)

}
