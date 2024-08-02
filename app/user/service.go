package user

import (
	"context"
	"database/sql"
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"go-rental/domain"
	"go-rental/enums"
	"go-rental/exceptions"
	"go-rental/utils"
)

type Service struct {
	rpo  domain.UserRepository
	db   *sql.DB
	mail *mailjet.Client
}

func (svc *Service) SaveRegisterBasicWithoutSSO(ctx context.Context, request *domain.RegisterBasicWithoutSSORequest) *domain.UserResponse {
	//log := config.CreateLoggers(nil)

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
		Id:               nil,
		Email:            request.Email,
		Password:         &request.Password,
		FirstName:        request.FirstName,
		LastName:         request.LastName,
		Role:             enums.ADMIN,
		Otp:              "",
		StatusOtp:        false,
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

	//messagesInfo := []mailjet.InfoMessagesV31{
	//	{
	//		From: &mailjet.RecipientV31{
	//			Email: viper.GetString("MJ_EMAIL"),
	//			Name:  "Sales App Super Admin",
	//		},
	//		To: &mailjet.RecipientsV31{
	//			mailjet.RecipientV31{
	//				Email: "oktaiscool@gmail.com",
	//				Name:  "passenger 1",
	//			},
	//		},
	//		Subject:  "Your email flight plan!",
	//		TextPart: "Dear passenger 1, welcome to Mailjet! May the delivery force be with you!",
	//		HTMLPart: "<h3>Dear passenger 1, welcome to <a href=\"https://www.mailjet.com/\">Mailjet</a>!</h3><br />May the delivery force be with you!",
	//	},
	//}

	//messages := mailjet.MessagesV31{Info: messagesInfo}

	//group := new(sync.WaitGroup)
	//
	//go func(group *sync.WaitGroup, messages *mailjet.MessagesV31) {
	//	defer group.Done()
	//
	//	_, err = svc.mail.SendMailV31(messages)
	//	if err != nil {
	//		log.Error(err)
	//	}
	//
	//	group.Add(1)
	//}(group, &messages)

	user = svc.rpo.Create(ctx, tx, user)

	//group.Wait()
	return &domain.UserResponse{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func (svc *Service) SaveRegisterBasicWithSSO(ctx context.Context, request *domain.RegisterBasicWithSSORequest) *domain.UserResponse {
	//TODO implement me
	panic("implement me")
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
