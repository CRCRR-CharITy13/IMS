package main

import (
	"fmt"

	"github.com/go-gorm/gorm"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

type Item struct {
	ID    uint `gorm: " primaryKey"`
	Name  string
	Stock int
}

func main() {
	// connect to the database
	db, err := gorm.Open(sqlite.Open(gik-ims-db), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// retrive items from the Items table
	var items []Item
	db.Find(&items)

	//display the list of items
	for _, item := range items {
		fmt.Printf("ID: %d, Name: %s, Stock: %d\n", item.ID, item.Name, item.Stock)
	}

}
