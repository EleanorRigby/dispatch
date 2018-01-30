///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/secret-store/gen/models"
)

// GetSecretOKCode is the HTTP code returned for type GetSecretOK
const GetSecretOKCode int = 200

/*GetSecretOK The secret identified by the secretName

swagger:response getSecretOK
*/
type GetSecretOK struct {

	/*
	  In: Body
	*/
	Payload *models.Secret `json:"body,omitempty"`
}

// NewGetSecretOK creates GetSecretOK with default headers values
func NewGetSecretOK() *GetSecretOK {
	return &GetSecretOK{}
}

// WithPayload adds the payload to the get secret o k response
func (o *GetSecretOK) WithPayload(payload *models.Secret) *GetSecretOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get secret o k response
func (o *GetSecretOK) SetPayload(payload *models.Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSecretOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSecretBadRequestCode is the HTTP code returned for type GetSecretBadRequest
const GetSecretBadRequestCode int = 400

/*GetSecretBadRequest Bad Request

swagger:response getSecretBadRequest
*/
type GetSecretBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSecretBadRequest creates GetSecretBadRequest with default headers values
func NewGetSecretBadRequest() *GetSecretBadRequest {
	return &GetSecretBadRequest{}
}

// WithPayload adds the payload to the get secret bad request response
func (o *GetSecretBadRequest) WithPayload(payload *models.Error) *GetSecretBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get secret bad request response
func (o *GetSecretBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSecretBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSecretNotFoundCode is the HTTP code returned for type GetSecretNotFound
const GetSecretNotFoundCode int = 404

/*GetSecretNotFound Resource Not Found if no secret exists with the given name

swagger:response getSecretNotFound
*/
type GetSecretNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSecretNotFound creates GetSecretNotFound with default headers values
func NewGetSecretNotFound() *GetSecretNotFound {
	return &GetSecretNotFound{}
}

// WithPayload adds the payload to the get secret not found response
func (o *GetSecretNotFound) WithPayload(payload *models.Error) *GetSecretNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get secret not found response
func (o *GetSecretNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSecretNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSecretDefault Standard error

swagger:response getSecretDefault
*/
type GetSecretDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSecretDefault creates GetSecretDefault with default headers values
func NewGetSecretDefault(code int) *GetSecretDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSecretDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get secret default response
func (o *GetSecretDefault) WithStatusCode(code int) *GetSecretDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get secret default response
func (o *GetSecretDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get secret default response
func (o *GetSecretDefault) WithPayload(payload *models.Error) *GetSecretDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get secret default response
func (o *GetSecretDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSecretDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
