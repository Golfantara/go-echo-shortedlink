package features

import (
	"shortlink/config"
	"shortlink/features/auth"
	"shortlink/features/auth/handler"
	"shortlink/features/auth/repository"
	"shortlink/features/auth/usecase"
	"shortlink/features/goly"
	"shortlink/helpers"
	"shortlink/utils"

	golyHandler "shortlink/features/goly/handler"
	golyRepo "shortlink/features/goly/repository"
	golyUsecase "shortlink/features/goly/usecase"
)

func UsersHandler() auth.Handler {
	config := config.InitConfig()

	db := utils.InitDB()
	jwt := helpers.New(config.Secret, config.RefreshSecret)

	repo := repository.New(db)
	uc := usecase.New(repo, jwt)
	return handler.New(uc)
}

func GolyHandler() goly.Handler {
    db := utils.InitDB()
	repo := golyRepo.New(db)
	uc := golyUsecase.New(repo)
	return golyHandler.New(uc)


}