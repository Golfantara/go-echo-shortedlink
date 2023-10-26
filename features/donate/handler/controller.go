package handler

import (
	"encoding/json"
	"net/http"
	"shortlink/features/donate"
	"shortlink/features/donate/dtos"
	"shortlink/helpers"
	"strings"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service donate.Usecase
}

func New(service donate.Usecase) donate.Handler {
	return &controller{
        service: service,
    }
}

func(ctl *controller) GetAllDonated() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)

		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helpers.Response("Please provide query `page` and `size` in number!"))
		}
		donate := ctl.service.FindAll(page, size)
		if donate == nil {
			return ctx.JSON(404, helpers.Response("there is no donated"))
		}
		return ctx.JSON(200, helpers.Response("succes!", map[string]any{
			"data":donate,
		}))
	}
}

func(ctl * controller) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		var transactionInput dtos.TransactionInput

		if err := c.Bind(&transactionInput); err!= nil {
            return c.JSON(http.StatusBadRequest, helpers.Response("Invalid request body!"))
        }

		result, err := ctl.service.Create(transactionInput)
		if err != nil {
			if strings.Contains(err.Error(),"validation fail"){
				return c.JSON(http.StatusBadRequest, helpers.Response(err.Error(), nil))
			}
			return c.JSON(http.StatusInternalServerError, helpers.Response(err.Error(), nil))
		}
		responseMap := map[string]any {
			"data": result,
		}
		return c.JSON(http.StatusCreated, helpers.Response("Success!", responseMap))
	}
}

func(ctl *controller) Notifications() echo.HandlerFunc{
	return func(c echo.Context) error {
		var notificationPayload map[string]any

		if err := json.NewDecoder(c.Request().Body).Decode(&notificationPayload); err != nil {
			return c.JSON(http.StatusBadRequest, helpers.Response("Error when parsing data"))
		}
		err := ctl.service.Notifications(notificationPayload)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.Response(err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helpers.Response("Success!", nil))
	}
}