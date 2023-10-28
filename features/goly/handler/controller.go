package handler

import (
	"fmt"
	"net/http"
	"shortlink/features/goly"
	"shortlink/features/goly/dtos"
	"shortlink/helpers"
	"strconv"
	"strings"

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

var validate *validator.Validate

func (ctl *controller) GetAllGoly() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helpers.Response("Please provide query `page` and `size` in number!"))
		}

		goly := ctl.service.FindAllGoly(page, size)

		if goly == nil {
			return ctx.JSON(404, helpers.Response("There is No goly!"))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": goly,
		}))
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

func (ctl *controller) GolyDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		golyID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response(err.Error()))
		}

		goly := ctl.service.FindGolyByID(golyID)

		if goly == nil {
			return ctx.JSON(404, helpers.Response("Goly Not Found!"))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": goly,
		}))
	}
}

func (ctl *controller) UpdateGoly() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.CreateGolyInput{}

		golyID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helpers.Response(errParam.Error()))
		}

		goly := ctl.service.FindGolyByID(golyID)

		if goly == nil {
			return ctx.JSON(404, helpers.Response("Goly Not Found!"))
		}
		
		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helpers.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		update := ctl.service.Modify(input, golyID)

		if !update {
			return ctx.JSON(500, helpers.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helpers.Response("Goly Success Updated!"))
	}
}

func (ctl *controller) DeleteGoly() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		golyID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response(err.Error()))
		}

		goly := ctl.service.FindGolyByID(golyID)

		if goly == nil {
			return ctx.JSON(404, helpers.Response("Goly Not Found!"))
		}

		delete := ctl.service.Remove(golyID)

		if !delete {
			return ctx.JSON(500, helpers.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helpers.Response("Goly Success Deleted!", nil))
	}
}

func (ctl *controller) SearchGoly() echo.HandlerFunc {
    return func(ctx echo.Context) error {
        golyname := ctx.Param("short")
        golies, err := ctl.service.SearchGoly(golyname)
        if err != nil {
            return ctx.JSON(400, helpers.Response(err.Error()))
        }
        return ctx.JSON(200, golies)
    }
}

func(ctl *controller) Redirect(c echo.Context) error {
	golyUrl := c.Param("redirect")
    goly, err := ctl.service.GetGolyByUrl(golyUrl)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "message": "could not find link in DB " + err.Error(),
        })
    }

	ip := c.RealIP()

	err = ctl.service.StoreIPAddress(goly, ip)
	if err != nil {
		fmt.Printf("error storing IP address: %v \n", err)
	}
    err = ctl.service.IncreaseClickAndRedirect(goly)
    if err != nil {
        fmt.Printf("error updating: %v\n", err)
    }

    if !strings.HasPrefix(goly.Redirect, "http://") && !strings.HasPrefix(goly.Redirect, "https://") {
        goly.Redirect = "http://" + goly.Redirect
    }

    return c.Redirect(http.StatusTemporaryRedirect, goly.Redirect)
}