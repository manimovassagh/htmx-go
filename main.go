package main

import (
	"github.com/labstack/echo/v4"
	"github.com/manimovassagh/htmx-app/handler"
)

func main() {
	app := echo.New()
	app.GET("/", handler.HomeHandler)
	app.Logger.Fatal(app.Start(":4000"))
}
