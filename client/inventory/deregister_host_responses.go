// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeregisterHostReader is a Reader for the DeregisterHost structure.
type DeregisterHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeregisterHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeregisterHostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeregisterHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeregisterHostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeregisterHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeregisterHostNoContent creates a DeregisterHostNoContent with default headers values
func NewDeregisterHostNoContent() *DeregisterHostNoContent {
	return &DeregisterHostNoContent{}
}

/*DeregisterHostNoContent handles this case with default header values.

Host deregistered
*/
type DeregisterHostNoContent struct {
}

func (o *DeregisterHostNoContent) Error() string {
	return fmt.Sprintf("[DELETE /hosts/{host_id}][%d] deregisterHostNoContent ", 204)
}

func (o *DeregisterHostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeregisterHostBadRequest creates a DeregisterHostBadRequest with default headers values
func NewDeregisterHostBadRequest() *DeregisterHostBadRequest {
	return &DeregisterHostBadRequest{}
}

/*DeregisterHostBadRequest handles this case with default header values.

Host in use
*/
type DeregisterHostBadRequest struct {
}

func (o *DeregisterHostBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /hosts/{host_id}][%d] deregisterHostBadRequest ", 400)
}

func (o *DeregisterHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeregisterHostNotFound creates a DeregisterHostNotFound with default headers values
func NewDeregisterHostNotFound() *DeregisterHostNotFound {
	return &DeregisterHostNotFound{}
}

/*DeregisterHostNotFound handles this case with default header values.

Host not found
*/
type DeregisterHostNotFound struct {
}

func (o *DeregisterHostNotFound) Error() string {
	return fmt.Sprintf("[DELETE /hosts/{host_id}][%d] deregisterHostNotFound ", 404)
}

func (o *DeregisterHostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeregisterHostInternalServerError creates a DeregisterHostInternalServerError with default headers values
func NewDeregisterHostInternalServerError() *DeregisterHostInternalServerError {
	return &DeregisterHostInternalServerError{}
}

/*DeregisterHostInternalServerError handles this case with default header values.

Internal server error
*/
type DeregisterHostInternalServerError struct {
}

func (o *DeregisterHostInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /hosts/{host_id}][%d] deregisterHostInternalServerError ", 500)
}

func (o *DeregisterHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}