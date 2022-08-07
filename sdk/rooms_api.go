package webexteams

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
)

// RoomsService is the service to communicate with the Rooms API endpoint
type RoomsService service

// RoomCreateRequest is the Room Create Request Parameters
type RoomCreateRequest struct {
	Title              string `json:"title,omitempty"`              // A user-friendly name for the room.
	TeamID             string `json:"teamId,omitempty"`             // The ID for the team with which this room is associated.
	ClassificationId   string `json:"classificationId,omitempty"`   // The classificationId for the room.
	IsLocked           bool   `json:"isLocked,omitempty"`           // Set the space as locked/moderated and the creator becomes a moderator.
	IsAnnouncementOnly bool   `json:"isAnnouncementOnly,omitempty"` // Sets the space into Announcement Mode.
}

// RoomUpdateRequest is the Room Update Request Parameters
type RoomUpdateRequest struct {
	Title string `json:"title,omitempty"` // Room title.
}

// Room is the Room definition
type Room struct {
	ID                 string    `json:"id,omitempty"`                 // Room ID.
	Title              string    `json:"title,omitempty"`              // Room title.
	RoomType           string    `json:"type,omitempty"`               // Room type (group or direct).
	IsLocked           bool      `json:"isLocked,omitempty"`           // Room is moderated.
	TeamID             string    `json:"teamId,omitempty"`             // Room Team ID.
	CreatorID          string    `json:"creatorId,omitempty"`          // Room creator Person ID.
	LastActivity       time.Time `json:"lastActivity,omitempty"`       // Room last activity date/time.
	Created            time.Time `json:"created,omitempty"`            // Room creation date/time.
	OwnerId            string    `json:"ownerId,omitempty"`            // The ID of the organization which owns this room.
	ClassificationId   string    `json:"classificationId,omitempty"`   // Space classification id represents the space's current classification.
	IsAnnouncementOnly bool      `json:"isAnnouncementOnly,omitempty"` // Indicates when a space is in Announcement Mode where only moderators can post messages
	IsReadOnly         bool      `json:"isReadOnly,omitempty"`         // Indicates when a space is in Read Only Mode (direct space blocked by a compliance officer).
}

// Rooms is the List of Rooms
type Rooms struct {
	Items []Room `json:"items,omitempty"`
}

// AddRoom is used to append a room to a slice of rooms
func (rooms *Rooms) AddRoom(item Room) []Room {
	rooms.Items = append(rooms.Items, item)
	return rooms.Items
}

func (s *RoomsService) roomsPagination(linkHeader string, size, max int) *Rooms {
	items := &Rooms{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := s.client.R().
				SetResult(&Rooms{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Rooms)
			if size != 0 {
				size = size + len(items.Items)
				if size < max {
					rooms := s.roomsPagination(response.Header().Get("Link"), size, max)
					for _, room := range rooms.Items {
						items.AddRoom(room)
					}
				}
			} else {
				rooms := s.roomsPagination(response.Header().Get("Link"), size, max)
				for _, room := range rooms.Items {
					items.AddRoom(room)
				}
			}

		}
	}

	return items
}

// CreateRoom Creates a room. The authenticated user is automatically added as a member of the room.
/* Creates a room. The authenticated user is automatically added as a member of the room. See the Memberships API to learn how to add more people to the room.
To create a 1-to-1 room, use the Create Messages endpoint to send a message directly to another person by using the toPersonID or toPersonEmail parameters.

 @param roomCreateRequest
 @return Room
*/
func (s *RoomsService) CreateRoom(roomCreateRequest *RoomCreateRequest) (*Room, *resty.Response, error) {

	path := "/rooms/"

	response, err := s.client.R().
		SetBody(roomCreateRequest).
		SetResult(&Room{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Room)
	return result, response, err

}

// DeleteRoom Deletes a room, by ID. Deleted rooms cannot be recovered.
/* Deletes a room, by ID. Deleted rooms cannot be recovered.
Deleting a room that is part of a team will archive the room instead.
Specify the room ID in the roomID parameter in the URI

 @param roomID Room ID.
 @return
*/
func (s *RoomsService) DeleteRoom(roomID string) (*resty.Response, error) {

	path := "/rooms/{roomId}"
	path = strings.Replace(path, "{"+"roomId"+"}", fmt.Sprintf("%v", roomID), -1)

	response, err := s.client.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetRoom Shows details for a room, by ID.
/* Shows details for a room, by ID.
The title of the room for 1-to-1 rooms will be the display name of the other person.
Specify the room ID in the roomID parameter in the URI.

 @param roomID Room ID.
 @return Room
*/
func (s *RoomsService) GetRoom(roomID string) (*Room, *resty.Response, error) {

	path := "/rooms/{roomId}"
	path = strings.Replace(path, "{"+"roomId"+"}", fmt.Sprintf("%v", roomID), -1)

	response, err := s.client.R().
		SetResult(&Room{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Room)
	return result, response, err

}

// ListRoomsQueryParams are the query params for the ListRooms API Call
type ListRoomsQueryParams struct {
	TeamID   string `url:"teamId,omitempty"` // Limit the rooms to those associated with a team, by ID.
	RoomType string `url:"type,omitempty"`   // direct returns all 1-to-1 rooms. group returns all group rooms.
	SortBy   string `url:"sortBy,omitempty"` // Sort results by room ID (id), most recent activity (lastactivity), or most recently created (created).
	Max      int    `url:"max,omitempty"`    // Limit the maximum number of items in the response.
	Paginate bool   // Indicates if pagination is needed
}

// ListRooms List rooms.
/* List rooms.
The title of the room for 1-to-1 rooms will be the display name of the other person.
By default, lists rooms to which the authenticated user belongs.
Long result sets will be split into pages.

 @param "teamId" (string) Limit the rooms to those associated with a team, by ID.
 @param "type_" (string) direct returns all 1-to-1 rooms. group returns all group rooms.
 @param "sortBy" (string) Sort results by room ID (id), most recent activity (lastactivity), or most recently created (created).
 @param "max" (int) Limit the maximum number of items in the response.
 @param paginate (bool) indicates if pagination is needed
 @return Rooms
*/
func (s *RoomsService) ListRooms(queryParams *ListRoomsQueryParams) (*Rooms, *resty.Response, error) {

	path := "/rooms/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := s.client.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Rooms{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Rooms)
	if queryParams.Paginate {
		items := s.roomsPagination(response.Header().Get("Link"), 0, 0)
		for _, room := range items.Items {
			result.AddRoom(room)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := s.roomsPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, room := range items.Items {
				result.AddRoom(room)
			}
		}
	}

	return result, response, err

}

// UpdateRoom Updates details for a room, by ID.
/* Updates details for a room, by ID.
Specify the room ID in the roomID parameter in the URI.

 @param roomID Room ID.
 @param roomUpdateRequest
 @return Room
*/
func (s *RoomsService) UpdateRoom(roomID string, roomUpdateRequest *RoomUpdateRequest) (*Room, *resty.Response, error) {

	path := "/rooms/{roomId}"
	path = strings.Replace(path, "{"+"roomId"+"}", fmt.Sprintf("%v", roomID), -1)

	response, err := s.client.R().
		SetBody(roomUpdateRequest).
		SetResult(&Room{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Room)
	return result, response, err

}
