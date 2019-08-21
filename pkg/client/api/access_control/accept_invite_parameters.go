// Code generated by go-swagger; DO NOT EDIT.

package access_control

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

// NewAcceptInviteParams creates a new AcceptInviteParams object
// with the default values initialized.
func NewAcceptInviteParams() *AcceptInviteParams {
	var ()
	return &AcceptInviteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptInviteParamsWithTimeout creates a new AcceptInviteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAcceptInviteParamsWithTimeout(timeout time.Duration) *AcceptInviteParams {
	var ()
	return &AcceptInviteParams{

		timeout: timeout,
	}
}

// NewAcceptInviteParamsWithContext creates a new AcceptInviteParams object
// with the default values initialized, and the ability to set a context for a request
func NewAcceptInviteParamsWithContext(ctx context.Context) *AcceptInviteParams {
	var ()
	return &AcceptInviteParams{

		Context: ctx,
	}
}

// NewAcceptInviteParamsWithHTTPClient creates a new AcceptInviteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAcceptInviteParamsWithHTTPClient(client *http.Client) *AcceptInviteParams {
	var ()
	return &AcceptInviteParams{
		HTTPClient: client,
	}
}

/*AcceptInviteParams contains all the parameters to send to the API endpoint
for the accept invite operation typically these are written to a http.Request
*/
type AcceptInviteParams struct {

	/*Body
	  The invitation to accept

	*/
	Body AcceptInviteBody
	/*InviteID
	  The invite ID to reference

	*/
	InviteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the accept invite params
func (o *AcceptInviteParams) WithTimeout(timeout time.Duration) *AcceptInviteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept invite params
func (o *AcceptInviteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept invite params
func (o *AcceptInviteParams) WithContext(ctx context.Context) *AcceptInviteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept invite params
func (o *AcceptInviteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept invite params
func (o *AcceptInviteParams) WithHTTPClient(client *http.Client) *AcceptInviteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept invite params
func (o *AcceptInviteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the accept invite params
func (o *AcceptInviteParams) WithBody(body AcceptInviteBody) *AcceptInviteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the accept invite params
func (o *AcceptInviteParams) SetBody(body AcceptInviteBody) {
	o.Body = body
}

// WithInviteID adds the inviteID to the accept invite params
func (o *AcceptInviteParams) WithInviteID(inviteID string) *AcceptInviteParams {
	o.SetInviteID(inviteID)
	return o
}

// SetInviteID adds the inviteId to the accept invite params
func (o *AcceptInviteParams) SetInviteID(inviteID string) {
	o.InviteID = inviteID
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptInviteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param inviteId
	if err := r.SetPathParam("inviteId", o.InviteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
