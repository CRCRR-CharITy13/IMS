package items

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"fmt"

	"github.com/gin-gonic/gin"
)

type suggestion struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

func GetAutoSuggest(c *gin.Context) {
	query := c.Query("query")

	items := []types.Item{}
	err := database.Database.Distinct("name", "sku").Where("name LIKE ?", fmt.Sprintf("%%%s%%", query)).Find(&items).Error
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to query items",
		})
		return
	}

	suggestions := []suggestion{}
	for _, item := range items {
		suggestions = append(suggestions, suggestion{item.Name, item.SKU})
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Auto suggested items",
		"data":    suggestions,
	})
}
