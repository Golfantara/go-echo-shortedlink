package usecase

import (
	"errors"
	"shortlink/features/auth"
	"shortlink/features/auth/dtos"
	"shortlink/helpers"
	"strconv"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model auth.Repository
	jwt helpers.JWTInterface
	hash helpers.HashInterface
}

func New(model auth.Repository, jwt helpers.JWTInterface, hash helpers.HashInterface) auth.UseCase {
	return &service{
        model: model,
        jwt: jwt,
		hash: hash,
    }
}

func (svc *service) FindAll(page, size int) []dtos.ResUsers {
	var user []dtos.ResUsers

	userEnt := svc.model.Paginate(page, size)

	for _, users := range userEnt {
		var data dtos.ResUsers

		if err := smapping.FillStruct(&data, smapping.MapFields(users)); err != nil {
			log.Error(err.Error())
		}

		user = append(user, data)
	}

	return user
}

func (svc *service) FindByID(userID int) *dtos.ResUsers {
	res := dtos.ResUsers{}
	user := svc.model.SelectByID(userID)

	if user == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(user))
	if err != nil {
		log.Error(err)
		return nil
	}
	return &res
}

func (svc *service) Create(newUsers dtos.InputUsers) (*dtos.ResRegister, error) {
	user := auth.Users{}
	err := smapping.FillStruct(&user, smapping.MapFields(newUsers))
	
	if err != nil {
		log.Error(err)
		return nil, errors.New("failed to create user")
	}
	user.Password = svc.hash.HashPassword(user.Password)

	
	// user. = helpers.GenerateUUID()
	userID := svc.model.Insert(&user)

	if userID == nil {
        return nil, errors.New("failed to insert user")
    }

	resUser := dtos.ResRegister{}
	errRes := smapping.FillStruct(&resUser, smapping.MapFields(user))
	if errRes != nil {
		log.Error(errRes)
		return nil, errors.New("failed to mapping users")
	}
	ID := strconv.Itoa(resUser.ID)
	tokenData := svc.jwt.GenerateJWT(ID)

	if tokenData == nil {
		log.Error("token procces fail")
	}

	resUser.Token = tokenData

	return &resUser, nil
}

func (svc *service) Modify(userData dtos.InputUsers, userID int) bool {
	newUser := auth.Users{}

	err := smapping.FillStruct(&newUser, smapping.MapFields(userData))
	if err != nil {
		log.Error(err)
		return false
	}

	newUser.ID = userID
	rowsAffected := svc.model.Update(newUser)

	if rowsAffected <= 0 {
		log.Error("there is no customer updated!")
		return false
	}
	return true
}

func (svc *service) Remove(userID int) bool {
	rowsAffected := svc.model.DeleteByID(userID)

	if rowsAffected <= 0 {
        log.Error("there is no user removed!")
        return false
    }
	return true
}

func (svc *service) Login(email, password string) (*dtos.ResLogin, error) {
	user, err := svc.model.Login(email, password)
	if err != nil {
		return nil, err
	}

	if !svc.hash.CompareHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	tokenData := svc.jwt.GenerateJWT(strconv.Itoa(user.ID))
	return &dtos.ResLogin{
		Name:  user.Fullname,
		Email: user.Email,
		Token: tokenData,
	}, nil
}