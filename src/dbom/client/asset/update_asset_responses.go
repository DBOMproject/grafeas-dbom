// Code generated by go-swagger; DO NOT EDIT.

package asset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateAssetReader is a Reader for the UpdateAsset structure.
type UpdateAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAssetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAssetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAssetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewUpdateAssetBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAssetOK creates a UpdateAssetOK with default headers values
func NewUpdateAssetOK() *UpdateAssetOK {
	return &UpdateAssetOK{}
}

/* UpdateAssetOK describes a response with status code 200, with default header values.

Asset was successfully created by the repository agent
*/
type UpdateAssetOK struct {
}

func (o *UpdateAssetOK) Error() string {
	return fmt.Sprintf("[PUT /repo/{repoID}/chan/{channelID}/asset/{assetID}][%d] updateAssetOK ", 200)
}

func (o *UpdateAssetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAssetBadRequest creates a UpdateAssetBadRequest with default headers values
func NewUpdateAssetBadRequest() *UpdateAssetBadRequest {
	return &UpdateAssetBadRequest{}
}

/* UpdateAssetBadRequest describes a response with status code 400, with default header values.

Expected schema for an asset was not matched
*/
type UpdateAssetBadRequest struct {
}

func (o *UpdateAssetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /repo/{repoID}/chan/{channelID}/asset/{assetID}][%d] updateAssetBadRequest ", 400)
}

func (o *UpdateAssetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAssetNotFound creates a UpdateAssetNotFound with default headers values
func NewUpdateAssetNotFound() *UpdateAssetNotFound {
	return &UpdateAssetNotFound{}
}

/* UpdateAssetNotFound describes a response with status code 404, with default header values.

No channel or repository found for that URI
*/
type UpdateAssetNotFound struct {
}

func (o *UpdateAssetNotFound) Error() string {
	return fmt.Sprintf("[PUT /repo/{repoID}/chan/{channelID}/asset/{assetID}][%d] updateAssetNotFound ", 404)
}

func (o *UpdateAssetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAssetBadGateway creates a UpdateAssetBadGateway with default headers values
func NewUpdateAssetBadGateway() *UpdateAssetBadGateway {
	return &UpdateAssetBadGateway{}
}

/* UpdateAssetBadGateway describes a response with status code 502, with default header values.

Error on repository agent
*/
type UpdateAssetBadGateway struct {
}

func (o *UpdateAssetBadGateway) Error() string {
	return fmt.Sprintf("[PUT /repo/{repoID}/chan/{channelID}/asset/{assetID}][%d] updateAssetBadGateway ", 502)
}

func (o *UpdateAssetBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
