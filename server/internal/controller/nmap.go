package controller

import (
	"context"
	"github.com/OVantsevich/security-utils/server/internal/model"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type NmapService interface {
	GetXML(ctx context.Context, parameters model.NmapScanParameters) ([]byte, error)
	Get(ctx context.Context, parameters model.NmapScanParameters) ([]byte, error)
}

type Nmap struct {
	nmapService NmapService
}

func NewNmap(nmapService NmapService) *Nmap {
	return &Nmap{nmapService: nmapService}
}

func (n *Nmap) Scan(c echo.Context) error {
	var request model.NmapScanParameters

	err := c.Bind(&request)
	if err != nil {
		log.Print(err)
		return &echo.HTTPError{
			Code: http.StatusNotAcceptable,
		}
	}

	response, err := n.nmapService.Get(c.Request().Context(), request)
	if err != nil {
		log.Printf("Nmap-Scan-n.nmapService.Get: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.HTML(http.StatusOK, string(response))
}

func (n *Nmap) Report(c echo.Context) error {
	var request model.NmapScanParameters

	err := c.Bind(&request)
	if err != nil {
		log.Print(err)
		return &echo.HTTPError{
			Code: http.StatusNotAcceptable,
		}
	}

	response, err := n.nmapService.GetXML(c.Request().Context(), request)
	if err != nil {
		log.Printf("Nmap-Report-n.nmapService.GetXML: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.HTML(http.StatusOK, string(response))
}
