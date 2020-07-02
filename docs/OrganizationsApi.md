# \OrganizationsApi

All URIs are relative to *https://webexapis.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganization**](OrganizationsApi.md#GetOrganization) | **Get** /organizations/{orgId} | Shows details for an organization, by ID.
[**ListOrganizations**](OrganizationsApi.md#ListOrganizations) | **Get** /organizations/ | List all organizations visible by your account. The results will not be paginated.


# **GetOrganization**
> Organization GetOrganization($orgID)

Shows details for an organization, by ID.

Shows details for an organization, by ID. Specify the org ID in the orgID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string**| Organization ID. | 

### Return type

[**Organization**](Organization.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrganizations**
> Organizations ListOrganizations($max)

List all organizations visible by your account. The results will not be paginated.

List all organizations visible by your account. The results will not be paginated.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Limit the maximum number of items in the response. | [optional] 

### Return type

[**Organizations**](Organizations.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

