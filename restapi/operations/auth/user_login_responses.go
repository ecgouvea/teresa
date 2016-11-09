package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/teresa-api/models"
)

/*UserLoginOK Login token

swagger:response userLoginOK
*/
type UserLoginOK struct {

	// In: body
	Payload *models.LoginToken `json:"body,omitempty"`
}

// NewUserLoginOK creates UserLoginOK with default headers values
func NewUserLoginOK() *UserLoginOK {
	return &UserLoginOK{}
}

// WithPayload adds the payload to the user login o k response
func (o *UserLoginOK) WithPayload(payload *models.LoginToken) *UserLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user login o k response
func (o *UserLoginOK) SetPayload(payload *models.LoginToken) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UserLoginUnauthorized Unauthorized

swagger:response userLoginUnauthorized
*/
type UserLoginUnauthorized struct {

	// In: body
	Payload *models.Forbidden `json:"body,omitempty"`
}

// NewUserLoginUnauthorized creates UserLoginUnauthorized with default headers values
func NewUserLoginUnauthorized() *UserLoginUnauthorized {
	return &UserLoginUnauthorized{}
}

// WithPayload adds the payload to the user login unauthorized response
func (o *UserLoginUnauthorized) WithPayload(payload *models.Forbidden) *UserLoginUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user login unauthorized response
func (o *UserLoginUnauthorized) SetPayload(payload *models.Forbidden) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserLoginUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UserLoginForbidden Forbidden

swagger:response userLoginForbidden
*/
type UserLoginForbidden struct {

	// In: body
	Payload *models.Forbidden `json:"body,omitempty"`
}

// NewUserLoginForbidden creates UserLoginForbidden with default headers values
func NewUserLoginForbidden() *UserLoginForbidden {
	return &UserLoginForbidden{}
}

// WithPayload adds the payload to the user login forbidden response
func (o *UserLoginForbidden) WithPayload(payload *models.Forbidden) *UserLoginForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user login forbidden response
func (o *UserLoginForbidden) SetPayload(payload *models.Forbidden) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserLoginForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UserLoginDefault Error

swagger:response userLoginDefault
*/
type UserLoginDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewUserLoginDefault creates UserLoginDefault with default headers values
func NewUserLoginDefault(code int) *UserLoginDefault {
	if code <= 0 {
		code = 500
	}

	return &UserLoginDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user login default response
func (o *UserLoginDefault) WithStatusCode(code int) *UserLoginDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user login default response
func (o *UserLoginDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user login default response
func (o *UserLoginDefault) WithPayload(payload *models.Error) *UserLoginDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user login default response
func (o *UserLoginDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserLoginDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
