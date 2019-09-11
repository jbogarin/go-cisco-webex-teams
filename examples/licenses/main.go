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

		Licenses

	*/

	// GET Licenses
	queryParams := &webexteams.ListLicensesQueryParams{
		Max: 2,
	}

	Licenses, _, err := Client.Licenses.ListLicenses(queryParams)
	if err != nil {
		log.Fatal(err)
	}

	LicenseID := ""
	for id, License := range Licenses.Items {
		fmt.Println("GET:", id, License.ID, License.Name, License.TotalUnits, License.ConsumedUnits)
		LicenseID = License.ID
	}

	// GET Licenses/<id>
	License, _, err := Client.Licenses.GetLicense(LicenseID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", License.ID, License.Name, License.TotalUnits, License.ConsumedUnits)

}
