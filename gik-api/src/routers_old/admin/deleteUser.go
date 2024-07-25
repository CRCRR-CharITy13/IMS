package admin

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	userId := c.Query("user_id")

	if userId == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user id",
		})
		return
	}

	// convert to int
	userIdInt, err := strconv.Atoi(userId)
	if err != nil || userIdInt < 1 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user id",
		})
		return
	}

	// try to find user
	user := types.User{}
	if err := database.Database.Where("id = ?", userIdInt).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user id",
		})
		return
	}
	/*
		user.Disabled = true

		if err := database.Database.Save(&user).Error; err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Error deleting user",
			})
			return
		}*/

	if err := database.Database.Where("designated_username = ?", user.Username).Delete(&types.SignupCode{}).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Error deleting user",
		})
		return
	}

	if err := database.Database.Where("ID = ?", userId).Delete(&types.User{}).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Error deleting user",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User Deleted",
	})
}
