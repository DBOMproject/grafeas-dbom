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

// NewAttachSubAssetParams creates a new AttachSubAssetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAttachSubAssetParams() *AttachSubAssetParams {
	return &AttachSubAssetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAttachSubAssetParamsWithTimeout creates a new AttachSubAssetParams object
// with the ability to set a timeout on a request.
func NewAttachSubAssetParamsWithTimeout(timeout time.Duration) *AttachSubAssetParams {
	return &AttachSubAssetParams{
		timeout: timeout,
	}
}

// NewAttachSubAssetParamsWithContext creates a new AttachSubAssetParams object
// with the ability to set a context for a request.
func NewAttachSubAssetParamsWithContext(ctx context.Context) *AttachSubAssetParams {
	return &AttachSubAssetParams{
		Context: ctx,
	}
}

// NewAttachSubAssetParamsWithHTTPClient creates a new AttachSubAssetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAttachSubAssetParamsWithHTTPClient(client *http.Client) *AttachSubAssetParams {
	return &AttachSubAssetParams{
		HTTPClient: client,
	}
}

/* AttachSubAssetParams contains all the parameters to send to the API endpoint
   for the attach sub asset operation.

   Typically these are written to a http.Request.
*/
type AttachSubAssetParams struct {

	// Body.
	Body *models.AttachSubAssetRequest

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

// WithDefaults hydrates default values in the attach sub asset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachSubAssetParams) WithDefaults() *AttachSubAssetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the attach sub asset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachSubAssetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the attach sub asset params
func (o *AttachSubAssetParams) WithTimeout(timeout time.Duration) *AttachSubAssetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attach sub asset params
func (o *AttachSubAssetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attach sub asset params
func (o *AttachSubAssetParams) WithContext(ctx context.Context) *AttachSubAssetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attach sub asset params
func (o *AttachSubAssetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attach sub asset params
func (o *AttachSubAssetParams) WithHTTPClient(client *http.Client) *AttachSubAssetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attach sub asset params
func (o *AttachSubAssetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the attach sub asset params
func (o *AttachSubAssetParams) WithBody(body *models.AttachSubAssetRequest) *AttachSubAssetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the attach sub asset params
func (o *AttachSubAssetParams) SetBody(body *models.AttachSubAssetRequest) {
	o.Body = body
}

// WithAssetID adds the assetID to the attach sub asset params
func (o *AttachSubAssetParams) WithAssetID(assetID string) *AttachSubAssetParams {
	o.SetAssetID(assetID)
	return o
}

// SetAssetID adds the assetId to the attach sub asset params
func (o *AttachSubAssetParams) SetAssetID(assetID string) {
	o.AssetID = assetID
}

// WithChannelID adds the channelID to the attach sub asset params
func (o *AttachSubAssetParams) WithChannelID(channelID string) *AttachSubAssetParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the attach sub asset params
func (o *AttachSubAssetParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WithRepoID adds the repoID to the attach sub asset params
func (o *AttachSubAssetParams) WithRepoID(repoID string) *AttachSubAssetParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the attach sub asset params
func (o *AttachSubAssetParams) SetRepoID(repoID string) {
	o.RepoID = repoID
}

// WriteToRequest writes these params to a swagger request
func (o *AttachSubAssetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
