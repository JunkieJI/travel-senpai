package handlers

import (
	"github.com/JunkieJI/travel-senpai/api/swagger/restapi/operations"
)

// RegisterHandlers registers all resource endpoint handlers
func RegisterHandlers(api *operations.TravelSenpaiAPI) {
	api.GetDestinationsHandler = operations.GetDestinationsHandlerFunc(GetDestinations)
	api.GetDestinationByIDHandler = operations.GetDestinationByIDHandlerFunc(GetDestinationByID)
}
