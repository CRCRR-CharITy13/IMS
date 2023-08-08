/*
toggleUser allows one to toggle the disabled value of a user
*/

package admin

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToggleUser(c *gin.Context) {
	// Get userId
	userId := c.Query("user_id")

	// Check if userID is not blank
	if userId == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user id",
		})
		return
	}

	// Convert userID to integer and check it is >= 1
	userIdInt, err := strconv.Atoi(userId)
	if err != nil || userIdInt < 1 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user id",
		})
		return
	}

	// Try to find the user
	user := types.User{}
	if err := database.Database.Where("id = ?", userIdInt).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user id",
		})
		return
	}

	// Flip the disabled value
	user.Disabled = !user.Disabled

	// Save the updated information
	if err := database.Database.Save(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Error saving user",
		})
		return
	}

	// Return the new value
	c.JSON(200, gin.H{
		"success": true,
		"message": "User toggled",
		"data":    user.Disabled,
	})

}
