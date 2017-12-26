// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/JunkieJI/travel-senpai/api/swagger/restapi/operations/destinations"
	"github.com/JunkieJI/travel-senpai/api/swagger/restapi/operations/transport"
)

// NewTravelSenpaiAPI creates a new TravelSenpai instance
func NewTravelSenpaiAPI(spec *loads.Document) *TravelSenpaiAPI {
	return &TravelSenpaiAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		DestinationsCreateDestinationHandler: destinations.CreateDestinationHandlerFunc(func(params destinations.CreateDestinationParams) middleware.Responder {
			return middleware.NotImplemented("operation DestinationsCreateDestination has not yet been implemented")
		}),
		DestinationsDeleteDestinationByIDHandler: destinations.DeleteDestinationByIDHandlerFunc(func(params destinations.DeleteDestinationByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DestinationsDeleteDestinationByID has not yet been implemented")
		}),
		DestinationsGetDestinationByIDHandler: destinations.GetDestinationByIDHandlerFunc(func(params destinations.GetDestinationByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DestinationsGetDestinationByID has not yet been implemented")
		}),
		DestinationsGetDestinationsHandler: destinations.GetDestinationsHandlerFunc(func(params destinations.GetDestinationsParams) middleware.Responder {
			return middleware.NotImplemented("operation DestinationsGetDestinations has not yet been implemented")
		}),
		TransportGetTransportByIDHandler: transport.GetTransportByIDHandlerFunc(func(params transport.GetTransportByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation TransportGetTransportByID has not yet been implemented")
		}),
		GetTripByIDHandler: GetTripByIDHandlerFunc(func(params GetTripByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetTripByID has not yet been implemented")
		}),
		GetTripsHandler: GetTripsHandlerFunc(func(params GetTripsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetTrips has not yet been implemented")
		}),
		DestinationsUpdateDestinationHandler: destinations.UpdateDestinationHandlerFunc(func(params destinations.UpdateDestinationParams) middleware.Responder {
			return middleware.NotImplemented("operation DestinationsUpdateDestination has not yet been implemented")
		}),
	}
}

/*TravelSenpaiAPI the travel senpai API */
type TravelSenpaiAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// DestinationsCreateDestinationHandler sets the operation handler for the create destination operation
	DestinationsCreateDestinationHandler destinations.CreateDestinationHandler
	// DestinationsDeleteDestinationByIDHandler sets the operation handler for the delete destination by ID operation
	DestinationsDeleteDestinationByIDHandler destinations.DeleteDestinationByIDHandler
	// DestinationsGetDestinationByIDHandler sets the operation handler for the get destination by ID operation
	DestinationsGetDestinationByIDHandler destinations.GetDestinationByIDHandler
	// DestinationsGetDestinationsHandler sets the operation handler for the get destinations operation
	DestinationsGetDestinationsHandler destinations.GetDestinationsHandler
	// TransportGetTransportByIDHandler sets the operation handler for the get transport by ID operation
	TransportGetTransportByIDHandler transport.GetTransportByIDHandler
	// GetTripByIDHandler sets the operation handler for the get trip by ID operation
	GetTripByIDHandler GetTripByIDHandler
	// GetTripsHandler sets the operation handler for the get trips operation
	GetTripsHandler GetTripsHandler
	// DestinationsUpdateDestinationHandler sets the operation handler for the update destination operation
	DestinationsUpdateDestinationHandler destinations.UpdateDestinationHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *TravelSenpaiAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TravelSenpaiAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TravelSenpaiAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TravelSenpaiAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TravelSenpaiAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TravelSenpaiAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TravelSenpaiAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TravelSenpaiAPI
func (o *TravelSenpaiAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.DestinationsCreateDestinationHandler == nil {
		unregistered = append(unregistered, "destinations.CreateDestinationHandler")
	}

	if o.DestinationsDeleteDestinationByIDHandler == nil {
		unregistered = append(unregistered, "destinations.DeleteDestinationByIDHandler")
	}

	if o.DestinationsGetDestinationByIDHandler == nil {
		unregistered = append(unregistered, "destinations.GetDestinationByIDHandler")
	}

	if o.DestinationsGetDestinationsHandler == nil {
		unregistered = append(unregistered, "destinations.GetDestinationsHandler")
	}

	if o.TransportGetTransportByIDHandler == nil {
		unregistered = append(unregistered, "transport.GetTransportByIDHandler")
	}

	if o.GetTripByIDHandler == nil {
		unregistered = append(unregistered, "GetTripByIDHandler")
	}

	if o.GetTripsHandler == nil {
		unregistered = append(unregistered, "GetTripsHandler")
	}

	if o.DestinationsUpdateDestinationHandler == nil {
		unregistered = append(unregistered, "destinations.UpdateDestinationHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TravelSenpaiAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TravelSenpaiAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *TravelSenpaiAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *TravelSenpaiAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *TravelSenpaiAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TravelSenpaiAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the travel senpai API
func (o *TravelSenpaiAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TravelSenpaiAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/destinations"] = destinations.NewCreateDestination(o.context, o.DestinationsCreateDestinationHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/destinations/{id}"] = destinations.NewDeleteDestinationByID(o.context, o.DestinationsDeleteDestinationByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/destinations/{id}"] = destinations.NewGetDestinationByID(o.context, o.DestinationsGetDestinationByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/destinations"] = destinations.NewGetDestinations(o.context, o.DestinationsGetDestinationsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/transport/{id}"] = transport.NewGetTransportByID(o.context, o.TransportGetTransportByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/trips/{id}"] = NewGetTripByID(o.context, o.GetTripByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/trips"] = NewGetTrips(o.context, o.GetTripsHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/destinations/{id}"] = destinations.NewUpdateDestination(o.context, o.DestinationsUpdateDestinationHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TravelSenpaiAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *TravelSenpaiAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
