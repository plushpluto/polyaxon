// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// RunsV1ListArchivedRunsReader is a Reader for the RunsV1ListArchivedRuns structure.
type RunsV1ListArchivedRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunsV1ListArchivedRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunsV1ListArchivedRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewRunsV1ListArchivedRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRunsV1ListArchivedRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunsV1ListArchivedRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRunsV1ListArchivedRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunsV1ListArchivedRunsOK creates a RunsV1ListArchivedRunsOK with default headers values
func NewRunsV1ListArchivedRunsOK() *RunsV1ListArchivedRunsOK {
	return &RunsV1ListArchivedRunsOK{}
}

/*RunsV1ListArchivedRunsOK handles this case with default header values.

A successful response.
*/
type RunsV1ListArchivedRunsOK struct {
	Payload *service_model.V1ListRunsResponse
}

func (o *RunsV1ListArchivedRunsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] runsV1ListArchivedRunsOK  %+v", 200, o.Payload)
}

func (o *RunsV1ListArchivedRunsOK) GetPayload() *service_model.V1ListRunsResponse {
	return o.Payload
}

func (o *RunsV1ListArchivedRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1ListArchivedRunsNoContent creates a RunsV1ListArchivedRunsNoContent with default headers values
func NewRunsV1ListArchivedRunsNoContent() *RunsV1ListArchivedRunsNoContent {
	return &RunsV1ListArchivedRunsNoContent{}
}

/*RunsV1ListArchivedRunsNoContent handles this case with default header values.

No content.
*/
type RunsV1ListArchivedRunsNoContent struct {
	Payload interface{}
}

func (o *RunsV1ListArchivedRunsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] runsV1ListArchivedRunsNoContent  %+v", 204, o.Payload)
}

func (o *RunsV1ListArchivedRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1ListArchivedRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1ListArchivedRunsForbidden creates a RunsV1ListArchivedRunsForbidden with default headers values
func NewRunsV1ListArchivedRunsForbidden() *RunsV1ListArchivedRunsForbidden {
	return &RunsV1ListArchivedRunsForbidden{}
}

/*RunsV1ListArchivedRunsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type RunsV1ListArchivedRunsForbidden struct {
	Payload interface{}
}

func (o *RunsV1ListArchivedRunsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] runsV1ListArchivedRunsForbidden  %+v", 403, o.Payload)
}

func (o *RunsV1ListArchivedRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1ListArchivedRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1ListArchivedRunsNotFound creates a RunsV1ListArchivedRunsNotFound with default headers values
func NewRunsV1ListArchivedRunsNotFound() *RunsV1ListArchivedRunsNotFound {
	return &RunsV1ListArchivedRunsNotFound{}
}

/*RunsV1ListArchivedRunsNotFound handles this case with default header values.

Resource does not exist.
*/
type RunsV1ListArchivedRunsNotFound struct {
	Payload interface{}
}

func (o *RunsV1ListArchivedRunsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] runsV1ListArchivedRunsNotFound  %+v", 404, o.Payload)
}

func (o *RunsV1ListArchivedRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1ListArchivedRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1ListArchivedRunsDefault creates a RunsV1ListArchivedRunsDefault with default headers values
func NewRunsV1ListArchivedRunsDefault(code int) *RunsV1ListArchivedRunsDefault {
	return &RunsV1ListArchivedRunsDefault{
		_statusCode: code,
	}
}

/*RunsV1ListArchivedRunsDefault handles this case with default header values.

An unexpected error response
*/
type RunsV1ListArchivedRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the runs v1 list archived runs default response
func (o *RunsV1ListArchivedRunsDefault) Code() int {
	return o._statusCode
}

func (o *RunsV1ListArchivedRunsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] RunsV1_ListArchivedRuns default  %+v", o._statusCode, o.Payload)
}

func (o *RunsV1ListArchivedRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *RunsV1ListArchivedRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}