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

// V1RunKind Run kind enum
// swagger:model v1RunKind
type V1RunKind string

const (

	// V1RunKindJob captures enum value "job"
	V1RunKindJob V1RunKind = "job"

	// V1RunKindService captures enum value "service"
	V1RunKindService V1RunKind = "service"

	// V1RunKindDag captures enum value "dag"
	V1RunKindDag V1RunKind = "dag"

	// V1RunKindParallel captures enum value "parallel"
	V1RunKindParallel V1RunKind = "parallel"
)

// for schema
var v1RunKindEnum []interface{}

func init() {
	var res []V1RunKind
	if err := json.Unmarshal([]byte(`["job","service","dag","parallel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1RunKindEnum = append(v1RunKindEnum, v)
	}
}

func (m V1RunKind) validateV1RunKindEnum(path, location string, value V1RunKind) error {
	if err := validate.Enum(path, location, value, v1RunKindEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 run kind
func (m V1RunKind) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1RunKindEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
