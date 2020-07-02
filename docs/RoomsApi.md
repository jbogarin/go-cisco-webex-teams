# \RoomsApi

All URIs are relative to *https://webexapis.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoom**](RoomsApi.md#CreateRoom) | **Post** /rooms/ | Creates a room. The authenticated user is automatically added as a member of the room.
[**DeleteRoom**](RoomsApi.md#DeleteRoom) | **Delete** /rooms/{roomId} | Deletes a room, by ID. Deleted rooms cannot be recovered.
[**GetRoom**](RoomsApi.md#GetRoom) | **Get** /rooms/{roomId} | Shows details for a room, by ID.
[**ListRooms**](RoomsApi.md#ListRooms) | **Get** /rooms/ | List rooms.
[**UpdateRoom**](RoomsApi.md#UpdateRoom) | **Put** /rooms/{roomId} | Updates details for a room, by ID.


# **CreateRoom**
> Room CreateRoom($roomCreateRequest)

Creates a room. The authenticated user is automatically added as a member of the room.

Creates a room. The authenticated user is automatically added as a member of the room. See the Memberships API to learn how to add more people to the room. To create a 1-to-1 room, use the Create Messages endpoint to send a message directly to another person by using the toPersonID or toPersonEmail parameters. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomCreateRequest** | [**RoomCreateRequest**](RoomCreateRequest.md)|  | 

### Return type

[**Room**](Room.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoom**
> DeleteRoom($roomID)

Deletes a room, by ID. Deleted rooms cannot be recovered.

Deletes a room, by ID. Deleted rooms cannot be recovered. Deleting a room that is part of a team will archive the room instead. Specify the room ID in the roomID parameter in the URI 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomId** | **string**| Room ID. | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoom**
> Room GetRoom($roomID)

Shows details for a room, by ID.

Shows details for a room, by ID. The title of the room for 1-to-1 rooms will be the display name of the other person. Specify the room ID in the roomID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomId** | **string**| Room ID. | 

### Return type

[**Room**](Room.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRooms**
> Rooms ListRooms($teamId, $type_, $sortBy, $max)

List rooms.

List rooms. The title of the room for 1-to-1 rooms will be the display name of the other person. By default, lists rooms to which the authenticated user belongs. Long result sets will be split into pages. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| Limit the rooms to those associated with a team, by ID. | [optional] 
 **type_** | **string**| direct returns all 1-to-1 rooms. group returns all group rooms. | [optional] 
 **sortBy** | **string**| Sort results by room ID (id), most recent activity (lastactivity), or most recently created (created). | [optional] 
 **max** | **int**| Limit the maximum number of items in the response. | [optional] 

### Return type

[**Rooms**](Rooms.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoom**
> Room UpdateRoom($roomId, $roomUpdateRequest)

Updates details for a room, by ID.

Updates details for a room, by ID. Specify the room ID in the roomID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomId** | **string**| Room ID. | 
 **roomUpdateRequest** | [**RoomUpdateRequest**](RoomUpdateRequest.md)|  | 

### Return type

[**Room**](Room.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

