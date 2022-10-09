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

		ROOMS

	*/

	// POST rooms

	roomRequest := &webexteams.RoomCreateRequest{
		Title: "Go Test Room",
	}

	newRoom, response, err := Client.Rooms.CreateRoom(roomRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST:", newRoom.ID, newRoom.Title, newRoom.IsLocked, newRoom.Created, response.StatusCode())

	// GET rooms
	roomsQueryParams := &webexteams.ListRoomsQueryParams{
		Max:       -1, // list all the rooms
		TeamID:    "",
		RequestBy: 200, // list rooms 200 by 200
	}

	rooms, _, err := Client.Rooms.ListRooms(roomsQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, room := range rooms.Items {
		fmt.Println("GET:", id, room.ID, room.IsLocked, room.Title)
	}

	// GET rooms/<id>

	room, _, err := Client.Rooms.GetRoom(newRoom.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET <ID>:", room.ID, room.Title, room.IsLocked, room.Created)

	updateRoomRequest := &webexteams.RoomUpdateRequest{
		Title: "Go Test Room 2",
	}

	updatedRoom, _, err := Client.Rooms.UpdateRoom(newRoom.ID, updateRoomRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PUT:", updatedRoom.ID, updatedRoom.Title, updatedRoom.IsLocked, updatedRoom.Created)

	// // DELETE

	resp, err := Client.Rooms.DeleteRoom(newRoom.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETE:", resp.StatusCode())

}
