// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"xsolla-sdk-backend/internal/server/models"
)

// LoginOKCode is the HTTP code returned for type LoginOK
const LoginOKCode int = 200

/*LoginOK User access token received

swagger:response loginOK
*/
type LoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.AccessToken `json:"body,omitempty"`
}

// NewLoginOK creates LoginOK with default headers values
func NewLoginOK() *LoginOK {

	return &LoginOK{}
}

// WithPayload adds the payload to the login o k response
func (o *LoginOK) WithPayload(payload *models.AccessToken) *LoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login o k response
func (o *LoginOK) SetPayload(payload *models.AccessToken) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginBadRequestCode is the HTTP code returned for type LoginBadRequest
const LoginBadRequestCode int = 400

/*LoginBadRequest Incorrect request parameters

swagger:response loginBadRequest
*/
type LoginBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Nr404 `json:"body,omitempty"`
}

// NewLoginBadRequest creates LoginBadRequest with default headers values
func NewLoginBadRequest() *LoginBadRequest {

	return &LoginBadRequest{}
}

// WithPayload adds the payload to the login bad request response
func (o *LoginBadRequest) WithPayload(payload *models.Nr404) *LoginBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login bad request response
func (o *LoginBadRequest) SetPayload(payload *models.Nr404) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginUnauthorizedCode is the HTTP code returned for type LoginUnauthorized
const LoginUnauthorizedCode int = 401

/*LoginUnauthorized Incorrect authorization data

swagger:response loginUnauthorized
*/
type LoginUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Nr401 `json:"body,omitempty"`
}

// NewLoginUnauthorized creates LoginUnauthorized with default headers values
func NewLoginUnauthorized() *LoginUnauthorized {

	return &LoginUnauthorized{}
}

// WithPayload adds the payload to the login unauthorized response
func (o *LoginUnauthorized) WithPayload(payload *models.Nr401) *LoginUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login unauthorized response
func (o *LoginUnauthorized) SetPayload(payload *models.Nr401) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}