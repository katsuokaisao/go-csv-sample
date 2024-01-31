package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/katsuokaisao/go-csv-sample/model"
	"github.com/katsuokaisao/go-csv-sample/util/csv"
	"golang.org/x/text/encoding/unicode"
)

func main() {
	dataList := []model.Data{
		{
			ID:      1,
			Name:    "taro",
			Age:     20,
			LoginAt: time.Now(),
		},
		{
			ID:      2,
			Name:    "jiro",
			Age:     30,
			LoginAt: time.Now().Add(1 * time.Hour),
		},
		{
			ID:      3,
			Name:    "saburo",
			Age:     40,
			LoginAt: time.Now().Add(2 * time.Hour),
		},
	}

	w, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}

	header := []string{
		"id",
		"name",
		"age",
		"login_at",
	}
	enc := unicode.UTF8.NewEncoder()
	useCRLF := true
	writer := csv.NewCSVWriter(
		header,
		useCRLF,
		enc,
	)

	if err := writer.WriteRows(w, convertDataListToRows(dataList)); err != nil {
		panic(err)
	}

	if err := w.Close(); err != nil {
		panic(err)
	}

	dataList, err = csv.ReadData("test.csv")
	if err != nil {
		panic(err)
	}

	for _, d := range dataList {
		fmt.Printf("id: %d\n", d.ID)
		fmt.Printf("name: %s\n", d.Name)
		fmt.Printf("age: %d\n", d.Age)
		fmt.Printf("login_at: %s\n", d.LoginAt)
		fmt.Println()
	}

	dataCh, err := csv.ReadDataAsCh("test.csv")
	if err != nil {
		panic(err)
	}

	outputStdOut(dataCh)
}

func convertDataListToRows(dataList []model.Data) []csv.Row {
	rows := make([]csv.Row, 0, len(dataList))

	for _, d := range dataList {
		rows = append(rows, convertDataToRow(d))
	}

	return rows
}

func convertDataToRow(data model.Data) csv.Row {
	// location, _ := time.LoadLocation("Asia/Tokyo")
	location, _ := time.LoadLocation("UTC")
	format := "2006-01-02 15:04:05"

	cells := make([]string, 4)
	cells[0] = strconv.Itoa(data.ID)
	cells[1] = data.Name
	cells[2] = strconv.Itoa(data.Age)
	cells[3] = data.LoginAt.In(location).Format(format)

	return csv.Row{Cells: cells}
}

func outputStdOut(dataCh <-chan model.Data) {
	for d := range dataCh {
		row := convertDataToRow(d)
		fmt.Println(strings.Join(row.Cells, ","))
	}
}
