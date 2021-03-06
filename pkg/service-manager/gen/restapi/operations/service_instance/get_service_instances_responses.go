///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package service_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetServiceInstancesOKCode is the HTTP code returned for type GetServiceInstancesOK
const GetServiceInstancesOKCode int = 200

/*GetServiceInstancesOK successful operation

swagger:response getServiceInstancesOK
*/
type GetServiceInstancesOK struct {

	/*
	  In: Body
	*/
	Payload []*v1.ServiceInstance `json:"body,omitempty"`
}

// NewGetServiceInstancesOK creates GetServiceInstancesOK with default headers values
func NewGetServiceInstancesOK() *GetServiceInstancesOK {

	return &GetServiceInstancesOK{}
}

// WithPayload adds the payload to the get service instances o k response
func (o *GetServiceInstancesOK) WithPayload(payload []*v1.ServiceInstance) *GetServiceInstancesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service instances o k response
func (o *GetServiceInstancesOK) SetPayload(payload []*v1.ServiceInstance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceInstancesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*v1.ServiceInstance, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetServiceInstancesBadRequestCode is the HTTP code returned for type GetServiceInstancesBadRequest
const GetServiceInstancesBadRequestCode int = 400

/*GetServiceInstancesBadRequest Invalid input

swagger:response getServiceInstancesBadRequest
*/
type GetServiceInstancesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetServiceInstancesBadRequest creates GetServiceInstancesBadRequest with default headers values
func NewGetServiceInstancesBadRequest() *GetServiceInstancesBadRequest {

	return &GetServiceInstancesBadRequest{}
}

// WithPayload adds the payload to the get service instances bad request response
func (o *GetServiceInstancesBadRequest) WithPayload(payload *v1.Error) *GetServiceInstancesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service instances bad request response
func (o *GetServiceInstancesBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceInstancesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServiceInstancesUnauthorizedCode is the HTTP code returned for type GetServiceInstancesUnauthorized
const GetServiceInstancesUnauthorizedCode int = 401

/*GetServiceInstancesUnauthorized Unauthorized Request

swagger:response getServiceInstancesUnauthorized
*/
type GetServiceInstancesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetServiceInstancesUnauthorized creates GetServiceInstancesUnauthorized with default headers values
func NewGetServiceInstancesUnauthorized() *GetServiceInstancesUnauthorized {

	return &GetServiceInstancesUnauthorized{}
}

// WithPayload adds the payload to the get service instances unauthorized response
func (o *GetServiceInstancesUnauthorized) WithPayload(payload *v1.Error) *GetServiceInstancesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service instances unauthorized response
func (o *GetServiceInstancesUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceInstancesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServiceInstancesForbiddenCode is the HTTP code returned for type GetServiceInstancesForbidden
const GetServiceInstancesForbiddenCode int = 403

/*GetServiceInstancesForbidden access to this resource is forbidden

swagger:response getServiceInstancesForbidden
*/
type GetServiceInstancesForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetServiceInstancesForbidden creates GetServiceInstancesForbidden with default headers values
func NewGetServiceInstancesForbidden() *GetServiceInstancesForbidden {

	return &GetServiceInstancesForbidden{}
}

// WithPayload adds the payload to the get service instances forbidden response
func (o *GetServiceInstancesForbidden) WithPayload(payload *v1.Error) *GetServiceInstancesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service instances forbidden response
func (o *GetServiceInstancesForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceInstancesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetServiceInstancesDefault Generic error response

swagger:response getServiceInstancesDefault
*/
type GetServiceInstancesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetServiceInstancesDefault creates GetServiceInstancesDefault with default headers values
func NewGetServiceInstancesDefault(code int) *GetServiceInstancesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServiceInstancesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get service instances default response
func (o *GetServiceInstancesDefault) WithStatusCode(code int) *GetServiceInstancesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get service instances default response
func (o *GetServiceInstancesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get service instances default response
func (o *GetServiceInstancesDefault) WithPayload(payload *v1.Error) *GetServiceInstancesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service instances default response
func (o *GetServiceInstancesDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceInstancesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
