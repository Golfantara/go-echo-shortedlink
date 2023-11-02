package handler

import (
	"shortlink/features/auth"
	"shortlink/features/auth/dtos"
	"shortlink/helpers"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service auth.UseCase
}

func New(service auth.UseCase) auth.Handler {
	return &controller{
        service: service,
    }
}

var validate *validator.Validate

func (ctl *controller) GetUsers() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helpers.Response("Please provide query `page` and `size` in number!"))
		}

		users := ctl.service.FindAll(page, size)

		if users == nil {
			return ctx.JSON(404, helpers.Response("There is No users!"))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": users,
		}))
	}
}

func (ctl *controller) UsersDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		userID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response(err.Error()))
		}

		user := ctl.service.FindByID(userID)

		if user == nil {
			return ctx.JSON(404, helpers.Response("Users Not Found!"))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": user,
		}))
	}
}

func (ctl *controller) CreateUsers() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := dtos.InputUsers{}

		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())

		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helpers.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		user, _ := ctl.service.Create(input)

		if user == nil {
			return ctx.JSON(500, helpers.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": user,
		}))
	}
}

func (ctl *controller) UpdateUsers() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputUsers{}

		userID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helpers.Response(errParam.Error()))
		}

		user := ctl.service.FindByID(userID)

		if user == nil {
			return ctx.JSON(404, helpers.Response("User Not Found!"))
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

		update := ctl.service.Modify(input, userID)

		if !update {
			return ctx.JSON(500, helpers.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helpers.Response("Customer Success Updated!"))
	}
}

func (ctl *controller) DeleteUsers() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		userID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response(err.Error()))
		}

		user := ctl.service.FindByID(userID)

		if user == nil {
			return ctx.JSON(404, helpers.Response("User Not Found!"))
		}

		delete := ctl.service.Remove(userID)

		if !delete {
			return ctx.JSON(500, helpers.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helpers.Response("User Success Deleted!", nil))
	}
}

func (ctl *controller) LoginUsers() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		loginData := dtos.LoginUsers{}

		if err := ctx.Bind(&loginData); err != nil {
			return ctx.JSON(400, helpers.Response("Invalid request body!"))
		}

		loginRes, err := ctl.service.Login(loginData.Email, loginData.Password)
		if err != nil {
			return ctx.JSON(401, helpers.Response("Invalid credentials!"))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any{
			"data": loginRes,
		}))
	}
}