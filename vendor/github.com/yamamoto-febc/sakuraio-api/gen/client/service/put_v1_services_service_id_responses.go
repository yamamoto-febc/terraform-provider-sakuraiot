package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutV1ServicesServiceIDReader is a Reader for the PutV1ServicesServiceID structure.
type PutV1ServicesServiceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1ServicesServiceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutV1ServicesServiceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPutV1ServicesServiceIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutV1ServicesServiceIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutV1ServicesServiceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPutV1ServicesServiceIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutV1ServicesServiceIDOK creates a PutV1ServicesServiceIDOK with default headers values
func NewPutV1ServicesServiceIDOK() *PutV1ServicesServiceIDOK {
	return &PutV1ServicesServiceIDOK{}
}

/*PutV1ServicesServiceIDOK handles this case with default header values.

Updated servic
*/
type PutV1ServicesServiceIDOK struct {
	Payload PutV1ServicesServiceIDOKBody
}

func (o *PutV1ServicesServiceIDOK) Error() string {
	return fmt.Sprintf("[PUT /v1/services/{serviceId}/][%d] putV1ServicesServiceIdOK  %+v", 200, o.Payload)
}

func (o *PutV1ServicesServiceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1ServicesServiceIDUnauthorized creates a PutV1ServicesServiceIDUnauthorized with default headers values
func NewPutV1ServicesServiceIDUnauthorized() *PutV1ServicesServiceIDUnauthorized {
	return &PutV1ServicesServiceIDUnauthorized{}
}

/*PutV1ServicesServiceIDUnauthorized handles this case with default header values.

Unauthenticated
*/
type PutV1ServicesServiceIDUnauthorized struct {
}

func (o *PutV1ServicesServiceIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/services/{serviceId}/][%d] putV1ServicesServiceIdUnauthorized ", 401)
}

func (o *PutV1ServicesServiceIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutV1ServicesServiceIDForbidden creates a PutV1ServicesServiceIDForbidden with default headers values
func NewPutV1ServicesServiceIDForbidden() *PutV1ServicesServiceIDForbidden {
	return &PutV1ServicesServiceIDForbidden{}
}

/*PutV1ServicesServiceIDForbidden handles this case with default header values.

Forbidden
*/
type PutV1ServicesServiceIDForbidden struct {
}

func (o *PutV1ServicesServiceIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/services/{serviceId}/][%d] putV1ServicesServiceIdForbidden ", 403)
}

func (o *PutV1ServicesServiceIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutV1ServicesServiceIDNotFound creates a PutV1ServicesServiceIDNotFound with default headers values
func NewPutV1ServicesServiceIDNotFound() *PutV1ServicesServiceIDNotFound {
	return &PutV1ServicesServiceIDNotFound{}
}

/*PutV1ServicesServiceIDNotFound handles this case with default header values.

Not found
*/
type PutV1ServicesServiceIDNotFound struct {
}

func (o *PutV1ServicesServiceIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/services/{serviceId}/][%d] putV1ServicesServiceIdNotFound ", 404)
}

func (o *PutV1ServicesServiceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutV1ServicesServiceIDUnprocessableEntity creates a PutV1ServicesServiceIDUnprocessableEntity with default headers values
func NewPutV1ServicesServiceIDUnprocessableEntity() *PutV1ServicesServiceIDUnprocessableEntity {
	return &PutV1ServicesServiceIDUnprocessableEntity{}
}

/*PutV1ServicesServiceIDUnprocessableEntity handles this case with default header values.

Validation error
*/
type PutV1ServicesServiceIDUnprocessableEntity struct {
}

func (o *PutV1ServicesServiceIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v1/services/{serviceId}/][%d] putV1ServicesServiceIdUnprocessableEntity ", 422)
}

func (o *PutV1ServicesServiceIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PutV1ServicesServiceIDOKBody Based on service type
swagger:model PutV1ServicesServiceIDOKBody
*/
type PutV1ServicesServiceIDOKBody interface{}