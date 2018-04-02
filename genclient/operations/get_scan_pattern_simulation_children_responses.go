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

// GetScanPatternSimulationChildrenReader is a Reader for the GetScanPatternSimulationChildren structure.
type GetScanPatternSimulationChildrenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScanPatternSimulationChildrenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetScanPatternSimulationChildrenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetScanPatternSimulationChildrenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetScanPatternSimulationChildrenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetScanPatternSimulationChildrenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetScanPatternSimulationChildrenOK creates a GetScanPatternSimulationChildrenOK with default headers values
func NewGetScanPatternSimulationChildrenOK() *GetScanPatternSimulationChildrenOK {
	return &GetScanPatternSimulationChildrenOK{}
}

/*GetScanPatternSimulationChildrenOK handles this case with default header values.

Successfully found the list of items
*/
type GetScanPatternSimulationChildrenOK struct {
	/*Contains paging information in json format - totalCount, totalPages
	 */
	XPagination string

	Payload []*models.ScanPatternSimulation
}

func (o *GetScanPatternSimulationChildrenOK) Error() string {
	return fmt.Sprintf("[GET /scanpatternsimulations/{id}/children][%d] getScanPatternSimulationChildrenOK  %+v", 200, o.Payload)
}

func (o *GetScanPatternSimulationChildrenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Pagination
	o.XPagination = response.GetHeader("X-Pagination")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanPatternSimulationChildrenUnauthorized creates a GetScanPatternSimulationChildrenUnauthorized with default headers values
func NewGetScanPatternSimulationChildrenUnauthorized() *GetScanPatternSimulationChildrenUnauthorized {
	return &GetScanPatternSimulationChildrenUnauthorized{}
}

/*GetScanPatternSimulationChildrenUnauthorized handles this case with default header values.

Not authorized
*/
type GetScanPatternSimulationChildrenUnauthorized struct {
	Payload *models.Error
}

func (o *GetScanPatternSimulationChildrenUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scanpatternsimulations/{id}/children][%d] getScanPatternSimulationChildrenUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScanPatternSimulationChildrenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanPatternSimulationChildrenForbidden creates a GetScanPatternSimulationChildrenForbidden with default headers values
func NewGetScanPatternSimulationChildrenForbidden() *GetScanPatternSimulationChildrenForbidden {
	return &GetScanPatternSimulationChildrenForbidden{}
}

/*GetScanPatternSimulationChildrenForbidden handles this case with default header values.

Forbidden
*/
type GetScanPatternSimulationChildrenForbidden struct {
	Payload *models.Error
}

func (o *GetScanPatternSimulationChildrenForbidden) Error() string {
	return fmt.Sprintf("[GET /scanpatternsimulations/{id}/children][%d] getScanPatternSimulationChildrenForbidden  %+v", 403, o.Payload)
}

func (o *GetScanPatternSimulationChildrenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanPatternSimulationChildrenDefault creates a GetScanPatternSimulationChildrenDefault with default headers values
func NewGetScanPatternSimulationChildrenDefault(code int) *GetScanPatternSimulationChildrenDefault {
	return &GetScanPatternSimulationChildrenDefault{
		_statusCode: code,
	}
}

/*GetScanPatternSimulationChildrenDefault handles this case with default header values.

unexpected error
*/
type GetScanPatternSimulationChildrenDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get scan pattern simulation children default response
func (o *GetScanPatternSimulationChildrenDefault) Code() int {
	return o._statusCode
}

func (o *GetScanPatternSimulationChildrenDefault) Error() string {
	return fmt.Sprintf("[GET /scanpatternsimulations/{id}/children][%d] getScanPatternSimulationChildren default  %+v", o._statusCode, o.Payload)
}

func (o *GetScanPatternSimulationChildrenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
