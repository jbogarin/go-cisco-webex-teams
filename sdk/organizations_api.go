package webexteams

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
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

func organizationsPagination(linkHeader string, size, max int) *Organizations {
	items := &Organizations{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := RestyClient.R().
				SetResult(&Organizations{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Organizations)
			if size != 0 {
				size = size + len(items.Items)
				if size < max {
					organizations := organizationsPagination(response.Header().Get("Link"), size, max)
					for _, organization := range organizations.Items {
						items.AddOrganization(organization)
					}
				}
			} else {
				organizations := organizationsPagination(response.Header().Get("Link"), size, max)
				for _, organization := range organizations.Items {
					items.AddOrganization(organization)
				}
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
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Organization)
	return result, response, err

}

// ListOrganizationsQueryParams are the query params for the ListOrganizations API Call
type ListOrganizationsQueryParams struct {
	Max      int  `url:"max,omitempty"` // Limit the maximum number of items in the response.
	Paginate bool // Indicates if pagination is needed
}

// ListOrganizations List all organizations visible by your account. The results will not be paginated.
/* List all organizations visible by your account. The results will not be paginated.
@param "max" (int) Limit the maximum number of items in the response.
@param paginate (bool) indicates if pagination is needed
@return Organizations
*/
func (s *OrganizationsService) ListOrganizations(queryParams *ListOrganizationsQueryParams) (*Organizations, *resty.Response, error) {

	path := "/organizations/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Organizations{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Organizations)
	if queryParams.Paginate == true {
		items := organizationsPagination(response.Header().Get("Link"), 0, 0)
		for _, organization := range items.Items {
			result.AddOrganization(organization)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := organizationsPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, organization := range items.Items {
				result.AddOrganization(organization)
			}
		}
	}
	return result, response, err

}
