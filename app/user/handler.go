package user

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go-rental/domain"
	"go-rental/response"
	"net/http"
)

type Handler struct {
	svc      domain.UserService
	validate *validator.Validate
}

// RegisterBasicWithoutSSO is an HTTP handler function that registers a new user without using SSO.
// It expects a JSON payload in the request body that conforms to the RegisterBasicWithoutSSORequest struct.
// It validates the request payload using the validator package.
// If the validation fails, it panics with the error.
// It saves the user data using the UserService and returns a JSON response with the user data.
// The response status code is set to 201 Created.
func (hdl *Handler) RegisterBasicWithoutSSO() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Decode the request payload into a RegisterBasicWithoutSSORequest struct
		req := new(domain.RegisterBasicWithoutSSORequest)

		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		// Validate the request payload
		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		// Save the user data using the UserService
		ctx := request.Context()
		result := hdl.svc.SaveRegisterBasicWithoutSSO(ctx, req)

		// Create a JSON response with the user data
		svcResponse := response.DefaultResponse{
			Code:   http.StatusCreated,
			Status: http.StatusText(http.StatusCreated),
			Data:   result,
		}

		// Set the response headers
		writer.Header().Set("Content-Type", "application/json")

		// Encode the response into the writer
		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

// RegisterBasicWithSSO is an HTTP handler function that registers a new user with SSO.
// It expects a JSON payload in the request body that conforms to the RegisterBasicWithSSORequest struct.
// It validates the request payload using the validator package.
// If the validation fails, it panics with the error.
// It saves the user data using the UserService and returns a JSON response with the user data.
// The response status code is set to 201 Created.
func (hdl *Handler) RegisterBasicWithSSO() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Decode the request payload into a RegisterBasicWithSSORequest struct
		req := new(domain.RegisterBasicWithSSORequest)

		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		// Validate the request payload
		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		// Save the user data using the UserService
		ctx := request.Context()
		result := hdl.svc.SaveRegisterBasicWithSSO(ctx, req)

		// Create a JSON response with the user data
		svcResponse := response.DefaultResponse{
			Code:   http.StatusCreated,
			Status: http.StatusText(http.StatusCreated),
			Data:   result,
		}

		// Set the response headers
		writer.Header().Set("Content-Type", "application/json")

		// Encode the response into the writer
		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

// GetByEmail is an HTTP handler function that retrieves a user by email.
// It expects the email to be passed as a query parameter in the request URL.
// It returns a JSON response with the user data.
// The response status code is set to 200 OK if the user is found, or 404 Not Found if the user is not found.
func (hdl *Handler) GetByEmail() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Get the email parameter from the request URL
		email := request.URL.Query().Get("email")

		// Create a context for the request
		ctx := request.Context()

		// Retrieve the user data from the service using the email
		user := hdl.svc.GetByEmail(ctx, email)

		// Create a default response with the user data
		svcResponse := response.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   user,
		}

		// Set the Content-Type header to indicate that the response is in JSON format
		writer.Header().Set("Content-Type", "application/json")

		// Encode the response into the writer
		encoder := json.NewEncoder(writer)

		err := encoder.Encode(svcResponse)
		if err != nil {
			// If there's an error encoding the response, panic with the error
			panic(err)
		}
	}
}

// OTPConfirmation is an HTTP handler function that confirms the OTP.
// It expects a JSON payload in the request body that conforms to the VerificationOTPRequest struct.
// It validates the request payload using the validator package.
// If the validation fails, it panics with the error.
// It checks the OTP using the UserService.
// If the OTP is valid, it returns a JSON response with the status code set to 201 Created.
func (hdl *Handler) VerificationOTP() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Decode the request payload into a VerificationOTPRequest struct
		req := new(domain.VerificationOTPRequest)

		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		// Validate the request payload
		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		// Check the OTP using the UserService
		ctx := request.Context()
		hdl.svc.CheckVerificationOTP(ctx, req)

		// Create a default response with the status code set to 200 OK
		svcResponse := response.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
		}

		// Set the response headers
		writer.Header().Set("Content-Type", "application/json")

		// Encode the response into the writer
		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

func (hdl *Handler) GenerateOTP() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(domain.GenerateOTPRequest)

		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		ctx := request.Context()
		hdl.svc.GenerateNewOTP(ctx, req)

		svcResponse := response.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}
