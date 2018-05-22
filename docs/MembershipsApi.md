# \MembershipsApi

All URIs are relative to *https://api.ciscospark.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMembership**](MembershipsApi.md#CreateMembership) | **Post** /memberships/ | Add someone to a room by Person ID or email address; optionally making them a moderator.
[**DeleteMembership**](MembershipsApi.md#DeleteMembership) | **Delete** /memberships/{membershipId} | Deletes a membership by ID.
[**GetMembership**](MembershipsApi.md#GetMembership) | **Get** /memberships/{membershipId} | Get details for a membership by ID.
[**ListMemberships**](MembershipsApi.md#ListMemberships) | **Get** /memberships/ | Lists all room memberships. By default, lists memberships for rooms to which the authenticated user belongs.
[**UpdateMembership**](MembershipsApi.md#UpdateMembership) | **Put** /memberships/{membershipId} | Updates properties for a membership by ID.


# **CreateMembership**
> Membership CreateMembership($membershipCreateRequest)

Add someone to a room by Person ID or email address; optionally making them a moderator.

Add someone to a room by Person ID or email address; optionally making them a moderator.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **membershipCreateRequest** | [**MembershipCreateRequest**](MembershipCreateRequest.md)|  | 

### Return type

[**Membership**](Membership.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMembership**
> DeleteMembership($membershipID)

Deletes a membership by ID.

Deletes a membership by ID. Specify the membership ID in the membershipID URI parameter. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **membershipId** | **string**| Membership ID. | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMembership**
> Membership GetMembership($membershipID)

Get details for a membership by ID.

Get details for a membership by ID. Specify the membership ID in the membershipID URI parameter. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **membershipId** | **string**| Membership ID. | 

### Return type

[**Membership**](Membership.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMemberships**
> Memberships ListMemberships($roomId, $personId, $personEmail, $max)

Lists all room memberships. By default, lists memberships for rooms to which the authenticated user belongs.

Lists all room memberships. By default, lists memberships for rooms to which the authenticated user belongs. Use query parameters to filter the response. Use roomID to list memberships for a room, by ID. Use either personID or personEmail to filter the results. Long result sets will be split into pages. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomId** | **string**| Room ID. | [optional] 
 **personId** | **string**| Person ID. | [optional] 
 **personEmail** | **string**| Person email. | [optional] 
 **max** | **int**| Limit the maximum number of items in the response. | [optional] 

### Return type

[**Memberships**](Memberships.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMembership**
> Membership UpdateMembership($membershipId, $membershipUpdateRequest)

Updates properties for a membership by ID.

Updates properties for a membership by ID. Specify the membership ID in the membershipID URI parameter. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **membershipId** | **string**| Membership ID. | 
 **membershipUpdateRequest** | [**MembershipUpdateRequest**](MembershipUpdateRequest.md)|  | 

### Return type

[**Membership**](Membership.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

