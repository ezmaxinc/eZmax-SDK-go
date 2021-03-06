# WebhookEzsignFolderCompleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjEzsignfolder** | [**EzsignfolderResponse**](EzsignfolderResponse.md) |  | 
**ObjWebhook** | [**WebhookResponse**](WebhookResponse.md) |  | 
**AObjAttempt** | [**[]AttemptResponse**](AttemptResponse.md) | An array containing details of previous attempts that were made to deliver the message. The array is empty if it&#39;s the first attempt. | 

## Methods

### NewWebhookEzsignFolderCompleted

`func NewWebhookEzsignFolderCompleted(objEzsignfolder EzsignfolderResponse, objWebhook WebhookResponse, aObjAttempt []AttemptResponse, ) *WebhookEzsignFolderCompleted`

NewWebhookEzsignFolderCompleted instantiates a new WebhookEzsignFolderCompleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEzsignFolderCompletedWithDefaults

`func NewWebhookEzsignFolderCompletedWithDefaults() *WebhookEzsignFolderCompleted`

NewWebhookEzsignFolderCompletedWithDefaults instantiates a new WebhookEzsignFolderCompleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjEzsignfolder

`func (o *WebhookEzsignFolderCompleted) GetObjEzsignfolder() EzsignfolderResponse`

GetObjEzsignfolder returns the ObjEzsignfolder field if non-nil, zero value otherwise.

### GetObjEzsignfolderOk

`func (o *WebhookEzsignFolderCompleted) GetObjEzsignfolderOk() (*EzsignfolderResponse, bool)`

GetObjEzsignfolderOk returns a tuple with the ObjEzsignfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfolder

`func (o *WebhookEzsignFolderCompleted) SetObjEzsignfolder(v EzsignfolderResponse)`

SetObjEzsignfolder sets ObjEzsignfolder field to given value.


### GetObjWebhook

`func (o *WebhookEzsignFolderCompleted) GetObjWebhook() WebhookResponse`

GetObjWebhook returns the ObjWebhook field if non-nil, zero value otherwise.

### GetObjWebhookOk

`func (o *WebhookEzsignFolderCompleted) GetObjWebhookOk() (*WebhookResponse, bool)`

GetObjWebhookOk returns a tuple with the ObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebhook

`func (o *WebhookEzsignFolderCompleted) SetObjWebhook(v WebhookResponse)`

SetObjWebhook sets ObjWebhook field to given value.


### GetAObjAttempt

`func (o *WebhookEzsignFolderCompleted) GetAObjAttempt() []AttemptResponse`

GetAObjAttempt returns the AObjAttempt field if non-nil, zero value otherwise.

### GetAObjAttemptOk

`func (o *WebhookEzsignFolderCompleted) GetAObjAttemptOk() (*[]AttemptResponse, bool)`

GetAObjAttemptOk returns a tuple with the AObjAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttempt

`func (o *WebhookEzsignFolderCompleted) SetAObjAttempt(v []AttemptResponse)`

SetAObjAttempt sets AObjAttempt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


