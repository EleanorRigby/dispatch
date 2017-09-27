///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/image-manager/gen/models"
)

// GetImageByNameOKCode is the HTTP code returned for type GetImageByNameOK
const GetImageByNameOKCode int = 200

/*GetImageByNameOK successful operation

swagger:response getImageByNameOK
*/
type GetImageByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.Image `json:"body,omitempty"`
}

// NewGetImageByNameOK creates GetImageByNameOK with default headers values
func NewGetImageByNameOK() *GetImageByNameOK {
	return &GetImageByNameOK{}
}

// WithPayload adds the payload to the get image by name o k response
func (o *GetImageByNameOK) WithPayload(payload *models.Image) *GetImageByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get image by name o k response
func (o *GetImageByNameOK) SetPayload(payload *models.Image) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImageByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetImageByNameBadRequestCode is the HTTP code returned for type GetImageByNameBadRequest
const GetImageByNameBadRequestCode int = 400

/*GetImageByNameBadRequest Invalid ID supplied

swagger:response getImageByNameBadRequest
*/
type GetImageByNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetImageByNameBadRequest creates GetImageByNameBadRequest with default headers values
func NewGetImageByNameBadRequest() *GetImageByNameBadRequest {
	return &GetImageByNameBadRequest{}
}

// WithPayload adds the payload to the get image by name bad request response
func (o *GetImageByNameBadRequest) WithPayload(payload *models.Error) *GetImageByNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get image by name bad request response
func (o *GetImageByNameBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImageByNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetImageByNameNotFoundCode is the HTTP code returned for type GetImageByNameNotFound
const GetImageByNameNotFoundCode int = 404

/*GetImageByNameNotFound Image not found

swagger:response getImageByNameNotFound
*/
type GetImageByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetImageByNameNotFound creates GetImageByNameNotFound with default headers values
func NewGetImageByNameNotFound() *GetImageByNameNotFound {
	return &GetImageByNameNotFound{}
}

// WithPayload adds the payload to the get image by name not found response
func (o *GetImageByNameNotFound) WithPayload(payload *models.Error) *GetImageByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get image by name not found response
func (o *GetImageByNameNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImageByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetImageByNameDefault Generic error response

swagger:response getImageByNameDefault
*/
type GetImageByNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetImageByNameDefault creates GetImageByNameDefault with default headers values
func NewGetImageByNameDefault(code int) *GetImageByNameDefault {
	if code <= 0 {
		code = 500
	}

	return &GetImageByNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get image by name default response
func (o *GetImageByNameDefault) WithStatusCode(code int) *GetImageByNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get image by name default response
func (o *GetImageByNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get image by name default response
func (o *GetImageByNameDefault) WithPayload(payload *models.Error) *GetImageByNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get image by name default response
func (o *GetImageByNameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImageByNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}