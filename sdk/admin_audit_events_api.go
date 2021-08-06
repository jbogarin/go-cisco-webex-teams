package webexteams

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// AdminAuditEventsService is the service to communicate with the AdminAuditEvents API endpoint
type AdminAuditEventsService service

// AuditEvents is the AuditEvents definition
type AuditEvents struct {
	Items []struct {
		Created    string `json:"created,omitempty"`    // The date and time the event took place.
		ActorOrgID string `json:"actorOrgId,omitempty"` // The orgId of the person who made the change.
		ID         string `json:"id,omitempty"`         // A unique identifier for the event.
		ActorID    string `json:"actorId,omitempty"`    // The personId of the person who made the change.
		Data       struct {
			ActorOrgName     string   `json:"actorOrgName,omitempty"`     // The display name of the organization.
			TargetName       string   `json:"targetName,omitempty"`       // The name of the resource being acted upon.
			EventDescription string   `json:"eventDescription,omitempty"` // A description for the event.
			ActorName        string   `json:"actorName,omitempty"`        // The name of the person who performed the action.
			ActorEmail       string   `json:"actorEmail,omitempty"`       // The email of the person who performed the action.
			AdminRoles       []string `json:"adminRoles,omitempty"`       // Admin roles for the person.
			TrackingID       string   `json:"trackingId,omitempty"`       // A tracking identifier for the event.
			TargetType       string   `json:"targetType,omitempty"`       // The type of resource changed by the event.
			TargetID         string   `json:"targetId,omitempty"`         // The identifier for the resource changed by the event.
			EventCategory    string   `json:"eventCategory,omitempty"`    // The category of resource changed by the event.
			ActorUserAgent   string   `json:"actorUserAgent,omitempty"`   // The browser user agent of the person who performed the action.
			ActorIP          string   `json:"actorIp,omitempty"`          // The IP address of the person who performed the action.
			TargetOrgID      string   `json:"targetOrgId,omitempty"`      // The orgId of the organization.
			ActionText       string   `json:"actionText,omitempty"`       // A more detailed description of the change made by the person.
			TargetOrgName    string   `json:"targetOrgName,omitempty"`    // The name of the organization being acted upon.
		} `json:"data,omitempty"` // data
	} `json:"items,omitempty"` // items
}

// ListAdminAuditEventsQueryParams defines the query parameters for this request
type ListAdminAuditEventsQueryParams struct {
	OrgID   string `url:"orgId,omitempty"`   // List events in this organization, by ID
	From    string `url:"from,omitempty"`    // List events which occurred after a specific date and tim
	To      string `url:"to,omitempty"`      // List events which occurred before a specific date and time
	ActorID string `url:"actorId,omitempty"` // List events performed by this person, by ID
	Max     int    `url:"max,omitempty"`     // Limit the maximum number of events in the response. The maximum value is 200 Default: 100
	Offset  int    `url:"offset,omitempty"`  // Offset from the first result that you want to fetch. Default: 0
}

// ListAdminAuditEvents List Admin Audit Events
/* List admin audit events in your organization. Several query parameters are available to filter the response.

@param orgId List events in this organization, by ID
@param from List events which occurred after a specific date and tim
@param to List events which occurred before a specific date and time
@param actorId List events performed by this person, by ID
@param max Limit the maximum number of events in the response. The maximum value is 200 Default: 100
@param offset Offset from the first result that you want to fetch. Default: 0
*/
func (s *AdminAuditEventsService) ListAdminAuditEvents(listAdminAuditEventsQueryParams *ListAdminAuditEventsQueryParams) (*AuditEvents, *resty.Response, error) {

	path := "/adminAudit/events"

	queryString, _ := query.Values(listAdminAuditEventsQueryParams)

	response, err := s.client.R().
		SetQueryString(queryString.Encode()).
		SetResult(&AuditEvents{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AuditEvents)
	return result, response, err

}
