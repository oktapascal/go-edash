package company

import (
	"github.com/go-chi/chi/v5"
	"go-edash/domain"
	"go-edash/middlewares"
)

type Router struct {
	hdl domain.CompanyHandler
}

func (router *Router) InitializeRoute(rtr *chi.Mux) {
	rtr.Route("/api/company", func(route chi.Router) {
		route.Use(middlewares.AuthorizationCheckMiddleware)
		route.Use(middlewares.VerifyTokenMiddleware)

		route.Get("/show", router.hdl.GetCompany())
		route.Post("/save", router.hdl.StoreCompany())
		route.Post("/update", router.hdl.UpdateCompany())
	})
}
