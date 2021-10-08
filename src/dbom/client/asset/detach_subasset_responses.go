// Code generated by go-swagger; DO NOT EDIT.

package asset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DetachSubassetReader is a Reader for the DetachSubasset structure.
type DetachSubassetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetachSubassetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetachSubassetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDetachSubassetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDetachSubassetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetachSubassetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewDetachSubassetBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetachSubassetOK creates a DetachSubassetOK with default headers values
func NewDetachSubassetOK() *DetachSubassetOK {
	return &DetachSubassetOK{}
}

/* DetachSubassetOK describes a response with status code 200, with default header values.

Asset was successfully unattached from the passed subasset by the repository agent
*/
type DetachSubassetOK struct {
}

func (o *DetachSubassetOK) Error() string {
	return fmt.Sprintf("[POST /repo/{repoID}/chan/{channelID}/asset/{assetID}/detach][%d] detachSubassetOK ", 200)
}

func (o *DetachSubassetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDetachSubassetBadRequest creates a DetachSubassetBadRequest with default headers values
func NewDetachSubassetBadRequest() *DetachSubassetBadRequest {
	return &DetachSubassetBadRequest{}
}

/* DetachSubassetBadRequest describes a response with status code 400, with default header values.

Expected schema for a subasset was not matched
*/
type DetachSubassetBadRequest struct {
}

func (o *DetachSubassetBadRequest) Error() string {
	return fmt.Sprintf("[POST /repo/{repoID}/chan/{channelID}/asset/{assetID}/detach][%d] detachSubassetBadRequest ", 400)
}

func (o *DetachSubassetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDetachSubassetForbidden creates a DetachSubassetForbidden with default headers values
func NewDetachSubassetForbidden() *DetachSubassetForbidden {
	return &DetachSubassetForbidden{}
}

/* DetachSubassetForbidden describes a response with status code 403, with default header values.

Forbidden to add this subasset (Possibly already linked)
*/
type DetachSubassetForbidden struct {
}

func (o *DetachSubassetForbidden) Error() string {
	return fmt.Sprintf("[POST /repo/{repoID}/chan/{channelID}/asset/{assetID}/detach][%d] detachSubassetForbidden ", 403)
}

func (o *DetachSubassetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDetachSubassetNotFound creates a DetachSubassetNotFound with default headers values
func NewDetachSubassetNotFound() *DetachSubassetNotFound {
	return &DetachSubassetNotFound{}
}

/* DetachSubassetNotFound describes a response with status code 404, with default header values.

No channel or repository found for that URI, or Subasset is not already attached to this asset
*/
type DetachSubassetNotFound struct {
}

func (o *DetachSubassetNotFound) Error() string {
	return fmt.Sprintf("[POST /repo/{repoID}/chan/{channelID}/asset/{assetID}/detach][%d] detachSubassetNotFound ", 404)
}

func (o *DetachSubassetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDetachSubassetBadGateway creates a DetachSubassetBadGateway with default headers values
func NewDetachSubassetBadGateway() *DetachSubassetBadGateway {
	return &DetachSubassetBadGateway{}
}

/* DetachSubassetBadGateway describes a response with status code 502, with default header values.

Error on repository agent
*/
type DetachSubassetBadGateway struct {
}

func (o *DetachSubassetBadGateway) Error() string {
	return fmt.Sprintf("[POST /repo/{repoID}/chan/{channelID}/asset/{assetID}/detach][%d] detachSubassetBadGateway ", 502)
}

func (o *DetachSubassetBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
