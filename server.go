package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/pxao", save)
	e.Logger.Fatal(e.Start(":1323"))

}

func save(c echo.Context) error {

	return c.HTML(http.StatusOK, "<h1 style={color: red}>Thank you!</h1>")
}
