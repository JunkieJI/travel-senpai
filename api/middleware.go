package api

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/rs/cors"

	"github.com/JunkieJI/travel-senpai/api/swagger/restapi/operations"
)

// RegisterMiddleware registers all request/response middlewares.
func RegisterMiddleware(api *operations.TravelSenpaiAPI) http.Handler {
	return setupGlobalMiddleware(api.Serve(setupMiddleware))
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddleware(handler http.Handler) http.Handler {
	limiter := tollbooth.NewLimiter(1000, nil)
	limiter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	return tollbooth.LimitFuncHandler(limiter, handler.ServeHTTP)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{"HEAD", "PUT", "GET", "POST", "DELETE", "OPTIONS"},
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Content-Type", "Accept-Language", "Authorization"},
	})
	return corsHandler.Handler(handler)
}
