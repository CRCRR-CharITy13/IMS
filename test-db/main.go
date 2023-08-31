package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Item has and belongs to many locations, `item_locations` is the join table
type Item struct {
	gorm.Model
	SKU           string
	Name          string
	StockTotal    int
	Price         float32
	Size          string
	Warehouses    []Warehouse `gorm:"many2many:warehouses;"`
	OrderItems    []OrderItem
	DonationItems []DonationItem
}

type Location struct {
	gorm.Model
	Name        string
	Description string
	Warehouses  []Warehouse `gorm:"many2many:warehouses"`
}

type Warehouse struct {
	gorm.Model
	ItemID     uint
	Item       Item `gorm:"foreignKey:ItemID;references:ID"`
	LocationID uint
	Location   Location `gorm:"foreignKey:LocationID;references:ID"`
	Stock      int
}

type User struct {
	gorm.Model
	Username        string
	Password        string
	RegisteredAt    int64
	Admin           bool
	Disabled        bool
	SignedOrders    []Order
	SignedDonations []Donation
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
	DonorID    uint
	DonorBy    Donor `gorm:"foreignKey:DonorID;references:ID"`
	UserID     uint
	SignedBy   User `gorm:"foreignKey:UserID;references:ID"`
	TotalValue float32
	TotalItems []DonationItem
}

type DonationItem struct {
	gorm.Model
	DonationID uint
	Donation   Donation `gorm:"foreignKey:DonationID;references:ID"`
	ItemID     uint
	Item       Item `gorm:"foreignKey:ItemID;references:ID"`
	count      uint
}

type Order struct {
	gorm.Model
	ClientID   uint
	Client     Client `gorm:"foreignKey:ClientID;references:ID"`
	UserID     uint
	SignedBy   User `gorm:"foreignKey:UserID;references:ID"`
	TotalCost  float32
	TotalItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID uint
	Order   Order `gorm:"foreignKey:OrderID;references:ID"`
	ItemID  uint
	Item    Item `gorm:"foreignKey:ItemID;references:ID"`
	count   int
}

type Session struct {
	gorm.Model
	UserID    uint  `json:"id"`
	User      User  `json:"user" gorm:"foreignKey:UserID; references:ID"`
	CreatedAt int64 `json:"createdAt"`
	ExpiresAt int64 `json:"expiresAt"`
}

type ListItemInLocationResponse struct {
	ItemName string `json:"item-name" binding: "required"`
	Stock    int    `json: "stock" binding : "required"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("../gik-api/assets/gik-ims-localdb.sqlite"), &gorm.Config{})

	db.AutoMigrate(&Item{}, &Location{}, &Warehouse{})
	db.AutoMigrate(&Donor{}, &Donation{}, &DonationItem{})
	db.AutoMigrate(&User{}, &Client{}, &Order{}, &OrderItem{})
	db.AutoMigrate(&Session{})
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	locationID := 3
	var location Location

	db.Preload("Warehouses").Where("location_id = ?", locationID).Find(&location.Warehouses)
	// fmt.Print(location)
	itemsInLocation := make([]ListItemInLocationResponse, len(location.Warehouses))
	for i, warehouse := range location.Warehouses {
		var item Item
		db.First(&item, warehouse.ItemID)
		itemsInLocation[i] = ListItemInLocationResponse{
			ItemName: item.Name,
			Stock:    warehouse.Stock,
		}
		fmt.Printf("item id: %s : %d\n", item.Name, warehouse.Stock)
	}
	jsonReturn, err := json.MarshalIndent(itemsInLocation, "", " ")
	if err != nil {
		fmt.Println("Cannot convert the result to json")
	}
	fmt.Println(string(jsonReturn))
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
	//
}
