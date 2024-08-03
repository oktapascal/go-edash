//go:build wireinject
// +build wireinject

package user

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/mailjet/mailjet-apiv3-go/v4"
)

func Wire(validate *validator.Validate, db *sql.DB, mail *mailjet.Client) *Router {
	panic(wire.Build(ProviderSet))
}
