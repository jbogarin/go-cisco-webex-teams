package webexteams

import (
	"os"

	"github.com/go-resty/resty/v2"
)

// RestyClient is the REST Client

const apiURL = "https://webexapis.com/v1"

// Client manages communication with the Webex Teams API API v1.0.0
// In most cases there should be only one, shared, APIClient.
type Client struct {
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	AdminAuditEvents  *AdminAuditEventsService
	AttachmentActions *AttachmentActionsService
	Contents          *ContentsService
	Events            *EventsService
	Devices           *DevicesService
	Licenses          *LicensesService
	Memberships       *MembershipsService
	Messages          *MessagesService
	Organizations     *OrganizationsService
	People            *PeopleService
	Recordings        *RecordingsService
	Roles             *RolesService
	Rooms             *RoomsService
	TeamMemberships   *TeamMembershipsService
	Teams             *TeamsService
	Webhooks          *WebhooksService
}

type service struct {
	client *resty.Client
}

// SetAuthToken defines the Authorization token sent in the request
func (s *Client) SetAuthToken(accessToken string) {
	s.common.client.SetAuthToken(accessToken)
}

// NewClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewClient() *Client {
	client := resty.New()
	c := &Client{}
	c.common.client = client
	c.common.client.SetHostURL(apiURL)
	if os.Getenv("WEBEX_TEAMS_ACCESS_TOKEN") != "" {
		c.common.client.SetAuthToken(os.Getenv("WEBEX_TEAMS_ACCESS_TOKEN"))
	}

	// API Services
	c.AdminAuditEvents = (*AdminAuditEventsService)(&c.common)
	c.AttachmentActions = (*AttachmentActionsService)(&c.common)
	c.Contents = (*ContentsService)(&c.common)
	c.Events = (*EventsService)(&c.common)
	c.Devices = (*DevicesService)(&c.common)
	c.Licenses = (*LicensesService)(&c.common)
	c.Memberships = (*MembershipsService)(&c.common)
	c.Messages = (*MessagesService)(&c.common)
	c.Organizations = (*OrganizationsService)(&c.common)
	c.People = (*PeopleService)(&c.common)
	c.Recordings = (*RecordingsService)(&c.common)
	c.Roles = (*RolesService)(&c.common)
	c.Rooms = (*RoomsService)(&c.common)
	c.TeamMemberships = (*TeamMembershipsService)(&c.common)
	c.Teams = (*TeamsService)(&c.common)
	c.Webhooks = (*WebhooksService)(&c.common)

	return c
}

// Error indicates an error from the invocation of a Webex API. See
// the following documentation for error context: https://developer.webex.com/docs/api/basics#api-errors.
type Error struct{}
