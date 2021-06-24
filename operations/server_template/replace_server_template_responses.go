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

package server_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// ReplaceServerTemplateOKCode is the HTTP code returned for type ReplaceServerTemplateOK
const ReplaceServerTemplateOKCode int = 200

/*ReplaceServerTemplateOK Server template replaced

swagger:response replaceServerTemplateOK
*/
type ReplaceServerTemplateOK struct {

	/*
	  In: Body
	*/
	Payload *models.ServerTemplate `json:"body,omitempty"`
}

// NewReplaceServerTemplateOK creates ReplaceServerTemplateOK with default headers values
func NewReplaceServerTemplateOK() *ReplaceServerTemplateOK {

	return &ReplaceServerTemplateOK{}
}

// WithPayload adds the payload to the replace server template o k response
func (o *ReplaceServerTemplateOK) WithPayload(payload *models.ServerTemplate) *ReplaceServerTemplateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server template o k response
func (o *ReplaceServerTemplateOK) SetPayload(payload *models.ServerTemplate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerTemplateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceServerTemplateAcceptedCode is the HTTP code returned for type ReplaceServerTemplateAccepted
const ReplaceServerTemplateAcceptedCode int = 202

/*ReplaceServerTemplateAccepted Configuration change accepted and reload requested

swagger:response replaceServerTemplateAccepted
*/
type ReplaceServerTemplateAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.ServerTemplate `json:"body,omitempty"`
}

// NewReplaceServerTemplateAccepted creates ReplaceServerTemplateAccepted with default headers values
func NewReplaceServerTemplateAccepted() *ReplaceServerTemplateAccepted {

	return &ReplaceServerTemplateAccepted{}
}

// WithReloadID adds the reloadId to the replace server template accepted response
func (o *ReplaceServerTemplateAccepted) WithReloadID(reloadID string) *ReplaceServerTemplateAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace server template accepted response
func (o *ReplaceServerTemplateAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace server template accepted response
func (o *ReplaceServerTemplateAccepted) WithPayload(payload *models.ServerTemplate) *ReplaceServerTemplateAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server template accepted response
func (o *ReplaceServerTemplateAccepted) SetPayload(payload *models.ServerTemplate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerTemplateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceServerTemplateBadRequestCode is the HTTP code returned for type ReplaceServerTemplateBadRequest
const ReplaceServerTemplateBadRequestCode int = 400

/*ReplaceServerTemplateBadRequest Bad request

swagger:response replaceServerTemplateBadRequest
*/
type ReplaceServerTemplateBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerTemplateBadRequest creates ReplaceServerTemplateBadRequest with default headers values
func NewReplaceServerTemplateBadRequest() *ReplaceServerTemplateBadRequest {

	return &ReplaceServerTemplateBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace server template bad request response
func (o *ReplaceServerTemplateBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceServerTemplateBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server template bad request response
func (o *ReplaceServerTemplateBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server template bad request response
func (o *ReplaceServerTemplateBadRequest) WithPayload(payload *models.Error) *ReplaceServerTemplateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server template bad request response
func (o *ReplaceServerTemplateBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerTemplateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceServerTemplateNotFoundCode is the HTTP code returned for type ReplaceServerTemplateNotFound
const ReplaceServerTemplateNotFoundCode int = 404

/*ReplaceServerTemplateNotFound The specified resource was not found

swagger:response replaceServerTemplateNotFound
*/
type ReplaceServerTemplateNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerTemplateNotFound creates ReplaceServerTemplateNotFound with default headers values
func NewReplaceServerTemplateNotFound() *ReplaceServerTemplateNotFound {

	return &ReplaceServerTemplateNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace server template not found response
func (o *ReplaceServerTemplateNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceServerTemplateNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server template not found response
func (o *ReplaceServerTemplateNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server template not found response
func (o *ReplaceServerTemplateNotFound) WithPayload(payload *models.Error) *ReplaceServerTemplateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server template not found response
func (o *ReplaceServerTemplateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerTemplateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*ReplaceServerTemplateDefault General Error

swagger:response replaceServerTemplateDefault
*/
type ReplaceServerTemplateDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerTemplateDefault creates ReplaceServerTemplateDefault with default headers values
func NewReplaceServerTemplateDefault(code int) *ReplaceServerTemplateDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceServerTemplateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace server template default response
func (o *ReplaceServerTemplateDefault) WithStatusCode(code int) *ReplaceServerTemplateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace server template default response
func (o *ReplaceServerTemplateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace server template default response
func (o *ReplaceServerTemplateDefault) WithConfigurationVersion(configurationVersion string) *ReplaceServerTemplateDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server template default response
func (o *ReplaceServerTemplateDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server template default response
func (o *ReplaceServerTemplateDefault) WithPayload(payload *models.Error) *ReplaceServerTemplateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server template default response
func (o *ReplaceServerTemplateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerTemplateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
