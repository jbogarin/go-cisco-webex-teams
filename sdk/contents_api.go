package webexteams

import (
	"fmt"
	"strings"

	"gopkg.in/resty.v1"
)

// ContentsService is the service to communicate with the Contents API endpoint
type ContentsService service

// GetContent Get File contents.
/* Get File contents by ID. Returns binary of file.
@param contentID Content ID.
@return
*/
func (s *ContentsService) GetContent(contentID string) (*resty.Response, error) {

	path := "/contents/{contentId}"
	path = strings.Replace(path, "{"+"contentId"+"}", fmt.Sprintf("%v", contentID), -1)

	response, err := RestyClient.R().
		Get(path)

	if err != nil {
		return nil, err
	}

	return response, err

}
