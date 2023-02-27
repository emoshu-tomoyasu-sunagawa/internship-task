package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv: ", err)
	}

	file, err := os.Create("member_list.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	cw := csv.NewWriter(file)
	defer cw.Flush()
}
