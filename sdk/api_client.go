package webexteams

import (
	"os"

	"github.com/go-resty/resty"
)

// RestyClient is the REST Client
var RestyClient *resty.Client

const apiURL = "https://api.ciscospark.com/v1"

// Client manages communication with the Webex Teams API API v1.0.0
// In most cases there should be only one, shared, APIClient.
type Client struct {
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	Contents        *ContentsService
	Licenses        *LicensesService
	Memberships     *MembershipsService
	Messages        *MessagesService
	Organizations   *OrganizationsService
	People          *PeopleService
	Roles           *RolesService
	Rooms           *RoomsService
	TeamMemberships *TeamMembershipsService
	Teams           *TeamsService
	Webhooks        *WebhooksService
}

type service struct {
	client *Client
}

// SetAuthToken defines the Authorization token sent in the request
func (s *Client) SetAuthToken(accessToken string) {
	RestyClient.SetAuthToken(accessToken)
}

// NewClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewClient() *Client {
	client := resty.New()
	c := &Client{}
	RestyClient = client
	RestyClient.SetHostURL(apiURL)
	if os.Getenv("WEBEX_TEAMS_ACCESS_TOKEN") != "" {
		RestyClient.SetAuthToken(os.Getenv("WEBEX_TEAMS_ACCESS_TOKEN"))
	}

	// API Services
	c.Contents = (*ContentsService)(&c.common)
	c.Licenses = (*LicensesService)(&c.common)
	c.Memberships = (*MembershipsService)(&c.common)
	c.Messages = (*MessagesService)(&c.common)
	c.Organizations = (*OrganizationsService)(&c.common)
	c.People = (*PeopleService)(&c.common)
	c.Roles = (*RolesService)(&c.common)
	c.Rooms = (*RoomsService)(&c.common)
	c.TeamMemberships = (*TeamMembershipsService)(&c.common)
	c.Teams = (*TeamsService)(&c.common)
	c.Webhooks = (*WebhooksService)(&c.common)

	return c
}
