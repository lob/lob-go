/*
Lob

The Lob API is organized around REST. Our API is designed to have predictable, resource-oriented URLs and uses HTTP response codes to indicate any API errors. <p> Looking for our [previous documentation](https://lob.github.io/legacy-docs/)? 

API version: 1.3.0
Contact: lob-openapi@lob.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lob

import (
	"encoding/json"
	
)

// TemplateVersionWritable struct for TemplateVersionWritable
type TemplateVersionWritable struct {
	// An internal description that identifies this resource. Must be no longer than 255 characters. 
	Description NullableString `json:"description,omitempty"`
	// An HTML string of less than 100,000 characters to be used as the `published_version` of this template. See [here](#section/HTML-Examples) for guidance on designing HTML templates. Please see endpoint specific documentation for any other product-specific HTML details: - [Postcards](https://docs.lob.com/#tag/Postcards/operation/postcard_create) - `front` and `back` - [Self Mailers](https://docs.lob.com/#tag/Self-Mailers/operation/self_mailer_create) - `inside` and `outside` - [Letters](https://docs.lob.com/#tag/Letters/operation/letter_create) - `file` - [Checks](https://docs.lob.com/#tag/Checks/operation/check_create) - `check_bottom` and `attachment` - [Cards](https://docs.lob.com/#tag/Cards/operation/card_create) - `front` and `back`  If there is a syntax error with your variable names within your HTML, then an error will be thrown, e.g. using a `{{#users}}` opening tag without the corresponding closing tag `{{/users}}`. 
	Html string `json:"html"`
	Engine NullableEngineHtml `json:"engine,omitempty"`
}

// NewTemplateVersionWritable instantiates a new TemplateVersionWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateVersionWritable(html string) *TemplateVersionWritable {
	this := TemplateVersionWritable{}
	this.Html = html
	var engine EngineHtml = ENGINEHTML_LEGACY
	this.Engine = *NewNullableEngineHtml(&engine)
	return &this
}

// NewTemplateVersionWritableWithDefaults instantiates a new TemplateVersionWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateVersionWritableWithDefaults() *TemplateVersionWritable {
	this := TemplateVersionWritable{}
	var engine EngineHtml = ENGINEHTML_LEGACY
	this.Engine = *NewNullableEngineHtml(&engine)
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateVersionWritable) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateVersionWritable) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateVersionWritable) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TemplateVersionWritable) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TemplateVersionWritable) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TemplateVersionWritable) UnsetDescription() {
	o.Description.Unset()
}

// GetHtml returns the Html field value
func (o *TemplateVersionWritable) GetHtml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Html
}

// GetHtmlOk returns a tuple with the Html field value
// and a boolean to check if the value has been set.
func (o *TemplateVersionWritable) GetHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Html, true
}

// SetHtml sets field value
func (o *TemplateVersionWritable) SetHtml(v string) {
	o.Html = v
}

// GetEngine returns the Engine field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateVersionWritable) GetEngine() EngineHtml {
	if o == nil || o.Engine.Get() == nil {
		var ret EngineHtml
		return ret
	}
	return *o.Engine.Get()
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateVersionWritable) GetEngineOk() (*EngineHtml, bool) {
	if o == nil {
		return nil, false
	}
	return o.Engine.Get(), o.Engine.IsSet()
}

// HasEngine returns a boolean if a field has been set.
func (o *TemplateVersionWritable) HasEngine() bool {
	if o != nil && o.Engine.IsSet() {
		return true
	}

	return false
}

// SetEngine gets a reference to the given NullableEngineHtml and assigns it to the Engine field.
func (o *TemplateVersionWritable) SetEngine(v EngineHtml) {
	o.Engine.Set(&v)
}
// SetEngineNil sets the value for Engine to be an explicit nil
func (o *TemplateVersionWritable) SetEngineNil() {
	o.Engine.Set(nil)
}

// UnsetEngine ensures that no value is present for Engine, not even an explicit nil
func (o *TemplateVersionWritable) UnsetEngine() {
	o.Engine.Unset()
}

func (o TemplateVersionWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["html"] = o.Html
	}
	if o.Engine.IsSet() {
		toSerialize["engine"] = o.Engine.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateVersionWritable struct {
	value *TemplateVersionWritable
	isSet bool
}

func (v NullableTemplateVersionWritable) Get() *TemplateVersionWritable {
	return v.value
}

func (v *NullableTemplateVersionWritable) Set(val *TemplateVersionWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateVersionWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateVersionWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateVersionWritable(val *TemplateVersionWritable) *NullableTemplateVersionWritable {
	return &NullableTemplateVersionWritable{value: val, isSet: true}
}

func (v NullableTemplateVersionWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateVersionWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


