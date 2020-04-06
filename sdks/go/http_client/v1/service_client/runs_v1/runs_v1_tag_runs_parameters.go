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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewRunsV1TagRunsParams creates a new RunsV1TagRunsParams object
// with the default values initialized.
func NewRunsV1TagRunsParams() *RunsV1TagRunsParams {
	var ()
	return &RunsV1TagRunsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunsV1TagRunsParamsWithTimeout creates a new RunsV1TagRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunsV1TagRunsParamsWithTimeout(timeout time.Duration) *RunsV1TagRunsParams {
	var ()
	return &RunsV1TagRunsParams{

		timeout: timeout,
	}
}

// NewRunsV1TagRunsParamsWithContext creates a new RunsV1TagRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunsV1TagRunsParamsWithContext(ctx context.Context) *RunsV1TagRunsParams {
	var ()
	return &RunsV1TagRunsParams{

		Context: ctx,
	}
}

// NewRunsV1TagRunsParamsWithHTTPClient creates a new RunsV1TagRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunsV1TagRunsParamsWithHTTPClient(client *http.Client) *RunsV1TagRunsParams {
	var ()
	return &RunsV1TagRunsParams{
		HTTPClient: client,
	}
}

/*RunsV1TagRunsParams contains all the parameters to send to the API endpoint
for the runs v1 tag runs operation typically these are written to a http.Request
*/
type RunsV1TagRunsParams struct {

	/*Body
	  Uuids of the entities

	*/
	Body *service_model.V1Uuids
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project under namesapce

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the runs v1 tag runs params
func (o *RunsV1TagRunsParams) WithTimeout(timeout time.Duration) *RunsV1TagRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the runs v1 tag runs params
func (o *RunsV1TagRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the runs v1 tag runs params
func (o *RunsV1TagRunsParams) WithContext(ctx context.Context) *RunsV1TagRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the runs v1 tag runs params
func (o *RunsV1TagRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the runs v1 tag runs params
func (o *RunsV1TagRunsParams) WithHTTPClient(client *http.Client) *RunsV1TagRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the runs v1 tag runs params
func (o *RunsV1TagRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the runs v1 tag runs params
func (o *RunsV1TagRunsParams) WithBody(body *service_model.V1Uuids) *RunsV1TagRunsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the runs v1 tag runs params
func (o *RunsV1TagRunsParams) SetBody(body *service_model.V1Uuids) {
	o.Body = body
}

// WithOwner adds the owner to the runs v1 tag runs params
func (o *RunsV1TagRunsParams) WithOwner(owner string) *RunsV1TagRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the runs v1 tag runs params
func (o *RunsV1TagRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the runs v1 tag runs params
func (o *RunsV1TagRunsParams) WithProject(project string) *RunsV1TagRunsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the runs v1 tag runs params
func (o *RunsV1TagRunsParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *RunsV1TagRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}