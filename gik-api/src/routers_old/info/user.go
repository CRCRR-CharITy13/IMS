package info

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsername(c *gin.Context) {
	id := c.Query("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to find user",
			"data":    "",
		})
		return
	}

	user := types.User{}

	if err := database.Database.Model(&types.User{}).Where("id = ?", idInt).Find(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to find user",
			"data":    "",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User found",
		"data":    user.Username,
	})
}

func GetCurrentUsername(c *gin.Context) {
	id := c.MustGet("userId").(uint)

	user := types.User{}

	if err := database.Database.Model(&types.User{}).Where("id = ?", id).Find(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to find user",
			"data":    "",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User found",
		"data":    user.Username,
	})
}
