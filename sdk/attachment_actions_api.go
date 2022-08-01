package webexteams

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

// AttachmentActionsService manages the interactions with the attachment actions API methods.
type AttachmentActionsService service

// AttachmentActionCreateRequest create request for AttachmentAction
type AttachmentActionCreateRequest struct {
	Type      string                 `json:"type"`             // The type of action.
	MessageID string                 `json:"messageId"`        // The unique identifier of the message.
	Inputs    map[string]interface{} `json:"inputs,omitempty"` // Action inputs
}

// AttachmentAction represents a Webex Teams attachment action.
type AttachmentAction struct {
	ID        string                 `json:"id,omitempty"`        // The unique identifier of the attachment action.
	Type      string                 `json:"type,omitempty"`      // The type of action.
	MessageID string                 `json:"messageId,omitempty"` // The ID of the message to which the attachment action belongs.
	Inputs    map[string]interface{} `json:"inputs,omitempty"`    // Action inputs
	PersonID  string                 `json:"personId,omitempty"`  // The person ID of the person who created the action.
	RoomID    string                 `json:"roomId,omitempty"`    // The room ID of the attachment action.
	Created   time.Time              `json:"created,omitempty"`   // Action creation date/time.
}

// CreateAttachmentAction creates an attachment action.
/*
Creates an attachment action
 @param attachActionCreateRequest
 @return AttachmentAction
*/
func (s *AttachmentActionsService) CreateAttachmentAction(attachActionCreateRequest *AttachmentActionCreateRequest) (*AttachmentAction, *resty.Response, error) {

	path := "/attachment/actions"

	response, err := s.client.R().
		SetBody(attachActionCreateRequest).
		SetResult(&AttachmentAction{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AttachmentAction)
	return result, response, err

}

// GetAttachmentAction Shows details for an action, by attachment action ID.
/* Shows details for an attachment action, by attachment action ID.
Specify the attachment action ID in the attachmentActionID parameter in the URI.

@param attachmentActionID The unique identifier for the attachment action.
*/
func (s *AttachmentActionsService) GetAttachmentAction(attachmentActionID string) (*AttachmentAction, *resty.Response, error) {

	path := "/attachment/actions/{attachmentActionID}"
	path = strings.Replace(path, "{"+"attachmentActionID"+"}", fmt.Sprintf("%v", attachmentActionID), -1)

	response, err := s.client.R().
		SetResult(&AttachmentAction{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AttachmentAction)
	return result, response, err

}
