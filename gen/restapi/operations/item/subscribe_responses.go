// Code generated by go-swagger; DO NOT EDIT.

package item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ustore/gen/models"
)

// SubscribeOKCode is the HTTP code returned for type SubscribeOK
const SubscribeOKCode int = 200

/*SubscribeOK Successful Subscription

swagger:response subscribeOK
*/
type SubscribeOK struct {

	/*
	  In: Body
	*/
	Payload *models.SubscriptionResponse `json:"body,omitempty"`
}

// NewSubscribeOK creates SubscribeOK with default headers values
func NewSubscribeOK() *SubscribeOK {

	return &SubscribeOK{}
}

// WithPayload adds the payload to the subscribe o k response
func (o *SubscribeOK) WithPayload(payload *models.SubscriptionResponse) *SubscribeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the subscribe o k response
func (o *SubscribeOK) SetPayload(payload *models.SubscriptionResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubscribeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SubscribeBadRequestCode is the HTTP code returned for type SubscribeBadRequest
const SubscribeBadRequestCode int = 400

/*SubscribeBadRequest Bad Request

swagger:response subscribeBadRequest
*/
type SubscribeBadRequest struct {
}

// NewSubscribeBadRequest creates SubscribeBadRequest with default headers values
func NewSubscribeBadRequest() *SubscribeBadRequest {

	return &SubscribeBadRequest{}
}

// WriteResponse to the client
func (o *SubscribeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// SubscribeNotFoundCode is the HTTP code returned for type SubscribeNotFound
const SubscribeNotFoundCode int = 404

/*SubscribeNotFound Item not found

swagger:response subscribeNotFound
*/
type SubscribeNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewSubscribeNotFound creates SubscribeNotFound with default headers values
func NewSubscribeNotFound() *SubscribeNotFound {

	return &SubscribeNotFound{}
}

// WithPayload adds the payload to the subscribe not found response
func (o *SubscribeNotFound) WithPayload(payload string) *SubscribeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the subscribe not found response
func (o *SubscribeNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubscribeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SubscribeInternalServerErrorCode is the HTTP code returned for type SubscribeInternalServerError
const SubscribeInternalServerErrorCode int = 500

/*SubscribeInternalServerError Server error

swagger:response subscribeInternalServerError
*/
type SubscribeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewSubscribeInternalServerError creates SubscribeInternalServerError with default headers values
func NewSubscribeInternalServerError() *SubscribeInternalServerError {

	return &SubscribeInternalServerError{}
}

// WithPayload adds the payload to the subscribe internal server error response
func (o *SubscribeInternalServerError) WithPayload(payload string) *SubscribeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the subscribe internal server error response
func (o *SubscribeInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubscribeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
