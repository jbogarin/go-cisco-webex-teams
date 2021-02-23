package webexteams

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
)

// DevicesService is the service to communicate with the Devices API endpoint
type DevicesService service

// Devices is the List of Devices
type Devices struct {
	Items []Device `json:"items,omitempty"`
}

// DeviceCodeRequest is the Create Device Activation Code Request Parameters
type DeviceCodeRequest struct {
	PlaceID string `json:"placeId,omitempty"` // The placeId of the place where the device will be activated.
}

// DeviceCode is the code to activate a device in a Place
type DeviceCode struct {
	ID      string    `json:"id,omitempty"`
	Code    string    `json:"code,omitempty"`
	PlaceID string    `json:"placeId,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Expires time.Time `json:"expires,omitempty"`
}

// Device is the Device definition
type Device struct {
	ID               string   `json:"id,omitempty"`               // A unique identifier for the device.
	DisplayName      string   `json:"displayName,omitempty"`      // A friendly name for the device.
	PlaceID          string   `json:"placeId,omitempty"`          // The place associated with the device.
	OrgID            string   `json:"orgId,omitempty"`            // The organization associated with the device.
	Capabilities     []string `json:"capabilities,omitempty"`     // The capabilities of the device.
	Permissions      []string `json:"permissions,omitempty"`      // The permissions the user has for this device. For example, xapi means this user is entitled to using the xapi against this device.
	ConnectionStatus string   `json:"connectionStatus,omitempty"` // The connection status of the device.
	Product          string   `json:"product,omitempty"`          // The product name.
	Tags             []string `json:"tags,omitempty"`             // Tags assigned to the device.
	IP               string   `json:"ip,omitempty"`               // The current IP address of the device.
	ActiveInterface  string   `json:"activeInterface,omitempty"`  // The current network connectivty for the device.
	MAC              string   `json:"mac,omitempty"`              // The unique address for the network adapter.
	Serial           string   `json:"serial,omitempty"`           // Serial number for the device.
	Software         string   `json:"software,omitempty"`         // The operating system name data and version tag.
	UpgradeChannel   string   `json:"upgradeChannel,omitempty"`   // The upgrade channel the device is assigned to.
	PrimarySIPURL    string   `json:"primarySipUrl,omitempty"`    // The primary SIP address to dial this device.
	SIPURLs          []string `json:"sipUrls,omitempty"`          // All SIP addresses to dial this device.
	ErrorCodes       []string `json:"errorcodes,omitempty"`       // Device Error codes
}

// AddDevice is used to append a device to a slice of Devices
func (devices *Devices) AddDevice(item Device) []Device {
	devices.Items = append(devices.Items, item)
	return devices.Items
}

func devicesPagination(linkHeader string, size, max int) *Devices {
	items := &Devices{}

	for _, l := range link.Parse(linkHeader) {
		if l.Rel == "next" {
			response, err := RestyClient.R().
				SetResult(&Devices{}).
				SetError(&Error{}).
				Get(l.URI)

			if err != nil {
				return nil
			}
			items = response.Result().(*Devices)
			if size != 0 {
				size = size + len(items.Items)
				if size < max {
					devices := devicesPagination(response.Header().Get("Link"), size, max)
					for _, device := range devices.Items {
						items.AddDevice(device)
					}
				}
			} else {
				devices := devicesPagination(response.Header().Get("Link"), size, max)
				for _, device := range devices.Items {
					items.AddDevice(device)
				}
			}

		}
	}

	return items
}

// CreateDeviceActivationCode Create a Device Activation Code
/* Generate an activation code for a device in a specific place by placeId.
Currently, activation codes may only be generated for shared places--personal mode is not supported.

 @param placeId (string) The placeId of the place where the device will be activated.
 @return DeviceCode
*/
func (s *DevicesService) CreateDeviceActivationCode(deviceCodeRequest *DeviceCodeRequest) (*DeviceCode, *resty.Response, error) {

	path := "/devices/activationCode"

	response, err := RestyClient.R().
		SetBody(deviceCodeRequest).
		SetResult(&DeviceCode{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DeviceCode)
	return result, response, err

}

// DeleteDevice Remove a device from the system. Only an admin can remove a device.
/* Remove a device from the system. Only an admin can remove a device.
Specify the device ID in the deviceID parameter in the URI.

 @param deviceID A unique identifier for the device.
 @return
*/
func (s *DevicesService) DeleteDevice(deviceID string) (*resty.Response, error) {

	path := "/devices/{deviceId}"
	path = strings.Replace(path, "{"+"deviceId"+"}", fmt.Sprintf("%v", deviceID), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

//GetDevice Shows details for a device, by ID.
/* Shows details for a device, by ID. Specify the device ID in the deviceId parameter in the URI.
@param deviceID (string) unique identifier for the device.
@return Device
*/
func (s *DevicesService) GetDevice(deviceID string) (*Device, *resty.Response, error) {

	path := "/devices/{deviceId}"
	path = strings.Replace(path, "{"+"deviceId"+"}", fmt.Sprintf("%v", deviceID), -1)

	response, err := RestyClient.R().
		SetResult(&Device{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Device)
	return result, response, err

}

// ListDevicesQueryParams are the query params for the ListDevices API Call
type ListDevicesQueryParams struct {
	PersonID         string `url:"personId,omitempty"`         // List devices by person ID.
	PlaceID          string `url:"placeId,omitempty"`          // List devices by place ID.
	OrgID            string `url:"orgId,omitempty"`            // List devices in this organization. Only admin users of another organization (such as partners) may use this parameter.
	DisplayName      string `url:"displayName,omitempty"`      // List devices with this display name.
	Product          string `url:"product,omitempty"`          // List devices with this product name. Possible values: DX-80, RoomKit, SX-80
	Tag              string `url:"tag,omitempty"`              // List devices which have a tag. Accepts multiple values separated by commas.
	ConnectionStatus string `url:"connectionStatus,omitempty"` // List devices with this connection status.
	Serial           string `url:"serial,omitempty"`           // List devices with this serial number.
	Software         string `url:"software,omitempty"`         // List devices with this software version.
	UpgradeChannel   string `url:"upgradeChannel,omitempty"`   // List devices with this upgrade channel.
	ErrorCode        string `url:"errorcode,omitempty"`        // List devices with this error code.
	Capability       string `url:"capability,omitempty"`       // List devices with this capability. Possible values: xapi
	Permission       string `url:"permission,omitempty"`       // List devices with this permission.
	Start            int    `url:"start,omitempty"`            // Offset. Default is 0.
	Max              int    `url:"max,omitempty"`              // Limit the maximum number of items in the response.
	Paginate         bool   // Indicates if pagination is needed
}

// ListDevices List devices in your organization.
/*  Lists all active Webex devices associated with the authenticated user, such as devices activated in personal mode.
Administrators can list all devices within an organization.
 @param personId (string) list devices by person ID.
 @param placeId (string ) list devices by place ID.
 @param orgId (string) list devices in this organization. Only admin users of another organization (such as partners) may use this parameter.
 @param displayName (string) list devices with this display name.
 @param product (string) list devices with this product name. Possible values: DX-80, RoomKit, SX-80
 @param tag (string) list devices which have a tag. Accepts multiple values separated by commas.
 @param connectionStatus (string) list devices with this connection status.
 @param serial (string) list devices with this serial number.
 @param software (string) list devices with this software version.
 @param upgradechannel (string) list devices with this upgrade channel.
 @param errorCode (string) list devices with this error code.
 @param capability (string) list devices with this capability. Possible values: xapi
 @param permission (string) list devices with this permission.
 @param start (int) offset. default is 0.
 @param max (int) limit the maximum number of items in the response.
 @param paginate (bool) indicates if pagination is needed
 @return Devices
*/
func (s *DevicesService) ListDevices(queryParams *ListDevicesQueryParams) (*Devices, *resty.Response, error) {

	path := "/devices/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Devices{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Devices)
	if queryParams.Paginate == true {
		items := devicesPagination(response.Header().Get("Link"), 0, 0)
		for _, device := range items.Items {
			result.AddDevice(device)
		}
	} else {
		if len(result.Items) < queryParams.Max {
			items := devicesPagination(response.Header().Get("Link"), len(result.Items), queryParams.Max)
			for _, device := range items.Items {
				result.AddDevice(device)
			}
		}
	}
	return result, response, err

}
