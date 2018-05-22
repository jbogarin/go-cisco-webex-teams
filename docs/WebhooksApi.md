# \WebhooksApi

All URIs are relative to *https://api.ciscospark.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhooksApi.md#CreateWebhook) | **Post** /webhooks/ | Creates a webhook.
[**DeleteWebhook**](WebhooksApi.md#DeleteWebhook) | **Delete** /webhooks/{webhookId} | Deletes a webhook, by ID.
[**GetWebhook**](WebhooksApi.md#GetWebhook) | **Get** /webhooks/{webhookId} | Shows details for a webhook, by ID.
[**ListWebhooks**](WebhooksApi.md#ListWebhooks) | **Get** /webhooks/ | Lists all of your webhooks.
[**UpdateWebhook**](WebhooksApi.md#UpdateWebhook) | **Put** /webhooks/{webhookId} | Updates a webhook, by ID.


# **CreateWebhook**
> Webhook CreateWebhook($webhookCreateRequest)

Creates a webhook.

Creates a webhook.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookCreateRequest** | [**WebhookCreateRequest**](WebhookCreateRequest.md)|  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebhook**
> DeleteWebhook($webhookID)

Deletes a webhook, by ID.

Deletes a webhook, by ID. Specify the webhook ID in the webhookID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookId** | **string**| Webhook ID. | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhook**
> Webhook GetWebhook($webhookID)

Shows details for a webhook, by ID.

Shows details for a webhook, by ID. Specify the webhook ID in the webhookID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookId** | **string**| Webhook ID. | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebhooks**
> Webhooks ListWebhooks($max)

Lists all of your webhooks.

Lists all of your webhooks.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Limit the maximum number of items in the response. | [optional] 

### Return type

[**Webhooks**](Webhooks.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhook**
> Webhook UpdateWebhook($webhookId, $webhookUpdateRequest)

Updates a webhook, by ID.

Updates a webhook, by ID. Specify the webhook ID in the webhookID parameter in the URI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookId** | **string**| Webhook ID. | 
 **webhookUpdateRequest** | [**WebhookUpdateRequest**](WebhookUpdateRequest.md)|  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

