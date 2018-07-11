package webexteams

import (
	"fmt"
	"strings"

	"time"

	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
	"gopkg.in/resty.v1"
)

// MembershipsService is the service to communicate with the Memberships API endpoint
type MembershipsService service

// Memberships is the List of Memberships
type Memberships struct {
	Items []Membership `json:"items,omitempty"`
}

// Membership is the Membership definition
type Membership struct {
	ID                string    `json:"id,omitempty"`                // Membership ID.
	RoomID            string    `json:"roomId,omitempty"`            // Room ID.
	PersonID          string    `json:"personId,omitempty"`          // Person ID.
	PersonEmail       string    `json:"personEmail,omitempty"`       // Person email.
	PersonDisplayName string    `json:"personDisplayName,omitempty"` // Person display name.
	IsModerator       bool      `json:"isModerator,omitempty"`       // Membership is moderator.
	IsMonitor         bool      `json:"isMonitor,omitempty"`         // Membership is monitor.
	Created           time.Time `json:"created,omitempty"`           // Membership creation date/time.
}

// MembershipUpdateRequest is the Update Membership Request object
type MembershipUpdateRequest struct {
	IsModerator bool `json:"isModerator,omitempty"` // Membership is a moderator.
}

// MembershipCreateRequest is the Create Membership Request Parameters
type MembershipCreateRequest struct {
	RoomID      string `json:"roomId,omitempty"`      // Room ID.
	PersonID    string `json:"personId,omitempty"`    // Person ID.
	PersonEmail string `json:"personEmail,omitempty"` // Person email.
	IsModerator bool   `json:"isModerator,omitempty"` // Membership is a moderator.
}

// AddMembership is used to append a membership to a slice of memberships
func (memberships *Memberships) AddMembership(item Membership) []Membership {
	memberships.Items = append(memberships.Items, item)
	return memberships.Items
}

func membershipLoop(linkHeader string) *Memberships {
	items := &Memberships{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := RestyClient.R().
				SetResult(&Memberships{}).
				Get(l.URI)

			if err != nil {
				fmt.Println("Error")
			}
			items = response.Result().(*Memberships)
			memberships := membershipLoop(response.Header().Get("Link"))
			for _, membership := range memberships.Items {
				items.AddMembership(membership)
			}
		}
	}

	return items
}

// CreateMembership Add someone to a room by Person ID or email address; optionally making them a moderator.
/* Add someone to a room by Person ID or email address; optionally making them a moderator.
@param membershipCreateRequest
@return Membership
*/
func (s *MembershipsService) CreateMembership(membershipCreateRequest *MembershipCreateRequest) (*Membership, *resty.Response, error) {

	path := "/memberships/"

	response, err := RestyClient.R().
		SetBody(membershipCreateRequest).
		SetResult(&Membership{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Membership)
	return result, response, err

}

// DeleteMembership Deletes a membership by ID.
/* Deletes a membership by ID.
Specify the membership ID in the membershipID URI parameter.

 @param membershipID Membership ID.
 @return
*/
func (s *MembershipsService) DeleteMembership(membershipID string) (*resty.Response, error) {

	path := "/memberships/{membershipId}"
	path = strings.Replace(path, "{"+"membershipId"+"}", fmt.Sprintf("%v", membershipID), -1)

	response, err := RestyClient.R().
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetMembership Get details for a membership by ID.
/* Get details for a membership by ID.
Specify the membership ID in the membershipID URI parameter.

 @param membershipID Membership ID.
 @return Membership
*/
func (s *MembershipsService) GetMembership(membershipID string) (*Membership, *resty.Response, error) {

	path := "/memberships/{membershipId}"
	path = strings.Replace(path, "{"+"membershipId"+"}", fmt.Sprintf("%v", membershipID), -1)

	response, err := RestyClient.R().
		SetResult(&Membership{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Membership)
	return result, response, err

}

// ListMembershipsQueryParams are the query params for the ListMemberships API Call
type ListMembershipsQueryParams struct {
	RoomID      string `url:"roomId,omitempty"`      // Room ID.
	PersonID    string `url:"personId,omitempty"`    // Person ID.
	PersonEmail string `url:"personEmail,omitempty"` // Person email.
	Max         int    `url:"max,omitempty"`         // Limit the maximum number of items in the response.
}

// ListMemberships Lists all room memberships. By default, lists memberships for rooms to which the authenticated user belongs.
/* Lists all room memberships. By default, lists memberships for rooms to which the authenticated user belongs.
Use query parameters to filter the response.
Use roomID to list memberships for a room, by ID.
Use either personID or personEmail to filter the results.
Long result sets will be split into pages.

 @param "roomId" (string) Room ID.
 @param "personId" (string) Person ID.
 @param "personEmail" (string) Person email.
 @param "max" (int) Limit the maximum number of items in the response.
 @return Memberships
*/
func (s *MembershipsService) ListMemberships(queryParams *ListMembershipsQueryParams) (*Memberships, *resty.Response, error) {

	path := "/memberships/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Memberships{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Memberships)
	items := membershipLoop(response.Header().Get("Link"))

	for _, membersip := range items.Items {
		result.AddMembership(membersip)
	}
	return result, response, err

}

// UpdateMembership Updates properties for a membership by ID.
/* Updates properties for a membership by ID.
Specify the membership ID in the membershipID URI parameter.

 @param membershipID Membership ID.
 @param membershipUpdateRequest
 @return Membership
*/
func (s *MembershipsService) UpdateMembership(membershipID string, membershipUpdateRequest *MembershipUpdateRequest) (*Membership, *resty.Response, error) {

	path := "/memberships/{membershipId}"
	path = strings.Replace(path, "{"+"membershipId"+"}", fmt.Sprintf("%v", membershipID), -1)

	response, err := RestyClient.R().
		SetBody(membershipUpdateRequest).
		SetResult(&Membership{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Membership)
	return result, response, err

}
