package user

import (
	"context"
	"database/sql"
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"go-rental/domain"
	"go-rental/exceptions"
	"go-rental/utils"
)

type Service struct {
	rpo  domain.UserRepository
	db   *sql.DB
	mail *mailjet.Client
}

func (svc *Service) SaveRegisterBasicWithoutSSO(ctx context.Context, request *domain.RegisterBasicWithoutSSORequest) *domain.UserResponse {
	//TODO implement me
	panic("implement me")
}

func (svc *Service) SaveRegisterBasicWithSSO(ctx context.Context, request *domain.RegisterBasicWithSSORequest) *domain.UserResponse {
	//TODO implement me
	panic("implement me")
}

func (svc *Service) GetByEmail(ctx context.Context, email string) *domain.UserResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer utils.CommitRollback(tx)

	user, errFind := svc.rpo.FindByEmail(ctx, tx, email)
	if errFind != nil {
		panic(exceptions.NewNotFoundError(errFind.Error()))
	}

	return &domain.UserResponse{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
