package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
	"go-rental/app/company"
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

	user.Wire(validate, db, mailjetClient).InitializeRoute(router)
	company.Wire(validate, db).InitializeRoute(router)

	router.Get("/", welcomeHandler.Welcome())
	router.NotFound(welcomeHandler.NotFoundApi())
	router.MethodNotAllowed(welcomeHandler.MethodNotAllowedApi())

	log.Info(fmt.Sprintf("%s Application Started", viper.GetString("APP_NAME")))
	err = http.ListenAndServe(":"+viper.GetString("APP_PORT"), router)
	if err != nil {
		log.Fatal(err)
	}
}
