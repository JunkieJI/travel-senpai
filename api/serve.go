package api

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/JunkieJI/travel-senpai/api/handlers"
	"github.com/JunkieJI/travel-senpai/api/swagger/restapi"
	"github.com/JunkieJI/travel-senpai/api/swagger/restapi/operations"
	"github.com/JunkieJI/travel-senpai/config"
	loads "github.com/go-openapi/loads"
)

// Serve : creates, configures and serves API.
func Serve() {
	// Load swagger specification.
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// Create API.
	api := operations.NewTravelSenpaiAPI(swaggerSpec)

	// Create Server.
	server := restapi.NewServer(api)

	// Configure server: set host and port.
	serverConf := config.GetServerConfig()
	server.Host = serverConf.Host
	server.Port = serverConf.Port

	// Configure server: register handlers on the API.
	handlers.RegisterHandlers(api)

	// Calls the default swagger generated config
	server.ConfigureAPI()

	// Configure server: register middlewares.
	server.SetHandler(RegisterMiddleware(api))

	defer server.Shutdown()

	log.Debug("Starting server...")
	log.Debug(fmt.Sprintf("Using config File: %s", serverConf.ConfigFile))

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
