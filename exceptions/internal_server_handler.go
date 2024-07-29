package exceptions

import (
	"encoding/json"
	"go-rental/config"
	"go-rental/response"
	"net/http"
)

// InternalServerHandler handles HTTP 500 Internal Server Error responses.
//
// It writes a JSON response with the appropriate status code and error details.
// If an error occurs while encoding the response, it logs the error and panics.
//
// Parameters:
// - writer: The http.ResponseWriter to write the response to.

// InternalServerHandler handles HTTP 500 Internal Server Error responses.
//
// It writes a JSON response with the appropriate status code and error details.
// If an error occurs while encoding the response, it logs the error and panics.
//
// Parameters:
// - writer: The http.ResponseWriter to write the response to.
// - err: The error interface containing the details of the error.
//
// Return:
// This function does not return any value.
func InternalServerHandler(writer http.ResponseWriter, err any) {
	// Create a logger for error logging
	log := config.CreateLoggers(nil)

	// Set the content type of the response to JSON
	writer.Header().Set("Content-Type", "application/json")

	// Set the status code of the response to Internal Server Error
	writer.WriteHeader(http.StatusInternalServerError)

	// Create an error response with the status code and error details
	errorResponse := response.ErrorResponse{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Errors: err,
	}

	// Encode the error response into JSON
	encoder := json.NewEncoder(writer)

	// Check if there was an error encoding the response
	if errEncoder := encoder.Encode(errorResponse); errEncoder != nil {
		// Log the error if there was an error encoding the response
		log.Error(errEncoder)
	}

	// Log the error details
	log.Error(err)
}
