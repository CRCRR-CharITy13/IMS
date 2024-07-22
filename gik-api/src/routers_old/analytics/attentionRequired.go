package analytics

import (
	"GIK_Web/types"
	"fmt"

	"github.com/gin-gonic/gin"

	"GIK_Web/database"
)

type AttentionItem struct {
	ID      int    `json:"ID" binding:"required"`
	SKU     string `json:"sku" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Size    string `json:"size" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func AttentionRequiredItem(c *gin.Context) {
	// get all the items
	items := []types.Item{}
	baseQuery := database.Database.Model(&types.Item{})
	baseQuery.Find(&items)

	var attentionItems []AttentionItem

	totalAttention := 0
	idx := 0
	for _, item := range items {
		var tmpItem types.Item
		database.Database.Preload("Warehouses").Where("item_id=?", item.ID).Find(&tmpItem.Warehouses)
		storedStock := 0
		for _, warehouse := range tmpItem.Warehouses {
			storedStock += warehouse.Stock
		}
		restStock := item.StockTotal - storedStock
		var msg string
		if restStock != 0 {
			if restStock > 0 {
				msg = fmt.Sprintf("%d pieces have not been stored yet", restStock)
			} else {
				msg = fmt.Sprintf("The total pieces stored in locations (%d) is greater than the recored total stock (%d). Verification need!", storedStock, item.StockTotal)
			}
			idx++
			attentionItems = append(attentionItems, AttentionItem{
				ID:      idx,
				SKU:     item.SKU,
				Name:    item.Name,
				Size:    item.Size,
				Message: msg,
			})
			totalAttention++
		}
	}

	c.JSON(200, gin.H{"success": true,
		"data": gin.H{
			"data":  attentionItems,
			"total": totalAttention,
		}})

}

type AttentionLocation struct {
	ID          int    `json:"ID" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Message     string `json:"message" binding:"required"`
}

func AttentionRequiredLocation(c *gin.Context) {
	// get all the locations
	locations := []types.Location{}
	baseQuery := database.Database.Model(&types.Location{})
	baseQuery.Find(&locations)

	var attentionLocations []AttentionLocation

	totalAttention := 0
	idx := 0
	for _, location := range locations {
		// var tmpItem types.Item
		database.Database.Preload("Warehouses").Where("location_id=?", location.ID).Find(&location.Warehouses)
		if len(location.Warehouses) == 0 {
			idx++
			attentionLocations = append(attentionLocations, AttentionLocation{
				ID:          idx,
				Name:        location.Name,
				Description: location.Description,
				Message:     "This location is currently empty",
			})
			totalAttention++
		}
	}

	c.JSON(200, gin.H{"success": true,
		"data": gin.H{
			"data":  attentionLocations,
			"total": totalAttention,
		}})

}
