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

package http_errors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// ReplaceHTTPErrorsSectionOKCode is the HTTP code returned for type ReplaceHTTPErrorsSectionOK
const ReplaceHTTPErrorsSectionOKCode int = 200

/*
ReplaceHTTPErrorsSectionOK http-error section updated

swagger:response replaceHttpErrorsSectionOK
*/
type ReplaceHTTPErrorsSectionOK struct {

	/*
	  In: Body
	*/
	Payload *models.HTTPErrorsSection `json:"body,omitempty"`
}

// NewReplaceHTTPErrorsSectionOK creates ReplaceHTTPErrorsSectionOK with default headers values
func NewReplaceHTTPErrorsSectionOK() *ReplaceHTTPErrorsSectionOK {

	return &ReplaceHTTPErrorsSectionOK{}
}

// WithPayload adds the payload to the replace Http errors section o k response
func (o *ReplaceHTTPErrorsSectionOK) WithPayload(payload *models.HTTPErrorsSection) *ReplaceHTTPErrorsSectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http errors section o k response
func (o *ReplaceHTTPErrorsSectionOK) SetPayload(payload *models.HTTPErrorsSection) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorsSectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceHTTPErrorsSectionAcceptedCode is the HTTP code returned for type ReplaceHTTPErrorsSectionAccepted
const ReplaceHTTPErrorsSectionAcceptedCode int = 202

/*
ReplaceHTTPErrorsSectionAccepted Configuration change accepted and reload requested

swagger:response replaceHttpErrorsSectionAccepted
*/
type ReplaceHTTPErrorsSectionAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.HTTPErrorsSection `json:"body,omitempty"`
}

// NewReplaceHTTPErrorsSectionAccepted creates ReplaceHTTPErrorsSectionAccepted with default headers values
func NewReplaceHTTPErrorsSectionAccepted() *ReplaceHTTPErrorsSectionAccepted {

	return &ReplaceHTTPErrorsSectionAccepted{}
}

// WithReloadID adds the reloadId to the replace Http errors section accepted response
func (o *ReplaceHTTPErrorsSectionAccepted) WithReloadID(reloadID string) *ReplaceHTTPErrorsSectionAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace Http errors section accepted response
func (o *ReplaceHTTPErrorsSectionAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace Http errors section accepted response
func (o *ReplaceHTTPErrorsSectionAccepted) WithPayload(payload *models.HTTPErrorsSection) *ReplaceHTTPErrorsSectionAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http errors section accepted response
func (o *ReplaceHTTPErrorsSectionAccepted) SetPayload(payload *models.HTTPErrorsSection) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorsSectionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceHTTPErrorsSectionBadRequestCode is the HTTP code returned for type ReplaceHTTPErrorsSectionBadRequest
const ReplaceHTTPErrorsSectionBadRequestCode int = 400

/*
ReplaceHTTPErrorsSectionBadRequest Bad request

swagger:response replaceHttpErrorsSectionBadRequest
*/
type ReplaceHTTPErrorsSectionBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPErrorsSectionBadRequest creates ReplaceHTTPErrorsSectionBadRequest with default headers values
func NewReplaceHTTPErrorsSectionBadRequest() *ReplaceHTTPErrorsSectionBadRequest {

	return &ReplaceHTTPErrorsSectionBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace Http errors section bad request response
func (o *ReplaceHTTPErrorsSectionBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceHTTPErrorsSectionBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Http errors section bad request response
func (o *ReplaceHTTPErrorsSectionBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Http errors section bad request response
func (o *ReplaceHTTPErrorsSectionBadRequest) WithPayload(payload *models.Error) *ReplaceHTTPErrorsSectionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http errors section bad request response
func (o *ReplaceHTTPErrorsSectionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorsSectionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceHTTPErrorsSectionNotFoundCode is the HTTP code returned for type ReplaceHTTPErrorsSectionNotFound
const ReplaceHTTPErrorsSectionNotFoundCode int = 404

/*
ReplaceHTTPErrorsSectionNotFound The specified resource was not found

swagger:response replaceHttpErrorsSectionNotFound
*/
type ReplaceHTTPErrorsSectionNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPErrorsSectionNotFound creates ReplaceHTTPErrorsSectionNotFound with default headers values
func NewReplaceHTTPErrorsSectionNotFound() *ReplaceHTTPErrorsSectionNotFound {

	return &ReplaceHTTPErrorsSectionNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace Http errors section not found response
func (o *ReplaceHTTPErrorsSectionNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceHTTPErrorsSectionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Http errors section not found response
func (o *ReplaceHTTPErrorsSectionNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Http errors section not found response
func (o *ReplaceHTTPErrorsSectionNotFound) WithPayload(payload *models.Error) *ReplaceHTTPErrorsSectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http errors section not found response
func (o *ReplaceHTTPErrorsSectionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorsSectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceHTTPErrorsSectionDefault General Error

swagger:response replaceHttpErrorsSectionDefault
*/
type ReplaceHTTPErrorsSectionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPErrorsSectionDefault creates ReplaceHTTPErrorsSectionDefault with default headers values
func NewReplaceHTTPErrorsSectionDefault(code int) *ReplaceHTTPErrorsSectionDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceHTTPErrorsSectionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace HTTP errors section default response
func (o *ReplaceHTTPErrorsSectionDefault) WithStatusCode(code int) *ReplaceHTTPErrorsSectionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace HTTP errors section default response
func (o *ReplaceHTTPErrorsSectionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace HTTP errors section default response
func (o *ReplaceHTTPErrorsSectionDefault) WithConfigurationVersion(configurationVersion string) *ReplaceHTTPErrorsSectionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace HTTP errors section default response
func (o *ReplaceHTTPErrorsSectionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace HTTP errors section default response
func (o *ReplaceHTTPErrorsSectionDefault) WithPayload(payload *models.Error) *ReplaceHTTPErrorsSectionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace HTTP errors section default response
func (o *ReplaceHTTPErrorsSectionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPErrorsSectionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
