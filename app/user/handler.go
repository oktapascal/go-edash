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

func (hdl *Handler) StoreUserWithoutSSO() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(domain.RegisterWithoutSSORequest)

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
		result := hdl.svc.SaveUserWithoutSSO(ctx, req)
		svcResponse := response.DefaultResponse{
			Code:   http.StatusCreated,
			Status: http.StatusText(http.StatusCreated),
			Data:   result,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

func (hdl *Handler) StoreUserWithSSO() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(domain.RegisterWithSSORequest)

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
		result := hdl.svc.SaveUserWithSSO(ctx, req)
		svcResponse := response.DefaultResponse{
			Code:   http.StatusCreated,
			Status: http.StatusText(http.StatusCreated),
			Data:   result,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

func (hdl *Handler) GetByEmail() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		email := request.URL.Query().Get("email")

		ctx := request.Context()
		user := hdl.svc.GetByEmail(ctx, email)
		svcResponse := response.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   user,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err := encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}
