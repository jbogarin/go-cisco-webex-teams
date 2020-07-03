package webexteams

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
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
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, err
	}

	return response, err

}
