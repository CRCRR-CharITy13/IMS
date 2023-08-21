package data_utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Location struct {
	Name        string
	Description string
}

type Item struct {
	SKU   string
	Name  string
	Stock int
	Price float32
	Size  string
}

type ItemLocation struct {
	SKU      string
	Location string
	Stock    int
}

func ReadItem(filename string) ([]Item, error) {
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

	var items []Item

	for _, record := range records[1:] {

		stock, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Println("Error: ", err)
			return nil, err
		}

		price, err := strconv.ParseFloat(record[3], 32)
		if err != nil {
			fmt.Println("Error: ", err)
			return nil, err
		}

		price32 := float32(price)

		item := Item{
			SKU:   record[0],
			Name:  record[1],
			Stock: stock,
			Price: price32,
			Size:  record[4],
		}
		items = append(items, item)
	}
	return items, err
}

func ReadLocation(filename string) ([]Location, error) {
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

func ReadItemLocation(filename string) ([]ItemLocation, error) {
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

	var itemLocations []ItemLocation

	for _, record := range records[1:] {
		stock, err := strconv.Atoi(record[2])

		if err != nil {
			fmt.Println("Error: ", err)
			return nil, err
		}

		itemLocation := ItemLocation{
			SKU:      record[0],
			Location: record[1],
			Stock:    stock,
		}
		itemLocations = append(itemLocations, itemLocation)
	}
	return itemLocations, err
}
