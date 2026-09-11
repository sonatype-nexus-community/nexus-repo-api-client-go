# ComponentVersionsPageXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ComponentVersionXO**](ComponentVersionXO.md) | The versions on this page, ordered as requested | [optional] 
**Page** | Pointer to **int32** | Zero-based index of this page | [optional] 
**Size** | Pointer to **int32** | Requested page size | [optional] 
**Total** | Pointer to **int64** | Total number of distinct versions matching the query | [optional] 

## Methods

### NewComponentVersionsPageXO

`func NewComponentVersionsPageXO() *ComponentVersionsPageXO`

NewComponentVersionsPageXO instantiates a new ComponentVersionsPageXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentVersionsPageXOWithDefaults

`func NewComponentVersionsPageXOWithDefaults() *ComponentVersionsPageXO`

NewComponentVersionsPageXOWithDefaults instantiates a new ComponentVersionsPageXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ComponentVersionsPageXO) GetItems() []ComponentVersionXO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ComponentVersionsPageXO) GetItemsOk() (*[]ComponentVersionXO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ComponentVersionsPageXO) SetItems(v []ComponentVersionXO)`

SetItems sets Items field to given value.

### HasItems

`func (o *ComponentVersionsPageXO) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPage

`func (o *ComponentVersionsPageXO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ComponentVersionsPageXO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ComponentVersionsPageXO) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ComponentVersionsPageXO) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *ComponentVersionsPageXO) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ComponentVersionsPageXO) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ComponentVersionsPageXO) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ComponentVersionsPageXO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotal

`func (o *ComponentVersionsPageXO) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ComponentVersionsPageXO) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ComponentVersionsPageXO) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ComponentVersionsPageXO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


