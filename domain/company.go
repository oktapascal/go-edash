package domain

import (
	"context"
	"database/sql"
	"go-rental/enums"
	"net/http"
)

type (
	Company struct {
		Id          string
		Name        string
		Description string
		Category    enums.CompanyCategory
	}

	SaveCompanyRequest struct {
		CompanyName        string                `validate:"required,min=1,max=50" json:"company_name"`
		CompanyDescription string                `validate:"required,min=1,max=50" json:"company_description"`
		CompanyCategory    enums.CompanyCategory `validate:"required,min=1,max=250" json:"category"`
	}

	UpdateCompanyRequest struct {
		CompanyName        string                `validate:"required,min=1,max=50" json:"company_name"`
		CompanyDescription string                `validate:"required,min=1,max=50" json:"company_description"`
		CompanyCategory    enums.CompanyCategory `validate:"required,min=1,max=250" json:"category"`
	}

	CompanyResponse struct {
		CompanyName        string                `json:"company_name"`
		CompanyDescription string                `json:"company_description"`
		CompanyCategory    enums.CompanyCategory `json:"category"`
	}

	CompanyRepository interface {
		Create(ctx context.Context, tx *sql.Tx, company *Company) *Company
		Update(ctx context.Context, tx *sql.Tx, company *Company) *Company
		FindById(ctx context.Context, tx *sql.Tx, id string) (*Company, error)
	}

	CompanyService interface {
		SaveCompany(ctx context.Context, request *SaveCompanyRequest) CompanyResponse
		UpdateCompany(ctx context.Context, request *UpdateCompanyRequest) CompanyResponse
		GetCompanyInformation(ctx context.Context) CompanyResponse
	}

	CompanyHandler interface {
		StoreCompany() http.HandlerFunc
		UpdateCompany() http.HandlerFunc
		GetCompany() http.HandlerFunc
	}
)
