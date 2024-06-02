package main

import (
	"vatansoft/config"
	"vatansoft/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()

	e := echo.New()
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
