package webexteams

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
)

// TeamsService is the service to communicate with the Teams API endpoint
type TeamsService service

// Teams is the List of Teams
type Teams struct {
	Items []Team `json:"items,omitempty"`
}

// Team is the Team definition
type Team struct {
	ID        string    `json:"id,omitempty"`        // Team ID.
	Name      string    `json:"name,omitempty"`      // Team Name.
	CreatorID string    `json:"creatorId,omitempty"` // Team creator ID.
	Created   time.Time `json:"created,omitempty"`   // Team creation date/time.
}

// TeamUpdateRequest is the Team Update Request Object
type TeamUpdateRequest struct {
	Name string `json:"name,omitempty"` // Team name.
}

// TeamCreateRequest is the Team Create Request Parameters
type TeamCreateRequest struct {
	Name string `json:"name,omitempty"` // Team name.
}

// AddTeam is used to append a team to a slice of teams
func (teams *Teams) AddTeam(item Team) []Team {
	teams.Items = append(teams.Items, item)
	return teams.Items
}

func (s *TeamsService) teamsPagination(linkHeader string, size, max int) *Teams {
	items := &Teams{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := s.client.R().
				SetResult(&Teams{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Teams)
			size = size + len(items.Items)
			if max < 0 || size < max {
				teams := s.teamsPagination(response.Header().Get("Link"), size, max)
				for _, team := range teams.Items {
					items.AddTeam(team)
				}
			}
		}
	}

	return items
}

// CreateTeam Creates a team.
/* Creates a team. The authenticated user is automatically added as a member of the team.
See the Team Memberships API to learn how to add more people to the team.

 @param teamCreateRequest
 @return Team
*/
func (s *TeamsService) CreateTeam(teamCreateRequest *TeamCreateRequest) (*Team, *resty.Response, error) {

	path := "/teams/"

	response, err := s.client.R().
		SetBody(teamCreateRequest).
		SetResult(&Team{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Team)
	return result, response, err

}

// DeleteTeam Deletes a team, by ID.
/* Deletes a team, by ID.
Specify the team ID in the teamID parameter in the URI.

 @param teamID Team ID.
 @return
*/
func (s *TeamsService) DeleteTeam(teamID string) (*resty.Response, error) {

	path := "/teams/{teamId}"
	path = strings.Replace(path, "{"+"teamId"+"}", fmt.Sprintf("%v", teamID), -1)

	response, err := s.client.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetTeam Shows details for a team, by ID.
/* Shows details for a team, by ID.
Specify the team ID in the teamID parameter in the URI.

 @param teamID Team ID.
 @return Team
*/
func (s *TeamsService) GetTeam(teamID string) (*Team, *resty.Response, error) {

	path := "/teams/{teamId}"
	path = strings.Replace(path, "{"+"teamId"+"}", fmt.Sprintf("%v", teamID), -1)

	response, err := s.client.R().
		SetResult(&Team{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Team)
	return result, response, err

}

// ListTeamsQueryParams are the query params for the ListTeams API Call
type ListTeamsQueryParams struct {
	Max       int `url:"max,omitempty"` // Limit the maximum number of items in the response. Negative value will list all items (use this carefully).
	RequestBy int `url:"-"`             // Number of items to retrieve by requests (Max if let at 0)
}

// ListTeams Lists teams to which the authenticated user belongs.
/* Lists teams to which the authenticated user belongs.
@param "max" (int) Limit the maximum number of items in the response. Negative value will list all items (use this carefully).
@param "requestBy" (int) Number of items by request
@return Teams
*/
func (s *TeamsService) ListTeams(queryParams *ListTeamsQueryParams) (*Teams, *resty.Response, error) {

	path := "/teams/"

	max := queryParams.Max

	if queryParams.RequestBy > 0 {
		queryParams.Max = queryParams.RequestBy
	} else if queryParams.Max < 0 {
		queryParams.Max = 0
	}

	queryParamsString, _ := query.Values(queryParams)

	response, err := s.client.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Teams{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Teams)

	if max < 0 || len(result.Items) < max {
		items := s.teamsPagination(response.Header().Get("Link"), len(result.Items), max)
		for _, team := range items.Items {
			result.AddTeam(team)
		}
	}

	return result, response, err

}

// UpdateTeam Updates details for a team, by ID.
/* Updates details for a team, by ID.
Specify the team ID in the teamID parameter in the URI.

 @param teamID Team ID.
 @param teamUpdateRequest
 @return Team
*/
func (s *TeamsService) UpdateTeam(teamID string, teamUpdateRequest *TeamUpdateRequest) (*Team, *resty.Response, error) {

	path := "/teams/{teamId}"
	path = strings.Replace(path, "{"+"teamId"+"}", fmt.Sprintf("%v", teamID), -1)

	response, err := s.client.R().
		SetBody(teamUpdateRequest).
		SetResult(&Team{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Team)
	return result, response, err

}
