package webexteams

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
)

// RecordingsService is the service to communicate with the Recordings API endpoint
type RecordingsService service

// Recordings is the List of Recordings
type Recordings struct {
	Items []Recording `json:"items,omitempty"`
}

// Recording is the Recording definition
type Recording struct {
	ID                 string    `json:"id,omitempty"`                 // A unique identifier for recording.
	MeetingId          string    `json:"meetingId,omitempty"`          // Unique identifier for the parent ended meeting instance which the recording belongs to.
	ScheduledMeetingId string    `json:"scheduledMeetingId,omitempty"` // Unique identifier for the parent scheduled meeting which the recording belongs to.
	MeetingSeriesId    string    `json:"meetingSeriesId,omitempty"`    // Unique identifier for the parent meeting series which the recording belongs to.
	Topic              string    `json:"topic,omitempty"`              // The recording's topic.
	CreateTime         time.Time `json:"createTime,omitempty"`         // The date and time recording was created in ISO 8601 compliant format.
	TimeRecorded       time.Time `json:"timeRecorded,omitempty"`       // The date and time recording started in ISO 8601 compliant format.
	SiteUrl            string    `json:"siteUrl,omitempty"`            // The site URL for the recording.
	DownloadURL        string    `json:"downloadUrl,omitempty"`        // The download link for recording.
	PlaybackURL        string    `json:"playbackUrl,omitempty"`        // The playback link for recording.
	Password           string    `json:"password,omitempty"`           // The recording's password.
	Format             string    `json:"format,omitempty"`             // MP4 or ARF
	ServiceType        string    `json:"serviceType,omitempty"`        // Record service type (MeetingCenter, EventCenter, SupportCenter or TrainingCenter).
	DurationSeconds    int       `json:"durationSeconds,omitempty"`    // The duration of the recording, in seconds.
	SizeBytes          int       `json:"sizeBytes,omitempty"`          // The size of the recording file, in bytes.
	ShareToMe          bool      `json:"shareToMe,omitempty"`          // Whether or not the recording has been shared to the current user.
	IntegrationTags    []string  `json:"integrationTags,omitempty"`    // External keys of the parent meeting created by an integration application.
}

// TemporaryDirectDownloadLinks definition
type TemporaryDirectDownloadLinks struct {
	RecordingDownloadLink  string `json:"recordingDownloadLink,omitempty"`  // The download link for recording MP4 file without HTML page rendering in browser or HTTP redirect. It expires in 3 hours after the API request.
	AudioDownloadLink      string `json:"audioDownloadLink,omitempty"`      // The download link for recording audio file without HTML page rendering in browser or HTTP redirect. It expires in 3 hours after the API request.
	TranscriptDownloadLink string `json:"transcriptDownloadLink,omitempty"` // The download link for recording transcript file without HTML page rendering in browser or HTTP redirect. It expires in 3 hours after the API request.
	Expiration             string `json:"expiration,omitempty"`             // The date and time when recordingDownloadLink, audioDownloadLink, and transcriptDownloadLink expire in ISO 8601 compliant format.
}

// RecordingDetails is the Recording definition
type RecordingDetails struct {
	ID                           string                       `json:"id,omitempty"`                           // A unique identifier for recording.
	Topic                        string                       `json:"topic,omitempty"`                        //The recording's topic.
	CreateTime                   time.Time                    `json:"createTime,omitempty"`                   // The date and time recording was created in ISO 8601 compliant format.
	DownloadURL                  string                       `json:"downloadUrl,omitempty"`                  // The download link for recording.
	PlaybackURL                  string                       `json:"playbackUrl,omitempty"`                  // The playback link for recording.
	Password                     string                       `json:"password,omitempty"`                     // The recording's password.
	TemporaryDirectDownloadLinks TemporaryDirectDownloadLinks `json:"temporaryDirectDownloadLinks,omitempty"` // The download links for MP4, audio, and transcript of the recording without HTML page rendering in browser or HTTP redirect. This attribute is only available for Meeting Center.
	Format                       string                       `json:"format,omitempty"`                       // MP4 or ARF
	DurationSeconds              int                          `json:"durationSeconds,omitempty"`              // The duration of the recording, in seconds.
	SizeBytes                    int                          `json:"sizeBytes,omitempty"`                    // The size of the recording file, in bytes.
	ShareToMe                    bool                         `json:"shareToMe,omitempty"`                    // Whether or not the recording has been shared to the current user.
}

// AddRecording is used to append a recording to a slice of Recordings
func (recordings *Recordings) AddRecording(item Recording) []Recording {
	recordings.Items = append(recordings.Items, item)
	return recordings.Items
}

func (s *RecordingsService) recordingsPagination(linkHeader string, size, max int) *Recordings {
	items := &Recordings{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {
			response, err := s.client.R().
				SetResult(&Recordings{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Recordings)
			if size != 0 {
				size = size + len(items.Items)
				if size < max {
					recordings := s.recordingsPagination(response.Header().Get("Link"), size, max)
					for _, recording := range recordings.Items {
						items.AddRecording(recording)
					}
				}
			} else {
				recordings := s.recordingsPagination(response.Header().Get("Link"), size, max)
				for _, recording := range recordings.Items {
					items.AddRecording(recording)
				}
			}

		}
	}

	return items
}

// ListRecordingsQueryParams are the query params for the ListRecordings API Call
type ListRecordingsQueryParams struct {
	From           string `url:"from,omitempty"`           // Starting date and time (inclusive) for recordings to return, in any ISO 8601 compliant format. from cannot be after current date and time or after to.
	To             string `url:"To,omitempty"`             // Ending date and time (exclusive) for List recordings to return, in any ISO 8601 compliant format. to cannot be after current date and time or before from.
	Max            int    `url:"max,omitempty"`            // Limit the maximum number of items in the response.
	MeetingId      string `url:"meetingId,omitempty"`      // Unique identifier for the parent meeting series, scheduled meeting, or meeting instance for which recordings are being requested.
	HostEmail      string `url:"hostEmail,omitempty"`      // Email address for the meeting host.
	SiteUrl        string `url:"siteUrl,omitempty"`        // URL of the Webex site which the API lists recordings from.
	IntegrationTag string `url:"integrationTag,omitempty"` // External key of the parent meeting created by an integration application.
	Topic          string `url:"topic,omitempty"`          // Recording's topic.
	Format         string `url:"format,omitempty"`         // Recoding's file format.
	ServiceType    string `url:"serviceType,omitempty"`    // Recording's service-type.
	Paginate       bool   // Indicates if pagination is needed
}

// ListRecordings List recordings.
/* Lists recordings. You can specify a date range and the maximum number of recordings to return.
Only recordings of meetings hosted by or shared with the authenticated user will be listed.
The list returned is sorted in descending order by the date and time that the recordings were created.
Long result sets are split into pages.
 @param from (string) Starting date and time (inclusive) for recordings to return, in any ISO 8601 compliant format.
 @param to (string) Ending date and time (exclusive) for List recordings to return, in any ISO 8601 compliant format.
 @param max (int) limit the maximum number of items in the response.
 @param meetingId (string) Unique identifier for the parent meeting series, scheduled meeting, or meeting instance.
 @param hostEmail (string) Email address for the meeting host.
 @param siteUrl (string) URL of the Webex site which the API lists recordings from.
 @param integrationTag (string) External key of the parent meeting created by an integration application.
 @param topic (string) Recording's topic.
 @param format (string) Recoding's file format.
 @param serviceType (string) Recording's service-type.
 @param paginate (bool) indicates if pagination is needed
 @return Recordings
*/
func (s *RecordingsService) ListRecordings(queryParams *ListRecordingsQueryParams) (*Recordings, *resty.Response, error) {

	path := "/recordings/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := s.client.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Recordings{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Recordings)
	if queryParams.Paginate {
		items := s.recordingsPagination(response.Header().Get("Link"), 0, 0)
		for _, recording := range items.Items {
			result.AddRecording(recording)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := s.recordingsPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, recording := range items.Items {
				result.AddRecording(recording)
			}
		}
	}
	return result, response, err

}

//GetRecording Shows details for a recording, by ID.
/* Shows details for a recording, by ID. Specify the recording ID in the recordingId parameter in the URI.
@param recordingID (string) unique identifier for the recording.
@return Recording
*/
func (s *RecordingsService) GetRecording(recordingID string) (*RecordingDetails, *resty.Response, error) {

	path := "/recordings/{recordingId}"
	path = strings.Replace(path, "{"+"recordingId"+"}", fmt.Sprintf("%v", recordingID), -1)

	response, err := s.client.R().
		SetResult(&RecordingDetails{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*RecordingDetails)
	return result, response, err

}

// DeleteRecording Removes a recording with a specified recording ID. The deleted recording cannot be recovered.
/* Removes a recording with a specified recording ID. The deleted recording cannot be recovered.

Only recordings of meetings hosted by the authenticated user can be deleted.

 @param recordingID A unique identifier for the recording.
 @return
*/
func (s *RecordingsService) DeleteRecording(recordingID string) (*resty.Response, error) {

	path := "/recordings/{recordingId}"
	path = strings.Replace(path, "{"+"recordingId"+"}", fmt.Sprintf("%v", recordingID), -1)

	response, err := s.client.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}
