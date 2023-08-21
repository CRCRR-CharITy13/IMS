package main

import (
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
}

type Location struct {
	gorm.Model
	Name        string
	Description string
}

type Warehouse struct {
	gorm.Model
	ItemId     string
	LocationId string
	Stock      int
}

type User struct {
	gorm.Model
	Username    string
	Password    string
	RegisterdAt int64
	Admin       bool
	Disabled    bool
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
	DonorBy    int
	SignedBy   int
	TotalValue float32
	TotalItem  int
}

type DonationItem struct {
	gorm.Model
	DonationId int
	ItemId     int
	count      int
}

type Order struct {
	gorm.Model
	ClientId   int
	SignedBy   int
	TotalCost  float32
	TotalItems int
}

type OrderItem struct {
	gorm.Model
	ItemId int
	count  int
}

func main() {
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
