package webexteams

import (
	"fmt"
	"strings"

	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// EventsService is the service to communicate with the Events API endpoint
type EventsService service

// Event is the Event definition
type Event struct {
	ID         string    `json:"id,omitempty"`         // A unique identifier for the event.
	Resource   string    `json:"resource,omitempty"`   // The type of resource in the event.
	Event      string    `json:"event,omitempty"`      // The type of event that was triggered.
	AppID      string    `json:"appId,omitempty"`      // The ID of the application for the event.
	ActorID    string    `json:"actorId,omitempty"`    // The personId of the person who made the change.
	OrgID      string    `json:"orgId,omitempty"`      // The ID of the organization for the event.
	Created    time.Time `json:"created,omitempty"`    // The date and time of the event.
	ActorOrgID string    `json:"actorOrgId,omitempty"` // The orgId of the person who made the change.
	Data       struct {
		ID          string `json:"id,omitempty"`          // Action ID.
		RoomID      string `json:"roomId,omitempty"`      // Room ID where the event happened.
		RoomType    string `json:"roomType,omitempty"`    // Room type where the event happened.
		Text        string `json:"text,omitempty"`        // Text related to the event, in the case of a message.
		PersonID    string `json:"personId,omitempty"`    // Person ID of the user who triggered the event.
		PersonEmail string `json:"personEmail,omitempty"` // Person Email of the user who triggered the event.
		Created     string `json:"created,omitempty"`     // The date and time of the event.
		Type        string `json:"type,omitempty"`        // The type of event.
	} `json:"data,omitempty"` // data
}

// Events is the Events definition
type Events struct {
	Items []Event `json:"items,omitempty"` //
}

// ListEventsQueryParams defines the query parameters for this request
type ListEventsQueryParams struct {
	Resource string `url:"resource,omitempty"` // List events with a specific resource type. Possible values: messages, memberships, tabs, rooms, attachmentActions
	Type     string `url:"type,omitempty"`     // List events with a specific event type. Possible values: created, updated, deleted
	ActorID  string `url:"actorId,omitempty"`  // List events performed by this person, by ID.
	From     string `url:"from,omitempty"`     // List events which occurred after a specific date and time.
	To       string `url:"to,omitempty"`       // List events which occurred before a specific date and time. If unspecified or set to a time in the future, lists events up to the present.
	Max      int    `url:"max,omitempty"`      // Limit the maximum number of events in the response. The maximum value is 200 Default: 100
}

// ListEvents List Events
/* List events in your organization. Several query parameters are available to filter the response.
Long result sets will be split into pages.

@param resource List events with a specific resource type. Possible values: messages, memberships, tabs, rooms, attachmentActions
@param type List events with a specific event type. Possible values: created, updated, deleted
@param actorId List events performed by this person, by ID.
@param from List events which occurred after a specific date and time.
@param to List events which occurred before a specific date and time. If unspecified or set to a time in the future, lists events up to the present.
@param max Limit the maximum number of events in the response. The maximum value is 200 Default: 100
*/
func (s *EventsService) ListEvents(listEventsQueryParams *ListEventsQueryParams) (*Events, *resty.Response, error) {

	path := "/events"

	queryString, _ := query.Values(listEventsQueryParams)

	response, err := s.client.R().
		SetQueryString(queryString.Encode()).
		SetResult(&Events{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Events)
	return result, response, err

}

// GetEvent Get Event Details
/* Shows details for an event, by event ID.
Specify the event ID in the eventId parameter in the URI.

@param eventId The unique identifier for the event.
*/
func (s *EventsService) GetEvent(eventID string) (*Event, *resty.Response, error) {

	path := "/events/{eventId}"
	path = strings.Replace(path, "{"+"eventId"+"}", fmt.Sprintf("%v", eventID), -1)

	response, err := s.client.R().
		SetResult(&Event{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Event)
	return result, response, err

}
