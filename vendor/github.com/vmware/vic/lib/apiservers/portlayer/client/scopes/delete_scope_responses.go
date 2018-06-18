package scopes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// DeleteScopeReader is a Reader for the DeleteScope structure.
type DeleteScopeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteScopeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteScopeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteScopeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteScopeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteScopeOK creates a DeleteScopeOK with default headers values
func NewDeleteScopeOK() *DeleteScopeOK {
	return &DeleteScopeOK{}
}

/*DeleteScopeOK handles this case with default header values.

OK
*/
type DeleteScopeOK struct {
}

func (o *DeleteScopeOK) Error() string {
	return fmt.Sprintf("[DELETE /scopes/{idName}][%d] deleteScopeOK ", 200)
}

func (o *DeleteScopeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteScopeNotFound creates a DeleteScopeNotFound with default headers values
func NewDeleteScopeNotFound() *DeleteScopeNotFound {
	return &DeleteScopeNotFound{}
}

/*DeleteScopeNotFound handles this case with default header values.

Not found
*/
type DeleteScopeNotFound struct {
	Payload *models.Error
}

func (o *DeleteScopeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /scopes/{idName}][%d] deleteScopeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteScopeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScopeInternalServerError creates a DeleteScopeInternalServerError with default headers values
func NewDeleteScopeInternalServerError() *DeleteScopeInternalServerError {
	return &DeleteScopeInternalServerError{}
}

/*DeleteScopeInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteScopeInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteScopeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /scopes/{idName}][%d] deleteScopeInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteScopeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}