# \RolesApi

All URIs are relative to *https://api.ciscospark.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRole**](RolesApi.md#GetRole) | **Get** /roles/{roleId} | Shows details for a role, by ID.
[**GetRoles**](RolesApi.md#GetRoles) | **Get** /roles/ | List all roles.


# **GetRole**
> Role GetRole($roleID)

Shows details for a role, by ID.

Shows details for a role, by ID. Specify the role ID in the roleID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleId** | **string**| Role ID. | 

### Return type

[**Role**](Role.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoles**
> Roles GetRoles($max)

List all roles.

List all roles.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Limit the maximum number of items in the response. | [optional] 

### Return type

[**Roles**](Roles.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

