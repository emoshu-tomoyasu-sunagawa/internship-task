package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type Employee struct {
	gorm.Model
	Id               int    `json:"id"`
	No               string `json:"no"`
	ProfileImg       string `json:"profile_img"`
	FullName         string `json:"full_name"`
	KanaName         string `json:"kana_name"`
	Motto            string `json:"motto"`
	Biography        string `json:"biography"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	EmploymentStatus string `json:"employment_status"`
	Status           string `json:"status"`
}

type Employees struct {
	Employees []Employee `json:"employee"`
}

func main() {
	e := echo.New()

	fmt.Printf("Air聴いてる？")
	// dsn := "root:password@tcp(127.0.0.1:3306)/emonavi_db?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("Failed to connect DB")
	// }

	// ルーティングへ飛ばす
	// initRouting(e)

	e.GET("/", hello)
	e.POST("/employee", createEmployee)

	e.Start(":3000")
	e.Logger.Fatal(e.Start(":3000"))
}

// ルーティングの一覧
// func initRouting(e *echo.Echo) {
// 	e.GET("/", hello)
// 	e.POST("/employee", createEmployee)
// }

func hello(c echo.Context) error {
	db := DBConnection()
	var member = Employee{No: "020", ProfileImg: "https://emoshu.co.jp", FullName: "John Doe"}
	db.Create(&member)

	return c.String(http.StatusOK, "新しい従業員を追加しました！")
}

func createEmployee(c echo.Context) error {
	db := DBConnection()
	no := c.FormValue("no")
	profile_img := c.FormValue("profile_img")
	full_name := c.FormValue("full_name")

	fmt.Printf(profile_img)

	var member = Employee{No: no, ProfileImg: profile_img, FullName: full_name}
	db.Create(&member)

	return c.String(http.StatusOK, full_name)
}

func DBConnection() *gorm.DB {
	dsn := "user:password@tcp(dbcontainer)/emonavi_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB")
	}

	return db
}
