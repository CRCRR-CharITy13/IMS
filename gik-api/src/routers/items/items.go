package items

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"GIK_Web/utils"
	"fmt"
	"math"
	"strconv"

	"GIK_Web/type_news"

	"github.com/gin-gonic/gin"
)

// ===== Structs and Methods of GIK version 2.0
type addNewItemRequest struct {
	SKU        string  `json:"sku" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Size       string  `json:"size" binding:"required"`
	Price      float32 `json:"price" binding:"required"`
	StockTotal int     `json:"stock_total" binding:"required"`
}

func AddItem(c *gin.Context) {
	json := addNewItemRequest{}

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("error in AddItem: ", err)
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	item := type_news.Item{}

	if json.Name != "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}
	item.Name = json.Name

	if json.SKU != "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}
	item.SKU = json.SKU

	if json.Size != "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return

	}

	item.Size = json.Size

	if json.Price < 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}
	item.Price = json.Price

	if json.StockTotal < 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
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
		"message": "new item added to the database",
	})
	utils.CreateSimpleLog(c, fmt.Sprintf("Added item %s", item.Name))
}

// 2. List items

func ListItem(c *gin.Context) {
	page := c.Query("page")
	name := c.Query("name")
	sku := c.Query("sku")

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

	baseQuery := database.Database.Model(&type_news.Item{})

	if name != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+name+"%")
	}
	if sku != "" {
		baseQuery = baseQuery.Where("sku LIKE ?", "%"+sku+"%")
	}

	var totalCount int64
	baseQuery.Count(&totalCount)

	baseQuery = baseQuery.Limit(limit).Offset(offset)

	items := []type_news.Item{}

	baseQuery.Find(&items)

	totalPages := math.Ceil(float64(totalCount) / float64(limit))

	c.JSON(200, gin.H{"success": true, "data": gin.H{
		"data":        items,
		"total":       totalCount,
		"currentPage": pageInt,
		"totalPages":  totalPages,
	}})

}

//
// 3. Update Items

type updateItemRequest struct {
	ID         string  `json:"id"`
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
			"message": "Invalid fields",
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

	item := type_news.Item{}
	if err := database.Database.Model(&types.Item{}).Where("ID = ?", jsonIdInt).First(&item).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid item ID",
		})
		return
	}

	///////////
	if json.Name != "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}
	item.Name = json.Name

	if json.SKU != "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}
	item.SKU = json.SKU

	if json.Size != "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return

	}

	item.Size = json.Size

	if json.Price < 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}
	item.Price = json.Price

	if json.StockTotal < 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
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

	utils.CreateSimpleLog(c, fmt.Sprintf("Updated item with id: %d, name: %s", item.ID, item.Name))
}

//

// 4. Delete items by id
func DeleteItem(c *gin.Context) {
	id := c.Query("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	item := types.Item{}
	if err := database.Database.Model(&types.Item{}).Where("id = ?", ID).First(&item).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid Item",
		})
		return
	}

	if err := database.Database.Model(&types.Item{}).Delete(&item).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to delete Item",
			"error":   err.Error(),
		})
		return
	}

	utils.CreateSimpleLog(c, "Deleted Item")

	c.JSON(200, gin.H{
		"success": true,
		"message": "Item successfully deleted.",
	})
}

//
// 5. List locations for an item, by id

type ListLocationForItemResponse struct {
	LocationName string `json:"location-name" binding: "required"`
	Stock        int    `json: "stock" binding : "required"`
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

	var item type_news.Item
	database.Database.Preload("Warehouses").Where("item_id = ?", idInt).Find(&item.Warehouses)
	// fmt.Print(location)
	locationsForItem := make([]ListLocationForItemResponse, len(item.Warehouses))
	for i, warehouse := range item.Warehouses {
		var location type_news.Location
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

//
//

type returnedItem struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
	//Category string  `json:"category"` // Clothes or not
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
	Size  string  `json:"size"`
}

type newItemRequest struct {
	Name     string  `json:"name" binding:"required"`
	SKU      string  `json:"sku" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Size     string  `json:"size" binding:"required"`
	Price    float32 `json:"price" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
}

func AddSize(c *gin.Context) {
	id := c.Query("id")
	size := c.Query("size")
	quantity := c.Query("quantity")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	data := newItemRequest{}
	if err := database.Database.Model(&types.Item{}).Where("id = ?", ID).First(&data).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid Item",
		})
		return
	}

	var count int64

	database.Database.Model(&types.Item{}).Where("id = ?", ID).Where("size = ?", size).Count(&count)

	if count != 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Size already exists",
		})
		return
	}

	data.Size = size
	data.Quantity, err = strconv.Atoi(quantity)

	item := types.Item{}

	item.Name = data.Name
	item.SKU = data.SKU
	item.Category = data.Category
	item.Size = data.Size
	item.Price = data.Price
	item.Quantity = int(data.Quantity)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	err = database.Database.Model(&types.Item{}).Create(&item).Error
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to create item",
			"error":   err.Error(),
		})
		return
	}

	utils.CreateSimpleLog(c, fmt.Sprintf("Added item %s", item.Name))
}

// Function to edit a single item in terms of name, SKU, size, stock and/or price
func EditItem(c *gin.Context) {
	// Get the id of the item
	id := c.Query("id")
	ID, err := strconv.Atoi(id)
	if (err != nil) || (ID < 0) {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
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
				"message": "Invalid fields",
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
				"message": "Invalid fields",
			})
			return
		}

		pricef32 := float32(Price)

		baseQuery.Update("price", pricef32)
	}
	// Update the item
	// database.Database.Model(&types.Item{}).Where("id = ?", ID).Updates(types.Item{Price: pricef32, Name: name, SKU: sku, Size: size, Quantity: Stock})

	// items := []types.Item{}
	// database.Database.Model(&types.Item{}).Select(&types.Item{}).
	// 	Where("id < ?", 0, "stock < ?", 0, "price < ?", 0, "size < ?",
	// 		strconv.Itoa(0), "SKU < ?", strconv.Itoa(0)).Find(&items)

	// if items != nil {
	// 	c.JSON(400, gin.H{
	// 		"success": false,
	// 		"message": "Invalid fields",
	// 	})
	// 	return
	// }

	// Return that you have successfully updated the item
	c.JSON(200, gin.H{
		"success": true,
		"message": "Item successfully edited.",
	})
}
