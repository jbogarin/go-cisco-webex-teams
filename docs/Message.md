# Message

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Message ID. | [optional] [default to null]
**RoomId** | **string** | Room ID. | [optional] [default to null]
**RoomType** | **string** | Room type (group or direct). | [optional] [default to null]
**ToPersonId** | **string** | Person ID (for type&#x3D;direct). | [optional] [default to null]
**ToPersonEmail** | **string** | Person email (for type&#x3D;direct). | [optional] [default to null]
**Text** | **string** | Message in plain text format. | [optional] [default to null]
**Markdown** | **string** | Message in markdown format. | [optional] [default to null]
**Files** | **[]string** | File URL array. | [optional] [default to null]
**PersonId** | **string** | Person ID. | [optional] [default to null]
**PersonEmail** | **string** | Person Email. | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) | Message creation date/time. | [optional] [default to null]
**MentionedPeople** | **[]string** | Person ID array. | [optional] [default to null]
**MentionedGroups** | **[]string** | Groups array. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


