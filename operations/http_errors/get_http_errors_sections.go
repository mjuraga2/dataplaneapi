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
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/client-native/v4/models"
)

// GetHTTPErrorsSectionsHandlerFunc turns a function with the right signature into a get HTTP errors sections handler
type GetHTTPErrorsSectionsHandlerFunc func(GetHTTPErrorsSectionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHTTPErrorsSectionsHandlerFunc) Handle(params GetHTTPErrorsSectionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetHTTPErrorsSectionsHandler interface for that can handle valid get HTTP errors sections params
type GetHTTPErrorsSectionsHandler interface {
	Handle(GetHTTPErrorsSectionsParams, interface{}) middleware.Responder
}

// NewGetHTTPErrorsSections creates a new http.Handler for the get HTTP errors sections operation
func NewGetHTTPErrorsSections(ctx *middleware.Context, handler GetHTTPErrorsSectionsHandler) *GetHTTPErrorsSections {
	return &GetHTTPErrorsSections{Context: ctx, Handler: handler}
}

/*
	GetHTTPErrorsSections swagger:route GET /services/haproxy/configuration/http_errors_sections HTTPErrors getHttpErrorsSections

# Return an array of http-error sections

Returns an array of all configured http-error sections.
*/
type GetHTTPErrorsSections struct {
	Context *middleware.Context
	Handler GetHTTPErrorsSectionsHandler
}

func (o *GetHTTPErrorsSections) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetHTTPErrorsSectionsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetHTTPErrorsSectionsOKBody get HTTP errors sections o k body
//
// swagger:model GetHTTPErrorsSectionsOKBody
type GetHTTPErrorsSectionsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.HTTPErrorsSections `json:"data"`
}

// Validate validates this get HTTP errors sections o k body
func (o *GetHTTPErrorsSectionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPErrorsSectionsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getHttpErrorsSectionsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getHttpErrorsSectionsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getHttpErrorsSectionsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get HTTP errors sections o k body based on the context it is used
func (o *GetHTTPErrorsSectionsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPErrorsSectionsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getHttpErrorsSectionsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getHttpErrorsSectionsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHTTPErrorsSectionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHTTPErrorsSectionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetHTTPErrorsSectionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
