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

// GetMicrostructureSimulationReader is a Reader for the GetMicrostructureSimulation structure.
type GetMicrostructureSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMicrostructureSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMicrostructureSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetMicrostructureSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetMicrostructureSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetMicrostructureSimulationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetMicrostructureSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMicrostructureSimulationOK creates a GetMicrostructureSimulationOK with default headers values
func NewGetMicrostructureSimulationOK() *GetMicrostructureSimulationOK {
	return &GetMicrostructureSimulationOK{}
}

/*GetMicrostructureSimulationOK handles this case with default header values.

Successfully retrieved simulation
*/
type GetMicrostructureSimulationOK struct {
	Payload *models.MicrostructureSimulation
}

func (o *GetMicrostructureSimulationOK) Error() string {
	return fmt.Sprintf("[GET /microstructuresimulations/{id}][%d] getMicrostructureSimulationOK  %+v", 200, o.Payload)
}

func (o *GetMicrostructureSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MicrostructureSimulation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicrostructureSimulationUnauthorized creates a GetMicrostructureSimulationUnauthorized with default headers values
func NewGetMicrostructureSimulationUnauthorized() *GetMicrostructureSimulationUnauthorized {
	return &GetMicrostructureSimulationUnauthorized{}
}

/*GetMicrostructureSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type GetMicrostructureSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *GetMicrostructureSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /microstructuresimulations/{id}][%d] getMicrostructureSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMicrostructureSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicrostructureSimulationForbidden creates a GetMicrostructureSimulationForbidden with default headers values
func NewGetMicrostructureSimulationForbidden() *GetMicrostructureSimulationForbidden {
	return &GetMicrostructureSimulationForbidden{}
}

/*GetMicrostructureSimulationForbidden handles this case with default header values.

Forbidden
*/
type GetMicrostructureSimulationForbidden struct {
	Payload *models.Error
}

func (o *GetMicrostructureSimulationForbidden) Error() string {
	return fmt.Sprintf("[GET /microstructuresimulations/{id}][%d] getMicrostructureSimulationForbidden  %+v", 403, o.Payload)
}

func (o *GetMicrostructureSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicrostructureSimulationNotFound creates a GetMicrostructureSimulationNotFound with default headers values
func NewGetMicrostructureSimulationNotFound() *GetMicrostructureSimulationNotFound {
	return &GetMicrostructureSimulationNotFound{}
}

/*GetMicrostructureSimulationNotFound handles this case with default header values.

Simulation not found (id invalid)
*/
type GetMicrostructureSimulationNotFound struct {
	Payload *models.Error
}

func (o *GetMicrostructureSimulationNotFound) Error() string {
	return fmt.Sprintf("[GET /microstructuresimulations/{id}][%d] getMicrostructureSimulationNotFound  %+v", 404, o.Payload)
}

func (o *GetMicrostructureSimulationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicrostructureSimulationDefault creates a GetMicrostructureSimulationDefault with default headers values
func NewGetMicrostructureSimulationDefault(code int) *GetMicrostructureSimulationDefault {
	return &GetMicrostructureSimulationDefault{
		_statusCode: code,
	}
}

/*GetMicrostructureSimulationDefault handles this case with default header values.

unexpected error
*/
type GetMicrostructureSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get microstructure simulation default response
func (o *GetMicrostructureSimulationDefault) Code() int {
	return o._statusCode
}

func (o *GetMicrostructureSimulationDefault) Error() string {
	return fmt.Sprintf("[GET /microstructuresimulations/{id}][%d] getMicrostructureSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *GetMicrostructureSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
