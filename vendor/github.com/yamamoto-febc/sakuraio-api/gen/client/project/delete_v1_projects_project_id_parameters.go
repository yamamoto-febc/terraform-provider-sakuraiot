package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteV1ProjectsProjectIDParams creates a new DeleteV1ProjectsProjectIDParams object
// with the default values initialized.
func NewDeleteV1ProjectsProjectIDParams() *DeleteV1ProjectsProjectIDParams {
	var ()
	return &DeleteV1ProjectsProjectIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ProjectsProjectIDParamsWithTimeout creates a new DeleteV1ProjectsProjectIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteV1ProjectsProjectIDParamsWithTimeout(timeout time.Duration) *DeleteV1ProjectsProjectIDParams {
	var ()
	return &DeleteV1ProjectsProjectIDParams{

		timeout: timeout,
	}
}

// NewDeleteV1ProjectsProjectIDParamsWithContext creates a new DeleteV1ProjectsProjectIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteV1ProjectsProjectIDParamsWithContext(ctx context.Context) *DeleteV1ProjectsProjectIDParams {
	var ()
	return &DeleteV1ProjectsProjectIDParams{

		Context: ctx,
	}
}

/*DeleteV1ProjectsProjectIDParams contains all the parameters to send to the API endpoint
for the delete v1 projects project ID operation typically these are written to a http.Request
*/
type DeleteV1ProjectsProjectIDParams struct {

	/*ProjectID
	  ID of project

	*/
	ProjectID int64

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the delete v1 projects project ID params
func (o *DeleteV1ProjectsProjectIDParams) WithTimeout(timeout time.Duration) *DeleteV1ProjectsProjectIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 projects project ID params
func (o *DeleteV1ProjectsProjectIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 projects project ID params
func (o *DeleteV1ProjectsProjectIDParams) WithContext(ctx context.Context) *DeleteV1ProjectsProjectIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 projects project ID params
func (o *DeleteV1ProjectsProjectIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithProjectID adds the projectID to the delete v1 projects project ID params
func (o *DeleteV1ProjectsProjectIDParams) WithProjectID(projectID int64) *DeleteV1ProjectsProjectIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete v1 projects project ID params
func (o *DeleteV1ProjectsProjectIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ProjectsProjectIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param projectId
	if err := r.SetPathParam("projectId", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}