package main

import (
	"echo-gorm-crud-api-example/database"

	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()
	err := sqlDB.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DBに接続できませんでした")
	} else {
		return c.String(http.StatusOK, "Hello, world!")
	}
}

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":3000"))
}
