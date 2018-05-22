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

		Roles

	*/

	// GET Roles
	queryParams := &webexteams.RolesListQueryParams{
		Max: 2,
	}

	Roles, _, err := Client.Roles.ListRoles(queryParams)
	if err != nil {
		log.Fatal(err)
	}

	RoleID := ""
	for id, Role := range Roles.Items {
		fmt.Println("GET:", id, Role.ID, Role.Name)
		RoleID = Role.ID
	}

	// GET Roles/<id>
	Role, _, err := Client.Roles.GetRole(RoleID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", Role.ID, Role.Name)

}
