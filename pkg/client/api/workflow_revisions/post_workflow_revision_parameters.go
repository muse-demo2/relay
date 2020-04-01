// Code generated by go-swagger; DO NOT EDIT.

package workflow_revisions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostWorkflowRevisionParams creates a new PostWorkflowRevisionParams object
// with the default values initialized.
func NewPostWorkflowRevisionParams() *PostWorkflowRevisionParams {
	var ()
	return &PostWorkflowRevisionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkflowRevisionParamsWithTimeout creates a new PostWorkflowRevisionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWorkflowRevisionParamsWithTimeout(timeout time.Duration) *PostWorkflowRevisionParams {
	var ()
	return &PostWorkflowRevisionParams{

		timeout: timeout,
	}
}

// NewPostWorkflowRevisionParamsWithContext creates a new PostWorkflowRevisionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWorkflowRevisionParamsWithContext(ctx context.Context) *PostWorkflowRevisionParams {
	var ()
	return &PostWorkflowRevisionParams{

		Context: ctx,
	}
}

// NewPostWorkflowRevisionParamsWithHTTPClient creates a new PostWorkflowRevisionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWorkflowRevisionParamsWithHTTPClient(client *http.Client) *PostWorkflowRevisionParams {
	var ()
	return &PostWorkflowRevisionParams{
		HTTPClient: client,
	}
}

/*PostWorkflowRevisionParams contains all the parameters to send to the API endpoint
for the post workflow revision operation typically these are written to a http.Request
*/
type PostWorkflowRevisionParams struct {

	/*Body
	  The workflow yaml content

	*/
	Body io.ReadCloser
	/*WorkflowName
	  Workflow name

	*/
	WorkflowName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post workflow revision params
func (o *PostWorkflowRevisionParams) WithTimeout(timeout time.Duration) *PostWorkflowRevisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workflow revision params
func (o *PostWorkflowRevisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workflow revision params
func (o *PostWorkflowRevisionParams) WithContext(ctx context.Context) *PostWorkflowRevisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workflow revision params
func (o *PostWorkflowRevisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workflow revision params
func (o *PostWorkflowRevisionParams) WithHTTPClient(client *http.Client) *PostWorkflowRevisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workflow revision params
func (o *PostWorkflowRevisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workflow revision params
func (o *PostWorkflowRevisionParams) WithBody(body io.ReadCloser) *PostWorkflowRevisionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workflow revision params
func (o *PostWorkflowRevisionParams) SetBody(body io.ReadCloser) {
	o.Body = body
}

// WithWorkflowName adds the workflowName to the post workflow revision params
func (o *PostWorkflowRevisionParams) WithWorkflowName(workflowName string) *PostWorkflowRevisionParams {
	o.SetWorkflowName(workflowName)
	return o
}

// SetWorkflowName adds the workflowName to the post workflow revision params
func (o *PostWorkflowRevisionParams) SetWorkflowName(workflowName string) {
	o.WorkflowName = workflowName
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkflowRevisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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