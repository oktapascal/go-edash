package exceptions

import (
	"encoding/json"
	"go-rental/config"
	"go-rental/response"
	"net/http"
)

type NotFoundError struct {
	Error string
}

// NewNotFoundError creates a new NotFoundError with the provided error message.
//
// Parameters:
// - error: The error message to be included in the NotFoundError.
//
// Returns:
// - NotFoundError: The newly created NotFoundError.
func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{
		Error: error, // Assign the provided error message to the Error field.
	}
}

// NotFoundHandler handles HTTP 404 Not Found responses.
// It writes a JSON response with the appropriate status code and error details.
// If an error occurs while encoding the response, it logs the error.
//
// Parameters:
// - writer: The http.ResponseWriter to write the response to.
// - err: The error interface containing the details of the error.
func NotFoundHandler(writer http.ResponseWriter, err any) {
	// Create a logger for error logging
	log := config.CreateLoggers(nil)

	// Set the content type of the response to JSON
	writer.Header().Set("Content-Type", "application/json")

	// Set the status code of the response to Not Found
	writer.WriteHeader(http.StatusNotFound)

	// Create an error response with the status code and error details
	errorResponse := response.ErrorResponse{

		Code:   http.StatusNotFound,                  // Set the status code to Not Found
		Status: http.StatusText(http.StatusNotFound), // Set the status text to the corresponding HTTP status text
		Errors: err,                                  // Set the error details to the provided error
	}

	// Encode the error response into JSON
	encoder := json.NewEncoder(writer)

	// Check if there was an error encoding the response
	if errEncoder := encoder.Encode(errorResponse); errEncoder != nil {
		// Log the error if there was an error encoding the response
		log.Error(errEncoder)
	}
}
