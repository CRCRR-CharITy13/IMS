package analytics

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type transaction struct {
	TotalQuantity int
}

func GraphTransactions(c *gin.Context) {
	var times []int64

	var labels []string

	typeS, _ := strconv.ParseBool(c.Query("type"))

	var samplePoints int = 7

	currentTime := time.Now()

	times = append(times, currentTime.Unix())
	labels = append(labels, currentTime.Format("2006-02-01"))

	for i := 0; i < samplePoints; i++ {
		currentTime = currentTime.AddDate(0, 0, -1)
		times = append(times, currentTime.Unix())
		labels = append(labels, currentTime.Format("2006-02-01"))
	}

	var data []int

	for i := 1; i < samplePoints+1; i++ {

		var week []transaction

		database.Database.Model(&types.Transaction{}).Where("Timestamp BETWEEN ? AND ?", times[i], times[i-1]).Where("type = ?", typeS).Scan(&week)
		var quantity int

		for _, k := range week {
			quantity += k.TotalQuantity
		}

		data = append(data, quantity)
	}

	// reverse data
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
		"labels":  labels[:6], // 1: to remove first date, :6 to remove last
	})

}
