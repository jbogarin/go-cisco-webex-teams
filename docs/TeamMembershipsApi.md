# \TeamMembershipsApi

All URIs are relative to *https://webexapis.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTeamMembership**](TeamMembershipsApi.md#CreateTeamMembership) | **Post** /team/memberships/ | Add someone to a team by Person ID or email address; optionally making them a moderator.
[**DeleteTeamMembership**](TeamMembershipsApi.md#DeleteTeamMembership) | **Delete** /team/memberships/{membershipId} | Deletes a team membership, by ID.
[**GetTeamMembership**](TeamMembershipsApi.md#GetTeamMembership) | **Get** /team/memberships/{membershipId} | Shows details for a team membership, by ID.
[**ListTeamMemberhips**](TeamMembershipsApi.md#ListTeamMemberhips) | **Get** /team/memberships/ | Lists all team memberships for a given team, specified by the teamID query parameter.
[**UpdateTeamMembership**](TeamMembershipsApi.md#UpdateTeamMembership) | **Put** /team/memberships/{membershipId} | Updates a team membership, by ID.


# **CreateTeamMembership**
> TeamMembership CreateTeamMembership($teamMemberhipCreateRequest)

Add someone to a team by Person ID or email address; optionally making them a moderator.

Add someone to a team by Person ID or email address; optionally making them a moderator.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamMemberhipCreateRequest** | [**TeamMembershipCreateRequest**](TeamMembershipCreateRequest.md)|  | 

### Return type

[**TeamMembership**](TeamMembership.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTeamMembership**
> DeleteTeamMembership($membershipID)

Deletes a team membership, by ID.

Deletes a team membership, by ID. Specify the team membership ID in the membershipID URI parameter. 


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

# **GetTeamMembership**
> TeamMembership GetTeamMembership($membershipID)

Shows details for a team membership, by ID.

Shows details for a team membership, by ID. Specify the team membership ID in the membershipID URI parameter. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **membershipId** | **string**| Membership ID. | 

### Return type

[**TeamMembership**](TeamMembership.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTeamMemberhips**
> TeamMemberships ListTeamMemberhips($teamId, $max)

Lists all team memberships for a given team, specified by the teamID query parameter.

Lists all team memberships for a given team, specified by the teamID query parameter. Use query parameters to filter the response. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| Team ID. | 
 **max** | **int**| Limit the maximum number of items in the response. | [optional] 

### Return type

[**TeamMemberships**](TeamMemberships.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeamMembership**
> TeamMembership UpdateTeamMembership($membershipId, $teamMembershipUpdateRequest)

Updates a team membership, by ID.

Updates a team membership, by ID. Specify the team membership ID in the membershipID URI parameter. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **membershipId** | **string**| Membership ID. | 
 **teamMembershipUpdateRequest** | [**TeamMembershipUpdateRequest**](TeamMembershipUpdateRequest.md)|  | 

### Return type

[**TeamMembership**](TeamMembership.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

