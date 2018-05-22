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

		TEAMS

	*/

	// POST teams
	teamRequest := &webexteams.TeamCreateRequest{
		Name: "Go Test Team",
	}

	newTeam, _, err := Client.Teams.CreateTeam(teamRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST:", newTeam.ID, newTeam.Name, newTeam.Created)

	// GET teams

	teamQueryParams := &webexteams.ListTeamsQueryParams{
		Max: 2,
	}

	teams, _, err := Client.Teams.ListTeams(teamQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, team := range teams.Items {
		fmt.Println("GET:", id, team.ID, team.Name, team.Created)
	}

	// GET teams/<id>
	team, _, err := Client.Teams.GetTeam(newTeam.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET <ID>:", team.ID, team.Created, team.Name)

	// PUT teams/<id>
	updateTeamRequest := &webexteams.TeamUpdateRequest{
		Name: "Go Test Team 2",
	}

	updatedTeam, _, err := Client.Teams.UpdateTeam(newTeam.ID, updateTeamRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PUT:", updatedTeam.ID, updatedTeam.Name, updatedTeam.Created)

	// DELETE teams/<id>
	resp, err := Client.Teams.DeleteTeam(newTeam.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETE:", resp.StatusCode())

}
