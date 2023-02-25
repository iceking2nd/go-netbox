// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

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
	"github.com/go-openapi/swag"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// NewExtrasSavedFiltersUpdateParams creates a new ExtrasSavedFiltersUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasSavedFiltersUpdateParams() *ExtrasSavedFiltersUpdateParams {
	return &ExtrasSavedFiltersUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasSavedFiltersUpdateParamsWithTimeout creates a new ExtrasSavedFiltersUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasSavedFiltersUpdateParamsWithTimeout(timeout time.Duration) *ExtrasSavedFiltersUpdateParams {
	return &ExtrasSavedFiltersUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasSavedFiltersUpdateParamsWithContext creates a new ExtrasSavedFiltersUpdateParams object
// with the ability to set a context for a request.
func NewExtrasSavedFiltersUpdateParamsWithContext(ctx context.Context) *ExtrasSavedFiltersUpdateParams {
	return &ExtrasSavedFiltersUpdateParams{
		Context: ctx,
	}
}

// NewExtrasSavedFiltersUpdateParamsWithHTTPClient creates a new ExtrasSavedFiltersUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasSavedFiltersUpdateParamsWithHTTPClient(client *http.Client) *ExtrasSavedFiltersUpdateParams {
	return &ExtrasSavedFiltersUpdateParams{
		HTTPClient: client,
	}
}

/*
ExtrasSavedFiltersUpdateParams contains all the parameters to send to the API endpoint

	for the extras saved filters update operation.

	Typically these are written to a http.Request.
*/
type ExtrasSavedFiltersUpdateParams struct {

	// Data.
	Data *models.SavedFilter

	/* ID.

	   A unique integer value identifying this saved filter.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras saved filters update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasSavedFiltersUpdateParams) WithDefaults() *ExtrasSavedFiltersUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras saved filters update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasSavedFiltersUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras saved filters update params
func (o *ExtrasSavedFiltersUpdateParams) WithTimeout(timeout time.Duration) *ExtrasSavedFiltersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras saved filters update params
func (o *ExtrasSavedFiltersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras saved filters update params
func (o *ExtrasSavedFiltersUpdateParams) WithContext(ctx context.Context) *ExtrasSavedFiltersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras saved filters update params
func (o *ExtrasSavedFiltersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras saved filters update params
func (o *ExtrasSavedFiltersUpdateParams) WithHTTPClient(client *http.Client) *ExtrasSavedFiltersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras saved filters update params
func (o *ExtrasSavedFiltersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras saved filters update params
func (o *ExtrasSavedFiltersUpdateParams) WithData(data *models.SavedFilter) *ExtrasSavedFiltersUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras saved filters update params
func (o *ExtrasSavedFiltersUpdateParams) SetData(data *models.SavedFilter) {
	o.Data = data
}

// WithID adds the id to the extras saved filters update params
func (o *ExtrasSavedFiltersUpdateParams) WithID(id int64) *ExtrasSavedFiltersUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras saved filters update params
func (o *ExtrasSavedFiltersUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasSavedFiltersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}