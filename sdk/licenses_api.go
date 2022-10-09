package webexteams

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
)

// LicensesService is the service to communicate with the Licenses API endpoint
type LicensesService service

// License is the License definition
type License struct {
	ID            string `json:"id,omitempty"`            // License ID.
	Name          string `json:"name,omitempty"`          // License Display Name.
	TotalUnits    int    `json:"totalUnits,omitempty"`    // License quantity total.
	ConsumedUnits int    `json:"consumedUnits,omitempty"` // License quantity consumed.
}

// Licenses is the List of Licenses
type Licenses struct {
	Items []License `json:"items,omitempty"`
}

// AddLicense is used to append a license to a slice of licenses
func (licenses *Licenses) AddLicense(item License) []License {
	licenses.Items = append(licenses.Items, item)
	return licenses.Items
}

func (s *LicensesService) licensesPagination(linkHeader string, size, max int) *Licenses {
	items := &Licenses{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := s.client.R().
				SetResult(&Licenses{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Licenses)
			size = size + len(items.Items)
			if max < 0 || size < max {
				licenses := s.licensesPagination(response.Header().Get("Link"), size, max)
				for _, license := range licenses.Items {
					items.AddLicense(license)
				}
			}

		}
	}

	return items
}

// GetLicense Shows details for a license, by ID.
/* Shows details for a license, by ID.
Specify the license ID in the licenseID parameter in the URI.

 @param licenseID License ID.
 @return License
*/
func (s *LicensesService) GetLicense(licenseID string) (*License, *resty.Response, error) {

	path := "/licenses/{licenseId}"
	path = strings.Replace(path, "{"+"licenseId"+"}", fmt.Sprintf("%v", licenseID), -1)

	response, err := s.client.R().
		SetResult(&License{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*License)
	return result, response, err

}

// ListLicensesQueryParams are the query params for the ListLicenses API Call
type ListLicensesQueryParams struct {
	Max       int `url:"max,omitempty"` // Limit the maximum number of items in the response. Negative value will list all items (use this carefully).
	RequestBy int `url:"-"`             // number of item to retrieve by requests (Max if let at 0)
}

// ListLicenses List all licenses for a given organization.
/* List all licenses for a given organization.
If no orgID is specified, the default is the organization of the authenticated user.

 @param "max" (int) Limit the maximum number of items in the response. Negative value will list all items (use this carefully).
 @param "requestBy" (int) Number of items by request
 @return Licenses
*/
func (s *LicensesService) ListLicenses(queryParams *ListLicensesQueryParams) (*Licenses, *resty.Response, error) {

	path := "/licenses/"

	max := queryParams.Max

	if queryParams.RequestBy > 0 {
		queryParams.Max = queryParams.RequestBy
	} else if queryParams.Max < 0 {
		queryParams.Max = 0
	}

	queryParamsString, _ := query.Values(queryParams)

	response, err := s.client.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Licenses{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Licenses)

	if max < 0 || len(result.Items) < max {
		items := s.licensesPagination(response.Header().Get("Link"), len(result.Items), max)
		for _, license := range items.Items {
			result.AddLicense(license)
		}
	}
	return result, response, err

}
