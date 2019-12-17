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

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// V1ArtifactFormat Artifact format
//
// - csv: Comma
//  - tsv: Tab
//  - psv: Pipe
//  - ssv: Space
// swagger:model v1ArtifactFormat
type V1ArtifactFormat string

const (

	// V1ArtifactFormatCsv captures enum value "csv"
	V1ArtifactFormatCsv V1ArtifactFormat = "csv"

	// V1ArtifactFormatTsv captures enum value "tsv"
	V1ArtifactFormatTsv V1ArtifactFormat = "tsv"

	// V1ArtifactFormatPsv captures enum value "psv"
	V1ArtifactFormatPsv V1ArtifactFormat = "psv"

	// V1ArtifactFormatSsv captures enum value "ssv"
	V1ArtifactFormatSsv V1ArtifactFormat = "ssv"
)

// for schema
var v1ArtifactFormatEnum []interface{}

func init() {
	var res []V1ArtifactFormat
	if err := json.Unmarshal([]byte(`["csv","tsv","psv","ssv"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ArtifactFormatEnum = append(v1ArtifactFormatEnum, v)
	}
}

func (m V1ArtifactFormat) validateV1ArtifactFormatEnum(path, location string, value V1ArtifactFormat) error {
	if err := validate.Enum(path, location, value, v1ArtifactFormatEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 artifact format
func (m V1ArtifactFormat) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ArtifactFormatEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
