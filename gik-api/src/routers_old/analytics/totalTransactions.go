package analytics

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"time"

	"github.com/gin-gonic/gin"
)

func GraphTotalTransactions(c *gin.Context) {
	var times []int64

	var samplePoints int = 7

	currentTime := time.Now()

	times = append(times, currentTime.Unix())

	for i := 0; i < samplePoints; i++ {
		currentTime = currentTime.AddDate(0, 0, -1)
		times = append(times, currentTime.Unix())
	}

	var exports []int

	for i := 1; i < samplePoints+1; i++ {

		var week []transaction

		database.Database.Model(&types.Order{}).Where("Timestamp BETWEEN ? AND ?", times[i], times[i-1]).Where("type = ?", false).Scan(&week)
		var quantity int

		for _, k := range week {
			quantity += k.TotalQuantity
		}

		exports = append(exports, quantity)
	}

	var imports []int

	for i := 1; i < samplePoints+1; i++ {

		var week []transaction

		database.Database.Model(&types.Order{}).Where("Timestamp BETWEEN ? AND ?", times[i], times[i-1]).Where("type = ?", true).Scan(&week)
		var quantity int

		for _, k := range week {
			quantity += k.TotalQuantity
		}

		imports = append(imports, quantity)
	}

	var data []int

	for i := 0; i < samplePoints; i++ {
		data = append(data, imports[i]-exports[i])
	}

	// reverse data
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})

}
