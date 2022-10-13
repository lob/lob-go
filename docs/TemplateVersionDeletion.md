# TemplateVersionDeletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier prefixed with &#x60;vrsn_&#x60;. | [optional] 
**Deleted** | Pointer to **bool** | Only returned if the resource has been successfully deleted. | [optional] 
**Object** | Pointer to **string** | Value is type of resource. | [optional] [default to "template_version_deleted"]

## Methods

### NewTemplateVersionDeletion

`func NewTemplateVersionDeletion() *TemplateVersionDeletion`

NewTemplateVersionDeletion instantiates a new TemplateVersionDeletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateVersionDeletionWithDefaults

`func NewTemplateVersionDeletionWithDefaults() *TemplateVersionDeletion`

NewTemplateVersionDeletionWithDefaults instantiates a new TemplateVersionDeletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateVersionDeletion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateVersionDeletion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateVersionDeletion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateVersionDeletion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *TemplateVersionDeletion) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *TemplateVersionDeletion) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *TemplateVersionDeletion) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *TemplateVersionDeletion) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetObject

`func (o *TemplateVersionDeletion) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TemplateVersionDeletion) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TemplateVersionDeletion) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *TemplateVersionDeletion) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


