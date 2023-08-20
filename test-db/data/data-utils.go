package data_utils

import (
	"encoding/csv"
	"os"
)

type Location struct {
	Name        string
	Description string
}

func ReadCSV(filename string) ([]Location, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var locations []Location

	for _, record := range records[1:] {
		location := Location{
			Name:        record[0],
			Description: record[1],
		}
		locations = append(locations, location)
	}
	return locations, err
}
