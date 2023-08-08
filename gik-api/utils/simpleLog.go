package utils

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateSimpleLog(c *gin.Context, action string) {
	newSimpleLog := types.SimpleLog{
		UserID:    c.MustGet("userId").(uint),
		Action:    action,
		Timestamp: time.Now().Unix(),
		IPAddress: c.ClientIP(),
	}

	database.Database.Create(&newSimpleLog)
}
