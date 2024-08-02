package exceptions

import (
	"encoding/json"
	"go-rental/config"
	"go-rental/response"
	"net/http"
)

type GoneError struct {
	Error string
}

// NewGoneError creates a new GoneError with the provided error message.
//
// error: the error message to be included in the GoneError.
//
// Returns a GoneError with the provided error message.
func NewGoneError(error string) GoneError {
	return GoneError{Error: error}
}

// GoneHandler is a function that handles HTTP 410 Gone responses.
// It writes a JSON response with the appropriate status code and error details.
// If an error occurs while encoding the response, it logs the error.
//
// Parameters:
// - writer: The http.ResponseWriter to write the response to.
// - err: The error interface containing the details of the error.
func GoneHandler(writer http.ResponseWriter, err any) {
	// Create a logger for error logging
	log := config.CreateLoggers(nil)

	// Set the content type of the response to JSON
	writer.Header().Set("Content-Type", "application/json")

	// Set the status code of the response to Gone
	writer.WriteHeader(http.StatusGone)

	// Create an error response with the status code and error details
	errorResponse := response.ErrorResponse{
		Code:   http.StatusGone,                  // Set the status code to Gone
		Status: http.StatusText(http.StatusGone), // Set the status text to the corresponding HTTP status text
		Errors: err,                              // Set the error details to the provided error
	}

	// Encode the error response into JSON
	encoder := json.NewEncoder(writer)

	// Check if there was an error encoding the response
	if errEncoder := encoder.Encode(errorResponse); errEncoder != nil {
		// Log the error if there was an error encoding the response
		log.Error(errEncoder)
	}
}
