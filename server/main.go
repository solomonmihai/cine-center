package main

import (
	"fmt"
	"main/data"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.CORS())

	if err := data.CreateScheduler(); err != nil {
		fmt.Printf("failed to start scheduler %s", err)
		return
	}

	if err := data.UpdateData(); err != nil {
		return
	}

	e.Static("/", "public")

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

	if err := e.Start(":" + string(port)); err != nil {
		fmt.Printf("failed to start server %s", err)
	}
}
