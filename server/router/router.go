package router

import (
	"main/data"
	"net/http"
	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	api := e.Group("/api")

	api.GET("/data", func(c echo.Context) error {
		return c.JSON(http.StatusOK, data.CinemaData)
	})

	api.GET("/all-films", func(c echo.Context) error {
		return c.JSON(http.StatusOK, data.AllFilmsByDate)
	})

	api.GET("/cinema-names", func(c echo.Context) error {
		return c.JSON(http.StatusOK, data.CinemaNames)
	})

	api.GET("/cinema/:name", func(c echo.Context) error {
		name := c.Param("name")
		cinema, err := data.GetCinemaData(name)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, cinema)
	})
}
