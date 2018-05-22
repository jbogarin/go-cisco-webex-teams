# \LicensesApi

All URIs are relative to *https://api.ciscospark.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLicense**](LicensesApi.md#GetLicense) | **Get** /licenses/{licenseId} | Shows details for a license, by ID.
[**ListLicenses**](LicensesApi.md#ListLicenses) | **Get** /licenses/ | List all licenses for a given organization.


# **GetLicense**
> License GetLicense($licenseID)

Shows details for a license, by ID.

Shows details for a license, by ID. Specify the license ID in the licenseID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licenseId** | **string**| License ID. | 

### Return type

[**License**](License.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLicenses**
> Licenses ListLicenses($max)

List all licenses for a given organization.

List all licenses for a given organization. If no orgID is specified, the default is the organization of the authenticated user. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Limit the maximum number of items in the response. | [optional] 

### Return type

[**Licenses**](Licenses.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

