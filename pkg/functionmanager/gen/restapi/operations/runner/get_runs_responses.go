///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/functionmanager/gen/models"
)

// GetRunsOKCode is the HTTP code returned for type GetRunsOK
const GetRunsOKCode int = 200

/*GetRunsOK List of function runs

swagger:response getRunsOK
*/
type GetRunsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Run `json:"body,omitempty"`
}

// NewGetRunsOK creates GetRunsOK with default headers values
func NewGetRunsOK() *GetRunsOK {
	return &GetRunsOK{}
}

// WithPayload adds the payload to the get runs o k response
func (o *GetRunsOK) WithPayload(payload []*models.Run) *GetRunsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runs o k response
func (o *GetRunsOK) SetPayload(payload []*models.Run) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRunsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Run, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetRunsNotFoundCode is the HTTP code returned for type GetRunsNotFound
const GetRunsNotFoundCode int = 404

/*GetRunsNotFound Function not found

swagger:response getRunsNotFound
*/
type GetRunsNotFound struct {
}

// NewGetRunsNotFound creates GetRunsNotFound with default headers values
func NewGetRunsNotFound() *GetRunsNotFound {
	return &GetRunsNotFound{}
}

// WriteResponse to the client
func (o *GetRunsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
