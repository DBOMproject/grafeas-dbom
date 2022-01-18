// Code generated by go-swagger; DO NOT EDIT.

package asset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddAssetReader is a Reader for the AddAsset structure.
type AddAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddAssetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddAssetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddAssetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewAddAssetBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddAssetCreated creates a AddAssetCreated with default headers values
func NewAddAssetCreated() *AddAssetCreated {
	return &AddAssetCreated{}
}

/* AddAssetCreated describes a response with status code 201, with default header values.

Asset was successfully created by the repository agent
*/
type AddAssetCreated struct {
}

func (o *AddAssetCreated) Error() string {
	return fmt.Sprintf("[POST /repo/{repoID}/chan/{channelID}/asset/{assetID}][%d] addAssetCreated ", 201)
}

func (o *AddAssetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddAssetBadRequest creates a AddAssetBadRequest with default headers values
func NewAddAssetBadRequest() *AddAssetBadRequest {
	return &AddAssetBadRequest{}
}

/* AddAssetBadRequest describes a response with status code 400, with default header values.

Expected schema for an asset was not matched
*/
type AddAssetBadRequest struct {
}

func (o *AddAssetBadRequest) Error() string {
	return fmt.Sprintf("[POST /repo/{repoID}/chan/{channelID}/asset/{assetID}][%d] addAssetBadRequest ", 400)
}

func (o *AddAssetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddAssetNotFound creates a AddAssetNotFound with default headers values
func NewAddAssetNotFound() *AddAssetNotFound {
	return &AddAssetNotFound{}
}

/* AddAssetNotFound describes a response with status code 404, with default header values.

No channel or repository found for that URI
*/
type AddAssetNotFound struct {
}

func (o *AddAssetNotFound) Error() string {
	return fmt.Sprintf("[POST /repo/{repoID}/chan/{channelID}/asset/{assetID}][%d] addAssetNotFound ", 404)
}

func (o *AddAssetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddAssetBadGateway creates a AddAssetBadGateway with default headers values
func NewAddAssetBadGateway() *AddAssetBadGateway {
	return &AddAssetBadGateway{}
}

/* AddAssetBadGateway describes a response with status code 502, with default header values.

Error on repository agent
*/
type AddAssetBadGateway struct {
}

func (o *AddAssetBadGateway) Error() string {
	return fmt.Sprintf("[POST /repo/{repoID}/chan/{channelID}/asset/{assetID}][%d] addAssetBadGateway ", 502)
}

func (o *AddAssetBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}