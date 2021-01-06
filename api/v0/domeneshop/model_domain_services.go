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

// DomainServices Domain services
type DomainServices struct {
	Registrar *bool `json:"registrar,omitempty"`
	Dns *bool `json:"dns,omitempty"`
	Email *bool `json:"email,omitempty"`
	Webhotel *string `json:"webhotel,omitempty"`
}

// NewDomainServices instantiates a new DomainServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainServices() *DomainServices {
	this := DomainServices{}
	return &this
}

// NewDomainServicesWithDefaults instantiates a new DomainServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainServicesWithDefaults() *DomainServices {
	this := DomainServices{}
	return &this
}

// GetRegistrar returns the Registrar field value if set, zero value otherwise.
func (o *DomainServices) GetRegistrar() bool {
	if o == nil || o.Registrar == nil {
		var ret bool
		return ret
	}
	return *o.Registrar
}

// GetRegistrarOk returns a tuple with the Registrar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainServices) GetRegistrarOk() (*bool, bool) {
	if o == nil || o.Registrar == nil {
		return nil, false
	}
	return o.Registrar, true
}

// HasRegistrar returns a boolean if a field has been set.
func (o *DomainServices) HasRegistrar() bool {
	if o != nil && o.Registrar != nil {
		return true
	}

	return false
}

// SetRegistrar gets a reference to the given bool and assigns it to the Registrar field.
func (o *DomainServices) SetRegistrar(v bool) {
	o.Registrar = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *DomainServices) GetDns() bool {
	if o == nil || o.Dns == nil {
		var ret bool
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainServices) GetDnsOk() (*bool, bool) {
	if o == nil || o.Dns == nil {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *DomainServices) HasDns() bool {
	if o != nil && o.Dns != nil {
		return true
	}

	return false
}

// SetDns gets a reference to the given bool and assigns it to the Dns field.
func (o *DomainServices) SetDns(v bool) {
	o.Dns = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *DomainServices) GetEmail() bool {
	if o == nil || o.Email == nil {
		var ret bool
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainServices) GetEmailOk() (*bool, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *DomainServices) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given bool and assigns it to the Email field.
func (o *DomainServices) SetEmail(v bool) {
	o.Email = &v
}

// GetWebhotel returns the Webhotel field value if set, zero value otherwise.
func (o *DomainServices) GetWebhotel() string {
	if o == nil || o.Webhotel == nil {
		var ret string
		return ret
	}
	return *o.Webhotel
}

// GetWebhotelOk returns a tuple with the Webhotel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainServices) GetWebhotelOk() (*string, bool) {
	if o == nil || o.Webhotel == nil {
		return nil, false
	}
	return o.Webhotel, true
}

// HasWebhotel returns a boolean if a field has been set.
func (o *DomainServices) HasWebhotel() bool {
	if o != nil && o.Webhotel != nil {
		return true
	}

	return false
}

// SetWebhotel gets a reference to the given string and assigns it to the Webhotel field.
func (o *DomainServices) SetWebhotel(v string) {
	o.Webhotel = &v
}

func (o DomainServices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Registrar != nil {
		toSerialize["registrar"] = o.Registrar
	}
	if o.Dns != nil {
		toSerialize["dns"] = o.Dns
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Webhotel != nil {
		toSerialize["webhotel"] = o.Webhotel
	}
	return json.Marshal(toSerialize)
}

type NullableDomainServices struct {
	value *DomainServices
	isSet bool
}

func (v NullableDomainServices) Get() *DomainServices {
	return v.value
}

func (v *NullableDomainServices) Set(val *DomainServices) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainServices) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainServices(val *DomainServices) *NullableDomainServices {
	return &NullableDomainServices{value: val, isSet: true}
}

func (v NullableDomainServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


