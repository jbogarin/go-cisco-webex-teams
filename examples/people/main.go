package main

import (
	"fmt"
	"log"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
)

// Client is Webex Teams API client
var Client *webexteams.Client
var Client2 *webexteams.Client

func main() {
	Client = webexteams.NewClient()

	/*

		PEOPLE

	*/

	// GET people
	queryParams := &webexteams.ListPeopleQueryParams{
		Email: "", // Change to the person you want to look for
		Max:   2,
	}

	people, _, err := Client.People.ListPeople(queryParams)
	if err != nil {
		log.Fatal(err)
	}

	personID := ""
	for id, person := range people.Items {
		fmt.Println("GET:", id, person.ID, person.DisplayName, person.Created)
		personID = person.ID
	}

	// GET people/<id>
	person, _, err := Client.People.GetPerson(personID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", person.ID, person.DisplayName, person.Created)

	// GET people/me
	me, _, err := Client.People.GetMe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET ME:", me.ID, me.DisplayName, me.Created)
}
