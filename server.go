package main

import (
	"shortlink/features"
	"shortlink/routes"

	"github.com/labstack/echo/v4"
)

var (
	userHandler = features.UsersHandler()
	golyhandler = features.GolyHandler()
)


func main() {
	e := echo.New()
	routes.Users(e, userHandler)
	routes.Goly(e, golyhandler)


	e.Start(":8000")
}