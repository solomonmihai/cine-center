package main

import (
	"main/data"
	"main/router"

	"fmt"
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

	router.SetupRouter(e)

	if err := e.Start(":" + string(port)); err != nil {
		fmt.Printf("failed to start server %s", err)
	}
}
