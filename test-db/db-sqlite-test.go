package main

import (
	"fmt"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

// Item has and belongs to many locations, `item_locations` is the join table
type Item struct {
	gorm.Model
	// ID        uint `gorm: "primaryKey"`
	Name      string
	Stock     int
	Size      string
	Price     float32
	SKU       string
	Locations []Location `gorm:"many2many:item_locations;"`
}

type Location struct {
	gorm.Model
	// ID          uint `gorm: "primaryKey"`
	Name        string
	Description string
}

// type Client struct {
// 	ID           uint `gorm: "primaryKey"`
// 	Contact_info string
// 	Balance      float32
// }

// type Donor struct {
// 	ID           uint `gorm: "primaryKey"`
// 	Contact_info string
// }

// type Users struct {
// 	ID       uint `gorm: "primaryKey"`
// 	Username string
// 	Password string
// 	Is_admin bool
// }

func main() {
	fmt.Println("Start to connect to the database")
	// connect to the database
	db, err := gorm.Open(sqlite.Open("gik-ims-testdb.sqlite"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(Item{}, Location{})

	// display the list of items
	// var items []Item
	// db.Find(&items)
	// for _, item := range items {
	// 	fmt.Printf("ID: %d, Name: %s, SKU: %s, Size: %s, Stock: %d, Price: %f\n", item.ID, item.Name, item.SKU, item.Size, item.Stock, item.Price)
	// }
	// //display the list of loctions
	// var locations []Location
	// db.Find(&locations)
	// for _, location := range locations {
	// 	fmt.Printf("ID: %d, Name: %s, Description: %s\n", location.ID, location.Name, location.Description)
	// }

	//
	// db.AutoMigrate(&Item{})
	// // retrive items from the Items table
	// var item Item
	// db.First(&item, 1)
	// // update this item
	// db.Model(&item).Update("Size", "XL")
	// //
	// // add an item
	// newItem := Item{Name: "Women Shprt T-Shirt", SKU: "W2234WST", Size: "L", Stock: 500, Price: 5.0}
	// result := db.Create(&newItem)
	// if result.Error != nil {
	// 	panic("Failed to insert new item to the database")
	// }
	// //
	// var items []Item
	// db.Find(&items)

	// //display the list of items
	// for _, item := range items {
	// 	fmt.Printf("ID: %d, Name: %s, SKU: %s, Size: %s, Stock: %d, Price: %f\n", item.ID, item.Name, item.SKU, item.Size, item.Stock, item.Price)
	// }
	//

}
