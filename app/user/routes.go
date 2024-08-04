package user

import (
	"github.com/go-chi/chi/v5"
	"go-rental/domain"
	"go-rental/middlewares"
)

type Router struct {
	hdl domain.UserHandler
}

func (router *Router) InitializeRoute(rtr *chi.Mux) {
	rtr.Route("/api/user", func(route chi.Router) {
		route.Post("/register/basic/without-sso", router.hdl.RegisterBasicWithoutSSO())
		route.Post("/register/basic/with-sso", router.hdl.RegisterBasicWithSSO())

		route.Group(func(secure chi.Router) {
			secure.Use(middlewares.AuthorizationCheckMiddleware)
			secure.Use(middlewares.VerifyTokenMiddleware)
			secure.Get("/check-email", router.hdl.GetByEmail())
			secure.Post("/verification-otp", router.hdl.VerificationOTP())
			secure.Post("/generate-otp", router.hdl.GenerateOTP())
		})
	})
}
