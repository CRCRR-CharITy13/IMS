package items

import (
	"GIK_Web/database"
	"GIK_Web/types"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

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

func ImportItems(c *gin.Context) {
	fileParent, err := c.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid file",
		})
		return
	}

	file, err := fileParent.Open()

	defer file.Close()

	items := []newItemRequest{}

	gocsv.Unmarshal(file, &items)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid file",
		})
		return
	}

	for _, item := range items {

		var count int64

		database.Database.Model(&types.Item{}).Where(types.Item{SKU: item.SKU, Size: item.Size}).Count(&count)

		if count == 0 {
			data := types.Item{}

			data.Name = item.Name
			data.SKU = item.SKU
			data.Size = item.Size
			data.Price = item.Price
			data.StockTotal = item.Quantity

			err := database.Database.Model(&types.Item{}).Create(&data).Error
			if err != nil {
				c.JSON(500, gin.H{
					"success": false,
					"message": "Unable to create item",
					"error":   err.Error(),
				})
				return
			}
		}
	}

	c.JSON(400, gin.H{
		"success": true,
		"message": "Items Added",
	})
}
