package main

import (
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Member struct {
	// gorm.Model
	Id               int     `json:"id"`
	No               *string `json:"no"`
	ProfileImg       string  `json:"profile_img"`
	FullName         string  `json:"full_name"`
	KanaName         *string `json:"kana_name"`
	Motto            *string `json:"motto"`
	Biography        *string `json:"biography"`
	StartDate        *string `json:"start_date"`
	EndDate          *string `json:"end_date"`
	EmploymentStatus *int    `json:"employment_status"`
	Status           *int    `json:"status"`
}

func main() {
	e := echo.New()

	e.POST("/member", createMember)  // 社員の新規登録
	e.GET("/members", getAllMembers) // 社員の一覧取得
	e.GET("/members/:id", getMember) // 社員の詳細情報取得
	e.Start(":3000")
}

// 社員の新規登録
func createMember(c echo.Context) error {
	var member Member
	err := c.Bind(&member)
	if err != nil {
		return c.String(http.StatusBadRequest, "It's a bad request!")
	}

	db := DBConnection()
	db.Create(&member)

	return c.String(http.StatusOK, member.FullName+"さんの社員情報を登録しました")
}

// 社員の一覧取得
func getAllMembers(c echo.Context) error {
	var members []Member
	db := DBConnection()
	db.Find(&members)

	return c.JSON(http.StatusOK, members)
}

// 社員の詳細情報取得
func getMember(c echo.Context) error {
	var member Member
	id := c.Param("id")
	db := DBConnection()
	db.First(&member, id)

	return c.JSON(http.StatusOK, member)
}

func DBConnection() *gorm.DB {
	err := godotenv.Load("./.env")
	if err != nil {
		panic("Error loading .env file")
	}

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	container_name := os.Getenv("CONTAINER_NAME")
	database := os.Getenv("MYSQL_DATABASE")
	dsn := user + ":" + password + "@tcp(" + container_name + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB")
	}

	return db
}
