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

		RECORDINGS

	*/

	// GET recordings
	queryParams := &webexteams.ListRecordingsQueryParams{
		From: "2019-01-01T00:00:00+00:00",
		To:   "2020-07-01T00:00:00+00:00",
		Max:  2,
	}

	recordings, resp, err := Client.Recordings.ListRecordings(queryParams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode())
	recordingID := ""
	for id, recording := range recordings.Items {
		fmt.Println("GET:", id, recording.ID, recording.Topic, recording.CreateTime, recording.PlaybackURL)
		recordingID = recording.ID
	}

	// GET recordings/<id>
	recording, _, err := Client.Recordings.GetRecording(recordingID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", recording.ID, recording.TemporaryDirectDownloadLinks.RecordingDownloadLink)

	// DELETE recordings/<id>
	resp, err = Client.Recordings.DeleteRecording(recordingID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DELETE:", resp.StatusCode())

}
