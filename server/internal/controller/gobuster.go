package controller

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
)

type Gobuster struct {
	defaultFilePath string
}

func NewGobuster(defaultFilePath string) *Gobuster {
	return &Gobuster{defaultFilePath: defaultFilePath}
}

func (g *Gobuster) Dns(c echo.Context) error {
	request := struct {
		Domain     string `json:"domain"`
		ShowOption string `json:"showOption"`
		Timeout    string `json:"timeout"`
		Threads    int    `json:"threads"`
	}{}
	var response []string

	err := c.Bind(&request)
	if err != nil {
		log.Print(err)
		return &echo.HTTPError{
			Code: http.StatusNotAcceptable,
		}
	}

	cmd := exec.Command("gobuster", "dns",
		"-d", request.Domain,
		"--timeout", request.Timeout,
		"-t", strconv.Itoa(request.Threads),
		"-w", g.defaultFilePath,
	)

	if request.ShowOption == "IP" {
		cmd.Args = append(cmd.Args, "-i", "true")
	} else if request.ShowOption == "CNAME" {
		cmd.Args = append(cmd.Args, "-c", "true")
	}

	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	err = cmd.Run()
	if err != nil {
		log.Print(err)
		return &echo.HTTPError{
			Code: http.StatusBadRequest,
		}
	}

	result := buf.String()

	re := regexp.MustCompile(`Found: (.*?)(?:\s|=$)`)
	matches := re.FindAllStringSubmatch(result, -1)
	for _, match := range matches {
		for _, submatch := range match {
			if submatch != "" {
				response = append(response, submatch[7:len(submatch)-2])
				break
			}
		}
	}

	log.Print(response)

	return c.JSON(http.StatusOK, response)
}
