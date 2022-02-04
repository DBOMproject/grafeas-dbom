// Code generated by go-swagger; DO NOT EDIT.

package asset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"dbom-grafeas/dbom/models"
)

// RetrieveAssetReader is a Reader for the RetrieveAsset structure.
type RetrieveAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveAssetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRetrieveAssetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewRetrieveAssetBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRetrieveAssetOK creates a RetrieveAssetOK with default headers values
func NewRetrieveAssetOK() *RetrieveAssetOK {
	return &RetrieveAssetOK{}
}

/* RetrieveAssetOK describes a response with status code 200, with default header values.

Asset has been retrieved and is in the body
*/
type RetrieveAssetOK struct {
	Payload *models.AssetDefinition
}

func (o *RetrieveAssetOK) Error() string {
	return fmt.Sprintf("[GET /repo/{repoID}/chan/{channelID}/asset/{assetID}][%d] retrieveAssetOK  %+v", 200, o.Payload)
}
func (o *RetrieveAssetOK) GetPayload() *models.AssetDefinition {
	return o.Payload
}

func (o *RetrieveAssetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssetDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveAssetNotFound creates a RetrieveAssetNotFound with default headers values
func NewRetrieveAssetNotFound() *RetrieveAssetNotFound {
	return &RetrieveAssetNotFound{}
}

/* RetrieveAssetNotFound describes a response with status code 404, with default header values.

No channel, repository or asset found for that URI
*/
type RetrieveAssetNotFound struct {
}

func (o *RetrieveAssetNotFound) Error() string {
	return fmt.Sprintf("[GET /repo/{repoID}/chan/{channelID}/asset/{assetID}][%d] retrieveAssetNotFound ", 404)
}

func (o *RetrieveAssetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrieveAssetBadGateway creates a RetrieveAssetBadGateway with default headers values
func NewRetrieveAssetBadGateway() *RetrieveAssetBadGateway {
	return &RetrieveAssetBadGateway{}
}

/* RetrieveAssetBadGateway describes a response with status code 502, with default header values.

Error on repository agent
*/
type RetrieveAssetBadGateway struct {
}

func (o *RetrieveAssetBadGateway) Error() string {
	return fmt.Sprintf("[GET /repo/{repoID}/chan/{channelID}/asset/{assetID}][%d] retrieveAssetBadGateway ", 502)
}

func (o *RetrieveAssetBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
