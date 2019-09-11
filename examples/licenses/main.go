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

		Licenses

	*/

	// GET Licenses
	queryParams := &webexteams.ListLicensesQueryParams{
		Max: 2,
	}

	Licenses, resp, err := Client.Licenses.ListLicenses(queryParams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode())
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
