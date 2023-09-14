package location

import (
	"GIK_Web/database"
	"GIK_Web/type_news"
	"GIK_Web/types"
	"GIK_Web/utils"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type location struct {
// 	ID          int
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	// Letter string `json:"letter"`
// 	// SKU    string `json:"sku"`
// }

// type lookupData struct {
// 	location
// 	Item types.Item `json:"product"`
// }

// type listData struct {
// 	location
// 	Item        types.Item `json:"product"`
// 	ProductName string     `json:"productName"`
// }

// start to implement methods

// 1. Add location
// To add a new, empty location to the list

type addRequest struct {
	Name        string `json:"name" binding: "required"`
	Description string `json:"description" binding: "required`
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
		locationResponse[i] = ListLocationResponse{
			Id:          int(location.ID),
			Name:        location.Name,
			Description: location.Description,
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

	err = database.Database.Model(&location).Where("id = ?", idInt).Delete(&location).Error
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Unable to delete location",
		})
		return
	}

	// err = database.Database.Delete(&location).Error
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"success": false,
	// 		"message": "Unable to delete location",
	// 	})
	// 	return
	// }

	c.JSON(200, gin.H{
		"success": true,
		"message": "Location deleted",
	})
	utils.CreateSimpleLog(c, "Deleted location "+id)
}

type addItemToLocationRequest struct {
	ItemID     uint `json:"item-id" binding: "required"`
	LocationID uint `json:"location-id" binding: "required`
	Stock      int  `json:"stock" binding: "required`
}

// 4. Add item to location
func AddItemToLocation(c *gin.Context) {
	json := addItemToLocationRequest{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	newWarehouse := type_news.Warehouse{
		ItemID:     json.ItemID,
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

	c.JSON(200, gin.H{
		"success": true,
		"message": "Successfully add an item to location",
	})

	utils.CreateSimpleLog(c, "Added item to location")

}

// 5. List items within location
type ListItemInLocationResponse struct {
	ItemName string `json:"item-name" binding: "required"`
	Stock    int    `json: "stock" binding : "required"`
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
			ItemName: item.Name,
			Stock:    warehouse.Stock,
		}
		fmt.Printf("item id: %s : %d\n", item.Name, warehouse.Stock)
	}
	// jsonReturn, err := json.MarshalIndent(itemsInLocation, "", " ")
	// if err != nil {
	// 	fmt.Println("Cannot convert the result to json")
	// }
	//fmt.Println(string(jsonReturn))
	c.JSON(200, gin.H{
		"success": true,
		"data":    itemsInLocation,
	})
}

//

func UpdateLocation(c *gin.Context) {
	json := type_news.Location{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	// database.Database.Model(&location{}).Where("name = ?", json.Name).Update("sku", json.SKU)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Updated location",
	})

	utils.CreateSimpleLog(c, "Updated location "+json.Name)

}

func AddSubLocation(c *gin.Context) {

	name := c.Query("name")

	data := types.Location{}

	database.Database.Model(&types.Location{}).Where(types.Location{Name: name}).First(&data)

	var count int64

	database.Database.Model(&types.Location{}).Where(types.Location{Name: name}).Count(&count)

	var letter string

	LETTERS := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if count == 0 {
		letter = ""
	} else if count == 1 {
		database.Database.Model(&types.Location{}).Where(types.Location{Name: name}).Update("letter", "A")
		letter = "B"
	} else {
		letter = LETTERS[count+1]
	}

	data.Letter = letter

	err := database.Database.Create(&data).Error
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Unable to create location",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Location created",
	})

	utils.CreateSimpleLog(c, "Added location "+name)
}

// func LookupLocation(c *gin.Context) {
// 	// product id
// 	// name := c.Query("name")
// 	// letter := c.Query("letter")

// 	//could remove check and use function also for list

// 	// if name == "" && letter == "" && itemID == 0 {
// 	// 	c.JSON(400, gin.H{
// 	// 		"success": false,
// 	// 		"message": "No fields provided",
// 	// 	})
// 	// 	return
// 	// }

// 	// var postData []location
// 	// database.Database.Model(&location{}).Where(&location{Name: name, Letter: letter}).Scan(&postData)

// 	// response := []lookupData{}

// 	// for _, location := range postData {
// 	// 	var item types.Item
// 	// 	err := database.Database.Model(&types.Item{}).Where("sku = ?", location.SKU).Scan(&item).Error

// 	// 	if err != nil {
// 	// 		continue
// 	// 	}

// 	// 	response = append(response, lookupData{
// 	// 		location: location,
// 	// 		Item:     item,
// 	// 	})

// 	// }

// 	c.JSON(200, gin.H{
// 		"success": true,
// 		"data":    response,
// 	})

// }

func GetScannedData(c *gin.Context) {
	name := c.Query("name")
	letter := c.Query("letter")

	var product types.Item

	location := types.Location{}
	if err := database.Database.Model(&types.Location{}).Where(&types.Location{Name: name, Letter: letter}).Scan(&location).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Unable to query location",
			"data":    gin.H{},
		})
		return
	}

	var item types.Item
	if err := database.Database.Model(&types.Item{}).Where("sku = ?", location.SKU).First(&item).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Unable to find item",
			"error":   err.Error(),
			"data":    gin.H{},
		})
		return
	}

	product = item

	c.JSON(200, gin.H{
		"success": true,
		"data":    product,
	})

	utils.CreateSimpleLog(c, "Scanned location "+name+" "+letter)

}

func ListLocationSKU(c *gin.Context) {

	name := c.Query("name")
	letter := c.Query("letter")

	var sku string

	database.Database.Model(&types.Location{}).Where("name = ?", name).Where("letter = ?", letter).Distinct("sku").Pluck("sku", &sku)

	c.JSON(200, gin.H{"success": true, "sku": sku})

}
