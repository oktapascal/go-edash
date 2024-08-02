package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
	"go-rental/app/user"
	"go-rental/app/welcome"
	"go-rental/config"
	"go-rental/middlewares"
	"net/http"
	"time"
)

// main is the entry point for the application.
// It does not take any parameters and does not return any value.
func main() {
	config.InitConfig()
	log := config.CreateLoggers(nil)
	mailjetClient := config.SetupMailjetClient()
	validate := config.CreateValidator()
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middlewares.LoggerMiddleware)
	router.Use(middlewares.RecoverMiddleware)
	router.Use(middleware.Timeout(60 * time.Second))

	welcomeHandler := welcome.Wire()
	userHandler := user.Wire(validate, db, mailjetClient)

	//router.Use(middlewares.AuthorizationCheckMiddleware)
	//router.Use(middlewares.VerifyTokenMiddleware)
	router.Get("/", welcomeHandler.Welcome())
	router.NotFound(welcomeHandler.NotFoundApi())
	router.MethodNotAllowed(welcomeHandler.MethodNotAllowedApi())

	router.Route("/api", func(route chi.Router) {
		route.Post("/secure/otp-verify", userHandler.OTPConfirmation())

		route.Route("/register", func(subroute chi.Router) {
			subroute.Post("/basic/without-sso", userHandler.RegisterBasicWithoutSSO())
			subroute.Post("/basic/with-sso", userHandler.RegisterBasicWithSSO())
		})
	})

	log.Info(fmt.Sprintf("%s Application Started", viper.GetString("APP_NAME")))
	err = http.ListenAndServe(":"+viper.GetString("APP_PORT"), router)
	if err != nil {
		log.Fatal(err)
	}
}
