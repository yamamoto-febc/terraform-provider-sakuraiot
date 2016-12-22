package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetV1ProjectsParams creates a new GetV1ProjectsParams object
// with the default values initialized.
func NewGetV1ProjectsParams() *GetV1ProjectsParams {
	var ()
	return &GetV1ProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ProjectsParamsWithTimeout creates a new GetV1ProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1ProjectsParamsWithTimeout(timeout time.Duration) *GetV1ProjectsParams {
	var ()
	return &GetV1ProjectsParams{

		timeout: timeout,
	}
}

// NewGetV1ProjectsParamsWithContext creates a new GetV1ProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1ProjectsParamsWithContext(ctx context.Context) *GetV1ProjectsParams {
	var ()
	return &GetV1ProjectsParams{

		Context: ctx,
	}
}

/*GetV1ProjectsParams contains all the parameters to send to the API endpoint
for the get v1 projects operation typically these are written to a http.Request
*/
type GetV1ProjectsParams struct {

	/*Name
	  Keyword in name

	*/
	Name *string
	/*Sort
	  Sort fields

	*/
	Sort *string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get v1 projects params
func (o *GetV1ProjectsParams) WithTimeout(timeout time.Duration) *GetV1ProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 projects params
func (o *GetV1ProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 projects params
func (o *GetV1ProjectsParams) WithContext(ctx context.Context) *GetV1ProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 projects params
func (o *GetV1ProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithName adds the name to the get v1 projects params
func (o *GetV1ProjectsParams) WithName(name *string) *GetV1ProjectsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get v1 projects params
func (o *GetV1ProjectsParams) SetName(name *string) {
	o.Name = name
}

// WithSort adds the sort to the get v1 projects params
func (o *GetV1ProjectsParams) WithSort(sort *string) *GetV1ProjectsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get v1 projects params
func (o *GetV1ProjectsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
