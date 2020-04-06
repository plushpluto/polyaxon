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

package projects_v1

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
)

// NewProjectsV1RestoreProjectParams creates a new ProjectsV1RestoreProjectParams object
// with the default values initialized.
func NewProjectsV1RestoreProjectParams() *ProjectsV1RestoreProjectParams {
	var ()
	return &ProjectsV1RestoreProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsV1RestoreProjectParamsWithTimeout creates a new ProjectsV1RestoreProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsV1RestoreProjectParamsWithTimeout(timeout time.Duration) *ProjectsV1RestoreProjectParams {
	var ()
	return &ProjectsV1RestoreProjectParams{

		timeout: timeout,
	}
}

// NewProjectsV1RestoreProjectParamsWithContext creates a new ProjectsV1RestoreProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsV1RestoreProjectParamsWithContext(ctx context.Context) *ProjectsV1RestoreProjectParams {
	var ()
	return &ProjectsV1RestoreProjectParams{

		Context: ctx,
	}
}

// NewProjectsV1RestoreProjectParamsWithHTTPClient creates a new ProjectsV1RestoreProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsV1RestoreProjectParamsWithHTTPClient(client *http.Client) *ProjectsV1RestoreProjectParams {
	var ()
	return &ProjectsV1RestoreProjectParams{
		HTTPClient: client,
	}
}

/*ProjectsV1RestoreProjectParams contains all the parameters to send to the API endpoint
for the projects v1 restore project operation typically these are written to a http.Request
*/
type ProjectsV1RestoreProjectParams struct {

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

// WithTimeout adds the timeout to the projects v1 restore project params
func (o *ProjectsV1RestoreProjectParams) WithTimeout(timeout time.Duration) *ProjectsV1RestoreProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects v1 restore project params
func (o *ProjectsV1RestoreProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects v1 restore project params
func (o *ProjectsV1RestoreProjectParams) WithContext(ctx context.Context) *ProjectsV1RestoreProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects v1 restore project params
func (o *ProjectsV1RestoreProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects v1 restore project params
func (o *ProjectsV1RestoreProjectParams) WithHTTPClient(client *http.Client) *ProjectsV1RestoreProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects v1 restore project params
func (o *ProjectsV1RestoreProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the projects v1 restore project params
func (o *ProjectsV1RestoreProjectParams) WithOwner(owner string) *ProjectsV1RestoreProjectParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the projects v1 restore project params
func (o *ProjectsV1RestoreProjectParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the projects v1 restore project params
func (o *ProjectsV1RestoreProjectParams) WithProject(project string) *ProjectsV1RestoreProjectParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects v1 restore project params
func (o *ProjectsV1RestoreProjectParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsV1RestoreProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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