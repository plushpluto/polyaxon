// Copyright 2019 Polyaxon, Inc.
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

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteOrganizationMemberReader is a Reader for the DeleteOrganizationMember structure.
type DeleteOrganizationMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOrganizationMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteOrganizationMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteOrganizationMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrganizationMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOrganizationMemberOK creates a DeleteOrganizationMemberOK with default headers values
func NewDeleteOrganizationMemberOK() *DeleteOrganizationMemberOK {
	return &DeleteOrganizationMemberOK{}
}

/*DeleteOrganizationMemberOK handles this case with default header values.

A successful response.
*/
type DeleteOrganizationMemberOK struct {
}

func (o *DeleteOrganizationMemberOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organizations/{owner}/members/{user}][%d] deleteOrganizationMemberOK ", 200)
}

func (o *DeleteOrganizationMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationMemberNoContent creates a DeleteOrganizationMemberNoContent with default headers values
func NewDeleteOrganizationMemberNoContent() *DeleteOrganizationMemberNoContent {
	return &DeleteOrganizationMemberNoContent{}
}

/*DeleteOrganizationMemberNoContent handles this case with default header values.

No content.
*/
type DeleteOrganizationMemberNoContent struct {
	Payload interface{}
}

func (o *DeleteOrganizationMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organizations/{owner}/members/{user}][%d] deleteOrganizationMemberNoContent  %+v", 204, o.Payload)
}

func (o *DeleteOrganizationMemberNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOrganizationMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationMemberForbidden creates a DeleteOrganizationMemberForbidden with default headers values
func NewDeleteOrganizationMemberForbidden() *DeleteOrganizationMemberForbidden {
	return &DeleteOrganizationMemberForbidden{}
}

/*DeleteOrganizationMemberForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DeleteOrganizationMemberForbidden struct {
	Payload interface{}
}

func (o *DeleteOrganizationMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organizations/{owner}/members/{user}][%d] deleteOrganizationMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrganizationMemberForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOrganizationMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationMemberNotFound creates a DeleteOrganizationMemberNotFound with default headers values
func NewDeleteOrganizationMemberNotFound() *DeleteOrganizationMemberNotFound {
	return &DeleteOrganizationMemberNotFound{}
}

/*DeleteOrganizationMemberNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteOrganizationMemberNotFound struct {
	Payload interface{}
}

func (o *DeleteOrganizationMemberNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organizations/{owner}/members/{user}][%d] deleteOrganizationMemberNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrganizationMemberNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOrganizationMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
