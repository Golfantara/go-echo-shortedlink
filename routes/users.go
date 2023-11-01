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

	// goly.Use(echojwt.JWT([]byte(cfg.Secret)))
	goly.POST("", handler.CreateGoly)
	goly.GET("", handler.GetAllGoly())
	goly.GET("/r/:redirect", handler.Redirect)
	goly.PUT("/:id",handler.UpdateGoly())
	goly.DELETE("/:id",handler.DeleteGoly())
	goly.GET("/:id",handler.GolyDetails())
	goly.GET("/search/:short", handler.SearchGoly())
	goly.GET("/ip", handler.GetAllIP())
}

func Donate(e *echo.Echo, handler donate.Handler, cfg config.ProgramConfig){
	donate := e.Group("/api/donate")

	donate.POST("", handler.Insert(), echojwt.JWT([]byte(cfg.Secret)))
	donate.POST("/notifications", handler.Notifications())
	donate.GET("", handler.GetAllDonated(), echojwt.JWT([]byte(cfg.Secret)))
}