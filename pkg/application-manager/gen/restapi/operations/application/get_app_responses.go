///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetAppOKCode is the HTTP code returned for type GetAppOK
const GetAppOKCode int = 200

/*GetAppOK Successful operation

swagger:response getAppOK
*/
type GetAppOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Application `json:"body,omitempty"`
}

// NewGetAppOK creates GetAppOK with default headers values
func NewGetAppOK() *GetAppOK {

	return &GetAppOK{}
}

// WithPayload adds the payload to the get app o k response
func (o *GetAppOK) WithPayload(payload *v1.Application) *GetAppOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app o k response
func (o *GetAppOK) SetPayload(payload *v1.Application) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAppBadRequestCode is the HTTP code returned for type GetAppBadRequest
const GetAppBadRequestCode int = 400

/*GetAppBadRequest Invalid Name supplied

swagger:response getAppBadRequest
*/
type GetAppBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetAppBadRequest creates GetAppBadRequest with default headers values
func NewGetAppBadRequest() *GetAppBadRequest {

	return &GetAppBadRequest{}
}

// WithPayload adds the payload to the get app bad request response
func (o *GetAppBadRequest) WithPayload(payload *v1.Error) *GetAppBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app bad request response
func (o *GetAppBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAppNotFoundCode is the HTTP code returned for type GetAppNotFound
const GetAppNotFoundCode int = 404

/*GetAppNotFound Application not found

swagger:response getAppNotFound
*/
type GetAppNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetAppNotFound creates GetAppNotFound with default headers values
func NewGetAppNotFound() *GetAppNotFound {

	return &GetAppNotFound{}
}

// WithPayload adds the payload to the get app not found response
func (o *GetAppNotFound) WithPayload(payload *v1.Error) *GetAppNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app not found response
func (o *GetAppNotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAppInternalServerErrorCode is the HTTP code returned for type GetAppInternalServerError
const GetAppInternalServerErrorCode int = 500

/*GetAppInternalServerError Internal error

swagger:response getAppInternalServerError
*/
type GetAppInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetAppInternalServerError creates GetAppInternalServerError with default headers values
func NewGetAppInternalServerError() *GetAppInternalServerError {

	return &GetAppInternalServerError{}
}

// WithPayload adds the payload to the get app internal server error response
func (o *GetAppInternalServerError) WithPayload(payload *v1.Error) *GetAppInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app internal server error response
func (o *GetAppInternalServerError) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}