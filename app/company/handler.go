package company

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go-edash/domain"
	"go-edash/response"
	"net/http"
)

type Handler struct {
	svc      domain.CompanyService
	validate *validator.Validate
}

func (hdl *Handler) StoreCompany() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(domain.SaveCompanyRequest)

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
		result := hdl.svc.SaveCompany(ctx, req)

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

func (hdl *Handler) UpdateCompany() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(domain.UpdateCompanyRequest)

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
		result := hdl.svc.UpdateCompany(ctx, req)

		svcResponse := response.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
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

func (hdl *Handler) GetCompany() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		company := hdl.svc.GetCompanyInformation(ctx)

		svcResponse := response.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   company,
		}

		writer.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)

		err := encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}
