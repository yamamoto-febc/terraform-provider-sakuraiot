package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetV1ServicesServiceIDReader is a Reader for the GetV1ServicesServiceID structure.
type GetV1ServicesServiceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ServicesServiceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV1ServicesServiceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetV1ServicesServiceIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetV1ServicesServiceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV1ServicesServiceIDOK creates a GetV1ServicesServiceIDOK with default headers values
func NewGetV1ServicesServiceIDOK() *GetV1ServicesServiceIDOK {
	return &GetV1ServicesServiceIDOK{}
}

/*GetV1ServicesServiceIDOK handles this case with default header values.

Service
*/
type GetV1ServicesServiceIDOK struct {
	Payload GetV1ServicesServiceIDOKBody
}

func (o *GetV1ServicesServiceIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/services/{serviceId}/][%d] getV1ServicesServiceIdOK  %+v", 200, o.Payload)
}

func (o *GetV1ServicesServiceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1ServicesServiceIDUnauthorized creates a GetV1ServicesServiceIDUnauthorized with default headers values
func NewGetV1ServicesServiceIDUnauthorized() *GetV1ServicesServiceIDUnauthorized {
	return &GetV1ServicesServiceIDUnauthorized{}
}

/*GetV1ServicesServiceIDUnauthorized handles this case with default header values.

Unauthenticated
*/
type GetV1ServicesServiceIDUnauthorized struct {
}

func (o *GetV1ServicesServiceIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/services/{serviceId}/][%d] getV1ServicesServiceIdUnauthorized ", 401)
}

func (o *GetV1ServicesServiceIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetV1ServicesServiceIDNotFound creates a GetV1ServicesServiceIDNotFound with default headers values
func NewGetV1ServicesServiceIDNotFound() *GetV1ServicesServiceIDNotFound {
	return &GetV1ServicesServiceIDNotFound{}
}

/*GetV1ServicesServiceIDNotFound handles this case with default header values.

Not found
*/
type GetV1ServicesServiceIDNotFound struct {
}

func (o *GetV1ServicesServiceIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/services/{serviceId}/][%d] getV1ServicesServiceIdNotFound ", 404)
}

func (o *GetV1ServicesServiceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetV1ServicesServiceIDOKBody Based on service type
swagger:model GetV1ServicesServiceIDOKBody
*/
type GetV1ServicesServiceIDOKBody interface{}
