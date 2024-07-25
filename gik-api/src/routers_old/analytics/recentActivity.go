package analytics

import (
	"GIK_Web/database"
	"GIK_Web/types"

	"github.com/gin-gonic/gin"
)

func GetRecentActivity(c *gin.Context) {
	// get the last 5 simple logs

	simpleLogs := []types.SimpleLog{}

	if err := database.Database.Model(&types.SimpleLog{}).Order("id desc").Limit(5).Find(&simpleLogs).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to find simple logs",
			"data":    []string{},
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Recent activity",
		"data":    simpleLogs,
	})
}
