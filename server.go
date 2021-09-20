package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/uuid", func(c echo.Context) error {
		type UUID struct {
			STATUS string `json:"status" xml:"status"`
			UUID   string `json:"uuid" xml:"uuid"`
		}
		id := uuid.New()
		uuid := &UUID{
			"200",
			id.String(),
		}
		return c.JSON(http.StatusOK, uuid)
	})
	e.GET("/")
	e.Logger.Fatal(e.Start(":1323"))
}
