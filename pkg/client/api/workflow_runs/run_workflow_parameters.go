// Code generated by go-swagger; DO NOT EDIT.

package workflow_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRunWorkflowParams creates a new RunWorkflowParams object
// with the default values initialized.
func NewRunWorkflowParams() *RunWorkflowParams {
	var ()
	return &RunWorkflowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunWorkflowParamsWithTimeout creates a new RunWorkflowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunWorkflowParamsWithTimeout(timeout time.Duration) *RunWorkflowParams {
	var ()
	return &RunWorkflowParams{

		timeout: timeout,
	}
}

// NewRunWorkflowParamsWithContext creates a new RunWorkflowParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunWorkflowParamsWithContext(ctx context.Context) *RunWorkflowParams {
	var ()
	return &RunWorkflowParams{

		Context: ctx,
	}
}

// NewRunWorkflowParamsWithHTTPClient creates a new RunWorkflowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunWorkflowParamsWithHTTPClient(client *http.Client) *RunWorkflowParams {
	var ()
	return &RunWorkflowParams{
		HTTPClient: client,
	}
}

/*RunWorkflowParams contains all the parameters to send to the API endpoint
for the run workflow operation typically these are written to a http.Request
*/
type RunWorkflowParams struct {

	/*Body
	  Workflow run to create

	*/
	Body RunWorkflowBody
	/*WorkflowName
	  Workflow name

	*/
	WorkflowName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the run workflow params
func (o *RunWorkflowParams) WithTimeout(timeout time.Duration) *RunWorkflowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the run workflow params
func (o *RunWorkflowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the run workflow params
func (o *RunWorkflowParams) WithContext(ctx context.Context) *RunWorkflowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the run workflow params
func (o *RunWorkflowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the run workflow params
func (o *RunWorkflowParams) WithHTTPClient(client *http.Client) *RunWorkflowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the run workflow params
func (o *RunWorkflowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the run workflow params
func (o *RunWorkflowParams) WithBody(body RunWorkflowBody) *RunWorkflowParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the run workflow params
func (o *RunWorkflowParams) SetBody(body RunWorkflowBody) {
	o.Body = body
}

// WithWorkflowName adds the workflowName to the run workflow params
func (o *RunWorkflowParams) WithWorkflowName(workflowName string) *RunWorkflowParams {
	o.SetWorkflowName(workflowName)
	return o
}

// SetWorkflowName adds the workflowName to the run workflow params
func (o *RunWorkflowParams) SetWorkflowName(workflowName string) {
	o.WorkflowName = workflowName
}

// WriteToRequest writes these params to a swagger request
func (o *RunWorkflowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param workflowName
	if err := r.SetPathParam("workflowName", o.WorkflowName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}