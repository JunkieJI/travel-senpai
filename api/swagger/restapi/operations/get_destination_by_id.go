// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDestinationByIDHandlerFunc turns a function with the right signature into a get destination by ID handler
type GetDestinationByIDHandlerFunc func(GetDestinationByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDestinationByIDHandlerFunc) Handle(params GetDestinationByIDParams) middleware.Responder {
	return fn(params)
}

// GetDestinationByIDHandler interface for that can handle valid get destination by ID params
type GetDestinationByIDHandler interface {
	Handle(GetDestinationByIDParams) middleware.Responder
}

// NewGetDestinationByID creates a new http.Handler for the get destination by ID operation
func NewGetDestinationByID(ctx *middleware.Context, handler GetDestinationByIDHandler) *GetDestinationByID {
	return &GetDestinationByID{Context: ctx, Handler: handler}
}

/*GetDestinationByID swagger:route GET /destinations/{id} destinations Destinations getDestinationById

Destination By ID

*/
type GetDestinationByID struct {
	Context *middleware.Context
	Handler GetDestinationByIDHandler
}

func (o *GetDestinationByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDestinationByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
