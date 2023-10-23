package goly

import (
	"shortlink/features/goly/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Insert(newGoly *Goly) *Goly
	SelectByID(golyID int) *Goly
	Update(goly Goly) int64
	UpdateButton(goly Goly) error
	DeleteByID(golyID int) int64
	Paginate(page, size int) []Goly
	FindByGolyUrl(url string) (Goly, error)
}

type UseCase interface {
	Create(newGoly dtos.CreateGolyInput) *dtos.GolyResponse
	FindAllGoly(page, size int) []dtos.GolyResponse
	FindGolyByID(golyID int) *dtos.GolyResponse
	Modify(golyData dtos.CreateGolyInput, golyID int) bool
	Remove(golyID int) bool
	IncreaseClickAndRedirect(goly Goly) error
	GetGolyByUrl(url string) (Goly, error)
}

type Handler interface {
	CreateGoly(c echo.Context) error
	GolyDetails() echo.HandlerFunc
	GetAllGoly() echo.HandlerFunc
	UpdateGoly() echo.HandlerFunc
	DeleteGoly() echo.HandlerFunc
	Redirect(c echo.Context) error
}