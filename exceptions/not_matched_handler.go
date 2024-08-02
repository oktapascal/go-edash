package exceptions

import (
	"encoding/json"
	"go-rental/config"
	"go-rental/response"
	"net/http"
)

type NotMatchedError struct {
	Error string
}

// NewNotMatchedError creates a new NotMatchedError instance.
//
// error: The error message to be associated with the NotMatchedError.
//
// Returns a NotMatchedError instance.
func NewNotMatchedError(error string) NotMatchedError {
	// Return a new NotMatchedError instance with the provided error message.
	return NotMatchedError{Error: error}
}

// NotMatchedHandler handles HTTP 400 Bad Request responses for not matched errors.
// It writes a JSON response with the appropriate status code and error details.
// If an error occurs while encoding the response, it logs the error.
//
// Parameters:
// - writer: The http.ResponseWriter to write the response to.
// - err: The error interface containing the details of the error.
func NotMatchedHandler(writer http.ResponseWriter, err any) {
	// Create a logger for error logging
	log := config.CreateLoggers(nil)

	// Set the content type of the response to JSON
	writer.Header().Set("Content-Type", "application/json")

	// Set the status code of the response to Bad Request
	writer.WriteHeader(http.StatusBadRequest)

	// Create an error response with the status code and error details
	errorResponse := response.ErrorResponse{
		Code:   http.StatusBadRequest,                  // Set the status code to Bad Request
		Status: http.StatusText(http.StatusBadRequest), // Set the status text to the corresponding HTTP status text
		Errors: err,                                    // Set the error details to the provided error
	}

	// Encode the error response into JSON
	encoder := json.NewEncoder(writer)

	// Check if there was an error encoding the response
	if errEncoder := encoder.Encode(errorResponse); errEncoder != nil {
		// Log the error if there was an error encoding the response
		log.Error(errEncoder)
	}
}
