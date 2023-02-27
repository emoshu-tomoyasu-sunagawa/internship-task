package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db *sql.DB
	var err error

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "user", "password", "localhost", "3306", "emonavi_db"))
	check_nil(err)
	defer db.Close()

	err = db.Ping()
	check_nil(err)

	file, err := os.Create("members.csv")
	check_nil(err)
	defer file.Close()

	cw := csv.NewWriter(file)
	defer cw.Flush()

	var id int
	var No, ProfileImg, FullName, KanaName, Motto, Biography, StartDate, EndDate, EmploymentStatus, Status, CreatedAt, UpdatedAt string
	rows, err := db.Query("select id, coalesce(no, '') as no, profile_img, full_name, coalesce(kana_name, '') as kana_name, coalesce(motto, '') as motto, coalesce(biography, '') as biography, coalesce(start_date, '') as start_date, coalesce(end_date, '') as end_date, coalesce(employment_status, '') as emplystatus, coalesce(status, '') as status, created_at, updated_at from members;")
	check_nil(err)
	defer rows.Close()

	col := []string{"no", "profile_img", "full_name", "kana_name", "motto", "biographt", "start_date", "end_date", "employment_status", "status", "created_at", "updated_at"}
	cw.Write(col)

	for rows.Next() {
		err := rows.Scan(&id, &No, &ProfileImg, &FullName, &KanaName, &Motto, &Biography, &StartDate, &EndDate, &EmploymentStatus, &Status, &CreatedAt, &UpdatedAt)
		check_nil(err)
		col := []string{No, ProfileImg, FullName, KanaName, Motto, Biography, StartDate, EndDate, EmploymentStatus, Status}
		cw.Write(col)
	}

	err = rows.Err()
	check_nil(err)
}

func check_nil(err error) {
	if err != nil {
		panic(err)
	}
}
