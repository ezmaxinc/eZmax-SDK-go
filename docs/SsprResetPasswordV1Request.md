# SsprResetPasswordV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PksCustomerCode** | **string** | The customer code assigned to your account | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**EUserTypeSSPR** | [**FieldEUserTypeSSPR**](FieldEUserTypeSSPR.md) |  | 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SUserLoginname** | Pointer to **string** | The Login name of the User. | [optional] 
**BinUserSSPRtoken** | **string** | Hex Encoded Secret SSPR token | 
**SPassword** | **string** | A Password.  Must meet complexity requirements | 

## Methods

### NewSsprResetPasswordV1Request

`func NewSsprResetPasswordV1Request(pksCustomerCode string, fkiLanguageID int32, eUserTypeSSPR FieldEUserTypeSSPR, binUserSSPRtoken string, sPassword string, ) *SsprResetPasswordV1Request`

NewSsprResetPasswordV1Request instantiates a new SsprResetPasswordV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsprResetPasswordV1RequestWithDefaults

`func NewSsprResetPasswordV1RequestWithDefaults() *SsprResetPasswordV1Request`

NewSsprResetPasswordV1RequestWithDefaults instantiates a new SsprResetPasswordV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPksCustomerCode

`func (o *SsprResetPasswordV1Request) GetPksCustomerCode() string`

GetPksCustomerCode returns the PksCustomerCode field if non-nil, zero value otherwise.

### GetPksCustomerCodeOk

`func (o *SsprResetPasswordV1Request) GetPksCustomerCodeOk() (*string, bool)`

GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPksCustomerCode

`func (o *SsprResetPasswordV1Request) SetPksCustomerCode(v string)`

SetPksCustomerCode sets PksCustomerCode field to given value.


### GetFkiLanguageID

`func (o *SsprResetPasswordV1Request) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *SsprResetPasswordV1Request) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *SsprResetPasswordV1Request) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetEUserTypeSSPR

`func (o *SsprResetPasswordV1Request) GetEUserTypeSSPR() FieldEUserTypeSSPR`

GetEUserTypeSSPR returns the EUserTypeSSPR field if non-nil, zero value otherwise.

### GetEUserTypeSSPROk

`func (o *SsprResetPasswordV1Request) GetEUserTypeSSPROk() (*FieldEUserTypeSSPR, bool)`

GetEUserTypeSSPROk returns a tuple with the EUserTypeSSPR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserTypeSSPR

`func (o *SsprResetPasswordV1Request) SetEUserTypeSSPR(v FieldEUserTypeSSPR)`

SetEUserTypeSSPR sets EUserTypeSSPR field to given value.


### GetSEmailAddress

`func (o *SsprResetPasswordV1Request) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *SsprResetPasswordV1Request) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *SsprResetPasswordV1Request) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *SsprResetPasswordV1Request) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSUserLoginname

`func (o *SsprResetPasswordV1Request) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *SsprResetPasswordV1Request) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *SsprResetPasswordV1Request) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.

### HasSUserLoginname

`func (o *SsprResetPasswordV1Request) HasSUserLoginname() bool`

HasSUserLoginname returns a boolean if a field has been set.

### GetBinUserSSPRtoken

`func (o *SsprResetPasswordV1Request) GetBinUserSSPRtoken() string`

GetBinUserSSPRtoken returns the BinUserSSPRtoken field if non-nil, zero value otherwise.

### GetBinUserSSPRtokenOk

`func (o *SsprResetPasswordV1Request) GetBinUserSSPRtokenOk() (*string, bool)`

GetBinUserSSPRtokenOk returns a tuple with the BinUserSSPRtoken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinUserSSPRtoken

`func (o *SsprResetPasswordV1Request) SetBinUserSSPRtoken(v string)`

SetBinUserSSPRtoken sets BinUserSSPRtoken field to given value.


### GetSPassword

`func (o *SsprResetPasswordV1Request) GetSPassword() string`

GetSPassword returns the SPassword field if non-nil, zero value otherwise.

### GetSPasswordOk

`func (o *SsprResetPasswordV1Request) GetSPasswordOk() (*string, bool)`

GetSPasswordOk returns a tuple with the SPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPassword

`func (o *SsprResetPasswordV1Request) SetSPassword(v string)`

SetSPassword sets SPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


