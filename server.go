package main

import (
	"fmt"
	"net/http"

	"github.com/byuoitav/common/db"
	"github.com/byuoitav/common/structs"
	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
)

func main() {

	port := ":9865"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	router.GET("/getPIList", getPIList)

	router.Static("/ui", "ui")

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)

}

func getPIList(context echo.Context) error {

	devices, err := db.GetDB().GetAllDevices()
	if err != nil {
		return fmt.Errorf("Failed trying to get all devices: %s", err)
	}

	var pis []structs.Device

	for _, d := range devices {
		if d.Type.ID == "Pi3" {
			pis = append(pis, d)
		}
	}

	context.JSON(http.StatusOK, pis)
	return nil

}
