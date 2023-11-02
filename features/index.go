package features

import (
	"shortlink/config"
	"shortlink/features/auth"
	"shortlink/features/auth/handler"
	"shortlink/features/auth/repository"
	"shortlink/features/auth/usecase"
	"shortlink/features/donate"
	"shortlink/features/goly"
	"shortlink/helpers"
	"shortlink/utils"

	golyHandler "shortlink/features/goly/handler"
	golyRepo "shortlink/features/goly/repository"
	golyUsecase "shortlink/features/goly/usecase"

	donateHandler "shortlink/features/donate/handler"
	donateRepo "shortlink/features/donate/repository"
	donateUsecase "shortlink/features/donate/usecase"

	"github.com/go-playground/validator/v10"
)

func UsersHandler() auth.Handler {
	config := config.InitConfig()

	db := utils.InitDB()
	jwt := helpers.New(config.Secret, config.RefreshSecret)
	hash := helpers.NewHash()

	repo := repository.New(db)
	uc := usecase.New(repo, jwt, hash)
	return handler.New(uc)
}

func GolyHandler() goly.Handler {
    db := utils.InitDB()
	repo := golyRepo.New(db)
	uc := golyUsecase.New(repo)
	return golyHandler.New(uc)
}

func DonateHandler() donate.Handler {
	config := config.InitConfig()
	snapClient := utils.MidtransSnapClient(config.MT_Server_Key)
	coreAPIClient := utils.MidtransCoreAPIClient(config.MT_Server_Key)
	validate := validator.New()
	db := utils.InitDB()
	repo := donateRepo.New(db, snapClient, coreAPIClient)
	uc := donateUsecase.New(repo, validate)
	return donateHandler.New(uc)
}