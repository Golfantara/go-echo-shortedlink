package main

import (
	"shortlink/config"
	"shortlink/features"
	"shortlink/routes"

	"github.com/labstack/echo/v4"
)

var (
	userHandler = features.UsersHandler()
	golyhandler = features.GolyHandler()
	cfg = config.InitConfig()
)


func main() {
	e := echo.New()
	routes.Users(e, userHandler, *cfg)
	routes.Goly(e, golyhandler, *cfg)


	e.Start(":8000")
}