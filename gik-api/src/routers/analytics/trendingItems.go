package analytics

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTrendingItems(c *gin.Context) {
	// gets the items with the most exports within the last 7 days and sends them back
	// no need to get times, just send the trending items in a slice

	var trendingItems []types.Item
	var times []int64

	var samplePoints int = 7

	currentTime := time.Now()

	times = append(times, currentTime.Unix())

	for i := 0; i < samplePoints; i++ {
		currentTime = currentTime.AddDate(0, 0, -1)
		times = append(times, currentTime.Unix())
	}

	for i := 1; i < samplePoints+1; i++ {
		var week []transaction

		database.Database.Model(&types.Order{}).Where("timestamp BETWEEN ? AND ?", times[i], times[i-1]).Where("type = ?", false).Find(&week)
		var quantity int

		for _, k := range week {
			quantity += k.TotalQuantity
		}

		trendingItems = append(trendingItems, types.Item{})
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Trending items",
		"data":    trendingItems,
	})
}
