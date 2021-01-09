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

// TLSA struct for TLSA
type TLSA struct {
	// ID of DNS record
	Id int32 `json:"id"`
	// The host/subdomain the DNS record applies to
	Host string `json:"host"`
	// TTL of DNS record in seconds. Must be a multiple of 60.
	Ttl *int32 `json:"ttl,omitempty"`
	Type string `json:"type"`
	// The value of the record.
	Data string `json:"data"`
	// TLSA record certificate usage.
	Usage int32 `json:"usage"`
	// TLSA record selector.
	Selector int32 `json:"selector"`
	// TLSA record matching type.
	Dtype int32 `json:"dtype"`
}

// NewTLSA instantiates a new TLSA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSA(id int32, host string, type_ string, data string, usage int32, selector int32, dtype int32, ) *TLSA {
	this := TLSA{}
	this.Id = id
	this.Host = host
	var ttl int32 = 3600
	this.Ttl = &ttl
	this.Type = type_
	this.Data = data
	this.Usage = usage
	this.Selector = selector
	this.Dtype = dtype
	return &this
}

// NewTLSAWithDefaults instantiates a new TLSA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSAWithDefaults() *TLSA {
	this := TLSA{}
	var ttl int32 = 3600
	this.Ttl = &ttl
	return &this
}

// GetId returns the Id field value
func (o *TLSA) GetId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TLSA) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TLSA) SetId(v int32) {
	o.Id = v
}

// GetHost returns the Host field value
func (o *TLSA) GetHost() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *TLSA) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *TLSA) SetHost(v string) {
	o.Host = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *TLSA) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSA) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *TLSA) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *TLSA) SetTtl(v int32) {
	o.Ttl = &v
}

// GetType returns the Type field value
func (o *TLSA) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TLSA) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TLSA) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value
func (o *TLSA) GetData() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TLSA) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TLSA) SetData(v string) {
	o.Data = v
}

// GetUsage returns the Usage field value
func (o *TLSA) GetUsage() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *TLSA) GetUsageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *TLSA) SetUsage(v int32) {
	o.Usage = v
}

// GetSelector returns the Selector field value
func (o *TLSA) GetSelector() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *TLSA) GetSelectorOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Selector, true
}

// SetSelector sets field value
func (o *TLSA) SetSelector(v int32) {
	o.Selector = v
}

// GetDtype returns the Dtype field value
func (o *TLSA) GetDtype() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Dtype
}

// GetDtypeOk returns a tuple with the Dtype field value
// and a boolean to check if the value has been set.
func (o *TLSA) GetDtypeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Dtype, true
}

// SetDtype sets field value
func (o *TLSA) SetDtype(v int32) {
	o.Dtype = v
}

func (o TLSA) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["host"] = o.Host
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
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

type NullableTLSA struct {
	value *TLSA
	isSet bool
}

func (v NullableTLSA) Get() *TLSA {
	return v.value
}

func (v *NullableTLSA) Set(val *TLSA) {
	v.value = val
	v.isSet = true
}

func (v NullableTLSA) IsSet() bool {
	return v.isSet
}

func (v *NullableTLSA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTLSA(val *TLSA) *NullableTLSA {
	return &NullableTLSA{value: val, isSet: true}
}

func (v NullableTLSA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTLSA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


