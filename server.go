package main

import (
	"net/http"

	"github.com/byuoitav/state-parsing/elk"
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

	query := `{
		"_source": [
			"room",
			"hostname"
		],
		"query": {
		  "bool": {
			"must": {
			  "match": {
				"_type": "control-processor"
			  }
			}
		  }
		},
		"size": 1000
	  }`

	body, err := elk.MakeELKRequest("POST", "/oit-static-av-devices/_search", []byte(query))

	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	} else {
		return context.JSONBlob(http.StatusOK, body)
	}

}
