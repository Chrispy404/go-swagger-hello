// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SayHelloHandlerFunc turns a function with the right signature into a say hello handler
type SayHelloHandlerFunc func(SayHelloParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SayHelloHandlerFunc) Handle(params SayHelloParams) middleware.Responder {
	return fn(params)
}

// SayHelloHandler interface for that can handle valid say hello params
type SayHelloHandler interface {
	Handle(SayHelloParams) middleware.Responder
}

// NewSayHello creates a new http.Handler for the say hello operation
func NewSayHello(ctx *middleware.Context, handler SayHelloHandler) *SayHello {
	return &SayHello{Context: ctx, Handler: handler}
}

/*SayHello swagger:route GET /hello hello sayHello

returns a hello world message

*/
type SayHello struct {
	Context *middleware.Context
	Handler SayHelloHandler
}

func (o *SayHello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSayHelloParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}