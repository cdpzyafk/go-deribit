// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/adampointer/go-deribit/models"
)

// GetPrivatePositionsReader is a Reader for the GetPrivatePositions structure.
type GetPrivatePositionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivatePositionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivatePositionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivatePositionsOK creates a GetPrivatePositionsOK with default headers values
func NewGetPrivatePositionsOK() *GetPrivatePositionsOK {
	return &GetPrivatePositionsOK{}
}

/*GetPrivatePositionsOK handles this case with default header values.

foo
*/
type GetPrivatePositionsOK struct {
	Payload *models.PrivatePositionsResponse
}

func (o *GetPrivatePositionsOK) Error() string {
	return fmt.Sprintf("[GET /private/positions][%d] getPrivatePositionsOK  %+v", 200, o.Payload)
}

func (o *GetPrivatePositionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivatePositionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}