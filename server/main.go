package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/OVantsevich/security-utils/server/internal/controller"
	"github.com/OVantsevich/security-utils/server/internal/nmap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	nmapCacheTTL = time.Minute * 5

	indexHTMl = "/index.html"

	apiGroupPrefix = "/api"
	port           = "12345"
)

func gracefulShutdown(stop func() error, FilesStorageDirectory string) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
	files, err := filepath.Glob(FilesStorageDirectory + "*.xml")
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
	filesStorageDirectory := "/bin/db/"
	frontEndFilesStorageDirectory := "/bin/static"

	e := echo.New()
	go gracefulShutdown(e.Close, filesStorageDirectory)

	nmapService := nmap.NewNmap(nmapCacheTTL, filesStorageDirectory)
	nmapController := controller.NewNmap(nmapService)
	gobusterController := controller.NewGobuster("/bin/db/wordlist.txt")

	apiGroup := e.Group(apiGroupPrefix)

	apiGroup.POST("/gobuster/scan", gobusterController.Dns)

	apiGroup.POST("/nmap/scan", nmapController.Scan)
	apiGroup.POST("/nmap/report", nmapController.Report)

	e.Static("/static", frontEndFilesStorageDirectory)
	e.GET("", func(c echo.Context) error {
		index, err := os.ReadFile(frontEndFilesStorageDirectory + indexHTMl)
		if err != nil {
			log.Printf("os.ReadFile: %v", err)
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.HTML(http.StatusOK, string(index))
	})

	e.Use(middleware.CORS())
	log.Fatal(e.Start("localhost:" + port))
}
