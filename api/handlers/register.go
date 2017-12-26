package handlers

import (
	"github.com/JunkieJI/travel-senpai/api/swagger/restapi/operations"
	"github.com/JunkieJI/travel-senpai/api/swagger/restapi/operations/destinations"
)

// RegisterHandlers registers all resource endpoint handlers
func RegisterHandlers(api *operations.TravelSenpaiAPI) {
	api.DestinationsGetDestinationsHandler = destinations.GetDestinationsHandlerFunc(GetDestinations)
	api.DestinationsGetDestinationByIDHandler = destinations.GetDestinationByIDHandlerFunc(GetDestinationByID)
}
