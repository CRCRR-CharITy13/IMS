package items

import (
	"GIK_Web/database"
	"GIK_Web/utils"
	"fmt"
	"math"
	"strconv"

	"GIK_Web/types"

	"github.com/gin-gonic/gin"
)

// ===== Structs and Methods of GIK version 2.0
type addNewItemRequest struct {
	SKU        string  `json:"sku" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Size       string  `json:"size" binding:"required"`
	Price      float32 `json:"price" binding:"required"`
	StockTotal int     `json:"quantity" binding:"required"`
}

func AddItem(c *gin.Context) {
	json := addNewItemRequest{}

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("error in AddItem: ", err)
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Missing values",
		})
		return
	}

	item := types.Item{}

	if json.Name == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Name",
		})
		return
	}
	item.Name = json.Name

	if json.SKU == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - SKU",
		})
		return
	}
	item.SKU = json.SKU

	if json.Size == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Size",
		})
		return

	}

	item.Size = json.Size

	if json.Price < 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Credits",
		})
		return
	}
	item.Price = json.Price

	if json.StockTotal < 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Quantity",
		})
		return
	}
	item.StockTotal = json.StockTotal

	err := database.Database.Create(&item).Error
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to create item",
			"error":   err.Error(),
		})
		return
	}

	// in case of success
	c.JSON(200, gin.H{
		"success": true,
		"message": "Item added to the database",
	})
	utils.CreateSimpleLog(c, fmt.Sprintf("Added item %s", item.Name))
}

// 2. List items

type ListItemResponse struct {
	Id       int     `json:"ID" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Sku      string  `json:"sku" binding:"required"`
	Size     string  `json:"size" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
	Price    float32 `json:"price" binding:"required"`
}

func ListItem(c *gin.Context) {
	page := c.Query("page")
	name := c.Query("name")
	sku := c.Query("sku")

	baseQuery := database.Database.Model(&types.Item{})

	if name != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+name+"%")
	}
	if sku != "" {
		baseQuery = baseQuery.Where("sku LIKE ?", "%"+sku+"%")
	}

	var totalCount int64
	baseQuery.Count(&totalCount)

	limit := 10 // Number of entries shown per page

	pageInt := 0

	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Invalid page number",
			})
			return
		}

		offset := (pageInt - 1) * limit

		baseQuery = baseQuery.Limit(limit).Offset(offset)
	}

	items := []types.Item{}

	baseQuery.Find(&items)

	totalPages := math.Ceil(float64(totalCount) / float64(limit))

	itemResponse := make([]ListItemResponse, len(items))
	for i, item := range items {
		itemResponse[i] = ListItemResponse{
			Id:       int(item.ID),
			Name:     item.Name,
			Sku:      item.SKU,
			Size:     item.Size,
			Quantity: item.StockTotal,
			Price:    item.Price,
		}
	}

	c.JSON(200, gin.H{"success": true, "data": gin.H{
		"data":        itemResponse,
		"total":       totalCount,
		"currentPage": pageInt,
		"totalPages":  totalPages,
	}})

}

//
// 3. Update Items

type updateItemRequest struct {
	ID         string  `json:"id" binding:"required"`
	SKU        string  `json:"sku" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Size       string  `json:"size" binding:"required"`
	Price      float32 `json:"price" binding:"required"`
	StockTotal int     `json:"stock_total" binding:"required"`
}

func UpdateItem(c *gin.Context) {
	json := updateItemRequest{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Missing values",
		})
		return
	}

	jsonIdInt, err := strconv.Atoi(json.ID)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - ID",
		})
		return
	}

	item := types.Item{}
	if err := database.Database.Model(&types.Item{}).Where("ID = ?", jsonIdInt).First(&item).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid item ID",
		})
		return
	}

	///////////
	if json.Name == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Name",
		})
		return
	}
	item.Name = json.Name

	if json.SKU == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - SKU",
		})
		return
	}
	item.SKU = json.SKU

	if json.Size == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Size",
		})
		return

	}

	item.Size = json.Size

	if json.Price < 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Credits",
		})
		return
	}
	item.Price = json.Price

	if json.StockTotal < 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Quantity",
		})
		return
	}
	item.StockTotal = json.StockTotal
	/////////

	database.Database.Save(item)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Item successfully updated",
	})

	utils.CreateSimpleLog(c, fmt.Sprintf("Updated item with id: %d, SKU: %s, and name: %s", item.ID, item.SKU, item.Name))
}

//

// 4. Delete items by id
func DeleteItem(c *gin.Context) {
	id := c.Query("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - ID",
		})
		return
	}

	item := types.Item{}
	if err := database.Database.Model(&types.Item{}).Where("id = ?", ID).First(&item).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "No such item found",
		})
		return
	}

	itemSKU := item.SKU
	itemName := item.Name

	if err := database.Database.Model(&types.Item{}).Delete(&item).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to delete item",
			"error":   err.Error(),
		})
		return
	}

	utils.CreateSimpleLog(c, fmt.Sprintf("Deleted Item, id: %d, SKU: %s, Name: %s", ID, itemSKU, itemName))

	c.JSON(200, gin.H{
		"success": true,
		"message": "Item successfully deleted.",
	})
}

//
// 5. List locations for an item, by id

type ListLocationForItemResponse struct {
	LocationName string `json:"location_name" binding:"required"`
	Stock        int    `json:"stock" binding:"required"`
}

func ListLocationForItem(c *gin.Context) {
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

	var item types.Item
	database.Database.Preload("Warehouses").Where("item_id = ?", idInt).Find(&item.Warehouses)
	// fmt.Print(location)
	locationsForItem := make([]ListLocationForItemResponse, len(item.Warehouses))
	for i, warehouse := range item.Warehouses {
		var location types.Location
		database.Database.First(&location, warehouse.LocationID)
		locationsForItem[i] = ListLocationForItemResponse{
			LocationName: location.Name,
			Stock:        warehouse.Stock,
		}
		fmt.Printf("location: %s : %d\n", location.Name, warehouse.Stock)
	}
	c.JSON(200, gin.H{
		"success": true,
		"data":    locationsForItem,
	},
	)
}

func GetUnstoredQuantity(c *gin.Context) {
	sku := c.Query("sku")

	var item types.Item
	baseQuery := database.Database.Model(&types.Item{})
	baseQuery.Find(&item, "sku=?", sku)
	database.Database.Preload("Warehouses").Where("item_id = ?", item.ID).Find(&item.Warehouses)
	// fmt.Print(location)
	storedQtt := 0
	for _, warehouse := range item.Warehouses {
		storedQtt += warehouse.Stock
	}
	restQtt := item.StockTotal - storedQtt
	c.JSON(200, gin.H{
		"success": true,
		"data":    restQtt,
	},
	)
}

//
//

// Function to edit a single item in terms of name, SKU, size, stock and/or price
func EditItem(c *gin.Context) {
	// Get the id of the item
	id := c.Query("id")
	ID, err := strconv.Atoi(id)
	if (err != nil) || (ID < 0) {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - ID",
		})
		return
	}

	baseQuery := database.Database.Model(&types.Item{}).Where("id = ?", ID)

	// Get any of price (name, SKU, size, stock or price)
	name := c.Query("name")
	if name != "" {
		baseQuery.Update("name", name)
	}

	sku := c.Query("SKU")
	if sku != "" {
		baseQuery.Update("sku", sku)
	}

	size := c.Query("size")
	if size != "" {
		baseQuery.Update("size", size)
	}

	stock := c.Query("stock")
	if stock != "" {
		Stock, err := strconv.Atoi(stock)

		if err != nil || Stock < 0 {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Invalid fields - Quantity",
			})
			return
		}

		baseQuery.Update("stock", Stock)
	}

	price := c.Query("price")
	if price != "" {
		Price, err := strconv.ParseFloat(price, 32)
		if err != nil || Price < 0 {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Invalid fields - Credits",
			})
			return
		}

		pricef32 := float32(Price)

		baseQuery.Update("price", pricef32)
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Item successfully edited.",
	})
}

func AddSize(c *gin.Context) {
	id := c.Query("id")
	ID, err := strconv.Atoi(id)
	if (err != nil) || (ID < 0) {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - ID",
		})
		return
	}

	size := c.Query("size")
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Size",
		})
		return
	}

	sku := c.Query("sku")
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - SKU",
		})
		return
	}

	item := types.Item{}

	item.SKU = sku
	item.Size = size
	item.StockTotal = 0

	// Fetch from old version
	itemOld := types.Item{}
	if err := database.Database.Model(&types.Item{}).Where("id = ?", ID).First(&itemOld).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "No such item found",
		})
		return
	}

	item.Price = itemOld.Price
	item.Name = itemOld.Name

	err = database.Database.Create(&item).Error
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to create item",
			"error":   err.Error(),
		})
		return
	}

	// in case of success
	c.JSON(200, gin.H{
		"success": true,
		"message": "Item added to the database",
	})
	utils.CreateSimpleLog(c, fmt.Sprintf("Added item %s", item.Name))

}
