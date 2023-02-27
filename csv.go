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
	var ProfileImg, FullName string
	rows, err := db.Query("select id, profile_img, full_name from members;")
	check_nil(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &ProfileImg, &FullName)
		check_nil(err)
		col := []string{ProfileImg, FullName}
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
