package main

import (
	"fmt"
	"main/data"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.CORS())

	if err := data.LoadData(); err != nil {
		fmt.Printf("failed to load data %s", err)
		return
	}

	e.GET("/data", func(c echo.Context) error {
		return c.JSON(http.StatusOK, data.CinemaData)
	})

	e.GET("/all-films", func(c echo.Context) error {
		return c.JSON(http.StatusOK, data.AllFilmsByDate)
	})

	e.GET("/cinema-names", func(c echo.Context) error {
		return c.JSON(http.StatusOK, data.CinemaNames)
	})

	e.GET("/cinema/:name", func(c echo.Context) error {
		name := c.Param("name")
		cinema, err := data.GetCinemaData(name)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, cinema)
	})

	if err := e.Start(":5555"); err != nil {
		fmt.Printf("failed to start server %s", err)
	}
}
