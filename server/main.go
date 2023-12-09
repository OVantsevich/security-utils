package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/OVantsevich/security-utils/server/internal/config"
	"github.com/OVantsevich/security-utils/server/internal/controller"
	"github.com/OVantsevich/security-utils/server/internal/nmap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func clean(stop func() error, FilesStorageDirectory string) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
	files, err := filepath.Glob(FilesStorageDirectory + "/*.xml")
	if err != nil {
		log.Fatalf("can't get files .xml: %v", err)
	}
	for _, f := range files {
		err = os.Remove(f)
		if err != nil {
			log.Fatalf("can't remove file: %v", err)
		}
	}
	err = stop()
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("can't parse configuration: %v", err)
	}
	e := echo.New()

	go clean(e.Close, cfg.FilesStorageDirectory)

	nmapService := nmap.NewNmap(cfg.NmapCacheTTl, cfg.FilesStorageDirectory)
	nmapController := controller.NewNmap(nmapService)

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

	e.POST("/api/nmap/scan", nmapController.Scan)
	e.POST("/api/nmap/report", nmapController.Report)

	e.Use(middleware.CORS())
	e.Static("/static", "server/static")
	e.GET("", func(c echo.Context) error {
		fil,
			c.HTML()
	})
	e.Start("localhost:12345")
}
