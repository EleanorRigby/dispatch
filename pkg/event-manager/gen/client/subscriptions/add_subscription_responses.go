///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// AddSubscriptionReader is a Reader for the AddSubscription structure.
type AddSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddSubscriptionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddSubscriptionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewAddSubscriptionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddSubscriptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddSubscriptionCreated creates a AddSubscriptionCreated with default headers values
func NewAddSubscriptionCreated() *AddSubscriptionCreated {
	return &AddSubscriptionCreated{}
}

/*AddSubscriptionCreated handles this case with default header values.

Subscription created
*/
type AddSubscriptionCreated struct {
	Payload *v1.Subscription
}

func (o *AddSubscriptionCreated) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] addSubscriptionCreated  %+v", 201, o.Payload)
}

func (o *AddSubscriptionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSubscriptionBadRequest creates a AddSubscriptionBadRequest with default headers values
func NewAddSubscriptionBadRequest() *AddSubscriptionBadRequest {
	return &AddSubscriptionBadRequest{}
}

/*AddSubscriptionBadRequest handles this case with default header values.

Invalid input
*/
type AddSubscriptionBadRequest struct {
	Payload *v1.Error
}

func (o *AddSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] addSubscriptionBadRequest  %+v", 400, o.Payload)
}

func (o *AddSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSubscriptionUnauthorized creates a AddSubscriptionUnauthorized with default headers values
func NewAddSubscriptionUnauthorized() *AddSubscriptionUnauthorized {
	return &AddSubscriptionUnauthorized{}
}

/*AddSubscriptionUnauthorized handles this case with default header values.

Unauthorized Request
*/
type AddSubscriptionUnauthorized struct {
	Payload *v1.Error
}

func (o *AddSubscriptionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] addSubscriptionUnauthorized  %+v", 401, o.Payload)
}

func (o *AddSubscriptionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSubscriptionConflict creates a AddSubscriptionConflict with default headers values
func NewAddSubscriptionConflict() *AddSubscriptionConflict {
	return &AddSubscriptionConflict{}
}

/*AddSubscriptionConflict handles this case with default header values.

Already Exists
*/
type AddSubscriptionConflict struct {
	Payload *v1.Error
}

func (o *AddSubscriptionConflict) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] addSubscriptionConflict  %+v", 409, o.Payload)
}

func (o *AddSubscriptionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSubscriptionInternalServerError creates a AddSubscriptionInternalServerError with default headers values
func NewAddSubscriptionInternalServerError() *AddSubscriptionInternalServerError {
	return &AddSubscriptionInternalServerError{}
}

/*AddSubscriptionInternalServerError handles this case with default header values.

Internal server error
*/
type AddSubscriptionInternalServerError struct {
	Payload *v1.Error
}

func (o *AddSubscriptionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] addSubscriptionInternalServerError  %+v", 500, o.Payload)
}

func (o *AddSubscriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSubscriptionDefault creates a AddSubscriptionDefault with default headers values
func NewAddSubscriptionDefault(code int) *AddSubscriptionDefault {
	return &AddSubscriptionDefault{
		_statusCode: code,
	}
}

/*AddSubscriptionDefault handles this case with default header values.

Unknown error
*/
type AddSubscriptionDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the add subscription default response
func (o *AddSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *AddSubscriptionDefault) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] addSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *AddSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}