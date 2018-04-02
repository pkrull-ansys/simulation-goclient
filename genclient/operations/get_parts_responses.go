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

// GetPartsReader is a Reader for the GetParts structure.
type GetPartsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPartsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPartsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetPartsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetPartsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPartsOK creates a GetPartsOK with default headers values
func NewGetPartsOK() *GetPartsOK {
	return &GetPartsOK{}
}

/*GetPartsOK handles this case with default header values.

Successfully found the list of parts
*/
type GetPartsOK struct {
	/*Contains paging information in json format - totalCount, totalPages
	 */
	XPagination string

	Payload []*models.Part
}

func (o *GetPartsOK) Error() string {
	return fmt.Sprintf("[GET /parts][%d] getPartsOK  %+v", 200, o.Payload)
}

func (o *GetPartsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Pagination
	o.XPagination = response.GetHeader("X-Pagination")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPartsForbidden creates a GetPartsForbidden with default headers values
func NewGetPartsForbidden() *GetPartsForbidden {
	return &GetPartsForbidden{}
}

/*GetPartsForbidden handles this case with default header values.

Forbidden
*/
type GetPartsForbidden struct {
	Payload *models.Error
}

func (o *GetPartsForbidden) Error() string {
	return fmt.Sprintf("[GET /parts][%d] getPartsForbidden  %+v", 403, o.Payload)
}

func (o *GetPartsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPartsDefault creates a GetPartsDefault with default headers values
func NewGetPartsDefault(code int) *GetPartsDefault {
	return &GetPartsDefault{
		_statusCode: code,
	}
}

/*GetPartsDefault handles this case with default header values.

unexpected error
*/
type GetPartsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get parts default response
func (o *GetPartsDefault) Code() int {
	return o._statusCode
}

func (o *GetPartsDefault) Error() string {
	return fmt.Sprintf("[GET /parts][%d] getParts default  %+v", o._statusCode, o.Payload)
}

func (o *GetPartsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
