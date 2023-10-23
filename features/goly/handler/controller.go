package handler

import (
	"shortlink/features/goly"
	"shortlink/features/goly/dtos"
	"shortlink/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service goly.UseCase
}

func New(service goly.UseCase) goly.Handler {
	return &controller {
		service: service,
	}
}

func (ctl *controller) CreateGoly(c echo.Context) error {
	input := dtos.CreateGolyInput{}
	c.Bind(&input)
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(input)
	if err != nil {
		errMap := helpers.ErrorMapValidation(err)
		return c.JSON(400, helpers.Response("Bad Request!", map[string]any {
			"error": errMap,
		}))
	}

	goly := ctl.service.Create(input)
	if goly == nil {
        return c.JSON(500, helpers.Response("Something whent wrong!", nil))
    }
	return c.JSON(200, helpers.Response("succes!", map[string]any {
		"data": goly,
	}))
}