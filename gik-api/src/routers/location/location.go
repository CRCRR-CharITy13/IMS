package location

import (
	"GIK_Web/database"
	"GIK_Web/type_news"
	"GIK_Web/utils"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// start to implement methods

// 1. Add location
// To add a new, empty location to the list

type addRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func AddLocation(c *gin.Context) {
	json := addRequest{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	newLocation := type_news.Location{
		Name:        json.Name,
		Description: json.Description,
	}

	err := database.Database.Create(&newLocation).Error
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Unable to create new location",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "New location created",
	})

	utils.CreateSimpleLog(c, "Added new location: "+json.Name)

}

type ListLocationResponse struct {
	Id          int    `json:"ID" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	TotalItem   int    `json:"total_item" binding:"required"`
}

// 2. List location: display the list of location (do not include the items within)
func ListLocation(c *gin.Context) {
	name := c.Query("name")
	description := c.Query("description")
	page := c.Query("page")

	if page == "" {
		page = "1"
	}

	pageInt, err := strconv.Atoi(page)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid page number",
		})
		return
	}

	limit := 10 // Number of entries shown per page
	offset := (pageInt - 1) * limit

	baseQuery := database.Database.Model(&type_news.Location{})

	if name != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+name+"%")
	}

	if description != "" {
		baseQuery = baseQuery.Where("description LIKE ?", "%"+description+"%")
	}

	var totalCount int64
	baseQuery.Count(&totalCount)

	baseQuery = baseQuery.Limit(limit).Offset(offset)

	locations := []type_news.Location{}

	err = baseQuery.Find(&locations).Error
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to query locations",
		})
		return
	}

	totalPages := math.Ceil(float64(totalCount) / float64(limit))

	locationResponse := make([]ListLocationResponse, len(locations))

	for i, location := range locations {
		itemCount := 0
		var tmpLocation type_news.Location
		database.Database.Preload("Warehouses").Where("location_id = ?", location.ID).Find(&tmpLocation.Warehouses)
		for _, warehouse := range tmpLocation.Warehouses {
			itemCount += warehouse.Stock
		}
		locationResponse[i] = ListLocationResponse{
			Id:          int(location.ID),
			Name:        location.Name,
			Description: location.Description,
			TotalItem:   itemCount,
		}
	}
	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"data":        locationResponse,
			"total":       totalCount,
			"currentPage": pageInt,
			"totalPages":  totalPages,
		},
	})

}

// 3. Delete location by id

func DeleteLocation(c *gin.Context) {
	id := c.Query("id")

	// conver to integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	location := type_news.Location{}

	if err := database.Database.Model(&type_news.Location{}).Where("id = ?", idInt).First(&location).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid Location",
		})
		return
	}

	locationName := location.Name
	locationDescription := location.Description

	err = database.Database.Model(&location).Where("id = ?", idInt).Delete(&location).Error
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Unable to delete location",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Location deleted",
	})
	utils.CreateSimpleLog(c, fmt.Sprintf("Deleted location, id: %d, Name:  %s, Description: %s", idInt, locationName, locationDescription))
}

type addItemToLocationRequest struct {
	ItemSKU    string `json:"itemSKU" binding:"required"`
	LocationID uint   `json:"locationID" binding:"required"`
	Stock      int    `json:"quantity" binding:"required"`
}

// 4. Add item to location by item's SKU
// Step-1 : look-up the id of item
// Step-2 : check if the record with foreign key pair (item_id, location_id) exists
// + Yes: add Stock to the stock field
// + No: create a new record
func AddItemToLocation(c *gin.Context) {
	json := addItemToLocationRequest{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}
	fmt.Println(json)
	var item type_news.Item

	database.Database.First(&item, "sku=?", json.ItemSKU)
	var warehouse type_news.Warehouse

	// look-up the record with foreign key pair = (item_id, location_id)
	result := database.Database.Where("item_id = ? AND location_id = ?", item.ID, json.LocationID).First(&warehouse)

	if result.Error == gorm.ErrRecordNotFound {
		// create a new record and add to the database
		fmt.Print("Not found, create new and add")
		newWarehouse := type_news.Warehouse{
			ItemID:     item.ID,
			LocationID: json.LocationID,
			Stock:      json.Stock,
		}
		err := database.Database.Create(&newWarehouse).Error

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Unable to add item to location",
			})
			return
		}
	} else {
		// add json.Stock to the current record stock
		fmt.Print("Found, add to the existence")
		warehouse.Stock += json.Stock
		database.Database.Save(warehouse)
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Successfully add an item to location",
	})

	utils.CreateSimpleLog(c, fmt.Sprintf("Added %d pieces of item with sku = %s to location id = %d", json.Stock, json.ItemSKU, json.LocationID))
}

// 5. List items within location
type ListItemInLocationResponse struct {
	ItemSKU  string `json:"item_sku" binding:"required"`
	ItemName string `json:"item_name" binding:"required"`
	Stock    int    `json:"stock" binding:"required"`
}

func ListItemInLocation(c *gin.Context) {
	id := c.Query("id")

	// conver to integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}
	var location type_news.Location
	database.Database.Preload("Warehouses").Where("location_id = ?", idInt).Find(&location.Warehouses)
	// fmt.Print(location)
	itemsInLocation := make([]ListItemInLocationResponse, len(location.Warehouses))
	for i, warehouse := range location.Warehouses {
		var item type_news.Item
		database.Database.First(&item, warehouse.ItemID)
		itemsInLocation[i] = ListItemInLocationResponse{
			ItemSKU:  item.SKU,
			ItemName: item.Name,
			Stock:    warehouse.Stock,
		}
		fmt.Printf("item id: %s : %d\n", item.Name, warehouse.Stock)
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    itemsInLocation,
	})
}

type updateLocationRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func UpdateLocation(c *gin.Context) {
	json := updateLocationRequest{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid json fields",
		})
		return
	}

	jsonIdInt, err := strconv.Atoi(json.ID)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	location := type_news.Location{}
	if err := database.Database.Model(&type_news.Location{}).Where("ID = ?", jsonIdInt).First(&location).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid location ID",
		})
		return
	}

	///////////
	if json.Name == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid name",
		})
		return
	}
	location.Name = json.Name

	if json.Description == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid description",
		})
		return
	}
	location.Description = json.Description

	database.Database.Save(location)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Location successfully updated",
	})

	utils.CreateSimpleLog(c, fmt.Sprintf("Updated location id: %d with name: %s and description: %s", location.ID, location.Name, location.Description))
}
