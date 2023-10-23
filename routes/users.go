package routes

import (
	"shortlink/features/auth"
	"shortlink/features/goly"

	"github.com/labstack/echo/v4"
)

func Users(e *echo.Echo, handler auth.Handler){
	users := e.Group("/users")

	users.GET("", handler.GetUsers())
	users.POST("", handler.CreateUsers())
	users.POST("/login", handler.LoginUsers())
	users.PUT("/:id", handler.UpdateUsers())
	users.DELETE("/:id", handler.DeleteUsers())
	users.GET("/:id", handler.UsersDetails())
}

func Goly(e *echo.Echo, handler goly.Handler){
	goly := e.Group("/goly")

	goly.POST("", handler.CreateGoly)
	goly.GET("", handler.GetAllGoly())
	goly.GET("/r/:redirect", handler.Redirect)
	goly.PUT("/:id",handler.UpdateGoly())
	goly.DELETE("/:id",handler.DeleteGoly())
	goly.GET("/:id",handler.GolyDetails())
}