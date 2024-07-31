package domain

import (
	"context"
	"database/sql"
	"net/http"
)

type (
	User struct {
		IdNumber    string
		Email       string
		Password    *string
		PhoneNumber string
		Address     string
		FirstName   string
		LastName    string
		Provider    *string
		ProviderId  *int8
		PhotoIdCard *string
	}

	UserResponse struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	RegisterAdminWithoutSSORequest struct {
		FirstName            string `validate:"required,min=1,max=50" json:"first_name"`
		LastName             string `validate:"required,min=1,max=50" json:"last_name"`
		Email                string `validate:"required,email" json:"email"`
		Password             string `validate:"required" json:"password"`
		PasswordConfirmation string `validate:"required,eqfield=Password" json:"password_confirmation"`
	}

	RegisterAdminWithSSORequest struct {
		FirstName string `validate:"required,min=1,max=50" json:"first_name"`
		LastName  string `validate:"required,min=1,max=50" json:"last_name"`
		Email     string `validate:"required,email" json:"email"`
	}

	RegisterWithoutSSORequest struct {
		IdNumber    string `validate:"required,min=1,max=16" json:"id_number"`
		Email       string `validate:"required,email" json:"email"`
		Password    string `validate:"required" json:"password"`
		PhoneNumber string `validate:"required,min=10,max=13" json:"phone_number"`
		Address     string `validate:"required" json:"address"`
		FirstName   string `validate:"required,min=1,max=50" json:"first_name"`
		LastName    string `validate:"required,min=1,max=50" json:"last_name"`
	}

	RegisterWithSSORequest struct {
		IdNumber    string `validate:"required,min=1,max=16" json:"id_number"`
		Email       string `validate:"required,email" json:"email"`
		PhoneNumber string `validate:"required,min=10,max=13" json:"phone_number"`
		Address     string `validate:"required" json:"address"`
		FirstName   string `validate:"required,min=1,max=50" json:"first_name"`
		LastName    string `validate:"required,min=1,max=50" json:"last_name"`
		Provider    string `validate:"required" json:"provider"`
		ProviderId  int8   `validate:"required" json:"provider_id"`
	}

	UserRepository interface {
		Create(ctx context.Context, tx *sql.Tx, user *User) *User
		FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*User, error)
	}

	UserService interface {
		SaveUserWithoutSSO(ctx context.Context, request *RegisterWithoutSSORequest) *UserResponse
		SaveUserWithSSO(ctx context.Context, request *RegisterWithSSORequest) *UserResponse
		SaveAdminWithoutSSO(ctx context.Context, request *RegisterAdminWithoutSSORequest) *UserResponse
		SaveAdminWithSSO(ctx context.Context, request *RegisterAdminWithSSORequest) *UserResponse
		GetByEmail(ctx context.Context, email string) *UserResponse
	}

	UserHandler interface {
		StoreUserWithoutSSO() http.HandlerFunc
		StoreUserWithSSO() http.HandlerFunc
		StoreAdminWithoutSSO() http.HandlerFunc
		StoreAdminWithSSO() http.HandlerFunc
		GetByEmail() http.HandlerFunc
	}
)
