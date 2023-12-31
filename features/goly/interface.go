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
	SearchingGoly(short string) ([]Goly, error)
	StoreIPForGoly(golyID uint64, ip string) error
	PaginateIP(page, size int) []IPAdresses
}

type UseCase interface {
	Create(newGoly dtos.CreateGolyInput) *dtos.GolyResponse
	FindAllGoly(page, size int) []dtos.GolyResponse
	FindGolyByID(golyID int) *dtos.GolyResponse
	Modify(golyData dtos.CreateGolyInput, golyID int) bool
	Remove(golyID int) bool
	IncreaseClickAndRedirect(goly Goly) error
	GetGolyByUrl(url string) (Goly, error)
	SearchGoly(short string) ([]Goly, error)
	StoreIPAddress(goly Goly, ip string) error
	FindAllIP(page, size int) []IPAdresses
	ExportIPToPDfAndSave() (string, error)
}

type Handler interface {
	CreateGoly(c echo.Context) error
	GolyDetails() echo.HandlerFunc
	GetAllGoly() echo.HandlerFunc
	UpdateGoly() echo.HandlerFunc
	DeleteGoly() echo.HandlerFunc
	Redirect(c echo.Context) error
	SearchGoly() echo.HandlerFunc
	GetAllIP() echo.HandlerFunc
	ExportIPToPDF(c echo.Context) (error)
}