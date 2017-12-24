package handlers

import (
	"github.com/JunkieJI/travel-senpai/api/swagger/models"
	"github.com/JunkieJI/travel-senpai/api/swagger/restapi/operations"
	"github.com/JunkieJI/travel-senpai/config"
	"github.com/JunkieJI/travel-senpai/store"
	"github.com/go-openapi/runtime/middleware"
)

var s = store.NewStore(config.GetDSN())

func init() {
	config.Init()
}

// GetDestinationByID : return destination details if it exists.
func GetDestinationByID(params operations.GetDestinationByIDParams) middleware.Responder {
	destination, err := s.GetDestinationByID(params.ID)
	if err != nil {
		return operations.NewGetDestinationByIDNotFound().WithPayload("Could not find destination")
	}
	payload, err := toSwaggerDestination(destination)
	if err != nil {
		return operations.NewGetDestinationByIDInternalServerError().WithPayload("Error handling destination")
	}
	return operations.NewGetDestinationByIDOK().WithPayload(payload)
}

// GetDestinations : return destinations details if it exists.
func GetDestinations(params operations.GetDestinationsParams) middleware.Responder {
	destinations, err := s.GetDestinations()
	if err != nil {
		return operations.NewGetDestinationsNotFound().WithPayload("Could not find any destinations")
	}

	var payload []*models.Destination
	for _, destination := range *destinations {
		p, _ := toSwaggerDestination(&destination)
		payload = append(payload, p)
	}
	return operations.NewGetDestinationsOK().WithPayload(payload)
}

func toSwaggerDestination(destination *store.Destination) (*models.Destination, error) {
	return &models.Destination{
		ID: destination.ID,
	}, nil
}
