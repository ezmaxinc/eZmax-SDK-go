# PhoneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiPhonetypeID** | **int32** | The unique ID of the Phonetype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| |3|Mobile| |4|Fax| |5|Pager| |6|Toll Free| | 
**EPhoneType** | [**FieldEPhoneType**](FieldEPhoneType.md) |  | 
**SPhoneRegion** | Pointer to **string** | The region of the phone number. (For a North America Number only)  The region is the \&quot;514\&quot; section in this sample phone number: (514) 990-1516 x123 | [optional] 
**SPhoneExchange** | Pointer to **string** | The exchange of the phone number. (For a North America Number only)  The exchange is the \&quot;990\&quot; section in this sample phone number: (514) 990-1516 x123 | [optional] 
**SPhoneNumber** | Pointer to **string** | The number of the phone number. (For a North America Number only)  The number is the \&quot;1516\&quot; section in this sample phone number: (514) 990-1516 x123 | [optional] 
**SPhoneInternational** | Pointer to **string** | The international phone number. (For phone numbers outside of North)  Do not specify the \&quot;011\&quot; part of the phone number used to dial an international phone number from North America.  For example for this sample phone number \&quot;+442071838750\&quot;, you would send \&quot;442071838750\&quot; without the \&quot;+\&quot; sign. | [optional] 
**SPhoneExtension** | Pointer to **string** | The extension of the phone number.  The extension is the \&quot;123\&quot; section in this sample phone number: (514) 990-1516 x123.  It can also be used with international phone numbers | [optional] 

## Methods

### NewPhoneRequest

`func NewPhoneRequest(fkiPhonetypeID int32, ePhoneType FieldEPhoneType, ) *PhoneRequest`

NewPhoneRequest instantiates a new PhoneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneRequestWithDefaults

`func NewPhoneRequestWithDefaults() *PhoneRequest`

NewPhoneRequestWithDefaults instantiates a new PhoneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiPhonetypeID

`func (o *PhoneRequest) GetFkiPhonetypeID() int32`

GetFkiPhonetypeID returns the FkiPhonetypeID field if non-nil, zero value otherwise.

### GetFkiPhonetypeIDOk

`func (o *PhoneRequest) GetFkiPhonetypeIDOk() (*int32, bool)`

GetFkiPhonetypeIDOk returns a tuple with the FkiPhonetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPhonetypeID

`func (o *PhoneRequest) SetFkiPhonetypeID(v int32)`

SetFkiPhonetypeID sets FkiPhonetypeID field to given value.


### GetEPhoneType

`func (o *PhoneRequest) GetEPhoneType() FieldEPhoneType`

GetEPhoneType returns the EPhoneType field if non-nil, zero value otherwise.

### GetEPhoneTypeOk

`func (o *PhoneRequest) GetEPhoneTypeOk() (*FieldEPhoneType, bool)`

GetEPhoneTypeOk returns a tuple with the EPhoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPhoneType

`func (o *PhoneRequest) SetEPhoneType(v FieldEPhoneType)`

SetEPhoneType sets EPhoneType field to given value.


### GetSPhoneRegion

`func (o *PhoneRequest) GetSPhoneRegion() string`

GetSPhoneRegion returns the SPhoneRegion field if non-nil, zero value otherwise.

### GetSPhoneRegionOk

`func (o *PhoneRequest) GetSPhoneRegionOk() (*string, bool)`

GetSPhoneRegionOk returns a tuple with the SPhoneRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneRegion

`func (o *PhoneRequest) SetSPhoneRegion(v string)`

SetSPhoneRegion sets SPhoneRegion field to given value.

### HasSPhoneRegion

`func (o *PhoneRequest) HasSPhoneRegion() bool`

HasSPhoneRegion returns a boolean if a field has been set.

### GetSPhoneExchange

`func (o *PhoneRequest) GetSPhoneExchange() string`

GetSPhoneExchange returns the SPhoneExchange field if non-nil, zero value otherwise.

### GetSPhoneExchangeOk

`func (o *PhoneRequest) GetSPhoneExchangeOk() (*string, bool)`

GetSPhoneExchangeOk returns a tuple with the SPhoneExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneExchange

`func (o *PhoneRequest) SetSPhoneExchange(v string)`

SetSPhoneExchange sets SPhoneExchange field to given value.

### HasSPhoneExchange

`func (o *PhoneRequest) HasSPhoneExchange() bool`

HasSPhoneExchange returns a boolean if a field has been set.

### GetSPhoneNumber

`func (o *PhoneRequest) GetSPhoneNumber() string`

GetSPhoneNumber returns the SPhoneNumber field if non-nil, zero value otherwise.

### GetSPhoneNumberOk

`func (o *PhoneRequest) GetSPhoneNumberOk() (*string, bool)`

GetSPhoneNumberOk returns a tuple with the SPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneNumber

`func (o *PhoneRequest) SetSPhoneNumber(v string)`

SetSPhoneNumber sets SPhoneNumber field to given value.

### HasSPhoneNumber

`func (o *PhoneRequest) HasSPhoneNumber() bool`

HasSPhoneNumber returns a boolean if a field has been set.

### GetSPhoneInternational

`func (o *PhoneRequest) GetSPhoneInternational() string`

GetSPhoneInternational returns the SPhoneInternational field if non-nil, zero value otherwise.

### GetSPhoneInternationalOk

`func (o *PhoneRequest) GetSPhoneInternationalOk() (*string, bool)`

GetSPhoneInternationalOk returns a tuple with the SPhoneInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneInternational

`func (o *PhoneRequest) SetSPhoneInternational(v string)`

SetSPhoneInternational sets SPhoneInternational field to given value.

### HasSPhoneInternational

`func (o *PhoneRequest) HasSPhoneInternational() bool`

HasSPhoneInternational returns a boolean if a field has been set.

### GetSPhoneExtension

`func (o *PhoneRequest) GetSPhoneExtension() string`

GetSPhoneExtension returns the SPhoneExtension field if non-nil, zero value otherwise.

### GetSPhoneExtensionOk

`func (o *PhoneRequest) GetSPhoneExtensionOk() (*string, bool)`

GetSPhoneExtensionOk returns a tuple with the SPhoneExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneExtension

`func (o *PhoneRequest) SetSPhoneExtension(v string)`

SetSPhoneExtension sets SPhoneExtension field to given value.

### HasSPhoneExtension

`func (o *PhoneRequest) HasSPhoneExtension() bool`

HasSPhoneExtension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


