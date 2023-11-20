package main

import (
	"github.com/OVantsevich/security-utils/server/internal/controller"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.GET("/api/tools", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []struct {
			Id   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		}{
			{
				Id:   "0",
				Name: "gobuster",
			},
			{
				Id:   "1",
				Name: "nmap",
			},
		})
	})
	gobuster := controller.NewGobuster("/home/oleg/GolandProjects/safqa/security-utils/server/wordlist.txt")
	e.POST("/api/gobuster-scan", gobuster.Dns)
	e.Use(middleware.CORS())
	e.Start("localhost:12345")
}
