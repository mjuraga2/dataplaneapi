// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetDefaultsSectionOKCode is the HTTP code returned for type GetDefaultsSectionOK
const GetDefaultsSectionOKCode int = 200

/*
GetDefaultsSectionOK Successful operation

swagger:response getDefaultsSectionOK
*/
type GetDefaultsSectionOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetDefaultsSectionOKBody `json:"body,omitempty"`
}

// NewGetDefaultsSectionOK creates GetDefaultsSectionOK with default headers values
func NewGetDefaultsSectionOK() *GetDefaultsSectionOK {

	return &GetDefaultsSectionOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get defaults section o k response
func (o *GetDefaultsSectionOK) WithConfigurationVersion(configurationVersion string) *GetDefaultsSectionOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get defaults section o k response
func (o *GetDefaultsSectionOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get defaults section o k response
func (o *GetDefaultsSectionOK) WithPayload(payload *GetDefaultsSectionOKBody) *GetDefaultsSectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get defaults section o k response
func (o *GetDefaultsSectionOK) SetPayload(payload *GetDefaultsSectionOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDefaultsSectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDefaultsSectionNotFoundCode is the HTTP code returned for type GetDefaultsSectionNotFound
const GetDefaultsSectionNotFoundCode int = 404

/*
GetDefaultsSectionNotFound The specified resource was not found

swagger:response getDefaultsSectionNotFound
*/
type GetDefaultsSectionNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDefaultsSectionNotFound creates GetDefaultsSectionNotFound with default headers values
func NewGetDefaultsSectionNotFound() *GetDefaultsSectionNotFound {

	return &GetDefaultsSectionNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get defaults section not found response
func (o *GetDefaultsSectionNotFound) WithConfigurationVersion(configurationVersion string) *GetDefaultsSectionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get defaults section not found response
func (o *GetDefaultsSectionNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get defaults section not found response
func (o *GetDefaultsSectionNotFound) WithPayload(payload *models.Error) *GetDefaultsSectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get defaults section not found response
func (o *GetDefaultsSectionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDefaultsSectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetDefaultsSectionDefault General Error

swagger:response getDefaultsSectionDefault
*/
type GetDefaultsSectionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDefaultsSectionDefault creates GetDefaultsSectionDefault with default headers values
func NewGetDefaultsSectionDefault(code int) *GetDefaultsSectionDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDefaultsSectionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get defaults section default response
func (o *GetDefaultsSectionDefault) WithStatusCode(code int) *GetDefaultsSectionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get defaults section default response
func (o *GetDefaultsSectionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get defaults section default response
func (o *GetDefaultsSectionDefault) WithConfigurationVersion(configurationVersion string) *GetDefaultsSectionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get defaults section default response
func (o *GetDefaultsSectionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get defaults section default response
func (o *GetDefaultsSectionDefault) WithPayload(payload *models.Error) *GetDefaultsSectionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get defaults section default response
func (o *GetDefaultsSectionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDefaultsSectionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}