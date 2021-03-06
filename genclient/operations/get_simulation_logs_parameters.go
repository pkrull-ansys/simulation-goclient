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

// NewGetSimulationLogsParams creates a new GetSimulationLogsParams object
// with the default values initialized.
func NewGetSimulationLogsParams() *GetSimulationLogsParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetSimulationLogsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSimulationLogsParamsWithTimeout creates a new GetSimulationLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSimulationLogsParamsWithTimeout(timeout time.Duration) *GetSimulationLogsParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetSimulationLogsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetSimulationLogsParamsWithContext creates a new GetSimulationLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSimulationLogsParamsWithContext(ctx context.Context) *GetSimulationLogsParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetSimulationLogsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetSimulationLogsParamsWithHTTPClient creates a new GetSimulationLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSimulationLogsParamsWithHTTPClient(client *http.Client) *GetSimulationLogsParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetSimulationLogsParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetSimulationLogsParams contains all the parameters to send to the API endpoint
for the get simulation logs operation typically these are written to a http.Request
*/
type GetSimulationLogsParams struct {

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

// WithTimeout adds the timeout to the get simulation logs params
func (o *GetSimulationLogsParams) WithTimeout(timeout time.Duration) *GetSimulationLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get simulation logs params
func (o *GetSimulationLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get simulation logs params
func (o *GetSimulationLogsParams) WithContext(ctx context.Context) *GetSimulationLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get simulation logs params
func (o *GetSimulationLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get simulation logs params
func (o *GetSimulationLogsParams) WithHTTPClient(client *http.Client) *GetSimulationLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get simulation logs params
func (o *GetSimulationLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get simulation logs params
func (o *GetSimulationLogsParams) WithID(id int32) *GetSimulationLogsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get simulation logs params
func (o *GetSimulationLogsParams) SetID(id int32) {
	o.ID = id
}

// WithLimit adds the limit to the get simulation logs params
func (o *GetSimulationLogsParams) WithLimit(limit *int32) *GetSimulationLogsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get simulation logs params
func (o *GetSimulationLogsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get simulation logs params
func (o *GetSimulationLogsParams) WithOffset(offset *int32) *GetSimulationLogsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get simulation logs params
func (o *GetSimulationLogsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSort adds the sort to the get simulation logs params
func (o *GetSimulationLogsParams) WithSort(sort []string) *GetSimulationLogsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get simulation logs params
func (o *GetSimulationLogsParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetSimulationLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
