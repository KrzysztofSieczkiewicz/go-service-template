// Code generated by go-swagger; DO NOT EDIT.

package meshes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteMeshReader is a Reader for the DeleteMesh structure.
type DeleteMeshReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMeshReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDeleteMeshCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteMeshCreated creates a DeleteMeshCreated with default headers values
func NewDeleteMeshCreated() *DeleteMeshCreated {
	return &DeleteMeshCreated{}
}

/* DeleteMeshCreated describes a response with status code 201, with default header values.

DeleteMeshCreated delete mesh created
*/
type DeleteMeshCreated struct {
}

func (o *DeleteMeshCreated) Error() string {
	return fmt.Sprintf("[DELETE /meshes/{id}][%d] deleteMeshCreated ", 201)
}

func (o *DeleteMeshCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}