package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/uuid", func(c echo.Context) error {
		uuid := genUUID()
		return c.JSON(http.StatusOK, uuid)
	})
	e.GET("/")
	e.Logger.Fatal(e.Start(":1323"))
}
