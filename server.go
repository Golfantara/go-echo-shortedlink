package main

import (
	"shortlink/config"
	"shortlink/features"
	"shortlink/helpers"
	"shortlink/routes"

	"github.com/labstack/echo/v4"
)

var (
	userHandler = features.UsersHandler()
	golyhandler = features.GolyHandler()
	donateHandler = features.DonateHandler()
	cfg = config.InitConfig()
)


func main() {
	e := echo.New()
	routes.Users(e, userHandler, *cfg)
	routes.Goly(e, golyhandler, *cfg)
	routes.Donate(e, donateHandler, *cfg)


	helpers.LogMiddlewares(e)
	e.Start(":8000")
}