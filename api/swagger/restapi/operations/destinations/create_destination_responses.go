// Code generated by go-swagger; DO NOT EDIT.

package destinations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/JunkieJI/travel-senpai/api/swagger/models"
)

// CreateDestinationCreatedCode is the HTTP code returned for type CreateDestinationCreated
const CreateDestinationCreatedCode int = 201

/*CreateDestinationCreated create destination created

swagger:response createDestinationCreated
*/
type CreateDestinationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Destination `json:"body,omitempty"`
}

// NewCreateDestinationCreated creates CreateDestinationCreated with default headers values
func NewCreateDestinationCreated() *CreateDestinationCreated {
	return &CreateDestinationCreated{}
}

// WithPayload adds the payload to the create destination created response
func (o *CreateDestinationCreated) WithPayload(payload *models.Destination) *CreateDestinationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create destination created response
func (o *CreateDestinationCreated) SetPayload(payload *models.Destination) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDestinationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateDestinationBadRequestCode is the HTTP code returned for type CreateDestinationBadRequest
const CreateDestinationBadRequestCode int = 400

/*CreateDestinationBadRequest create destination bad request

swagger:response createDestinationBadRequest
*/
type CreateDestinationBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateDestinationBadRequest creates CreateDestinationBadRequest with default headers values
func NewCreateDestinationBadRequest() *CreateDestinationBadRequest {
	return &CreateDestinationBadRequest{}
}

// WithPayload adds the payload to the create destination bad request response
func (o *CreateDestinationBadRequest) WithPayload(payload string) *CreateDestinationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create destination bad request response
func (o *CreateDestinationBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDestinationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// CreateDestinationNotFoundCode is the HTTP code returned for type CreateDestinationNotFound
const CreateDestinationNotFoundCode int = 404

/*CreateDestinationNotFound create destination not found

swagger:response createDestinationNotFound
*/
type CreateDestinationNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateDestinationNotFound creates CreateDestinationNotFound with default headers values
func NewCreateDestinationNotFound() *CreateDestinationNotFound {
	return &CreateDestinationNotFound{}
}

// WithPayload adds the payload to the create destination not found response
func (o *CreateDestinationNotFound) WithPayload(payload string) *CreateDestinationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create destination not found response
func (o *CreateDestinationNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDestinationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// CreateDestinationInternalServerErrorCode is the HTTP code returned for type CreateDestinationInternalServerError
const CreateDestinationInternalServerErrorCode int = 500

/*CreateDestinationInternalServerError create destination internal server error

swagger:response createDestinationInternalServerError
*/
type CreateDestinationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateDestinationInternalServerError creates CreateDestinationInternalServerError with default headers values
func NewCreateDestinationInternalServerError() *CreateDestinationInternalServerError {
	return &CreateDestinationInternalServerError{}
}

// WithPayload adds the payload to the create destination internal server error response
func (o *CreateDestinationInternalServerError) WithPayload(payload string) *CreateDestinationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create destination internal server error response
func (o *CreateDestinationInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDestinationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}