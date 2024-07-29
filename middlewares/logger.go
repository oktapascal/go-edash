package middlewares

import (
	"go-rental/config"
	"net/http"
)

// LoggerMiddleware is a middleware function that logs incoming HTTP requests.
//
// It wraps the provided http.Handler and logs a message using the CreateLogEntry function from the libs package.
//
// Parameters:
// - next: The http.Handler to be wrapped by the middleware.
//
// Returns:
// - An http.Handler that logs incoming requests before passing them to the next handler.
func LoggerMiddleware(next http.Handler) http.Handler {
	// Create a new http.HandlerFunc that wraps the provided http.Handler
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Create a logger using the CreateLoggers function from the config package
		log := config.CreateLoggers(request)

		// Log an "Incoming Request" message
		log.Info("Incoming Request")

		// Call the next handler in the chain
		next.ServeHTTP(writer, request)
	})
}
