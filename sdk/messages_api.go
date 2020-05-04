package webexteams

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
)

// MessagesService is the service to communicate with the Messages API endpoint
type MessagesService service

// File is the struct used to define a file that needs to be sent. A file can either be a remote URI 
// or an io.Reader. If RemoteFileURI is set, it takes precedence over the Reader.  
type File struct {
	Name        	string    `json:"fileName,omitempty"`        // File Name.
	Reader      	io.Reader `json:"fileReader,omitempty"`   	 // File io.Reader.
	ContentType 	string    `json:"contentType,omitempty"`     // File Content Type.
	RemoteFileURI   string    `json:"remoteFileURI,omitempty"`   // Remote file URI.
}

// MessageCreateRequest is the Create Message Request Parameters
type MessageCreateRequest struct {
	RoomID        string   			`json:"roomId,omitempty"`        // Room ID.
	ToPersonID    string   			`json:"toPersonId,omitempty"`    // Person ID (for type=direct).
	ToPersonEmail string   			`json:"toPersonEmail,omitempty"` // Person email (for type=direct).
	Text          string   			`json:"text,omitempty"`          // Message in plain text format.
	Markdown      string   			`json:"markdown,omitempty"`      // Message in markdown format.
	Files    	  []File 			`json:"files,omitempty"` 	 	 // files array.
}

// Message is the Message definition
type Message struct {
	ID              string    `json:"id,omitempty"`              // Message ID.
	RoomID          string    `json:"roomId,omitempty"`          // Room ID.
	RoomType        string    `json:"roomType,omitempty"`        // Room type (group or direct).
	ToPersonID      string    `json:"toPersonId,omitempty"`      // Person ID (for type=direct).
	ToPersonEmail   string    `json:"toPersonEmail,omitempty"`   // Person email (for type=direct).
	Text            string    `json:"text,omitempty"`            // Message in plain text format.
	Markdown        string    `json:"markdown,omitempty"`        // Message in markdown format.
	Files           []string `json:"files,omitempty"`            // File array.
	PersonID        string    `json:"personId,omitempty"`        // Person ID.
	PersonEmail     string    `json:"personEmail,omitempty"`     // Person Email.
	Created         time.Time `json:"created,omitempty"`         // Message creation date/time.
	MentionedPeople []string  `json:"mentionedPeople,omitempty"` // Person ID array.
	MentionedGroups []string  `json:"mentionedGroups,omitempty"` // Groups array.
}

// Messages is the List of Messages
type Messages struct {
	Items []Message `json:"items,omitempty"`
}

// AddMessage is used to append a message to a slice of messages
func (messages *Messages) AddMessage(item Message) []Message {
	messages.Items = append(messages.Items, item)
	return messages.Items
}

func messagesPagination(linkHeader string, size, max int) *Messages {
	items := &Messages{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := RestyClient.R().
				SetResult(&Messages{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Messages)
			if size != 0 {
				size = size + len(items.Items)
				if size < max {
					messages := messagesPagination(response.Header().Get("Link"), size, max)
					for _, message := range messages.Items {
						items.AddMessage(message)
					}
				}
			} else {
				messages := messagesPagination(response.Header().Get("Link"), size, max)
				for _, message := range messages.Items {
					items.AddMessage(message)
				}
			}

		}
	}

	return items
}

// CreateMessage Post a plain text or rich text message, and optionally, a media content attachment, to a room.
/* Post a plain text or rich text message, and optionally, a media content attachment, to a room.
The files parameter is an array, which accepts multiple values to allow for future expansion, but currently only one file may be included with the message.

 @param messageCreateRequest
 @return Message
*/
func (s *MessagesService) CreateMessage(messageCreateRequest *MessageCreateRequest) (*Message, *resty.Response, error) {

	path := "/messages/"

	responsePart := RestyClient.R()

	if messageCreateRequest.RoomID != "" {
		responsePart.SetFormDataFromValues(url.Values{"roomId": []string{messageCreateRequest.RoomID}})
	}

	if messageCreateRequest.Markdown != "" {
		responsePart.SetFormDataFromValues(url.Values{"markdown": []string{messageCreateRequest.Markdown}})
	}

	if messageCreateRequest.Text != "" {
		responsePart.SetFormDataFromValues(url.Values{"text": []string{messageCreateRequest.Text}})
	}

	if messageCreateRequest.ToPersonEmail != "" {
		responsePart.SetFormDataFromValues(url.Values{"toPersonEmail": []string{messageCreateRequest.ToPersonEmail}})
	}

	if messageCreateRequest.ToPersonID != "" {
		responsePart.SetFormDataFromValues(url.Values{"toPersonId": []string{messageCreateRequest.ToPersonID}})
	}

	if len(messageCreateRequest.Files) > 1 {
		return nil, nil, errors.New("Multi file attachment is not supported yet.")
	}

	for _, fileToSend := range messageCreateRequest.Files {
		if fileToSend.RemoteFileURI != "" {
			responsePart.SetFormDataFromValues(url.Values{"files": []string{fileToSend.RemoteFileURI}})
		} else if fileToSend.Reader != nil {
			responsePart.SetMultipartField("files", fileToSend.Name, fileToSend.ContentType, fileToSend.Reader)
		}
	}

	response, err := responsePart.SetResult(&Message{}).Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Message)
	return result, response, err

}

// DeleteMessage Delete a Message.
/* Deletes a message by ID.
@param messageID Message ID.
@return
*/
func (s *MessagesService) DeleteMessage(messageID string) (*resty.Response, error) {

	path := "/messages/{messageId}"
	path = strings.Replace(path, "{"+"messageId"+"}", fmt.Sprintf("%v", messageID), -1)

	response, err := RestyClient.R().
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetMessage Shows details for a message, by message ID.
/* Shows details for a message, by message ID.
Specify the message ID in the messageID parameter in the URI.

 @param messageID Message ID.
 @return Message
*/
func (s *MessagesService) GetMessage(messageID string) (*Message, *resty.Response, error) {

	path := "/messages/{messageId}"
	path = strings.Replace(path, "{"+"messageId"+"}", fmt.Sprintf("%v", messageID), -1)

	response, err := RestyClient.R().
		SetResult(&Message{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Message)
	return result, response, err

}

// DirectMessagesQueryParams are the query params for the ListMessages API Call
type DirectMessagesQueryParams struct {
	ParentID    string `url:"parentId,omitempty"`    // List messages with a parent, by ID.
	PersonID    string `url:"personId,omitempty"`    // List messages in a 1:1 room, by person ID.
	PersonEmail string `url:"personEmail,omitempty"` // List messages in a 1:1 room, by person email.
	Max         int    `url:"max,omitempty"`         // Limit the maximum number of items in the response.
	Paginate    bool   // Indicates if pagination is needed
}

// GetDirectMessages Lists all messages in a 1:1 (direct) room.
/* Lists all messages in a 1:1 (direct) room.
Use the personId or personEmail query parameter to specify the room.
 @param parentId Parent Message ID.
 @param personId Person ID.
 @param personEmail Person Email.
 @return a list of Messages
*/
func (s *MessagesService) GetDirectMessages(queryParams *DirectMessagesQueryParams) (*Messages, *resty.Response, error) {
	path := "/messages/direct"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Messages{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Messages)
	if queryParams.Paginate == true {
		items := messagesPagination(response.Header().Get("Link"), 0, 0)
		for _, message := range items.Items {
			result.AddMessage(message)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := messagesPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, message := range items.Items {
				result.AddMessage(message)
			}
		}
	}

	return result, response, err

}

// ListMessagesQueryParams are the query params for the ListMessages API Call
type ListMessagesQueryParams struct {
	RoomID          string    `url:"roomId,omitempty"`          // List messages for a room, by ID.
	MentionedPeople string    `url:"mentionedPeople,omitempty"` // List messages where the caller is mentioned by specifying *me* or the caller personId.
	Before          time.Time `url:"before,omitempty"`          // List messages sent before a date and time, in ISO8601 format. Format: yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ
	BeforeMessage   string    `url:"beforeMessage,omitempty"`   // List messages sent before a message, by ID.
	Max             int       `url:"max,omitempty"`             // Limit the maximum number of items in the response.
	Paginate        bool      // Indicates if pagination is needed
}

// ListMessages Lists all messages in a room. Each message will include content attachments if present.
/* Lists all messages in a room. Each message will include content attachments if present.
The list sorts the messages in descending order by creation date.
Long result sets will be split into pages.

 @param roomID List messages for a room, by ID.
 @param "mentionedPeople" (string) List messages where the caller is mentioned by specifying *me* or the caller personId.
 @param "before" (time.Time) List messages sent before a date and time, in ISO8601 format. Format: yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ
 @param "beforeMessage" (string) List messages sent before a message, by ID.
 @param "max" (int) Limit the maximum number of items in the response.
 @param "paginate" (bool) Indicates if pagination is needed
 @return Messages
*/
func (s *MessagesService) ListMessages(queryParams *ListMessagesQueryParams) (*Messages, *resty.Response, error) {

	path := "/messages/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Messages{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Messages)
	if queryParams.Paginate == true {
		items := messagesPagination(response.Header().Get("Link"), 0, 0)
		for _, message := range items.Items {
			result.AddMessage(message)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := messagesPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, message := range items.Items {
				result.AddMessage(message)
			}
		}
	}

	return result, response, err

}
