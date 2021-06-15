// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// UpdateDynOAuth2ClientReader is a Reader for the UpdateDynOAuth2Client structure.
type UpdateDynOAuth2ClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDynOAuth2ClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDynOAuth2ClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateDynOAuth2ClientInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDynOAuth2ClientOK creates a UpdateDynOAuth2ClientOK with default headers values
func NewUpdateDynOAuth2ClientOK() *UpdateDynOAuth2ClientOK {
	return &UpdateDynOAuth2ClientOK{}
}

/* UpdateDynOAuth2ClientOK describes a response with status code 200, with default header values.

oAuth2Client
*/
type UpdateDynOAuth2ClientOK struct {
	Payload *models.OAuth2Client
}

func (o *UpdateDynOAuth2ClientOK) Error() string {
	return fmt.Sprintf("[PUT /dyn-clients/{id}][%d] updateDynOAuth2ClientOK  %+v", 200, o.Payload)
}
func (o *UpdateDynOAuth2ClientOK) GetPayload() *models.OAuth2Client {
	return o.Payload
}

func (o *UpdateDynOAuth2ClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2Client)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDynOAuth2ClientInternalServerError creates a UpdateDynOAuth2ClientInternalServerError with default headers values
func NewUpdateDynOAuth2ClientInternalServerError() *UpdateDynOAuth2ClientInternalServerError {
	return &UpdateDynOAuth2ClientInternalServerError{}
}

/* UpdateDynOAuth2ClientInternalServerError describes a response with status code 500, with default header values.

genericError
*/
type UpdateDynOAuth2ClientInternalServerError struct {
	Payload *models.GenericError
}

func (o *UpdateDynOAuth2ClientInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /dyn-clients/{id}][%d] updateDynOAuth2ClientInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateDynOAuth2ClientInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateDynOAuth2ClientInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
