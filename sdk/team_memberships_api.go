package webexteams

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
)

// TeamMembershipsService is the service to communicate with the TeamMemberships API endpoint
type TeamMembershipsService service

// TeamMembershipCreateRequest is the Team Membership Create Request Parameters
type TeamMembershipCreateRequest struct {
	TeamID      string `json:"teamId,omitempty"`      // Team ID.
	PersonID    string `json:"personId,omitempty"`    // Person ID.
	PersonEmail string `json:"personEmail,omitempty"` // Person Email.
	IsModerator bool   `json:"isModerator,omitempty"` // Team Membership is a moderator.
}

// TeamMembershipUpdateRequest is the Team Membership Update Request object
type TeamMembershipUpdateRequest struct {
	IsModerator bool `json:"isModerator,omitempty"` // Team Membership is a moderator.
}

// TeamMembership is the Team Membership definition
type TeamMembership struct {
	ID                string    `json:"id,omitempty"`                // Team Membership ID.
	TeamID            string    `json:"teamId,omitempty"`            // Team ID.
	PersonID          string    `json:"personId,omitempty"`          // Person ID.
	PersonEmail       string    `json:"personEmail,omitempty"`       // Person email.
	PersonDisplayName string    `json:"personDisplayName,omitempty"` // Person display name.
	IsModerator       bool      `json:"isModerator,omitempty"`       // Team Membership is moderator.
	Created           time.Time `json:"created,omitempty"`           // Team Membership creation date/time.
}

// TeamMemberships is the List of Team Memberships
type TeamMemberships struct {
	Items []TeamMembership `json:"items,omitempty"`
}

// AddTeamMembership is used to append a teamMembership to a slice of teamMemberships
func (teamMembership *TeamMemberships) AddTeamMembership(item TeamMembership) []TeamMembership {
	teamMembership.Items = append(teamMembership.Items, item)
	return teamMembership.Items
}

func teamMembershipsPagination(linkHeader string, size, max int) *TeamMemberships {
	items := &TeamMemberships{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := RestyClient.R().
				SetResult(&TeamMemberships{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*TeamMemberships)
			if size != 0 {
				size = size + len(items.Items)
				if size < max {
					teamMemberships := teamMembershipsPagination(response.Header().Get("Link"), size, max)
					for _, teamMembership := range teamMemberships.Items {
						items.AddTeamMembership(teamMembership)
					}
				}
			} else {
				teamMemberships := teamMembershipsPagination(response.Header().Get("Link"), size, max)
				for _, teamMembership := range teamMemberships.Items {
					items.AddTeamMembership(teamMembership)
				}
			}

		}
	}

	return items
}

// CreateTeamMembership Add someone to a team by Person ID or email address; optionally making them a moderator.
/* Add someone to a team by Person ID or email address; optionally making them a moderator.
@param teamMemberhipCreateRequest
@return TeamMembership
*/
func (s *TeamMembershipsService) CreateTeamMembership(teamMemberhipCreateRequest *TeamMembershipCreateRequest) (*TeamMembership, *resty.Response, error) {

	path := "/team/memberships/"

	response, err := RestyClient.R().
		SetBody(teamMemberhipCreateRequest).
		SetResult(&TeamMembership{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TeamMembership)
	return result, response, err

}

// DeleteTeamMembership Deletes a team membership, by ID.
/* Deletes a team membership, by ID.
Specify the team membership ID in the membershipID URI parameter.

 @param membershipID Membership ID.
 @return
*/
func (s *TeamMembershipsService) DeleteTeamMembership(membershipID string) (*resty.Response, error) {

	path := "/team/memberships/{membershipId}"
	path = strings.Replace(path, "{"+"membershipId"+"}", fmt.Sprintf("%v", membershipID), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetTeamMembership Shows details for a team membership, by ID.
/* Shows details for a team membership, by ID.
Specify the team membership ID in the membershipID URI parameter.

 @param membershipID Membership ID.
 @return TeamMembership
*/
func (s *TeamMembershipsService) GetTeamMembership(membershipID string) (*TeamMembership, *resty.Response, error) {

	path := "/team/memberships/{membershipId}"
	path = strings.Replace(path, "{"+"membershipId"+"}", fmt.Sprintf("%v", membershipID), -1)

	response, err := RestyClient.R().
		SetResult(&TeamMembership{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TeamMembership)
	return result, response, err

}

// ListTeamMemberhipsQueryParams are the query params for the ListTeamMemberhips API Call
type ListTeamMemberhipsQueryParams struct {
	TeamID   string `url:"teamId,omitempty"` // Team ID.
	Max      int    `url:"max,omitempty"`    // Limit the maximum number of items in the response.
	Paginate bool   // Indicates if pagination is needed
}

// ListTeamMemberhips Lists all team memberships for a given team, specified by the teamID query parameter.
/* Lists all team memberships for a given team, specified by the teamID query parameter.
Use query parameters to filter the response.

 @param teamID Team ID.
 @param "max" (int) Limit the maximum number of items in the response.
 @param paginate (bool) indicates if pagination is needed
 @return TeamMemberships
*/
func (s *TeamMembershipsService) ListTeamMemberhips(queryParams *ListTeamMemberhipsQueryParams) (*TeamMemberships, *resty.Response, error) {

	path := "/team/memberships/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&TeamMemberships{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TeamMemberships)
	if queryParams.Paginate == true {
		items := teamMembershipsPagination(response.Header().Get("Link"), 0, 0)
		for _, teamMembership := range items.Items {
			result.AddTeamMembership(teamMembership)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := teamMembershipsPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, teamMembership := range items.Items {
				result.AddTeamMembership(teamMembership)
			}
		}
	}
	return result, response, err

}

// UpdateTeamMembership Updates a team membership, by ID.
/* Updates a team membership, by ID.
Specify the team membership ID in the membershipID URI parameter.

 @param membershipID Membership ID.
 @param teamMembershipUpdateRequest
 @return TeamMembership
*/
func (s *TeamMembershipsService) UpdateTeamMembership(membershipID string, teamMembershipUpdateRequest *TeamMembershipUpdateRequest) (*TeamMembership, *resty.Response, error) {

	path := "/team/memberships/{membershipId}"
	path = strings.Replace(path, "{"+"membershipId"+"}", fmt.Sprintf("%v", membershipID), -1)

	response, err := RestyClient.R().
		SetBody(teamMembershipUpdateRequest).
		SetResult(&TeamMembership{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TeamMembership)
	return result, response, err

}
