package routes

import (
	"shortlink/config"
	"shortlink/features/auth"
	"shortlink/features/donate"
	"shortlink/features/goly"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Users(e *echo.Echo, handler auth.Handler, cfg config.ProgramConfig){
	users := e.Group("/api/users")

	users.GET("", handler.GetUsers())
	users.POST("", handler.CreateUsers())
	users.POST("/login", handler.LoginUsers())
	users.PUT("/:id", handler.UpdateUsers(), echojwt.JWT([]byte(cfg.Secret)))
	users.DELETE("/:id", handler.DeleteUsers(), echojwt.JWT([]byte(cfg.Secret)))
	users.GET("/:id", handler.UsersDetails(), echojwt.JWT([]byte(cfg.Secret)))
}

func Goly(e *echo.Echo, handler goly.Handler, cfg config.ProgramConfig){
	goly := e.Group("/api/goly")

	goly.POST("", handler.CreateGoly, echojwt.JWT([]byte(cfg.Secret)))
	goly.GET("", handler.GetAllGoly())
	goly.GET("/r/:redirect", handler.Redirect, echojwt.JWT([]byte(cfg.Secret)))
	goly.PUT("/:id",handler.UpdateGoly(), echojwt.JWT([]byte(cfg.Secret)))
	goly.DELETE("/:id",handler.DeleteGoly(), echojwt.JWT([]byte(cfg.Secret)))
	goly.GET("/:id",handler.GolyDetails(), echojwt.JWT([]byte(cfg.Secret)))
	goly.GET("/search/:short", handler.SearchGoly())
}

func Donate(e *echo.Echo, handler donate.Handler, cfg config.ProgramConfig){
	donate := e.Group("/api/donate")

	donate.POST("", handler.Insert())
	donate.POST("/transactions/notifications", handler.Notifications())
	donate.GET("", handler.GetAllDonated())
}