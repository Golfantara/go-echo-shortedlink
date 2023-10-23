package auth

import (
	"shortlink/features/auth/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Users
	Insert(newUsers *Users) *Users
	SelectByID(userID int) *Users
	Update(user Users) int64
	DeleteByID(userID int) int64
	Login(email, password string) (*Users, error)
}

type UseCase interface {
	FindAll(page, size int) []dtos.ResUsers
	FindByID(userID int) *dtos.ResUsers
	Create(newUsers dtos.InputUsers) *dtos.ResRegister
	Modify(userData dtos.InputUsers, userID int) bool
	Remove(userID int) bool
	Login(email, password string) (*dtos.ResLogin, error)
}

type Handler interface {
	GetUsers() echo.HandlerFunc
	UsersDetails() echo.HandlerFunc
	CreateUsers() echo.HandlerFunc
	UpdateUsers() echo.HandlerFunc
	DeleteUsers() echo.HandlerFunc
	LoginUsers() echo.HandlerFunc
}