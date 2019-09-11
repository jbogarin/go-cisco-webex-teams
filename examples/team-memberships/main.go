package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty"
	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
)

// Client is Webex Teams API client
var Client *webexteams.Client

func main() {
	client := resty.New()
	token := "" // Change to your test token
	client.SetAuthToken(token)
	Client = webexteams.NewClient(client)

	/*

		TEAM MEMBERSHIPS

	*/

	myTeamID := ""      // Change to your test team
	newTeamMember := "" // Change to the person email you want to add to the team

	// POST team-memberships
	teamMembershipRequest := &webexteams.TeamMembershipCreateRequest{
		TeamID:      myTeamID,
		PersonEmail: newTeamMember,
		IsModerator: true,
	}

	newTeamMembership, _, err := Client.TeamMemberships.CreateTeamMembership(teamMembershipRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST:", newTeamMembership.ID, newTeamMembership.PersonEmail, newTeamMembership.IsModerator, newTeamMembership.Created)

	// GET team-memberships
	teamMembershipsQueryParams := &webexteams.ListTeamMemberhipsQueryParams{
		Max:    2,
		TeamID: myTeamID,
	}

	teamMemberships, _, err := Client.TeamMemberships.ListTeamMemberhips(teamMembershipsQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, teamMembership := range teamMemberships.Items {
		fmt.Println("GET:", id, teamMembership.ID, teamMembership.PersonEmail, teamMembership.IsModerator, teamMembership.Created)
	}

	// GET team-memberships/<id>
	teamMembership, _, err := Client.TeamMemberships.GetTeamMembership(newTeamMembership.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET <ID>:", teamMembership.ID, teamMembership.PersonEmail, teamMembership.IsModerator, teamMembership.Created)

	// PUT team-memberships/<id>
	updateTeamMembershipRequest := &webexteams.TeamMembershipUpdateRequest{
		IsModerator: false,
	}

	updatedTeamMembership, _, err := Client.TeamMemberships.UpdateTeamMembership(newTeamMembership.ID, updateTeamMembershipRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedTeamMembership.ID, updatedTeamMembership.PersonEmail, updatedTeamMembership.IsModerator, updatedTeamMembership.Created)

	// DELETE team-memberships/<id>
	resp, err := Client.TeamMemberships.DeleteTeamMembership(newTeamMembership.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETE:", resp.StatusCode())

}
