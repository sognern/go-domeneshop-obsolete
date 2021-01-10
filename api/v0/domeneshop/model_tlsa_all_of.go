/*
 * Domeneshop API Documentation
 *
 * # Overview  Domeneshop offers a simple, REST-based API, which currently supports the following features:  - List domains - List invoices - Create, read, update and delete DNS records for domains - Create, read, update and delete HTTP forwards (\"WWW forwarding\") for domains - Dynamic DNS (DDNS) update endpoints for use in consumer routers  More features are planned, including:  - Web hosting administration - Email address and email user/account administration  # Testing period  The API service is in version 0, which means it is possible that the interface will change rapidly during the testing period. For that reason, **the documentation on these pages may sometimes be outdated.**  Additionally, we make no guarantees about the stability of the API service during this testing period, and therefore ask customers to be careful with using the service for business critical purposes.  # Authentication  The Domeneshop API currently supports only one method of authentication, **HTTP Basic Auth**. More authentication methods may be added in the future.  To generate credentials, visit <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">this page</a> after logging in to the control panel on our website:  <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">https://www.domeneshop.no/admin?view=api</a>  # Libraries  Domeneshop maintains multiple API libraries to simplify using the API. Please note that these libraries have the same stability guarantees to the API while the API is in version 0.  The libraries may be found in our [Github repository](https://github.com/domeneshop/).  Domeneshop also maintains a plugin for [EFF's Certbot](https://certbot.eff.org/), which automates issuance and renewal of SSL-certificates on your own servers for domains that use Domeneshop's DNS service. This plugin is found in our Github repository [here](https://github.com/domeneshop/certbot-dns-domeneshop).  <SecurityDefinitions /> 
 *
 * API version: v0
 * Contact: kundeservice@domeneshop.no
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package domeneshop

import (
	"encoding/json"
)

// TLSAAllOf struct for TLSAAllOf
type TLSAAllOf struct {
	Type string `json:"type"`
	// The value of the record.
	Data string `json:"data"`
	// TLSA record certificate usage.
	Usage string `json:"usage"`
	// TLSA record selector.
	Selector string `json:"selector"`
	// TLSA record matching type.
	Dtype string `json:"dtype"`
}

// NewTLSAAllOf instantiates a new TLSAAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSAAllOf(type_ string, data string, usage string, selector string, dtype string, ) *TLSAAllOf {
	this := TLSAAllOf{}
	this.Type = type_
	this.Data = data
	this.Usage = usage
	this.Selector = selector
	this.Dtype = dtype
	return &this
}

// NewTLSAAllOfWithDefaults instantiates a new TLSAAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSAAllOfWithDefaults() *TLSAAllOf {
	this := TLSAAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *TLSAAllOf) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TLSAAllOf) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TLSAAllOf) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value
func (o *TLSAAllOf) GetData() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TLSAAllOf) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TLSAAllOf) SetData(v string) {
	o.Data = v
}

// GetUsage returns the Usage field value
func (o *TLSAAllOf) GetUsage() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *TLSAAllOf) GetUsageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *TLSAAllOf) SetUsage(v string) {
	o.Usage = v
}

// GetSelector returns the Selector field value
func (o *TLSAAllOf) GetSelector() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *TLSAAllOf) GetSelectorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Selector, true
}

// SetSelector sets field value
func (o *TLSAAllOf) SetSelector(v string) {
	o.Selector = v
}

// GetDtype returns the Dtype field value
func (o *TLSAAllOf) GetDtype() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Dtype
}

// GetDtypeOk returns a tuple with the Dtype field value
// and a boolean to check if the value has been set.
func (o *TLSAAllOf) GetDtypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Dtype, true
}

// SetDtype sets field value
func (o *TLSAAllOf) SetDtype(v string) {
	o.Dtype = v
}

func (o TLSAAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["usage"] = o.Usage
	}
	if true {
		toSerialize["selector"] = o.Selector
	}
	if true {
		toSerialize["dtype"] = o.Dtype
	}
	return json.Marshal(toSerialize)
}

type NullableTLSAAllOf struct {
	value *TLSAAllOf
	isSet bool
}

func (v NullableTLSAAllOf) Get() *TLSAAllOf {
	return v.value
}

func (v *NullableTLSAAllOf) Set(val *TLSAAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTLSAAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTLSAAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTLSAAllOf(val *TLSAAllOf) *NullableTLSAAllOf {
	return &NullableTLSAAllOf{value: val, isSet: true}
}

func (v NullableTLSAAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTLSAAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


