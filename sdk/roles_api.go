package webexteams

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"gopkg.in/resty.v1"
)

// RolesService is the service to communicate with the Roles API endpoint
type RolesService service

// Role is the Role definition
type Role struct {
	ID   string `json:"id,omitempty"`   // Role ID.
	Name string `json:"name,omitempty"` // Role Display Name.
}

// Roles is the List of Roles
type Roles struct {
	Items []Role `json:"items,omitempty"`
}

// GetRole Shows details for a role, by ID.
/* Shows details for a role, by ID.
Specify the role ID in the roleID parameter in the URI.

 @param roleID Role ID.
 @return Role
*/
func (s *RolesService) GetRole(roleID string) (*Role, *resty.Response, error) {

	path := "/roles/{roleId}"
	path = strings.Replace(path, "{"+"roleId"+"}", fmt.Sprintf("%v", roleID), -1)

	response, err := RestyClient.R().
		SetResult(&Role{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Role)
	return result, response, err

}

// RolesListQueryParams are the query params for the GetRoles API Call
type RolesListQueryParams struct {
	Max int `url:"max,omitempty"` // Limit the maximum number of items in the response.
}

// ListRoles List all roles.
/* List all roles.
@param "max" (int) Limit the maximum number of items in the response.
@return Roles
*/
func (s *RolesService) ListRoles(queryParams *RolesListQueryParams) (*Roles, *resty.Response, error) {

	path := "/roles/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Roles{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Roles)
	return result, response, err

}
