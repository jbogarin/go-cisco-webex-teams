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

		PEOPLE

	*/

	// GET people
	queryParams := &webexteams.ListPeopleQueryParams{
		DisplayName: "", // Change to the person you want to look for
		Max:         2,
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
