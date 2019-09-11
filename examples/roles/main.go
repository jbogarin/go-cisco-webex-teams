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
