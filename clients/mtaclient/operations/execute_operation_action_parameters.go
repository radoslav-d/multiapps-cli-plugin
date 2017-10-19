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

	strfmt "github.com/go-openapi/strfmt"
)

// NewExecuteOperationActionParams creates a new ExecuteOperationActionParams object
// with the default values initialized.
func NewExecuteOperationActionParams() *ExecuteOperationActionParams {
	var ()
	return &ExecuteOperationActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteOperationActionParamsWithTimeout creates a new ExecuteOperationActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExecuteOperationActionParamsWithTimeout(timeout time.Duration) *ExecuteOperationActionParams {
	var ()
	return &ExecuteOperationActionParams{

		timeout: timeout,
	}
}

// NewExecuteOperationActionParamsWithContext creates a new ExecuteOperationActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewExecuteOperationActionParamsWithContext(ctx context.Context) *ExecuteOperationActionParams {
	var ()
	return &ExecuteOperationActionParams{

		Context: ctx,
	}
}

// NewExecuteOperationActionParamsWithHTTPClient creates a new ExecuteOperationActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExecuteOperationActionParamsWithHTTPClient(client *http.Client) *ExecuteOperationActionParams {
	var ()
	return &ExecuteOperationActionParams{
		HTTPClient: client,
	}
}

/*ExecuteOperationActionParams contains all the parameters to send to the API endpoint
for the execute operation action operation typically these are written to a http.Request
*/
type ExecuteOperationActionParams struct {

	/*ActionID*/
	ActionID string
	/*OperationID*/
	OperationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the execute operation action params
func (o *ExecuteOperationActionParams) WithTimeout(timeout time.Duration) *ExecuteOperationActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute operation action params
func (o *ExecuteOperationActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute operation action params
func (o *ExecuteOperationActionParams) WithContext(ctx context.Context) *ExecuteOperationActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute operation action params
func (o *ExecuteOperationActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute operation action params
func (o *ExecuteOperationActionParams) WithHTTPClient(client *http.Client) *ExecuteOperationActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute operation action params
func (o *ExecuteOperationActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the execute operation action params
func (o *ExecuteOperationActionParams) WithActionID(actionID string) *ExecuteOperationActionParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the execute operation action params
func (o *ExecuteOperationActionParams) SetActionID(actionID string) {
	o.ActionID = actionID
}

// WithOperationID adds the operationID to the execute operation action params
func (o *ExecuteOperationActionParams) WithOperationID(operationID string) *ExecuteOperationActionParams {
	o.SetOperationID(operationID)
	return o
}

// SetOperationID adds the operationId to the execute operation action params
func (o *ExecuteOperationActionParams) SetOperationID(operationID string) {
	o.OperationID = operationID
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteOperationActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param actionId
	qrActionID := o.ActionID
	qActionID := qrActionID
	if qActionID != "" {
		if err := r.SetQueryParam("actionId", qActionID); err != nil {
			return err
		}
	}

	// path param operationId
	if err := r.SetPathParam("operationId", o.OperationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
