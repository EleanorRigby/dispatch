///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/function-manager/gen/models"
)

// GetRunReader is a Reader for the GetRun structure.
type GetRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRunBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetRunInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRunOK creates a GetRunOK with default headers values
func NewGetRunOK() *GetRunOK {
	return &GetRunOK{}
}

/*GetRunOK handles this case with default header values.

Function Run
*/
type GetRunOK struct {
	Payload *models.Run
}

func (o *GetRunOK) Error() string {
	return fmt.Sprintf("[GET /{functionName}/runs/{runName}][%d] getRunOK  %+v", 200, o.Payload)
}

func (o *GetRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunBadRequest creates a GetRunBadRequest with default headers values
func NewGetRunBadRequest() *GetRunBadRequest {
	return &GetRunBadRequest{}
}

/*GetRunBadRequest handles this case with default header values.

Bad Request
*/
type GetRunBadRequest struct {
	Payload *models.Error
}

func (o *GetRunBadRequest) Error() string {
	return fmt.Sprintf("[GET /{functionName}/runs/{runName}][%d] getRunBadRequest  %+v", 400, o.Payload)
}

func (o *GetRunBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunNotFound creates a GetRunNotFound with default headers values
func NewGetRunNotFound() *GetRunNotFound {
	return &GetRunNotFound{}
}

/*GetRunNotFound handles this case with default header values.

Function or Run not found
*/
type GetRunNotFound struct {
	Payload *models.Error
}

func (o *GetRunNotFound) Error() string {
	return fmt.Sprintf("[GET /{functionName}/runs/{runName}][%d] getRunNotFound  %+v", 404, o.Payload)
}

func (o *GetRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunInternalServerError creates a GetRunInternalServerError with default headers values
func NewGetRunInternalServerError() *GetRunInternalServerError {
	return &GetRunInternalServerError{}
}

/*GetRunInternalServerError handles this case with default header values.

Internal error
*/
type GetRunInternalServerError struct {
	Payload *models.Error
}

func (o *GetRunInternalServerError) Error() string {
	return fmt.Sprintf("[GET /{functionName}/runs/{runName}][%d] getRunInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRunInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
