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
	Title                    string           `json:"title,omitempty"`                    // Meeting title.
	Agenda                   string           `json:"agenda,omitempty"`                   // Meeting agenda. The agenda can be a maximum of 2500 characters long.
	Password                 string           `json:"password,omitempty"`                 // Meeting password.
	Start                    time.Time        `json:"start,omitempty"`                    // Date and time for the start of meeting in any ISO 8601 compliant format. start cannot be before current date and time or after end.
	End                      time.Time        `json:"end,omitempty"`                      // Date and time for the end of meeting in any ISO 8601 compliant format. end cannot be before current date and time or before start.
	Timezone                 string           `json:"timezone,omitempty"`                 // Time zone in which meeting was originally scheduled (conforming with the IANA time zone database).
	Recurrence               string           `json:"recurrence,omitempty"`               // Meeting series recurrence rule (conforming with RFC 2445), applying only to meeting series.
	EnabledAutoRecordMeeting string           `json:"enabledAutoRecordMeeting,omitempty"` // Whether or not meeting is recorded automatically.
	AllowAnyUserToBeCoHost   string           `json:"allowAnyUserToBeCoHost,omitempty"`   // Whether or not to allow any invitee to be a cohost.
	Invitees                 []MeetingInvitee `json:"invitees,omitempty"`                 // Invitees for meeting.
}

// MeetingTelephony is the Meeting Telephony definition
type MeetingTelephony struct {
	AccessCode    string `json:"accessCode,omitempty"` // Code for authenticating a user to join teleconference.
	CallInNumbers struct {
		Label        string `json:"label,omitempty"`
		CallInNumber string `json:"callInNumber,omitempty"`
		TollType     string `json:"tollType,omitempty"`
	} `json:",omitempty"` // Array of call-in numbers for joining teleconference from a phone.
	Links struct {
		Rel    string `json:"rel,omitempty"`    // Link relation describing how the target resource is related to the current context (conforming with RFC5998).
		HREF   string `json:"href,omitempty"`   // Target resource URI (conforming with RFC5998).
		Method string `json:"method,omitempty"` // Target resource method (conforming with RFC5998).
	} `json:"links,omitempty"` // HATEOAS information of global call-in numbers for joining teleconference from a phone.
}

// Meeting is the Meeting definition
type Meeting struct {
	ID                       string             `json:"id,omitempty"`                          // Unique identifier for meeting
	MeetingSeriesID          string             `json:"meetingSeriesId,omitempty"`             // Unique identifier for meeting series.
	MeetingNumber            string             `json:"meetingNumber,omitempty"`               // Meeting number.
	Title                    string             `json:"title,omitempty"`                       // Meeting title.
	Agenda                   string             `json:"agenda,omitempty"`                      // Meeting agenda. The agenda can be a maximum of 2500 characters long.
	Password                 string             `json:"password,omitempty"`                    //  Meeting password.
	MeetingType              string             `json:"meetingType,omitempty"`                 // One Of: meetingSeries, scheduledMeeting o meeting
	State                    string             `json:"state,omitempty"`                       // Meeting state.
	Timezone                 string             `json:"timezone,omitempty"`                    // Time zone of start and end, conforming with the IANA time zone database.
	Start                    time.Time          `json:"start,omitempty"`                       // Start time for meeting in ISO 8601 compliant format.
	End                      time.Time          `json:"end,omitempty"`                         // End time for meeting in ISO 8601 compliant format.
	Recurrence               string             `json:"recurrence,omitempty"`                  // Meeting series recurrence rule (conforming with RFC 2445), applying only to recurring meeting series.
	HostUserID               string             `json:"hostUserId,omitempty"`                  // Unique identifier for meeting host.
	HostDisplayName          string             `json:"hostDisplayName,omitempty"`             // Display name for meeting host.
	HostEmail                string             `json:"hostEmail,omitempty"`                   // Email address for meeting host.
	HostKey                  string             `json:"hostKey,omitempty"`                     // Key for joining meeting as host.
	WebLink                  string             `json:"webLink,omitempty"`                     // Link to meeting information page where meeting client will be launched if the meeting is ready for start or join.
	SIPAddress               string             `json:"sipAddress,omitempty"`                  // SIP address for callback from a video system.
	DialInIPAddress          string             `json:"dialInIpAddress,omitempty"`             // IP address for callback from a video system.
	EnabledAutoRecordMeeting bool               `json:"enabledAutoRecordingMeeting,omitempty"` // Whether or not meeting is recorded automatically.
	AllowAnyUserToBeCoHost   bool               `json:"allowAnyUserToBeCoHost,omitempty"`      // Whether or not to allow any invitee to be a cohost
	Telephony                []MeetingTelephony `json:"telephony,omitempty"`                   // Information for callbacks from meeting to phone or for joining a teleconference using a phone.
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

func (s *MeetingsService) meetingsPagination(linkHeader string, size, max int) *Meetings {
	items := &Meetings{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {

			response, err := s.client.R().
				SetResult(&Meetings{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Meetings)
			if size != 0 {
				size = size + len(items.Items)
				if size < max {
					meetings := s.meetingsPagination(response.Header().Get("Link"), size, max)
					for _, meeting := range meetings.Items {
						items.AddMeeting(meeting)
					}
				}
			} else {
				meetings := s.meetingsPagination(response.Header().Get("Link"), size, max)
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

	response, err := s.client.R().
		SetBody(meetingCreateRequest).
		SetResult(&Meeting{}).
		SetError(&Error{}).
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

	response, err := s.client.R().
		SetError(&Error{}).
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

	response, err := s.client.R().
		SetResult(&Meeting{}).
		SetError(&Error{}).
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

	response, err := s.client.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Meetings{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Meetings)
	if queryParams.Paginate {
		items := s.meetingsPagination(response.Header().Get("Link"), 0, 0)
		for _, meeting := range items.Items {
			result.AddMeeting(meeting)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := s.meetingsPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, meeting := range items.Items {
				result.AddMeeting(meeting)
			}
		}
	}

	return result, response, err

}
