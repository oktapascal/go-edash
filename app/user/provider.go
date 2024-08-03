package user

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/mailjet/mailjet-apiv3-go/v4"
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

	rpo     *Repository
	rpoOnce sync.Once

	ProviderSet = wire.NewSet(
		ProvideRouter,
		ProvideHandler,
		ProvideService,
		ProvideRepository,
		wire.Bind(new(domain.UserHandler), new(*Handler)),
		wire.Bind(new(domain.UserService), new(*Service)),
		wire.Bind(new(domain.UserRepository), new(*Repository)),
	)
)

func ProvideRouter(hdl domain.UserHandler) *Router {
	routeOnce.Do(func() {
		route = &Router{
			hdl: hdl,
		}
	})

	return route
}

func ProvideHandler(validate *validator.Validate, svc domain.UserService) *Handler {
	hdlOnce.Do(func() {
		hdl = &Handler{
			svc:      svc,
			validate: validate,
		}
	})

	return hdl
}

func ProvideService(rpo domain.UserRepository, db *sql.DB, mail *mailjet.Client) *Service {
	svcOnce.Do(func() {
		svc = &Service{
			rpo:  rpo,
			db:   db,
			mail: mail,
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
