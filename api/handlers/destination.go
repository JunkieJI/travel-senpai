package handlers

import (
	"github.com/JunkieJI/travel-senpai/api/swagger/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// GetDestinationByID : return destination details if it exists.
func GetDestinationByID(params operations.GetDestinationByIDParams) middleware.Responder {
	return operations.NewGetDestinationByIDOK()
}

// GetDestinations : return destinations details if it exists.
func GetDestinations(params operations.GetDestinationsParams) middleware.Responder {
	return operations.NewGetDestinationsOK()
}
