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

	models "github.com/vmware/dispatch/pkg/event-manager/gen/models"
)

// UpdateSubscriptionReader is a Reader for the UpdateSubscription structure.
type UpdateSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateSubscriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateSubscriptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSubscriptionOK creates a UpdateSubscriptionOK with default headers values
func NewUpdateSubscriptionOK() *UpdateSubscriptionOK {
	return &UpdateSubscriptionOK{}
}

/*UpdateSubscriptionOK handles this case with default header values.

Successful operation
*/
type UpdateSubscriptionOK struct {
	Payload *models.Subscription
}

func (o *UpdateSubscriptionOK) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionName}][%d] updateSubscriptionOK  %+v", 200, o.Payload)
}

func (o *UpdateSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubscriptionBadRequest creates a UpdateSubscriptionBadRequest with default headers values
func NewUpdateSubscriptionBadRequest() *UpdateSubscriptionBadRequest {
	return &UpdateSubscriptionBadRequest{}
}

/*UpdateSubscriptionBadRequest handles this case with default header values.

Invalid Name supplied
*/
type UpdateSubscriptionBadRequest struct {
	Payload *models.Error
}

func (o *UpdateSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionName}][%d] updateSubscriptionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubscriptionNotFound creates a UpdateSubscriptionNotFound with default headers values
func NewUpdateSubscriptionNotFound() *UpdateSubscriptionNotFound {
	return &UpdateSubscriptionNotFound{}
}

/*UpdateSubscriptionNotFound handles this case with default header values.

Subscription not found
*/
type UpdateSubscriptionNotFound struct {
	Payload *models.Error
}

func (o *UpdateSubscriptionNotFound) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionName}][%d] updateSubscriptionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSubscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubscriptionInternalServerError creates a UpdateSubscriptionInternalServerError with default headers values
func NewUpdateSubscriptionInternalServerError() *UpdateSubscriptionInternalServerError {
	return &UpdateSubscriptionInternalServerError{}
}

/*UpdateSubscriptionInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateSubscriptionInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateSubscriptionInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionName}][%d] updateSubscriptionInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSubscriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubscriptionDefault creates a UpdateSubscriptionDefault with default headers values
func NewUpdateSubscriptionDefault(code int) *UpdateSubscriptionDefault {
	return &UpdateSubscriptionDefault{
		_statusCode: code,
	}
}

/*UpdateSubscriptionDefault handles this case with default header values.

Unknown error
*/
type UpdateSubscriptionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update subscription default response
func (o *UpdateSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSubscriptionDefault) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionName}][%d] updateSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}