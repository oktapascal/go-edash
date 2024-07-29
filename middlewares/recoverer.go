package middlewares

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"go-rental/exceptions"
	"net/http"
)

// RecoverMiddleware is a middleware function that recovers from panics and logs the error.

// It wraps the provided handler and recovers from any panics that occur during its execution.
// If a panic occurs, the function logs the error using the CreateLoggerConsole and CreateLoggerFile functions,
// and then calls the InternalServerHandler function to send an error response to the client.

// RecoverMiddleware is a middleware function that recovers from panics and logs the error.
// It wraps the provided handler and recovers from any panics that occur during its execution.
// If a panic occurs, the function logs the error using the CreateLoggerConsole and CreateLoggerFile functions,
// and then calls the InternalServerHandler function to send an error response to the client.
//
// Parameters:
// - next: The http.Handler to be wrapped by the middleware.
//
// Returns:
// - A http.Handler that recovers from panics and logs the error.
func RecoverMiddleware(next http.Handler) http.Handler {
	// Return a new http.Handler that wraps the provided handler
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Recover from panics
		defer func() {
			// Check if a panic occurred
			if err := recover(); err != nil {
				// Check if the error is a NotFoundError
				if str, ok := err.(exceptions.NotFoundError); ok {
					// Call NotFoundHandler to send error response
					exceptions.NotFoundHandler(writer, str)
					return
				}

				// Check if the error is a DuplicateError
				if str, ok := err.(exceptions.DuplicateError); ok {
					exceptions.DuplicateHandler(writer, str)
					return
				}

				// Check if the error is a string
				if str, ok := err.(string); ok {
					// Call InternalServerHandler to send error response
					exceptions.InternalServerHandler(writer, str)
					return
				}

				// Check if the error is a ValidationErrors
				var validationErrors validator.ValidationErrors
				if errors.As(err.(error), &validationErrors) {
					// Format validation errors and call BadRequestHandler
					badRequestHandler := exceptions.FormatErrors(validationErrors)
					exceptions.BadRequestHandler(writer, badRequestHandler)
					return
				}

				// Call InternalServerHandler to send error response
				exceptions.InternalServerHandler(writer, err)
			}
		}()

		// Call the next handler
		next.ServeHTTP(writer, request)
	})
}
