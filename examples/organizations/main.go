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

		Organizations

	*/

	// GET Organizations
	queryParams := &webexteams.ListOrganizationsQueryParams{
		Max: 2,
	}

	Organizations, _, err := Client.Organizations.ListOrganizations(queryParams)
	if err != nil {
		log.Fatal(err)
	}

	for id, Organization := range Organizations.Items {
		fmt.Println("GET:", id, Organization.ID, Organization.DisplayName, Organization.Created)
	}

	OrganizationID := "Y2lzY29zcGFyazovL3VzL09SR0FOSVpBVElPTi9hMTI1NGM2ZS0xMTdkLTQ5ZmYtOTE5Ny1kZDUyYjQzOWY2OWM" // Your test organization ID

	// GET Organizations/<id>
	Organization, _, err := Client.Organizations.GetOrganization(OrganizationID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", Organization.ID, Organization.DisplayName, Organization.Created)

}
