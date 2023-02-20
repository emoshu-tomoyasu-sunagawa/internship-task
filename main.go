package main

import (
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type Member struct {
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

type Members struct {
	Members []Member `json:"member"`
}

func main() {
	e := echo.New()

	e.GET("/", hello)
	e.POST("/member", createMember)

	e.Start(":3000")
}

func hello(c echo.Context) error {
	db := DBConnection()
	var member = Member{No: "020", ProfileImg: "https://emoshu.co.jp", FullName: "John Doe"}
	db.Create(&member)

	return c.String(http.StatusOK, "新しい従業員を追加しました！")
}

func createMember(c echo.Context) error {
	db := DBConnection()
	no := c.FormValue("no")
	profile_img := c.FormValue("profile_img")
	full_name := c.FormValue("full_name")

	var member = Member{No: no, ProfileImg: profile_img, FullName: full_name}
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
