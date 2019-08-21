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

// NewGetWorkflowRunsParams creates a new GetWorkflowRunsParams object
// with the default values initialized.
func NewGetWorkflowRunsParams() *GetWorkflowRunsParams {
	var ()
	return &GetWorkflowRunsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowRunsParamsWithTimeout creates a new GetWorkflowRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkflowRunsParamsWithTimeout(timeout time.Duration) *GetWorkflowRunsParams {
	var ()
	return &GetWorkflowRunsParams{

		timeout: timeout,
	}
}

// NewGetWorkflowRunsParamsWithContext creates a new GetWorkflowRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkflowRunsParamsWithContext(ctx context.Context) *GetWorkflowRunsParams {
	var ()
	return &GetWorkflowRunsParams{

		Context: ctx,
	}
}

// NewGetWorkflowRunsParamsWithHTTPClient creates a new GetWorkflowRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkflowRunsParamsWithHTTPClient(client *http.Client) *GetWorkflowRunsParams {
	var ()
	return &GetWorkflowRunsParams{
		HTTPClient: client,
	}
}

/*GetWorkflowRunsParams contains all the parameters to send to the API endpoint
for the get workflow runs operation typically these are written to a http.Request
*/
type GetWorkflowRunsParams struct {

	/*WorkflowName
	  Workflow name

	*/
	WorkflowName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workflow runs params
func (o *GetWorkflowRunsParams) WithTimeout(timeout time.Duration) *GetWorkflowRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow runs params
func (o *GetWorkflowRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow runs params
func (o *GetWorkflowRunsParams) WithContext(ctx context.Context) *GetWorkflowRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow runs params
func (o *GetWorkflowRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow runs params
func (o *GetWorkflowRunsParams) WithHTTPClient(client *http.Client) *GetWorkflowRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow runs params
func (o *GetWorkflowRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkflowName adds the workflowName to the get workflow runs params
func (o *GetWorkflowRunsParams) WithWorkflowName(workflowName string) *GetWorkflowRunsParams {
	o.SetWorkflowName(workflowName)
	return o
}

// SetWorkflowName adds the workflowName to the get workflow runs params
func (o *GetWorkflowRunsParams) SetWorkflowName(workflowName string) {
	o.WorkflowName = workflowName
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workflowName
	if err := r.SetPathParam("workflowName", o.WorkflowName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
