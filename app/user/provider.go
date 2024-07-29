package user

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go-rental/domain"
	"sync"
)

var (
	hdl     *Handler
	hdlOnce sync.Once

	svc     *Service
	svcOnce sync.Once

	rpo     *Repository
	rpoOnce sync.Once

	ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,
		wire.Bind(new(domain.UserHandler), new(*Handler)),
		wire.Bind(new(domain.UserService), new(*Service)),
		wire.Bind(new(domain.UserRepository), new(*Repository)),
	)
)

func ProvideHandler(validate *validator.Validate, svc domain.UserService) *Handler {
	hdlOnce.Do(func() {
		hdl = &Handler{
			svc:      svc,
			validate: validate,
		}
	})

	return hdl
}

func ProvideService(rpo domain.UserRepository, db *sql.DB) *Service {
	svcOnce.Do(func() {
		svc = &Service{
			rpo: rpo,
			db:  db,
		}
	})

	return svc
}

func ProvideRepository() *Repository {
	rpoOnce.Do(func() {
		rpo = new(Repository)
	})

	return rpo
}
