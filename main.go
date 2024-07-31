package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"go-rental/app/user"
	"go-rental/app/welcome"
	"go-rental/config"
	"go-rental/middlewares"
	"net/http"
)

// main is the entry point for the application.
// It does not take any parameters and does not return any value.
func main() {
	config.InitConfig()

	log := config.CreateLoggers(nil)

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	validate := config.CreateValidator()

	router := chi.NewRouter()

	router.Use(middlewares.LoggerMiddleware)
	router.Use(middlewares.RecoverMiddleware)

	welcomeHandler := welcome.Wire()
	userHandler := user.Wire(validate, db)

	//router.Use(middlewares.AuthorizationCheckMiddleware)
	//router.Use(middlewares.VerifyTokenMiddleware)
	router.Get("/", welcomeHandler.Welcome())
	router.NotFound(welcomeHandler.NotFoundApi())
	router.MethodNotAllowed(welcomeHandler.MethodNotAllowedApi())

	router.Route("/admin", func(route chi.Router) {
		route.Post("/", userHandler.StoreAdminWithoutSSO())
	})

	router.Route("/user", func(route chi.Router) {
		route.Get("/email", userHandler.GetByEmail())
		route.Post("/", userHandler.StoreUserWithoutSSO())
		route.Post("/sso", userHandler.StoreUserWithSSO())
	})

	log.Info(fmt.Sprintf("%s Application Started", viper.GetString("APP_NAME")))
	err = http.ListenAndServe(":"+viper.GetString("APP_PORT"), router)
	if err != nil {
		log.Fatal(err)
	}
}
