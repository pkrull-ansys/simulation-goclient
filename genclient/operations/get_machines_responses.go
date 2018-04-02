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

// GetMachinesReader is a Reader for the GetMachines structure.
type GetMachinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMachinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetMachinesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetMachinesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetMachinesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMachinesOK creates a GetMachinesOK with default headers values
func NewGetMachinesOK() *GetMachinesOK {
	return &GetMachinesOK{}
}

/*GetMachinesOK handles this case with default header values.

Successfully retrieved list of machines
*/
type GetMachinesOK struct {
	/*Contains paging information in json format - totalCount, totalPages
	 */
	XPagination string

	Payload []*models.Machine
}

func (o *GetMachinesOK) Error() string {
	return fmt.Sprintf("[GET /machines][%d] getMachinesOK  %+v", 200, o.Payload)
}

func (o *GetMachinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Pagination
	o.XPagination = response.GetHeader("X-Pagination")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachinesUnauthorized creates a GetMachinesUnauthorized with default headers values
func NewGetMachinesUnauthorized() *GetMachinesUnauthorized {
	return &GetMachinesUnauthorized{}
}

/*GetMachinesUnauthorized handles this case with default header values.

Not authorized
*/
type GetMachinesUnauthorized struct {
	Payload *models.Error
}

func (o *GetMachinesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /machines][%d] getMachinesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMachinesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachinesForbidden creates a GetMachinesForbidden with default headers values
func NewGetMachinesForbidden() *GetMachinesForbidden {
	return &GetMachinesForbidden{}
}

/*GetMachinesForbidden handles this case with default header values.

Forbidden
*/
type GetMachinesForbidden struct {
	Payload *models.Error
}

func (o *GetMachinesForbidden) Error() string {
	return fmt.Sprintf("[GET /machines][%d] getMachinesForbidden  %+v", 403, o.Payload)
}

func (o *GetMachinesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachinesDefault creates a GetMachinesDefault with default headers values
func NewGetMachinesDefault(code int) *GetMachinesDefault {
	return &GetMachinesDefault{
		_statusCode: code,
	}
}

/*GetMachinesDefault handles this case with default header values.

unexpected error
*/
type GetMachinesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get machines default response
func (o *GetMachinesDefault) Code() int {
	return o._statusCode
}

func (o *GetMachinesDefault) Error() string {
	return fmt.Sprintf("[GET /machines][%d] getMachines default  %+v", o._statusCode, o.Payload)
}

func (o *GetMachinesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
