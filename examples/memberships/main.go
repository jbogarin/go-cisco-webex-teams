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

		MEMBERSHIPS

	*/

	// This works if you create a room where you are a moderator (paid feature). I tested with a room that it is part of a team.

	myRoomID := ""       // Change to your testing room
	myTestingEmail := "" // Change to your email

	// GET memberships

	membershipQueryParams := &webexteams.ListMembershipsQueryParams{
		Max:         10,
		PersonEmail: myTestingEmail,
	}

	memberships, _, err := Client.Memberships.ListMemberships(membershipQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, membership := range memberships.Items {
		fmt.Println("GET:", id, membership.ID, membership.PersonEmail, membership.IsModerator, membership.Created)
	}

	// POST memberships

	membershipRequest := &webexteams.MembershipCreateRequest{
		RoomID:      myRoomID,
		PersonEmail: myTestingEmail,
		IsModerator: true,
	}

	testMembership, _, err := Client.Memberships.CreateMembership(membershipRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST:", testMembership.ID, testMembership.PersonEmail, testMembership.IsModerator, testMembership.Created)

	// GET memberships/<ID>

	membership, _, err := Client.Memberships.GetMembership(testMembership.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET <ID>:", membership.ID, membership.PersonEmail, membership.IsModerator, membership.Created)

	// PUT memberships/<ID>

	updateMembershipRequest := &webexteams.MembershipUpdateRequest{
		IsModerator: false,
	}

	updatedMembership, _, err := Client.Memberships.UpdateMembership(testMembership.ID, updateMembershipRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PUT:", updatedMembership.ID, updatedMembership.PersonEmail, updatedMembership.IsModerator, updatedMembership.Created)

	// DELETE memberships<ID>

	resp, err := Client.Memberships.DeleteMembership(testMembership.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETE:", resp.StatusCode())

}
