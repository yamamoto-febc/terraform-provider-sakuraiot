package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteV1ServicesServiceIDReader is a Reader for the DeleteV1ServicesServiceID structure.
type DeleteV1ServicesServiceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1ServicesServiceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteV1ServicesServiceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewDeleteV1ServicesServiceIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteV1ServicesServiceIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteV1ServicesServiceIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteV1ServicesServiceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteV1ServicesServiceIDOK creates a DeleteV1ServicesServiceIDOK with default headers values
func NewDeleteV1ServicesServiceIDOK() *DeleteV1ServicesServiceIDOK {
	return &DeleteV1ServicesServiceIDOK{}
}

/*DeleteV1ServicesServiceIDOK handles this case with default header values.

Successful
*/
type DeleteV1ServicesServiceIDOK struct {
}

func (o *DeleteV1ServicesServiceIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/services/{serviceId}/][%d] deleteV1ServicesServiceIdOK ", 200)
}

func (o *DeleteV1ServicesServiceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV1ServicesServiceIDNoContent creates a DeleteV1ServicesServiceIDNoContent with default headers values
func NewDeleteV1ServicesServiceIDNoContent() *DeleteV1ServicesServiceIDNoContent {
	return &DeleteV1ServicesServiceIDNoContent{}
}

/*DeleteV1ServicesServiceIDNoContent handles this case with default header values.

No Content
*/
type DeleteV1ServicesServiceIDNoContent struct {
}

func (o *DeleteV1ServicesServiceIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/services/{serviceId}/][%d] deleteV1ServicesServiceIdNoContent ", 204)
}

func (o *DeleteV1ServicesServiceIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV1ServicesServiceIDUnauthorized creates a DeleteV1ServicesServiceIDUnauthorized with default headers values
func NewDeleteV1ServicesServiceIDUnauthorized() *DeleteV1ServicesServiceIDUnauthorized {
	return &DeleteV1ServicesServiceIDUnauthorized{}
}

/*DeleteV1ServicesServiceIDUnauthorized handles this case with default header values.

Unauthenticated
*/
type DeleteV1ServicesServiceIDUnauthorized struct {
}

func (o *DeleteV1ServicesServiceIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/services/{serviceId}/][%d] deleteV1ServicesServiceIdUnauthorized ", 401)
}

func (o *DeleteV1ServicesServiceIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV1ServicesServiceIDForbidden creates a DeleteV1ServicesServiceIDForbidden with default headers values
func NewDeleteV1ServicesServiceIDForbidden() *DeleteV1ServicesServiceIDForbidden {
	return &DeleteV1ServicesServiceIDForbidden{}
}

/*DeleteV1ServicesServiceIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteV1ServicesServiceIDForbidden struct {
}

func (o *DeleteV1ServicesServiceIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/services/{serviceId}/][%d] deleteV1ServicesServiceIdForbidden ", 403)
}

func (o *DeleteV1ServicesServiceIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV1ServicesServiceIDNotFound creates a DeleteV1ServicesServiceIDNotFound with default headers values
func NewDeleteV1ServicesServiceIDNotFound() *DeleteV1ServicesServiceIDNotFound {
	return &DeleteV1ServicesServiceIDNotFound{}
}

/*DeleteV1ServicesServiceIDNotFound handles this case with default header values.

Not found
*/
type DeleteV1ServicesServiceIDNotFound struct {
}

func (o *DeleteV1ServicesServiceIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/services/{serviceId}/][%d] deleteV1ServicesServiceIdNotFound ", 404)
}

func (o *DeleteV1ServicesServiceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}