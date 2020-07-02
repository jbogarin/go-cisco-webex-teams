package webexteams

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
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

// AddRole is used to append a role to a slice of roles
func (roles *Roles) AddRole(item Role) []Role {
	roles.Items = append(roles.Items, item)
	return roles.Items
}

func rolesPagination(linkHeader string, size, max int) *Roles {
	items := &Roles{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := RestyClient.R().
				SetResult(&Roles{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Roles)
			if size != 0 {
				size = size + len(items.Items)
				if size < max {
					roles := rolesPagination(response.Header().Get("Link"), size, max)
					for _, role := range roles.Items {
						items.AddRole(role)
					}
				}
			} else {
				roles := rolesPagination(response.Header().Get("Link"), size, max)
				for _, role := range roles.Items {
					items.AddRole(role)
				}
			}

		}
	}

	return items
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
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Role)
	return result, response, err

}

// RolesListQueryParams are the query params for the GetRoles API Call
type RolesListQueryParams struct {
	Max      int  `url:"max,omitempty"` // Limit the maximum number of items in the response.
	Paginate bool // Indicates if pagination is needed
}

// ListRoles List all roles.
/* List all roles.
@param "max" (int) Limit the maximum number of items in the response.
@param paginate (bool) indicates if pagination is needed
@return Roles
*/
func (s *RolesService) ListRoles(queryParams *RolesListQueryParams) (*Roles, *resty.Response, error) {

	path := "/roles/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Roles{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Roles)
	if queryParams.Paginate == true {
		items := rolesPagination(response.Header().Get("Link"), 0, 0)
		for _, role := range items.Items {
			result.AddRole(role)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := rolesPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, role := range items.Items {
				result.AddRole(role)
			}
		}
	}
	return result, response, err

}
