package webexteams

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
)

// MeetingsService is the service to communicate with the Meetings API endpoint
type MeetingsService service

// MeetingInvitee is the invitee definition
type MeetingInvitee struct {
	Email       string `json:"email,omitempty"`
	DisplayName string `json:"DisplayName,omitempty"`
	CoHost      string `json:"CoHost,omitempty"`
}

// MeetingCreateRequest is the Create Meeting Request Parameters
type MeetingCreateRequest struct {
	Title                    string           `json:"title,omitempty"`
	Agenda                   string           `json:"agenda,omitempty"`
	Password                 string           `json:"password,omitempty"`
	Start                    time.Time        `json:"start,omitempty"`
	End                      time.Time        `json:"end,omitempty"`
	Timezone                 string           `json:"timezone,omitempty"`
	Recurrence               string           `json:"recurrence,omitempty"`
	EnabledAutoRecordMeeting string           `json:"enabledAutoRecordMeeting,omitempty"`
	AllowAnyUserToBeCoHost   string           `json:"allowAnyUserToBeCoHost,omitempty"`
	Invitees                 []MeetingInvitee `json:"invitees,omitempty"`
}

// MeetingTelephony is the Meeting Telephony definition
type MeetingTelephony struct {
	AccessCode    string `json:"accessCode,omitempty"`
	CallInNumbers struct {
	} `json:",omitempty"`
	Links struct{} `json:"links,omitempty"`
}

// Meeting is the Meeting definition
type Meeting struct {
	ID                       string             `json:"id,omitempty"`
	MeetingSeriesID          string             `json:"meetingSeriesId,omitempty"`
	MeetingNumber            string             `json:"meetingNumber,omitempty"`
	Title                    string             `json:"title,omitempty"`
	Agenda                   string             `json:"agenda,omitempty"`
	Password                 string             `json:"password,omitempty"`
	MeetingType              string             `json:"meetingType,omitempty"`
	State                    string             `json:"state,omitempty"`
	Timezone                 string             `json:"timezone,omitempty"`
	Start                    time.Time          `json:"start,omitempty"`
	End                      time.Time          `json:"end,omitempty"`
	Recurrence               string             `json:"recurrence,omitempty"`
	HostUserID               string             `json:"hostUserId,omitempty"`
	HostDisplayName          string             `json:"hostDisplayName,omitempty"`
	HostEmail                string             `json:"hostEmail,omitempty"`
	HostKey                  string             `json:"hostKey,omitempty"`
	WebLink                  string             `json:"webLink,omitempty"`
	SIPAddress               string             `json:"sipAddress,omitempty"`
	DialInIPAddress          string             `json:"dialInIpAddress,omitempty"`
	EnabledAutoRecordMeeting bool               `json:"enabledAutoRecordingMeeting,omitempty"`
	AllowAnyUserToBeCoHost   bool               `json:"allowAnyUserToBeCoHost,omitempty"`
	Telephony                []MeetingTelephony `json:"telephony,omitempty"`
}

// Meetings is the List of Meetings
type Meetings struct {
	Items []Meeting `json:"items,omitempty"`
}

// AddMeeting is used to append a meeting to a slice of meetings
func (meetings *Meetings) AddMeeting(item Meeting) []Meeting {
	meetings.Items = append(meetings.Items, item)
	return meetings.Items
}

func meetingsPagination(linkHeader string, size, max int) *Meetings {
	items := &Meetings{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := RestyClient.R().
				SetResult(&Meetings{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Meetings)
			if size != 0 {
				size = size + len(items.Items)
				if size < max {
					meetings := meetingsPagination(response.Header().Get("Link"), size, max)
					for _, meeting := range meetings.Items {
						items.AddMeeting(meeting)
					}
				}
			} else {
				meetings := meetingsPagination(response.Header().Get("Link"), size, max)
				for _, meeting := range meetings.Items {
					items.AddMeeting(meeting)
				}
			}

		}
	}

	return items
}

// CreateMeeting Post a plain text or rich text meeting, and optionally, a media content attachment, to a room.
/* Post a plain text or rich text meeting, and optionally, a media content attachment, to a room.
The files parameter is an array, which accepts multiple values to allow for future expansion, but currently only one file may be included with the meeting.

 @param meetingCreateRequest
 @return Meeting
*/
func (s *MeetingsService) CreateMeeting(meetingCreateRequest *MeetingCreateRequest) (*Meeting, *resty.Response, error) {

	path := "/meetings/"

	response, err := RestyClient.R().
		SetBody(meetingCreateRequest).
		SetResult(&Meeting{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Meeting)
	return result, response, err

}

// DeleteMeeting Delete a Meeting.
/* Deletes a meeting by ID.
@param meetingID Meeting ID.
@return
*/
func (s *MeetingsService) DeleteMeeting(meetingID string) (*resty.Response, error) {

	path := "/meetings/{meetingId}"
	path = strings.Replace(path, "{"+"meetingId"+"}", fmt.Sprintf("%v", meetingID), -1)

	response, err := RestyClient.R().
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetMeeting Shows details for a meeting, by meeting ID.
/* Shows details for a meeting, by meeting ID.
Specify the meeting ID in the meetingID parameter in the URI.

 @param meetingID Meeting ID.
 @return Meeting
*/
func (s *MeetingsService) GetMeeting(meetingID string) (*Meeting, *resty.Response, error) {

	path := "/meetings/{meetingId}"
	path = strings.Replace(path, "{"+"meetingId"+"}", fmt.Sprintf("%v", meetingID), -1)

	response, err := RestyClient.R().
		SetResult(&Meeting{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Meeting)
	return result, response, err

}

// ListMeetingsQueryParams are the query params for the ListMeetings API Call
type ListMeetingsQueryParams struct {
	RoomID          string    `url:"roomId,omitempty"`          // List meetings for a room, by ID.
	MentionedPeople string    `url:"mentionedPeople,omitempty"` // List meetings where the caller is mentioned by specifying *me* or the caller personId.
	Before          time.Time `url:"before,omitempty"`          // List meetings sent before a date and time, in ISO8601 format. Format: yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ
	BeforeMeeting   string    `url:"beforeMeeting,omitempty"`   // List meetings sent before a meeting, by ID.
	Max             int       `url:"max,omitempty"`             // Limit the maximum number of items in the response.
	Paginate        bool      // Indicates if pagination is needed
}

// ListMeetings Lists all meetings in a room. Each meeting will include content attachments if present.
/* Lists all meetings in a room. Each meeting will include content attachments if present.
The list sorts the meetings in descending order by creation date.
Long result sets will be split into pages.

 @param roomID List meetings for a room, by ID.
 @param "mentionedPeople" (string) List meetings where the caller is mentioned by specifying *me* or the caller personId.
 @param "before" (time.Time) List meetings sent before a date and time, in ISO8601 format. Format: yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ
 @param "beforeMeeting" (string) List meetings sent before a meeting, by ID.
 @param "max" (int) Limit the maximum number of items in the response.
 @param "paginate" (bool) Indicates if pagination is needed
 @return Meetings
*/
func (s *MeetingsService) ListMeetings(queryParams *ListMeetingsQueryParams) (*Meetings, *resty.Response, error) {

	path := "/meetings/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Meetings{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Meetings)
	if queryParams.Paginate == true {
		items := meetingsPagination(response.Header().Get("Link"), 0, 0)
		for _, meeting := range items.Items {
			result.AddMeeting(meeting)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := meetingsPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, meeting := range items.Items {
				result.AddMeeting(meeting)
			}
		}
	}

	return result, response, err

}
