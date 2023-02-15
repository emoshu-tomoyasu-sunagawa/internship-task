package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// ルーティング
	e.GET("/", hello)
	e.POST("/employee", createEmployee)

	e.Logger.Fatal(e.Start(":3000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func createEmployee(c echo.Context) error {
	return c.String(http.StatusOK, "EMoshU")
}
