package webexteams

import (
	"fmt"
	"strings"

	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
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
	Created           time.Time `json:"created,omitempty"`           // Membership creation date/time.
	IsModerator       bool      `json:"isModerator,omitempty"`       // Membership is moderator.
	IsMonitor         bool      `json:"isMonitor,omitempty"`         // Membership is monitor.
	IsRoomHidden      bool      `json:"isRoomHidden,omitempty"`      // Whether or not the direct type room is hidden in the Webex clients.
	PersonDisplayName string    `json:"personDisplayName,omitempty"` // Person display name.
	PersonEmail       string    `json:"personEmail,omitempty"`       // Person email.
	PersonID          string    `json:"personId,omitempty"`          // Person ID.
	PersonOrgID       string    `json:"personOrgId,omitempty"`       // The organization ID of the person.
	RoomID            string    `json:"roomId,omitempty"`            // Room ID.
	RoomType          string    `json:"roomType,omitempty"`          // Room Type.
}

// MembershipUpdateRequest is the Update Membership Request object
type MembershipUpdateRequest struct {
	IsModerator  bool `json:"isModerator,omitempty"`  // Membership is a moderator.
	IsRoomHidden bool `json:"isRoomHidden,omitempty"` // Whether or not the direct type room is hidden in the Webex clients.
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

func (s *MembershipsService) membershipsPagination(linkHeader string, size, max int) *Memberships {
	items := &Memberships{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := s.client.R().
				SetResult(&Memberships{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Memberships)
			size = size + len(items.Items)
			if max < 0 || size < max {
				memberships := s.membershipsPagination(response.Header().Get("Link"), size, max)
				for _, membership := range memberships.Items {
					items.AddMembership(membership)
				}
			}

		}
	}

	return items
}

// CreateMembership Add someone to a membership by Person ID or email address; optionally making them a moderator.
/* Add someone to a membership by Person ID or email address; optionally making them a moderator.
@param membershipCreateRequest
@return Membership
*/
func (s *MembershipsService) CreateMembership(membershipCreateRequest *MembershipCreateRequest) (*Membership, *resty.Response, error) {

	path := "/memberships/"

	response, err := s.client.R().
		SetBody(membershipCreateRequest).
		SetResult(&Membership{}).
		SetError(&Error{}).
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

	response, err := s.client.R().
		SetError(&Error{}).
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

	response, err := s.client.R().
		SetResult(&Membership{}).
		SetError(&Error{}).
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
	Max         int    `url:"max,omitempty"`         // Limit the maximum number of items in the response. Negative value will list all items (use this carefully).
	RequestBy   int    `url:"-"`                     // Number of items to retrieve by requests (Max if let at 0)
}

// ListMemberships Lists all membership memberships. By default, lists memberships for Memberships to which the authenticated user belongs.
/* Lists all membership memberships. By default, lists memberships for Memberships to which the authenticated user belongs.
Use query parameters to filter the response.
Use roomID to list memberships for a membership, by ID.
Use either personID or personEmail to filter the results.
Long result sets will be split into pages.

 @param "roomId" (string) Room ID.
 @param "personId" (string) Person ID.
 @param "personEmail" (string) Person email.
 @param "max" (int) Limit the maximum number of items in the response. Negative value will list all items (use this carefully).
 @param "requestBy" (int) Number of items by request
 @return Memberships
*/
func (s *MembershipsService) ListMemberships(queryParams *ListMembershipsQueryParams) (*Memberships, *resty.Response, error) {

	path := "/memberships/"

	max := queryParams.Max

	if queryParams.RequestBy > 0 {
		queryParams.Max = queryParams.RequestBy
	} else if queryParams.Max < 0 {
		queryParams.Max = 0
	}

	queryParamsString, _ := query.Values(queryParams)

	response, err := s.client.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Memberships{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Memberships)

	if max < 0 || len(result.Items) < max {
		items := s.membershipsPagination(response.Header().Get("Link"), len(result.Items), max)
		for _, membership := range items.Items {
			result.AddMembership(membership)
		}
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

	response, err := s.client.R().
		SetBody(membershipUpdateRequest).
		SetResult(&Membership{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Membership)
	return result, response, err

}
