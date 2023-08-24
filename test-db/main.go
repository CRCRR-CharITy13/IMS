package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Item has and belongs to many locations, `item_locations` is the join table
type Item struct {
	gorm.Model
	SKU        string
	Name       string
	StockTotal int
	Price      float32
	Size       string
	Warehouses []Warehouse `gorm:"many2many:warehouses;"`
}

type Location struct {
	gorm.Model
	Name        string
	Description string
	Warehouses  []Warehouse `gorm:"many2many:warehouses;"`
}

type Warehouse struct {
	gorm.Model
	ItemID     uint
	LocationID uint
	Stock      int
}

type User struct {
	gorm.Model
	Username     string
	Password     string
	RegisteredAt int64
	Admin        bool
	Disabled     bool
}
type Client struct {
	gorm.Model
	OrgName     string
	Address     string
	PhoneNumber string
	Email       string
	Contact     string
	Balance     float32
}

type Donor struct {
	gorm.Model
	Name        string
	PhoneNumber string
	Address     string
	Email       string
}

type Donation struct {
	gorm.Model
	DonorBy    uint `gorm:"foreignKey:DonorID;references:ID"`
	SignedBy   uint `gorm:"foreignKey:UserID;references:ID"`
	TotalValue float32
	TotalItems []DonationItem
}

type DonationItem struct {
	gorm.Model
	DonationId uint
	ItemID     uint
	count      uint
}

type Order struct {
	gorm.Model
	ClientID   uint `gorm:"foreignKey:ClientID;references:ID"`
	SignedBy   uint `gorm:"foreignKey:UserID;references:ID"`
	TotalCost  float32
	TotalItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID uint
	ItemID  uint
	count   int
}

type Session struct {
	gorm.Model
	ID        string `json:"id"`
	UserID    uint   `json:"userId"`
	User      User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt int64  `json:"createdAt"`
	ExpiresAt int64  `json:"expiresAt"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("gik-ims-localdb.sqlite"), &gorm.Config{})
	db.AutoMigrate(&Item{}, &Location{}, &Warehouse{})
	db.AutoMigrate(&Donor{}, &Donation{}, &DonationItem{})
	db.AutoMigrate(&User{}, &Client{}, &Order{}, &OrderItem{})
	db.AutoMigrate(&Session{})
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// fmt.Printf("==== To Test The GIK-IMS Database\n")
	// itemFileName := "data/gik-ims-items.csv"
	// locationFileName := "data/gik-ims-locations.csv"
	// itemLocationFileName := "data/gik-ims-items-locations.csv"

	// data_items, err := data_utils.ReadItem(itemFileName)
	// data_locations, err := data_utils.ReadLocation(locationFileName)
	// data_itemLocations, err := data_utils.ReadItemLocation(itemLocationFileName)

	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	// fmt.Println("List of items:\n")

	// fmt.Println("Start to connect to the database")
	// // connect to the database
	// db, err := gorm.Open(sqlite.Open("gik-ims-testdb.sqlite"), &gorm.Config{})
	// if err != nil {
	// 	panic("Failed to connect to database")
	// }
	// db.AutoMigrate(&Item{}, &Location{}, &ItemLocation{})

	// for _, item := range data_items {
	// 	fmt.Printf("%s \t %s \t %d \t %0.2f \t %s \n", item.SKU, item.Name, item.Stock, item.Price, item.Size)
	// 	item_record := Item{SKU: item.SKU, Name: item.Name, Stock: item.Stock, Price: item.Price, Size: item.Size}
	// 	db.Create(&item_record)
	// }

	// fmt.Println("List of locations:\n")
	// for _, location := range data_locations {
	// 	fmt.Printf("%s \t %s\n", location.Name, location.Description)
	// 	location_record := Location{Name: location.Name, Description: location.Description}
	// 	db.Create(&location_record)
	// }
	// fmt.Println("List of item - location :\n")
	// for _, itemLocation := range data_itemLocations {
	// 	fmt.Printf("%s : \t %s: \t %d\n", itemLocation.SKU, itemLocation.Location, itemLocation.Stock)
	// 	itemLocation_item := ItemLocation{ItemSKU: itemLocation.SKU, LocationName: itemLocation.Location, Stock: itemLocation.Stock}
	// 	db.Create(&itemLocation_item)
	// }
	//

}
