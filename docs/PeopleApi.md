# \PeopleApi

All URIs are relative to *https://api.ciscospark.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePerson**](PeopleApi.md#CreatePerson) | **Post** /people/ | Create a new user account for a given organization.
[**DeletePerson**](PeopleApi.md#DeletePerson) | **Delete** /people/{personId} | Remove a person from the system. Only an admin can remove a person.
[**GetMe**](PeopleApi.md#GetMe) | **Get** /people/me | Show the profile for the authenticated user.
[**GetPerson**](PeopleApi.md#GetPerson) | **Get** /people/{personId} | Shows details for a person, by ID.
[**ListPeople**](PeopleApi.md#ListPeople) | **Get** /people/ | List people in your organization.
[**Update**](PeopleApi.md#Update) | **Put** /people/{personId} | Update details for a person, by ID.


# **CreatePerson**
> Person CreatePerson($personRequest)

Create a new user account for a given organization.

Create a new user account for a given organization. Only an admin can create a new user account. Currently, users may have only one email address associated with their account. The emails parameter is an array, which accepts multiple values to allow for future expansion, but currently only one email address will be used for the new user. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personRequest** | [**PersonRequest**](PersonRequest.md)|  | 

### Return type

[**Person**](Person.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePerson**
> DeletePerson($personID)

Remove a person from the system. Only an admin can remove a person.

Remove a person from the system. Only an admin can remove a person. Specify the person ID in the personID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personId** | **string**| Person ID. | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMe**
> Person GetMe()

Show the profile for the authenticated user.

Show the profile for the authenticated user. This is the same as GET /people/:id using the Person ID associated with your Auth token.


### Parameters
This endpoint does not need any parameter.

### Return type

[**Person**](Person.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPerson**
> Person GetPerson($personID)

Shows details for a person, by ID.

Shows details for a person, by ID. Certain fields, such as status or lastActivity, will only be displayed for people within your organization or an organzation you manage. Specify the person ID in the personID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personId** | **string**| Person ID. | 

### Return type

[**Person**](Person.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPeople**
> People ListPeople($id, $email, $displayName, $max, $orgID)

List people in your organization.

List people in your organization. For most users, either the email or displayName parameter is required.  Admin users can omit these fields and list all users in their organization. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| List people by ID. Accepts up to 85 person IDs separated by commas. | [optional] 
 **email** | **string**| List people with this email address. For non-admin requests, either this or displayName are required. | [optional] 
 **displayName** | **string**| List people whose name starts with this string. For non-admin requests, either this or email are required. | [optional] 
 **max** | **int**| Limit the maximum number of items in the response. | [optional] 
 **orgId** | **string**| List people in this organization. Only admin users of another organization (such as partners) may use this parameter. | [optional] 

### Return type

[**People**](People.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> Person Update($personId, $personRequest)

Update details for a person, by ID.

Update details for a person, by ID. Specify the person ID in the personID parameter in the URI. Only an admin can update a person details. Email addresses for a person cannot be changed via the Cisco Spark API. Include all details for the person. This action expects all user details to be present in the request. A common approach is to first GET the person's details, make changes, then PUT both the changed and unchanged values.       


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personId** | **string**| Person ID. | 
 **personRequest** | [**PersonRequest**](PersonRequest.md)|  | 

### Return type

[**Person**](Person.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

