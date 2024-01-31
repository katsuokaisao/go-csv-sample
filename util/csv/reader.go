package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/katsuokaisao/go-csv-sample/model"
)

func ReadData(filename string) ([]model.Data, error) {
	res := make([]model.Data, 0, 1024)
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error in Open: %v", err)
	}

	r := csv.NewReader(f)

	record, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("error in Read: %v", err)
	}
	header := make(map[string]int)
	for i, v := range record {
		header[v] = i
	}

	IDIndex, ok := header["id"]
	if !ok {
		return nil, errors.New("id column not found")
	}
	nameIndex, ok := header["name"]
	if !ok {
		return nil, errors.New("name column not found")
	}
	ageIndex, ok := header["age"]
	if !ok {
		return nil, errors.New("age column not found")
	}
	loginAtIndex, ok := header["login_at"]
	if !ok {
		return nil, errors.New("login at column not found")
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error in Read: %v", err)
		}

		inputData := model.Data{}

		id, err := strconv.Atoi(record[IDIndex])
		if err != nil {
			return nil, fmt.Errorf("error in Atoi: %v", err)
		}
		inputData.ID = id

		inputData.Name = record[nameIndex]

		age, err := strconv.Atoi(record[ageIndex])
		if err != nil {
			return nil, fmt.Errorf("error in Atoi: %v", err)
		}
		inputData.Age = age

		location, _ := time.LoadLocation("UTC")
		loginAt, err := time.ParseInLocation("2006-01-02 15:04:05", record[loginAtIndex], location)
		if err != nil {
			return nil, fmt.Errorf("error in Parse: %v", err)
		}
		inputData.LoginAt = loginAt

		res = append(res, inputData)
	}

	return res, nil
}

func ReadDataAsCh(filename string) (<-chan model.Data, error) {
	res := make(chan model.Data, 1024)

	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error in Open: %v", err)
	}

	r := csv.NewReader(f)

	record, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("error in Read: %v", err)
	}
	header := make(map[string]int)
	for i, v := range record {
		header[v] = i
	}

	IDIndex, ok := header["id"]
	if !ok {
		return nil, errors.New("id column not found")
	}
	nameIndex, ok := header["name"]
	if !ok {
		return nil, errors.New("name column not found")
	}
	ageIndex, ok := header["age"]
	if !ok {
		return nil, errors.New("age column not found")
	}
	loginAtIndex, ok := header["login_at"]
	if !ok {
		return nil, errors.New("login at column not found")
	}

	go func() {
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(fmt.Errorf("error in Read: %v", err))
			}

			inputData := model.Data{}

			id, err := strconv.Atoi(record[IDIndex])
			if err != nil {
				panic(fmt.Errorf("error in Atoi: %v", err))
			}
			inputData.ID = id

			inputData.Name = record[nameIndex]

			age, err := strconv.Atoi(record[ageIndex])
			if err != nil {
				panic(fmt.Errorf("error in Atoi: %v", err))
			}
			inputData.Age = age

			location, _ := time.LoadLocation("UTC")
			loginAt, err := time.ParseInLocation("2006-01-02 15:04:05", record[loginAtIndex], location)
			if err != nil {
				panic(fmt.Errorf("error in Parse: %v", err))
			}
			inputData.LoginAt = loginAt

			res <- inputData
		}
		close(res)
	}()

	return res, nil
}
