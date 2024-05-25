package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/manimovassagh/htmx-app/handle"
	"github.com/manimovassagh/htmx-app/utilities"
	"net/http"
)

func HomeHandler(c echo.Context) error {
	return utilities.Render(c, http.StatusOK, handle.Home())
}
func FetcherHandler(c echo.Context) error {

	return utilities.Render(c, http.StatusOK, handle.Fetcher())
}
