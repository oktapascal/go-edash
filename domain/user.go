package domain

import (
	"context"
	"database/sql"
	"go-rental/enums"
	"net/http"
)

type (
	User struct {
		Id               *string
		Email            string
		Password         *string
		PhoneNumber      string
		FirstName        string
		LastName         string
		Role             enums.Role
		Provider         *string
		ProviderId       *int8
		Otp              string
		StatusOtp        bool
		RegistrationStep int8
		StatusTrial      *bool
		TrialStartDate   *string
	}

	UserResponse struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Token     string `json:"token"`
	}

	RegisterBasicWithoutSSORequest struct {
		FirstName            string `validate:"required,min=1,max=50" json:"first_name"`
		LastName             string `validate:"required,min=1,max=50" json:"last_name"`
		Email                string `validate:"required,email" json:"email"`
		Password             string `validate:"required" json:"password"`
		PasswordConfirmation string `validate:"required,eqfield=Password" json:"password_confirmation"`
	}

	RegisterBasicWithSSORequest struct {
		FirstName string `validate:"required,min=1,max=50" json:"first_name"`
		LastName  string `validate:"required,min=1,max=50" json:"last_name"`
		Email     string `validate:"required,email" json:"email"`
	}

	UserRepository interface {
		Create(ctx context.Context, tx *sql.Tx, user *User) *User
		FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*User, error)
	}

	UserService interface {
		SaveRegisterBasicWithoutSSO(ctx context.Context, request *RegisterBasicWithoutSSORequest) *UserResponse
		SaveRegisterBasicWithSSO(ctx context.Context, request *RegisterBasicWithSSORequest) *UserResponse
		GetByEmail(ctx context.Context, email string) *UserResponse
	}

	UserHandler interface {
		RegisterBasicWithoutSSO() http.HandlerFunc
		RegisterBasicWithSSO() http.HandlerFunc
		GetByEmail() http.HandlerFunc
	}
)
