package main

import (
	"joshdstockdale/grounded/handler"
	"joshdstockdale/grounded/middleware"

	"github.com/labstack/echo/v4"
	// "joshdstockdale/grounded/handler"
)

type DB struct {}

func main() {
	app := echo.New()

	//var db DB
	pageHandler := handler.PageHandler{}
	app.Use(middleware.WithUser)
	app.GET("/home", pageHandler.HandlePageShow)

	app.Start(":3000")
}
