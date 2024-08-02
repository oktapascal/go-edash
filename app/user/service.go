package user

import (
	"context"
	"database/sql"
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/spf13/viper"
	"go-rental/config"
	"go-rental/domain"
	"go-rental/enums"
	"go-rental/exceptions"
	"go-rental/utils"
	"sync"
	"time"
)

type Service struct {
	rpo  domain.UserRepository
	db   *sql.DB
	mail *mailjet.Client
}

func (svc *Service) SaveRegisterBasicWithoutSSO(ctx context.Context, request *domain.RegisterBasicWithoutSSORequest) *domain.UserResponse {
	log := config.CreateLoggers(nil)

	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer utils.CommitRollback(tx)

	_, err = svc.rpo.FindByEmail(ctx, tx, request.Email)
	if err == nil {
		panic(exceptions.NewDuplicateError("email already exists"))
	}

	user := &domain.User{
		Email:            request.Email,
		Password:         &request.Password,
		FirstName:        request.FirstName,
		LastName:         request.LastName,
		Role:             enums.ADMIN,
		RegistrationStep: 0,
	}

	hash, errHash := utils.Hash(*user.Password)
	if errHash == nil {
		user.Password = &hash
	}

	otp, errOtp := utils.OTPGenerator(6)
	if errOtp != nil {
		panic(errOtp)
	}

	user.Otp = otp

	timeNow := time.Now()

	expiredOtp := timeNow.Add(10 * time.Minute)
	formatExpiredOtp := expiredOtp.Format("15:04:05")

	user.OtpExpiredTime = formatExpiredOtp

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: viper.GetString("MJ_EMAIL"),
				Name:  "EDash Admin",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: user.Email,
					Name:  user.FirstName + " " + user.LastName,
				},
			},
			Subject:          "Kode Autentikasi EDash",
			TemplateID:       6184340,
			TemplateLanguage: true,
			Variables: map[string]interface{}{
				"name":  user.FirstName + " " + user.LastName,
				"email": user.Email,
				"otp":   user.Otp,
			},
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}

	group := new(sync.WaitGroup)

	go func(group *sync.WaitGroup, messages *mailjet.MessagesV31) {
		defer group.Done()

		_, err = svc.mail.SendMailV31(messages)
		if err != nil {
			log.Error(err)
		}

		group.Add(1)
	}(group, &messages)

	user = svc.rpo.Create(ctx, tx, user)

	group.Wait()
	return &domain.UserResponse{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func (svc *Service) SaveRegisterBasicWithSSO(ctx context.Context, request *domain.RegisterBasicWithSSORequest) *domain.UserResponse {
	log := config.CreateLoggers(nil)

	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer utils.CommitRollback(tx)

	_, err = svc.rpo.FindByEmail(ctx, tx, request.Email)
	if err == nil {
		panic(exceptions.NewDuplicateError("email already exists"))
	}

	user := &domain.User{
		Email:            request.Email,
		FirstName:        request.FirstName,
		LastName:         request.LastName,
		Role:             enums.ADMIN,
		RegistrationStep: 0,
	}

	hash, errHash := utils.Hash(*user.Password)
	if errHash == nil {
		user.Password = &hash
	}

	otp, errOtp := utils.OTPGenerator(6)
	if errOtp != nil {
		panic(errOtp)
	}

	user.Otp = otp

	timeNow := time.Now()

	expiredOtp := timeNow.Add(10 * time.Minute)
	formatExpiredOtp := expiredOtp.Format("15:04:05")

	user.OtpExpiredTime = formatExpiredOtp

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: viper.GetString("MJ_EMAIL"),
				Name:  "EDash Admin",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: user.Email,
					Name:  user.FirstName + " " + user.LastName,
				},
			},
			Subject:          "Kode Autentikasi EDash",
			TemplateID:       6184340,
			TemplateLanguage: true,
			Variables: map[string]interface{}{
				"name":  user.FirstName + " " + user.LastName,
				"email": user.Email,
				"otp":   user.Otp,
			},
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}

	group := new(sync.WaitGroup)

	go func(group *sync.WaitGroup, messages *mailjet.MessagesV31) {
		defer group.Done()

		_, err = svc.mail.SendMailV31(messages)
		if err != nil {
			log.Error(err)
		}

		group.Add(1)
	}(group, &messages)

	user = svc.rpo.Create(ctx, tx, user)

	group.Wait()
	return &domain.UserResponse{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func (svc *Service) GetByEmail(ctx context.Context, email string) *domain.UserResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer utils.CommitRollback(tx)

	user, errFind := svc.rpo.FindByEmail(ctx, tx, email)
	if errFind != nil {
		panic(exceptions.NewNotFoundError(errFind.Error()))
	}

	return &domain.UserResponse{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
