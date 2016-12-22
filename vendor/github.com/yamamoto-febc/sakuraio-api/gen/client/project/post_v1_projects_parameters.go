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

	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

// NewPostV1ProjectsParams creates a new PostV1ProjectsParams object
// with the default values initialized.
func NewPostV1ProjectsParams() *PostV1ProjectsParams {
	var ()
	return &PostV1ProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ProjectsParamsWithTimeout creates a new PostV1ProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostV1ProjectsParamsWithTimeout(timeout time.Duration) *PostV1ProjectsParams {
	var ()
	return &PostV1ProjectsParams{

		timeout: timeout,
	}
}

// NewPostV1ProjectsParamsWithContext creates a new PostV1ProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostV1ProjectsParamsWithContext(ctx context.Context) *PostV1ProjectsParams {
	var ()
	return &PostV1ProjectsParams{

		Context: ctx,
	}
}

/*PostV1ProjectsParams contains all the parameters to send to the API endpoint
for the post v1 projects operation typically these are written to a http.Request
*/
type PostV1ProjectsParams struct {

	/*Project
	  Project to create

	*/
	Project *models.ProjectUpdate

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post v1 projects params
func (o *PostV1ProjectsParams) WithTimeout(timeout time.Duration) *PostV1ProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 projects params
func (o *PostV1ProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 projects params
func (o *PostV1ProjectsParams) WithContext(ctx context.Context) *PostV1ProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 projects params
func (o *PostV1ProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithProject adds the project to the post v1 projects params
func (o *PostV1ProjectsParams) WithProject(project *models.ProjectUpdate) *PostV1ProjectsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the post v1 projects params
func (o *PostV1ProjectsParams) SetProject(project *models.ProjectUpdate) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Project == nil {
		o.Project = new(models.ProjectUpdate)
	}

	if err := r.SetBodyParam(o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}