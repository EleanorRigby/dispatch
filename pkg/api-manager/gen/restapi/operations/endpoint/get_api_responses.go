///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetAPIOKCode is the HTTP code returned for type GetAPIOK
const GetAPIOKCode int = 200

/*GetAPIOK Successful operation

swagger:response getApiOK
*/
type GetAPIOK struct {

	/*
	  In: Body
	*/
	Payload *v1.API `json:"body,omitempty"`
}

// NewGetAPIOK creates GetAPIOK with default headers values
func NewGetAPIOK() *GetAPIOK {

	return &GetAPIOK{}
}

// WithPayload adds the payload to the get Api o k response
func (o *GetAPIOK) WithPayload(payload *v1.API) *GetAPIOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api o k response
func (o *GetAPIOK) SetPayload(payload *v1.API) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAPIBadRequestCode is the HTTP code returned for type GetAPIBadRequest
const GetAPIBadRequestCode int = 400

/*GetAPIBadRequest Invalid Name supplied

swagger:response getApiBadRequest
*/
type GetAPIBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetAPIBadRequest creates GetAPIBadRequest with default headers values
func NewGetAPIBadRequest() *GetAPIBadRequest {

	return &GetAPIBadRequest{}
}

// WithPayload adds the payload to the get Api bad request response
func (o *GetAPIBadRequest) WithPayload(payload *v1.Error) *GetAPIBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api bad request response
func (o *GetAPIBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAPINotFoundCode is the HTTP code returned for type GetAPINotFound
const GetAPINotFoundCode int = 404

/*GetAPINotFound API not found

swagger:response getApiNotFound
*/
type GetAPINotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetAPINotFound creates GetAPINotFound with default headers values
func NewGetAPINotFound() *GetAPINotFound {

	return &GetAPINotFound{}
}

// WithPayload adds the payload to the get Api not found response
func (o *GetAPINotFound) WithPayload(payload *v1.Error) *GetAPINotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api not found response
func (o *GetAPINotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPINotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAPIInternalServerErrorCode is the HTTP code returned for type GetAPIInternalServerError
const GetAPIInternalServerErrorCode int = 500

/*GetAPIInternalServerError Internal error

swagger:response getApiInternalServerError
*/
type GetAPIInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetAPIInternalServerError creates GetAPIInternalServerError with default headers values
func NewGetAPIInternalServerError() *GetAPIInternalServerError {

	return &GetAPIInternalServerError{}
}

// WithPayload adds the payload to the get Api internal server error response
func (o *GetAPIInternalServerError) WithPayload(payload *v1.Error) *GetAPIInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api internal server error response
func (o *GetAPIInternalServerError) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}