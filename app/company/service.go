package company

import (
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"go-edash/domain"
	"go-edash/exceptions"
	"go-edash/utils"
)

type Service struct {
	crpo domain.CompanyRepository
	urpo domain.UserRepository
	db   *sql.DB
}

func (svc *Service) SaveCompany(ctx context.Context, request *domain.SaveCompanyRequest) domain.CompanyResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer utils.CommitRollback(tx)

	claims := ctx.Value("claims").(jwt.MapClaims)
	sub := claims["sub"]

	company := &domain.Company{
		Name:        request.CompanyName,
		Description: request.CompanyDescription,
		Category:    request.CompanyCategory,
	}

	company = svc.crpo.Create(ctx, tx, company)

	user, errUser := svc.urpo.FindByEmail(ctx, tx, sub.(string))
	if errUser != nil {
		panic(exceptions.NewNotFoundError(errUser.Error()))
	}

	user.CompanyId = company.Id

	svc.urpo.Update(ctx, tx, user)

	return domain.CompanyResponse{
		CompanyName:        company.Name,
		CompanyDescription: company.Description,
		CompanyCategory:    company.Category,
	}
}

func (svc *Service) UpdateCompany(ctx context.Context, request *domain.UpdateCompanyRequest) domain.CompanyResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer utils.CommitRollback(tx)

	claims := ctx.Value("claims").(jwt.MapClaims)
	sub := claims["sub"]

	user, errUser := svc.urpo.FindByEmail(ctx, tx, sub.(string))
	if errUser != nil {
		panic(exceptions.NewNotFoundError(errUser.Error()))
	}

	companyId := user.CompanyId

	company, errCompany := svc.crpo.FindById(ctx, tx, companyId)
	if errCompany != nil {
		panic(exceptions.NewNotFoundError(errCompany.Error()))
	}

	company.Name = request.CompanyName
	company.Description = request.CompanyDescription
	company.Category = request.CompanyCategory

	company = svc.crpo.Update(ctx, tx, company)

	return domain.CompanyResponse{
		CompanyName:        company.Name,
		CompanyDescription: company.Description,
		CompanyCategory:    company.Category,
	}
}

func (svc *Service) GetCompanyInformation(ctx context.Context) domain.CompanyResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer utils.CommitRollback(tx)

	claims := ctx.Value("claims").(jwt.MapClaims)
	sub := claims["sub"]

	user, errUser := svc.urpo.FindByEmail(ctx, tx, sub.(string))
	if errUser != nil {
		panic(exceptions.NewNotFoundError(errUser.Error()))
	}

	companyId := user.CompanyId

	company, errCompany := svc.crpo.FindById(ctx, tx, companyId)
	if errCompany != nil {
		panic(exceptions.NewNotFoundError(errCompany.Error()))
	}

	return domain.CompanyResponse{
		CompanyName:        company.Name,
		CompanyDescription: company.Description,
		CompanyCategory:    company.Category,
	}
}
