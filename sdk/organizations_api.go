package webexteams

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
	"gopkg.in/resty.v1"
)

// OrganizationsService is the service to communicate with the Organizations API endpoint
type OrganizationsService service

// Organization is the Organization definition
type Organization struct {
	ID          string    `json:"id,omitempty"`          // Organization ID.
	DisplayName string    `json:"displayName,omitempty"` // Organization Display Name.
	Created     time.Time `json:"created,omitempty"`     // Organization creation date/time.
}

// Organizations is the List of Organizations
type Organizations struct {
	Items []Organization `json:"items,omitempty"`
}

// AddOrganization is used to append a organization to a slice of organizations
func (organizations *Organizations) AddOrganization(item Organization) []Organization {
	organizations.Items = append(organizations.Items, item)
	return organizations.Items
}

func organizationLoop(linkHeader string) *Organizations {
	items := &Organizations{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := RestyClient.R().
				SetResult(&Organizations{}).
				Get(l.URI)

			if err != nil {
				fmt.Println("Error")
			}
			items = response.Result().(*Organizations)
			organizations := organizationLoop(response.Header().Get("Link"))
			for _, organization := range organizations.Items {
				items.AddOrganization(organization)
			}
		}
	}

	return items
}

// GetOrganization Shows details for an organization, by ID.
/* Shows details for an organization, by ID.
Specify the org ID in the orgID parameter in the URI.

 @param orgID Organization ID.
 @return Organization
*/
func (s *OrganizationsService) GetOrganization(orgID string) (*Organization, *resty.Response, error) {

	path := "/organizations/{orgId}"
	path = strings.Replace(path, "{"+"orgId"+"}", fmt.Sprintf("%v", orgID), -1)

	response, err := RestyClient.R().
		SetResult(&Organization{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Organization)
	return result, response, err

}

// ListOrganizationsQueryParams are the query params for the ListOrganizations API Call
type ListOrganizationsQueryParams struct {
	Max int `url:"max,omitempty"` // Limit the maximum number of items in the response.
}

// ListOrganizations List all organizations visible by your account. The results will not be paginated.
/* List all organizations visible by your account. The results will not be paginated.
@param "max" (int) Limit the maximum number of items in the response.
@return Organizations
*/
func (s *OrganizationsService) ListOrganizations(queryParams *ListOrganizationsQueryParams) (*Organizations, *resty.Response, error) {

	path := "/organizations/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Organizations{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Organizations)
	items := organizationLoop(response.Header().Get("Link"))

	for _, organization := range items.Items {
		result.AddOrganization(organization)
	}
	return result, response, err

}
