package webexteams

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
)

// WebhooksService is the service to communicate with the Webhooks API endpoint
type WebhooksService service

// WebhookCreateRequest is the Webhook Create Request Parameters
type WebhookCreateRequest struct {
	Name      string `json:"name,omitempty"`      // Webhook name.
	TargetURL string `json:"targetUrl,omitempty"` // Webhook target URL.
	Resource  string `json:"resource,omitempty"`  // Webhook resource.
	Event     string `json:"event,omitempty"`     // Webhook event.
	Filter    string `json:"filter,omitempty"`    // Webhook filter.
	Secret    string `json:"secret,omitempty"`    // Webhook secret.
}

// WebhookUpdateRequest is the Update Webhook Request Parameters
type WebhookUpdateRequest struct {
	Name      string `json:"name,omitempty"`      // Webhook name.
	TargetURL string `json:"targetUrl,omitempty"` // Webhook target URL.
	Secret    string `json:"secret,omitempty"`    // Webhook secret.
	Status    string `json:"status,omitempty"`    // Webhook status.
}

// Webhook is the Webhook definition
type Webhook struct {
	ID        string    `json:"id,omitempty"`        // Webhook ID.
	Name      string    `json:"name,omitempty"`      // Webhook name.
	TargetURL string    `json:"targetUrl,omitempty"` // Webhook target URL.
	Resource  string    `json:"resource,omitempty"`  // Webhook resource.
	Event     string    `json:"event,omitempty"`     // Webhook event.
	OrgID     string    `json:"orgId,omitempty"`     // Webhook organization ID.
	CreatedBy string    `json:"createdBy,omitempty"` // Webhook created by Person ID.
	AppID     string    `json:"appId,omitempty"`     // Webhook application ID.
	OwnedBy   string    `json:"ownedBy,omitempty"`   // Webhook owner Person ID.
	Filter    string    `json:"filter,omitempty"`    // Webhook filter.
	Status    string    `json:"status,omitempty"`    // Webhook status.
	Secret    string    `json:"secret,omitempty"`    // Webhook secret.
	Created   time.Time `json:"created,omitempty"`   // Webhook creation date/time.
}

// WebhookRequestData is the Webhook trigger request
type WebhookRequestData struct {
	ID          string    `json:"id,omitempty"`
	RoomID      string    `json:"roomId,omitempty"`
	PersonID    string    `json:"personId,omitempty"`
	PersonEmail string    `json:"personEmail,omitempty"`
	Created     time.Time `json:"created,omitempty"`
}

// WebhookRequest is the Webhook trigger request
type WebhookRequest struct {
	ID        string             `json:"id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Resource  string             `json:"resource,omitempty"`
	Event     string             `json:"event,omitempty"`
	Filter    string             `json:"filter,omitempty"`
	OrgID     string             `json:"orgId,omitempty"`
	CreatedBy string             `json:"createdBy,omitempty"`
	AppID     string             `json:"appId,omitempty"`
	OwnedBy   string             `json:"ownedBy,omitempty"`
	Status    string             `json:"status,omitempty"`
	ActorID   string             `json:"actorId,omitempty"`
	Data      WebhookRequestData `json:"data,omitempty"`
}

// Webhooks is the List of Webhooks
type Webhooks struct {
	Items []Webhook `json:"items,omitempty"`
}

// AddWebhook is used to append a webhook to a slice of webhooks
func (webhooks *Webhooks) AddWebhook(item Webhook) []Webhook {
	webhooks.Items = append(webhooks.Items, item)
	return webhooks.Items
}

func webhooksPagination(linkHeader string, size, max int) *Webhooks {
	items := &Webhooks{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := RestyClient.R().
				SetResult(&Webhooks{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Webhooks)
			if size != 0 {
				size = size + len(items.Items)
				if size < max {
					webhooks := webhooksPagination(response.Header().Get("Link"), size, max)
					for _, webhook := range webhooks.Items {
						items.AddWebhook(webhook)
					}
				}
			} else {
				webhooks := webhooksPagination(response.Header().Get("Link"), size, max)
				for _, webhook := range webhooks.Items {
					items.AddWebhook(webhook)
				}
			}

		}
	}

	return items
}

// CreateWebhook Creates a webhook.
/* Creates a webhook.
@param webhookCreateRequest
@return Webhook
*/
func (s *WebhooksService) CreateWebhook(webhookCreateRequest *WebhookCreateRequest) (*Webhook, *resty.Response, error) {

	path := "/webhooks/"

	response, err := RestyClient.R().
		SetBody(webhookCreateRequest).
		SetResult(&Webhook{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Webhook)
	return result, response, err

}

// DeleteWebhook Deletes a webhook, by ID.
/* Deletes a webhook, by ID.
Specify the webhook ID in the webhookID parameter in the URI.

 @param webhookID Webhook ID.
 @return
*/
func (s *WebhooksService) DeleteWebhook(webhookID string) (*resty.Response, error) {

	path := "/webhooks/{webhookId}"
	path = strings.Replace(path, "{"+"webhookId"+"}", fmt.Sprintf("%v", webhookID), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetWebhook Shows details for a webhook, by ID.
/* Shows details for a webhook, by ID.
Specify the webhook ID in the webhookID parameter in the URI.

 @param webhookID Webhook ID.
 @return Webhook
*/
func (s *WebhooksService) GetWebhook(webhookID string) (*Webhook, *resty.Response, error) {

	path := "/webhooks/{webhookId}"
	path = strings.Replace(path, "{"+"webhookId"+"}", fmt.Sprintf("%v", webhookID), -1)

	response, err := RestyClient.R().
		SetResult(&Webhook{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Webhook)
	return result, response, err

}

// ListWebhooksQueryParams are the query params for the ListWebhooks API Call
type ListWebhooksQueryParams struct {
	Max      int  `url:"max,omitempty"` // Limit the maximum number of items in the response.
	Paginate bool // Indicates if pagination is needed
}

// ListWebhooks Lists all of your webhooks.
/* Lists all of your webhooks.
@param "max" (int) Limit the maximum number of items in the response.
@param paginate (bool) indicates if pagination is needed
@return Webhooks
*/
func (s *WebhooksService) ListWebhooks(queryParams *ListWebhooksQueryParams) (*Webhooks, *resty.Response, error) {

	path := "/webhooks/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Webhooks{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Webhooks)
	if queryParams.Paginate == true {
		items := webhooksPagination(response.Header().Get("Link"), 0, 0)
		for _, webhook := range items.Items {
			result.AddWebhook(webhook)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := webhooksPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, webhook := range items.Items {
				result.AddWebhook(webhook)
			}
		}
	}

	return result, response, err

}

// UpdateWebhook Updates a webhook, by ID.
/* Updates a webhook, by ID.
Specify the webhook ID in the webhookID parameter in the URI.

 @param webhookID Webhook ID.
 @param webhookUpdateRequest
 @return Webhook
*/
func (s *WebhooksService) UpdateWebhook(webhookID string, webhookUpdateRequest *WebhookUpdateRequest) (*Webhook, *resty.Response, error) {

	path := "/webhooks/{webhookId}"
	path = strings.Replace(path, "{"+"webhookId"+"}", fmt.Sprintf("%v", webhookID), -1)

	response, err := RestyClient.R().
		SetBody(webhookUpdateRequest).
		SetResult(&Webhook{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Webhook)
	return result, response, err

}
