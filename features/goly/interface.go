package goly

import (
	"shortlink/features/goly/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Insert(newGoly *Goly) *Goly
	SelectByID(golyID int) *Goly
	Update(goly Goly) int64
	DeleteByID(golyID int) int64
}

type UseCase interface {
	Create(newGoly dtos.CreateGolyInput) *dtos.GolyResponse
}

type Handler interface {
	CreateGoly(c echo.Context) error
}