package user

import (
	"github.com/go-chi/chi/v5"
	"go-rental/domain"
)

type Router struct {
	hdl domain.UserHandler
}

func (router *Router) InitializeRoute(rtr *chi.Mux) {
	rtr.Route("/api/user", func(route chi.Router) {
		route.Post("/register/basic/without-sso", router.hdl.RegisterBasicWithoutSSO())
		route.Post("/register/basic/with-sso", router.hdl.RegisterBasicWithSSO())
		route.Post("/secure/verification-otp", router.hdl.VerificationOTP())
		route.Post("/secure/generate-otp", router.hdl.GenerateOTP())
	})
}
