package webexteams

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
)

// PlacesService is the service to communicate with the Places API endpoint
type PlacesService service

// PlaceCreateRequest is the Place Create Request Parameters
type PlaceCreateRequest struct {
	Title  string `json:"title,omitempty"`  // A user-friendly name for the place.
	TeamID string `json:"teamId,omitempty"` // The ID for the team with which this place is associated.
}

// PlaceUpdateRequest is the Place Update Request Parameters
type PlaceUpdateRequest struct {
	Title string `json:"title,omitempty"` // Place title.
}

// Place is the Place definition
type Place struct {
	ID           string    `json:"id,omitempty"`           // Place ID.
	Title        string    `json:"title,omitempty"`        // Place title.
	PlaceType    string    `json:"type,omitempty"`         // Place type (group or direct).
	IsLocked     bool      `json:"isLocked,omitempty"`     // Place is moderated.
	TeamID       string    `json:"teamId,omitempty"`       // Place Team ID.
	CreatorID    string    `json:"creatorId,omitempty"`    // Place creator Person ID.
	LastActivity time.Time `json:"lastActivity,omitempty"` // Place last activity date/time.
	Created      time.Time `json:"created,omitempty"`      // Place creation date/time.
}

// Places is the List of Places
type Places struct {
	Items []Place `json:"items,omitempty"`
}

// AddPlace is used to append a place to a slice of places
func (places *Places) AddPlace(item Place) []Place {
	places.Items = append(places.Items, item)
	return places.Items
}

func (s *PlacesService) placesPagination(linkHeader string, size, max int) *Places {
	items := &Places{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := s.client.R().
				SetResult(&Places{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Places)
			if size != 0 {
				size = size + len(items.Items)
				if size < max {
					places := s.placesPagination(response.Header().Get("Link"), size, max)
					for _, place := range places.Items {
						items.AddPlace(place)
					}
				}
			} else {
				places := s.placesPagination(response.Header().Get("Link"), size, max)
				for _, place := range places.Items {
					items.AddPlace(place)
				}
			}

		}
	}

	return items
}

// CreatePlace Creates a place. The authenticated user is automatically added as a member of the place.
/* Creates a place. The authenticated user is automatically added as a member of the place. See the Memberships API to learn how to add more people to the place.
To create a 1-to-1 place, use the Create Messages endpoint to send a message directly to another person by using the toPersonID or toPersonEmail parameters.

 @param placeCreateRequest
 @return Place
*/
func (s *PlacesService) CreatePlace(placeCreateRequest *PlaceCreateRequest) (*Place, *resty.Response, error) {

	path := "/places/"

	response, err := s.client.R().
		SetBody(placeCreateRequest).
		SetResult(&Place{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Place)
	return result, response, err

}

// DeletePlace Deletes a place, by ID. Deleted places cannot be recovered.
/* Deletes a place, by ID. Deleted places cannot be recovered.
Deleting a place that is part of a team will archive the place instead.
Specify the place ID in the placeID parameter in the URI

 @param placeID Place ID.
 @return
*/
func (s *PlacesService) DeletePlace(placeID string) (*resty.Response, error) {

	path := "/places/{placeId}"
	path = strings.Replace(path, "{"+"placeId"+"}", fmt.Sprintf("%v", placeID), -1)

	response, err := s.client.R().
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetPlace Shows details for a place, by ID.
/* Shows details for a place, by ID.
The title of the place for 1-to-1 places will be the display name of the other person.
Specify the place ID in the placeID parameter in the URI.

 @param placeID Place ID.
 @return Place
*/
func (s *PlacesService) GetPlace(placeID string) (*Place, *resty.Response, error) {

	path := "/places/{placeId}"
	path = strings.Replace(path, "{"+"placeId"+"}", fmt.Sprintf("%v", placeID), -1)

	response, err := s.client.R().
		SetResult(&Place{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Place)
	return result, response, err

}

// ListPlacesQueryParams are the query params for the ListPlaces API Call
type ListPlacesQueryParams struct {
	TeamID    string `url:"teamId,omitempty"` // Limit the places to those associated with a team, by ID.
	PlaceType string `url:"type,omitempty"`   // direct returns all 1-to-1 places. group returns all group places.
	SortBy    string `url:"sortBy,omitempty"` // Sort results by place ID (id), most recent activity (lastactivity), or most recently created (created).
	Max       int    `url:"max,omitempty"`    // Limit the maximum number of items in the response.
	Paginate  bool   // Indicates if pagination is needed
}

// ListPlaces List places.
/* List places.
The title of the place for 1-to-1 places will be the display name of the other person.
By default, lists places to which the authenticated user belongs.
Long result sets will be split into pages.

 @param "teamId" (string) Limit the places to those associated with a team, by ID.
 @param "type_" (string) direct returns all 1-to-1 places. group returns all group places.
 @param "sortBy" (string) Sort results by place ID (id), most recent activity (lastactivity), or most recently created (created).
 @param "max" (int) Limit the maximum number of items in the response.
 @param paginate (bool) indicates if pagination is needed
 @return Places
*/
func (s *PlacesService) ListPlaces(queryParams *ListPlacesQueryParams) (*Places, *resty.Response, error) {

	path := "/places/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := s.client.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Places{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Places)
	if queryParams.Paginate {
		items := s.placesPagination(response.Header().Get("Link"), 0, 0)
		for _, place := range items.Items {
			result.AddPlace(place)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := s.placesPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, place := range items.Items {
				result.AddPlace(place)
			}
		}
	}

	return result, response, err

}

// UpdatePlace Updates details for a place, by ID.
/* Updates details for a place, by ID.
Specify the place ID in the placeID parameter in the URI.

 @param placeID Place ID.
 @param placeUpdateRequest
 @return Place
*/
func (s *PlacesService) UpdatePlace(placeID string, placeUpdateRequest *PlaceUpdateRequest) (*Place, *resty.Response, error) {

	path := "/places/{placeId}"
	path = strings.Replace(path, "{"+"placeId"+"}", fmt.Sprintf("%v", placeID), -1)

	response, err := s.client.R().
		SetBody(placeUpdateRequest).
		SetResult(&Place{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Place)
	return result, response, err

}
