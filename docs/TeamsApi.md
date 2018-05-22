# \TeamsApi

All URIs are relative to *https://api.ciscospark.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTeam**](TeamsApi.md#CreateTeam) | **Post** /teams/ | Creates a team.
[**DeleteTeam**](TeamsApi.md#DeleteTeam) | **Delete** /teams/{teamId} | Deletes a team, by ID.
[**GetTeam**](TeamsApi.md#GetTeam) | **Get** /teams/{teamId} | Shows details for a team, by ID.
[**ListTeams**](TeamsApi.md#ListTeams) | **Get** /teams/ | Lists teams to which the authenticated user belongs.
[**UpdateTeam**](TeamsApi.md#UpdateTeam) | **Put** /teams/{teamId} | Updates details for a team, by ID.


# **CreateTeam**
> Team CreateTeam($teamCreateRequest)

Creates a team.

Creates a team. The authenticated user is automatically added as a member of the team. See the Team Memberships API to learn how to add more people to the team. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamCreateRequest** | [**TeamCreateRequest**](TeamCreateRequest.md)|  | 

### Return type

[**Team**](Team.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTeam**
> DeleteTeam($teamID)

Deletes a team, by ID.

Deletes a team, by ID. Specify the team ID in the teamID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| Team ID. | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeam**
> Team GetTeam($teamID)

Shows details for a team, by ID.

Shows details for a team, by ID. Specify the team ID in the teamID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| Team ID. | 

### Return type

[**Team**](Team.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTeams**
> Teams ListTeams($max)

Lists teams to which the authenticated user belongs.

Lists teams to which the authenticated user belongs.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Limit the maximum number of items in the response. | [optional] 

### Return type

[**Teams**](Teams.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeam**
> Team UpdateTeam($teamId, $teamUpdateRequest)

Updates details for a team, by ID.

Updates details for a team, by ID. Specify the team ID in the teamID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| Team ID. | 
 **teamUpdateRequest** | [**TeamUpdateRequest**](TeamUpdateRequest.md)|  | 

### Return type

[**Team**](Team.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

