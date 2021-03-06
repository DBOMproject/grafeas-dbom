// Code generated by go-swagger; DO NOT EDIT.

package asset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"dbom-grafeas/dbom/models"
)

// NewDetachSubassetParams creates a new DetachSubassetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDetachSubassetParams() *DetachSubassetParams {
	return &DetachSubassetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDetachSubassetParamsWithTimeout creates a new DetachSubassetParams object
// with the ability to set a timeout on a request.
func NewDetachSubassetParamsWithTimeout(timeout time.Duration) *DetachSubassetParams {
	return &DetachSubassetParams{
		timeout: timeout,
	}
}

// NewDetachSubassetParamsWithContext creates a new DetachSubassetParams object
// with the ability to set a context for a request.
func NewDetachSubassetParamsWithContext(ctx context.Context) *DetachSubassetParams {
	return &DetachSubassetParams{
		Context: ctx,
	}
}

// NewDetachSubassetParamsWithHTTPClient creates a new DetachSubassetParams object
// with the ability to set a custom HTTPClient for a request.
func NewDetachSubassetParamsWithHTTPClient(client *http.Client) *DetachSubassetParams {
	return &DetachSubassetParams{
		HTTPClient: client,
	}
}

/* DetachSubassetParams contains all the parameters to send to the API endpoint
   for the detach subasset operation.

   Typically these are written to a http.Request.
*/
type DetachSubassetParams struct {

	// Body.
	Body *models.DetachSubassetRequest

	/* AssetID.

	   Asset ID
	*/
	AssetID string

	/* ChannelID.

	   Channel ID
	*/
	ChannelID string

	/* RepoID.

	   Repository ID
	*/
	RepoID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the detach subasset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetachSubassetParams) WithDefaults() *DetachSubassetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the detach subasset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetachSubassetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the detach subasset params
func (o *DetachSubassetParams) WithTimeout(timeout time.Duration) *DetachSubassetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detach subasset params
func (o *DetachSubassetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detach subasset params
func (o *DetachSubassetParams) WithContext(ctx context.Context) *DetachSubassetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detach subasset params
func (o *DetachSubassetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detach subasset params
func (o *DetachSubassetParams) WithHTTPClient(client *http.Client) *DetachSubassetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detach subasset params
func (o *DetachSubassetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the detach subasset params
func (o *DetachSubassetParams) WithBody(body *models.DetachSubassetRequest) *DetachSubassetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the detach subasset params
func (o *DetachSubassetParams) SetBody(body *models.DetachSubassetRequest) {
	o.Body = body
}

// WithAssetID adds the assetID to the detach subasset params
func (o *DetachSubassetParams) WithAssetID(assetID string) *DetachSubassetParams {
	o.SetAssetID(assetID)
	return o
}

// SetAssetID adds the assetId to the detach subasset params
func (o *DetachSubassetParams) SetAssetID(assetID string) {
	o.AssetID = assetID
}

// WithChannelID adds the channelID to the detach subasset params
func (o *DetachSubassetParams) WithChannelID(channelID string) *DetachSubassetParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the detach subasset params
func (o *DetachSubassetParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WithRepoID adds the repoID to the detach subasset params
func (o *DetachSubassetParams) WithRepoID(repoID string) *DetachSubassetParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the detach subasset params
func (o *DetachSubassetParams) SetRepoID(repoID string) {
	o.RepoID = repoID
}

// WriteToRequest writes these params to a swagger request
func (o *DetachSubassetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param assetID
	if err := r.SetPathParam("assetID", o.AssetID); err != nil {
		return err
	}

	// path param channelID
	if err := r.SetPathParam("channelID", o.ChannelID); err != nil {
		return err
	}

	// path param repoID
	if err := r.SetPathParam("repoID", o.RepoID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
