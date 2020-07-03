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

func licensesPagination(linkHeader string, size, max int) *Licenses {
	items := &Licenses{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := RestyClient.R().
				SetResult(&Licenses{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Licenses)
			if size != 0 {
				size = size + len(items.Items)
				if size < max {
					licenses := licensesPagination(response.Header().Get("Link"), size, max)
					for _, license := range licenses.Items {
						items.AddLicense(license)
					}
				}
			} else {
				licenses := licensesPagination(response.Header().Get("Link"), size, max)
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

	response, err := RestyClient.R().
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
	Max      int  `url:"max,omitempty"` // Limit the maximum number of items in the response.
	Paginate bool // Indicates if pagination is needed
}

// ListLicenses List all licenses for a given organization.
/* List all licenses for a given organization.
If no orgID is specified, the default is the organization of the authenticated user.

 @param "max" (int) Limit the maximum number of items in the response.
 @param paginate (bool) indicates if pagination is needed
 @return Licenses
*/
func (s *LicensesService) ListLicenses(queryParams *ListLicensesQueryParams) (*Licenses, *resty.Response, error) {

	path := "/licenses/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Licenses{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Licenses)
	if queryParams.Paginate == true {
		items := licensesPagination(response.Header().Get("Link"), 0, 0)
		for _, license := range items.Items {
			result.AddLicense(license)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := licensesPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, license := range items.Items {
				result.AddLicense(license)
			}
		}
	}
	return result, response, err

}
