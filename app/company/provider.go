package company

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go-rental/app/user"
	"go-rental/domain"
	"sync"
)

var (
	route     *Router
	routeOnce sync.Once

	hdl     *Handler
	hdlOnce sync.Once

	svc     *Service
	svcOnce sync.Once

	crpo     *Repository
	crpoOnce sync.Once

	urpo     *user.Repository
	urpoOnce sync.Once

	ProviderSet = wire.NewSet(
		ProvideRouter,
		ProvideHandler,
		ProvideService,
		ProvideCompanyRepository,
		ProvideUserRepository,
		wire.Bind(new(domain.CompanyHandler), new(*Handler)),
		wire.Bind(new(domain.CompanyService), new(*Service)),
		wire.Bind(new(domain.CompanyRepository), new(*Repository)),
		wire.Bind(new(domain.UserRepository), new(*user.Repository)),
	)
)

func ProvideRouter(hdl domain.CompanyHandler) *Router {
	routeOnce.Do(func() {
		route = &Router{
			hdl: hdl,
		}
	})

	return route
}

func ProvideHandler(validate *validator.Validate, svc domain.CompanyService) *Handler {
	hdlOnce.Do(func() {
		hdl = &Handler{
			svc:      svc,
			validate: validate,
		}
	})

	return hdl
}

func ProvideService(urpo domain.UserRepository, crpo domain.CompanyRepository, db *sql.DB) *Service {
	svcOnce.Do(func() {
		svc = &Service{
			urpo: urpo,
			crpo: crpo,
			db:   db,
		}
	})

	return svc
}

func ProvideCompanyRepository() *Repository {
	crpoOnce.Do(func() {
		crpo = new(Repository)
	})

	return crpo
}

func ProvideUserRepository() *user.Repository {
	urpoOnce.Do(func() {
		urpo = new(user.Repository)
	})

	return urpo
}
