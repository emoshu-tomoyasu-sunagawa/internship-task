package main

import (
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type Member struct {
	// gorm.Model
	Id               int       `json:"id"`
	No               string    `json:"no"`
	ProfileImg       string    `json:"profile_img"`
	FullName         string    `json:"full_name"`
	KanaName         string    `json:"kana_name"`
	Motto            string    `json:"motto"`
	Biography        string    `json:"biography"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	EmploymentStatus int       `json:"employment_status"`
	Status           int       `json:"status"`
}

type Members struct {
	Members []Member `json:"member"`
}

func main() {
	e := echo.New()

	e.GET("/", hello)
	e.POST("/member", createMember) // ユーザーの新規登録

	e.Start(":3000")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "新しい従業員を追加しました！")
}

func createMember(c echo.Context) error {
	db := DBConnection()

	no := c.FormValue("no")
	profile_img := c.FormValue("profile_img")
	full_name := c.FormValue("full_name")
	kana_name := c.FormValue("kana_name")
	motto := c.FormValue("motto")
	biography := c.FormValue("biography")
	start_date, _ := time.Parse("2006-01-02 15:04:05", c.FormValue("start_date"))
	end_date, _ := time.Parse("2006-01-02 15:04:05", c.FormValue("end_date"))
	employment_status, _ := strconv.Atoi(c.FormValue("employment_status"))
	status, _ := strconv.Atoi(c.FormValue("status"))

	var member = Member{
		No:               no,
		ProfileImg:       profile_img,
		FullName:         full_name,
		KanaName:         kana_name,
		Motto:            motto,
		Biography:        biography,
		StartDate:        start_date,
		EndDate:          end_date,
		EmploymentStatus: employment_status,
		Status:           status,
	}
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
