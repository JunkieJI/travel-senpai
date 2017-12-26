package handlers

import (
	"github.com/JunkieJI/travel-senpai/api/swagger/models"
	"github.com/JunkieJI/travel-senpai/api/swagger/restapi/operations/destinations"
	"github.com/JunkieJI/travel-senpai/config"
	"github.com/JunkieJI/travel-senpai/store"
	"github.com/go-openapi/runtime/middleware"
)

var s = store.NewStore(config.GetDSN())

func init() {
	config.Init()
}

// GetDestinationByID : return destination details if it exists.
func GetDestinationByID(params destinations.GetDestinationByIDParams) middleware.Responder {
	destination, err := s.GetDestinationByID(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		return destinations.NewGetDestinationByIDNotFound().WithPayload("Could not find destination")
	}
	payload, err := toSwaggerDestination(destination)
	if err != nil {
		return destinations.NewGetDestinationByIDInternalServerError().WithPayload("Error handling destination")
	}
	return destinations.NewGetDestinationByIDOK().WithPayload(payload)
}

// GetDestinations : return destinations details if it exists.
func GetDestinations(params destinations.GetDestinationsParams) middleware.Responder {
	res, err := s.GetDestinations(params.HTTPRequest.Context())
	if err != nil {
		return destinations.NewGetDestinationsNotFound().WithPayload("Could not find any destinations")
	}

	var payload []*models.Destination
	for _, dest := range *res {
		p, _ := toSwaggerDestination(&dest)
		payload = append(payload, p)
	}
	return destinations.NewGetDestinationsOK().WithPayload(payload)
}

// AddDestination : creates a destination
func AddDestination(params destinations.CreateDestinationParams) middleware.Responder {
	res, err := s.AddDestination(params.HTTPRequest.Context(), toStoreDestination(params.Body))
	if err != nil {
		return destinations.NewCreateDestinationInternalServerError().WithPayload("Could not create destination")
	}
	payload, err := toSwaggerDestination(res)
	if err != nil {
		return destinations.NewCreateDestinationInternalServerError().WithPayload("Error creating destination")
	}
	return destinations.NewCreateDestinationCreated().WithPayload(payload)
}

// UpdateDestination : updates a destination
func UpdateDestination(params destinations.UpdateDestinationParams) middleware.Responder {
	res, err := s.UpdateDestination(params.HTTPRequest.Context(), toStoreDestination(params.Body))
	if err != nil {
		return destinations.NewUpdateDestinationInternalServerError().WithPayload("Could not update destination")
	}
	payload, err := toSwaggerDestination(res)
	if err != nil {
		return destinations.NewUpdateDestinationInternalServerError().WithPayload("Error updating destination")
	}
	return destinations.NewUpdateDestinationOK().WithPayload(payload)
}

// DeleteDestination : deletes a destination
func DeleteDestination(params destinations.DeleteDestinationByIDParams) middleware.Responder {
	err := s.DeleteDestination(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		return destinations.NewDeleteDestinationByIDInternalServerError().WithPayload("Unable to delete destination")
	}
	return destinations.NewDeleteDestinationByIDNoContent().WithPayload("Successfully deleted destination")
}

func toSwaggerDestination(destination *store.Destination) (*models.Destination, error) {
	return &models.Destination{
		ID:       destination.ID,
		Budget:   destination.Budget,
		City:     destination.City,
		Country:  destination.Country,
		Currency: destination.Currency,
	}, nil
}

func toStoreDestination(destination *models.Destination) *store.Destination {
	return &store.Destination{
		Budget:   destination.Budget,
		City:     destination.City,
		Country:  destination.Country,
		Currency: destination.Currency,
	}
}
