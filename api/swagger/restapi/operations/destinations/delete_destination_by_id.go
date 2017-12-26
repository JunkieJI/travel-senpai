// Code generated by go-swagger; DO NOT EDIT.

package destinations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteDestinationByIDHandlerFunc turns a function with the right signature into a delete destination by ID handler
type DeleteDestinationByIDHandlerFunc func(DeleteDestinationByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteDestinationByIDHandlerFunc) Handle(params DeleteDestinationByIDParams) middleware.Responder {
	return fn(params)
}

// DeleteDestinationByIDHandler interface for that can handle valid delete destination by ID params
type DeleteDestinationByIDHandler interface {
	Handle(DeleteDestinationByIDParams) middleware.Responder
}

// NewDeleteDestinationByID creates a new http.Handler for the delete destination by ID operation
func NewDeleteDestinationByID(ctx *middleware.Context, handler DeleteDestinationByIDHandler) *DeleteDestinationByID {
	return &DeleteDestinationByID{Context: ctx, Handler: handler}
}

/*DeleteDestinationByID swagger:route DELETE /destinations/{id} Destinations deleteDestinationById

Delete Destination By ID

*/
type DeleteDestinationByID struct {
	Context *middleware.Context
	Handler DeleteDestinationByIDHandler
}

func (o *DeleteDestinationByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteDestinationByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}