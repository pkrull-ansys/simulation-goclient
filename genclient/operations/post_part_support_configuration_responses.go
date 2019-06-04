// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// PostPartSupportConfigurationReader is a Reader for the PostPartSupportConfiguration structure.
type PostPartSupportConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPartSupportConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPartSupportConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPostPartSupportConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostPartSupportConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostPartSupportConfigurationOK creates a PostPartSupportConfigurationOK with default headers values
func NewPostPartSupportConfigurationOK() *PostPartSupportConfigurationOK {
	return &PostPartSupportConfigurationOK{}
}

/*PostPartSupportConfigurationOK handles this case with default header values.

Successfully added a part support configuration
*/
type PostPartSupportConfigurationOK struct {
	Payload *models.PartSupportConfiguration
}

func (o *PostPartSupportConfigurationOK) Error() string {
	return fmt.Sprintf("[POST /parts/{partId}/supportconfigurations][%d] postPartSupportConfigurationOK  %+v", 200, o.Payload)
}

func (o *PostPartSupportConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PartSupportConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPartSupportConfigurationForbidden creates a PostPartSupportConfigurationForbidden with default headers values
func NewPostPartSupportConfigurationForbidden() *PostPartSupportConfigurationForbidden {
	return &PostPartSupportConfigurationForbidden{}
}

/*PostPartSupportConfigurationForbidden handles this case with default header values.

Forbidden
*/
type PostPartSupportConfigurationForbidden struct {
	Payload *models.Error
}

func (o *PostPartSupportConfigurationForbidden) Error() string {
	return fmt.Sprintf("[POST /parts/{partId}/supportconfigurations][%d] postPartSupportConfigurationForbidden  %+v", 403, o.Payload)
}

func (o *PostPartSupportConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPartSupportConfigurationDefault creates a PostPartSupportConfigurationDefault with default headers values
func NewPostPartSupportConfigurationDefault(code int) *PostPartSupportConfigurationDefault {
	return &PostPartSupportConfigurationDefault{
		_statusCode: code,
	}
}

/*PostPartSupportConfigurationDefault handles this case with default header values.

unexpected error
*/
type PostPartSupportConfigurationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post part support configuration default response
func (o *PostPartSupportConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *PostPartSupportConfigurationDefault) Error() string {
	return fmt.Sprintf("[POST /parts/{partId}/supportconfigurations][%d] postPartSupportConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *PostPartSupportConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
