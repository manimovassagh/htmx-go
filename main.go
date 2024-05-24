package main

import (
	"github.com/manimovassagh/htmx-app/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manimovassagh/htmx-app/handle"
)

func main() {
	app := echo.New()
	app.GET("/", HomeHandler)
	app.Logger.Fatal(app.Start(":4000"))
}

func HomeHandler(c echo.Context) error {
	return utils.Render(c, http.StatusOK, handle.Home())
}
