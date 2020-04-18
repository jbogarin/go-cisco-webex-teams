# \MessagesApi

All URIs are relative to *https://api.ciscospark.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessage**](MessagesApi.md#CreateMessage) | **Post** /messages/ | Post a plain text or rich text message, and optionally, a media content attachment, to a room.
[**DeleteMessage**](MessagesApi.md#DeleteMessage) | **Delete** /messages/{messageId} | Delete a Message.
[**GetMessage**](MessagesApi.md#GetMessage) | **Get** /messages/{messageId} | Shows details for a message, by message ID.
[**ListMessages**](MessagesApi.md#ListMessages) | **Get** /messages/ | Lists all messages in a room. Each message will include content attachments if present.


# **CreateMessage**
> Message CreateMessage($messageCreateRequest)

Post a plain text or rich text message, and optionally, a media content attachment, to a room.

Post a plain text or rich text message, and optionally, a media content attachment, to a room. The files parameter is an array of File struct, which accepts multiple values to allow for future expansion, but currently only one file may be included with the message. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageCreateRequest** | [**MessageCreateRequest**](MessageCreateRequest.md)|  | 

### Return type

[**Message**](Message.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMessage**
> DeleteMessage($messageID)

Delete a Message.

Deletes a message by ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageId** | **string**| Message ID. | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessage**
> Message GetMessage($messageID)

Shows details for a message, by message ID.

Shows details for a message, by message ID. Specify the message ID in the messageID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageId** | **string**| Message ID. | 

### Return type

[**Message**](Message.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMessages**
> Messages ListMessages($roomId, $mentionedPeople, $before, $beforeMessage, $max)

Lists all messages in a room. Each message will include content attachments if present.

Lists all messages in a room. Each message will include content attachments if present. The list sorts the messages in descending order by creation date. Long result sets will be split into pages. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomId** | **string**| List messages for a room, by ID. | 
 **mentionedPeople** | **string**| List messages where the caller is mentioned by specifying *me* or the caller personId. | [optional] 
 **before** | **time.Time**| List messages sent before a date and time, in ISO8601 format. Format: yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ | [optional] 
 **beforeMessage** | **string**| List messages sent before a message, by ID. | [optional] 
 **max** | **int**| Limit the maximum number of items in the response. | [optional] 

### Return type

[**Messages**](Messages.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

