package info

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetClientInfo(c *gin.Context) {
	// id
	id := c.Query("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Invalid id",
		})
		return
	}

	// get the client
	var client types.Client
	err = database.Database.Where("id = ?", idInt).First(&client).Error
	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Client not found",
		})
		return
	}

	// get the client's name

	c.JSON(200, gin.H{
		"success": true,
		"message": "Got username",
		"data":    client.OrgName,
	})

}
