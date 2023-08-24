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

type item struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	SKU      string  `json:"sku"`
	Category string  `json:"category"` // Clothes or not
	Price    float32 `json:"price"`
	Quantity int     `json:"stock"`
	Size     string  `json:"size"`
}

type returnedItem struct {
	Name     string  `json:"name"`
	SKU      string  `json:"sku"`
	Category string  `json:"category"` // Clothes or not
	Price    float32 `json:"price"`
	Stock    int     `json:"stock"`
	Size     string  `json:"size"`
}

func ListItem(c *gin.Context) {
	page := c.Query("page")
	name := c.Query("name")
	sku := c.Query("sku")
	//tags := strings.Split(c.Query("tags"), ",")

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
	//offset := (pageInt - 1) * limit

	baseQuery := database.Database.Model(&type_news.Item{})

	baseQuery = baseQuery.Order("sku, FIELD(size, 'XXL',  'XL', 'L', 'M', 'S', 'XS', 'XXS'), size")

	// for _, tag := range tags {
	// 	baseQuery = baseQuery.Where("category LIKE ?", "%"+tag+"%")
	// }
	if name != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+name+"%")
	}
	if sku != "" {
		baseQuery = baseQuery.Where("sku LIKE ?", "%"+sku+"%")
	}

	var totalCount int64
	baseQuery.Count(&totalCount)

	//baseQuery = baseQuery.Limit(limit).Offset(offset)

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

func UpdateItem(c *gin.Context) {
	json := item{}
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

	item := types.Item{}
	if err := database.Database.Model(&types.Item{}).Where("ID = ?", jsonIdInt).First(&item).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid username or password",
		})
		return
	}

	item.Name = json.Name
	item.SKU = json.SKU
	item.Category = json.Category
	item.Size = json.Size
	item.Price = json.Price
	item.Quantity = json.Quantity

	database.Database.Save(item)
}

type newItemRequest struct {
	Name     string  `json:"name" binding:"required"`
	SKU      string  `json:"sku" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Size     string  `json:"size" binding:"required"`
	Price    float32 `json:"price" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
}

func AddItem(c *gin.Context) {
	json := newItemRequest{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	item := types.Item{}

	item.Name = json.Name
	item.SKU = json.SKU
	item.Category = json.Category
	item.Size = json.Size
	item.Price = json.Price
	item.Quantity = int(json.Quantity)

	err := database.Database.Create(&item).Error
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
