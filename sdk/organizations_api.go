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

func (s *OrganizationsService) organizationsPagination(linkHeader string, size, max int) *Organizations {
	items := &Organizations{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := s.client.R().
				SetResult(&Organizations{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Organizations)
			size = size + len(items.Items)
			if max < 0 || size < max {
				organizations := s.organizationsPagination(response.Header().Get("Link"), size, max)
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

	response, err := s.client.R().
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
	Max       int `url:"max,omitempty"` // Limit the maximum number of items in the response. Negative value will list all items (use this carefully).
	RequestBy int `url:"-"`             // Number of items to retrieve by requests (Max if let at 0)
}

// ListOrganizations List all organizations visible by your account. The results will not be paginated.
/* List all organizations visible by your account. The results will not be paginated.
@param "max" (int) Limit the maximum number of items in the response. Negative value will list all items (use this carefully).
@param "requestBy" (int) Number of items by request
@return Organizations
*/
func (s *OrganizationsService) ListOrganizations(queryParams *ListOrganizationsQueryParams) (*Organizations, *resty.Response, error) {

	path := "/organizations/"

	max := queryParams.Max

	if queryParams.RequestBy > 0 {
		queryParams.Max = queryParams.RequestBy
	} else if queryParams.Max < 0 {
		queryParams.Max = 0
	}

	queryParamsString, _ := query.Values(queryParams)

	response, err := s.client.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Organizations{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Organizations)

	if max < 0 || len(result.Items) < max {
		items := s.organizationsPagination(response.Header().Get("Link"), len(result.Items), max)
		for _, organization := range items.Items {
			result.AddOrganization(organization)
		}
	}
	return result, response, err

}
