package main

import (
	"fmt"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

type Item struct {
	ID    uint `gorm: " primaryKey"`
	Name  string
	Stock int
	Size  string
	Price float32
	SKU   string
}

func main() {
	fmt.Println("Start to connect to the database")
	// connect to the database
	db, err := gorm.Open(sqlite.Open("gik-ims-db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	//
	db.AutoMigrate(&Item{})
	// retrive items from the Items table
	var item Item
	db.First(&item, 1)
	// update this item
	db.Model(&item).Update("Size", "XL")
	//
	// add an item
	newItem := Item{Name: "Women Shprt T-Shirt", SKU: "W2234WST", Size: "L", Stock: 500, Price: 5.0}
	result := db.Create(&newItem)
	if result.Error != nil {
		panic("Failed to insert new item to the database")
	}
	//
	var items []Item
	db.Find(&items)

	//display the list of items
	for _, item := range items {
		fmt.Printf("ID: %d, Name: %s, SKU: %s, Size: %s, Stock: %d, Price: %f\n", item.ID, item.Name, item.SKU, item.Size, item.Stock, item.Price)
	}

	//

}
