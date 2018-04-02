// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSimulationOutputsParams creates a new GetSimulationOutputsParams object
// with the default values initialized.
func NewGetSimulationOutputsParams() *GetSimulationOutputsParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetSimulationOutputsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSimulationOutputsParamsWithTimeout creates a new GetSimulationOutputsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSimulationOutputsParamsWithTimeout(timeout time.Duration) *GetSimulationOutputsParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetSimulationOutputsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetSimulationOutputsParamsWithContext creates a new GetSimulationOutputsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSimulationOutputsParamsWithContext(ctx context.Context) *GetSimulationOutputsParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetSimulationOutputsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetSimulationOutputsParamsWithHTTPClient creates a new GetSimulationOutputsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSimulationOutputsParamsWithHTTPClient(client *http.Client) *GetSimulationOutputsParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetSimulationOutputsParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetSimulationOutputsParams contains all the parameters to send to the API endpoint
for the get simulation outputs operation typically these are written to a http.Request
*/
type GetSimulationOutputsParams struct {

	/*ID
	  simulation identifier

	*/
	ID int32
	/*Limit
	  number of items to return within the query

	*/
	Limit *int32
	/*Offset
	  starting paging count; ex. offset of 60 will skip the first 60 items in the list

	*/
	Offset *int32
	/*Sort
	  key:direction pairs for one or multiple field sort orders.  e.g. sort=key1:desc,key2:asc

	*/
	Sort []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get simulation outputs params
func (o *GetSimulationOutputsParams) WithTimeout(timeout time.Duration) *GetSimulationOutputsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get simulation outputs params
func (o *GetSimulationOutputsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get simulation outputs params
func (o *GetSimulationOutputsParams) WithContext(ctx context.Context) *GetSimulationOutputsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get simulation outputs params
func (o *GetSimulationOutputsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get simulation outputs params
func (o *GetSimulationOutputsParams) WithHTTPClient(client *http.Client) *GetSimulationOutputsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get simulation outputs params
func (o *GetSimulationOutputsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get simulation outputs params
func (o *GetSimulationOutputsParams) WithID(id int32) *GetSimulationOutputsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get simulation outputs params
func (o *GetSimulationOutputsParams) SetID(id int32) {
	o.ID = id
}

// WithLimit adds the limit to the get simulation outputs params
func (o *GetSimulationOutputsParams) WithLimit(limit *int32) *GetSimulationOutputsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get simulation outputs params
func (o *GetSimulationOutputsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get simulation outputs params
func (o *GetSimulationOutputsParams) WithOffset(offset *int32) *GetSimulationOutputsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get simulation outputs params
func (o *GetSimulationOutputsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSort adds the sort to the get simulation outputs params
func (o *GetSimulationOutputsParams) WithSort(sort []string) *GetSimulationOutputsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get simulation outputs params
func (o *GetSimulationOutputsParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetSimulationOutputsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesSort := o.Sort

	joinedSort := swag.JoinByFormat(valuesSort, "csv")
	// query array param sort
	if err := r.SetQueryParam("sort", joinedSort...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
