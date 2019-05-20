// Code generated by go-swagger; DO NOT EDIT.

package integrations_v1

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

	models "github.com/puppetlabs/nebula/pkg/client/api/models"
)

// NewCreateIntegrationParams creates a new CreateIntegrationParams object
// with the default values initialized.
func NewCreateIntegrationParams() *CreateIntegrationParams {
	var ()
	return &CreateIntegrationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateIntegrationParamsWithTimeout creates a new CreateIntegrationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateIntegrationParamsWithTimeout(timeout time.Duration) *CreateIntegrationParams {
	var ()
	return &CreateIntegrationParams{

		timeout: timeout,
	}
}

// NewCreateIntegrationParamsWithContext creates a new CreateIntegrationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateIntegrationParamsWithContext(ctx context.Context) *CreateIntegrationParams {
	var ()
	return &CreateIntegrationParams{

		Context: ctx,
	}
}

// NewCreateIntegrationParamsWithHTTPClient creates a new CreateIntegrationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateIntegrationParamsWithHTTPClient(client *http.Client) *CreateIntegrationParams {
	var ()
	return &CreateIntegrationParams{
		HTTPClient: client,
	}
}

/*CreateIntegrationParams contains all the parameters to send to the API endpoint
for the create integration operation typically these are written to a http.Request
*/
type CreateIntegrationParams struct {

	/*Accept
	  The version of the API, in this case should be "application/nebula-api.v1+json"

	*/
	Accept string
	/*Body
	  Integration to create

	*/
	Body *models.CreateIntegrationSubmission

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create integration params
func (o *CreateIntegrationParams) WithTimeout(timeout time.Duration) *CreateIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create integration params
func (o *CreateIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create integration params
func (o *CreateIntegrationParams) WithContext(ctx context.Context) *CreateIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create integration params
func (o *CreateIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create integration params
func (o *CreateIntegrationParams) WithHTTPClient(client *http.Client) *CreateIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create integration params
func (o *CreateIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccept adds the accept to the create integration params
func (o *CreateIntegrationParams) WithAccept(accept string) *CreateIntegrationParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the create integration params
func (o *CreateIntegrationParams) SetAccept(accept string) {
	o.Accept = accept
}

// WithBody adds the body to the create integration params
func (o *CreateIntegrationParams) WithBody(body *models.CreateIntegrationSubmission) *CreateIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create integration params
func (o *CreateIntegrationParams) SetBody(body *models.CreateIntegrationSubmission) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Accept
	if err := r.SetHeaderParam("Accept", o.Accept); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}