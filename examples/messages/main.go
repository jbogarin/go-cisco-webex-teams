package main

import (
	"fmt"
	"log"

	"github.com/jbogarin/go-cisco-webex-teams/sdk"
	resty "gopkg.in/resty.v1"
)

// Client is Webex Teams API client
var Client *webexteams.Client

func main() {
	client := resty.New()
	token := "" // Change to your test token
	client.SetAuthToken(token)
	Client = webexteams.NewClient(client)

	/*

		MESSAGES

	*/

	myRoomID := "" // Change to your testing room

	// POST messages - Text Message

	message := &webexteams.MessageCreateRequest{
		Text:   "This is a text message",
		RoomID: myRoomID,
	}
	newTextMessage, _, err := Client.Messages.CreateMessage(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("POST:", newTextMessage.ID, newTextMessage.Text, newTextMessage.Created)

	// POST messages - Markdown Message

	markDownMessage := &webexteams.MessageCreateRequest{
		Markdown: "This is a markdown message. *Italic*, **bold** and ***italic/bold***.",
		RoomID:   myRoomID,
	}
	newMarkDownMessage, _, err := Client.Messages.CreateMessage(markDownMessage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("POST:", newMarkDownMessage.ID, newMarkDownMessage.Markdown, newMarkDownMessage.Created)

	// GET messages
	messageQueryParams := &webexteams.ListMessagesQueryParams{
		Max:    5,
		RoomID: myRoomID,
	}

	messages, _, err := Client.Messages.ListMessages(messageQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, message := range messages.Items {
		fmt.Println("GET:", id, message.ID, message.Text, message.Created)
	}

	// GET messages/<ID>

	htmlMessageGet, _, err := Client.Messages.GetMessage(newMarkDownMessage.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", htmlMessageGet.ID, htmlMessageGet.Text, htmlMessageGet.Created)

	// DELETE messages<ID>

	resp, err := Client.Messages.DeleteMessage(newTextMessage.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DELETE:", resp.StatusCode())

}
