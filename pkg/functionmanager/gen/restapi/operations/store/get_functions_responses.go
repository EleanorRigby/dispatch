///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/functionmanager/gen/models"
)

// GetFunctionsOKCode is the HTTP code returned for type GetFunctionsOK
const GetFunctionsOKCode int = 200

/*GetFunctionsOK Successful operation

swagger:response getFunctionsOK
*/
type GetFunctionsOK struct {

	/*
	  In: Body
	*/
	Payload models.GetFunctionsOKBody `json:"body,omitempty"`
}

// NewGetFunctionsOK creates GetFunctionsOK with default headers values
func NewGetFunctionsOK() *GetFunctionsOK {
	return &GetFunctionsOK{}
}

// WithPayload adds the payload to the get functions o k response
func (o *GetFunctionsOK) WithPayload(payload models.GetFunctionsOKBody) *GetFunctionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get functions o k response
func (o *GetFunctionsOK) SetPayload(payload models.GetFunctionsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFunctionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetFunctionsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetFunctionsDefault Custom error

swagger:response getFunctionsDefault
*/
type GetFunctionsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetFunctionsDefault creates GetFunctionsDefault with default headers values
func NewGetFunctionsDefault(code int) *GetFunctionsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetFunctionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get functions default response
func (o *GetFunctionsDefault) WithStatusCode(code int) *GetFunctionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get functions default response
func (o *GetFunctionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get functions default response
func (o *GetFunctionsDefault) WithPayload(payload models.Error) *GetFunctionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get functions default response
func (o *GetFunctionsDefault) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFunctionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
