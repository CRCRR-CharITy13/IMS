package admin

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteSignupCodes(c *gin.Context) {
	codeID := c.Query("user_id")

	if codeID == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user ID",
		})
		return
	}

	// convert to int
	codeIdInt, err := strconv.Atoi(codeID)
	if err != nil || codeIdInt < 1 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user ID",
		})
		return
	}

	// try to find user
	code := types.SignupCode{}
	if err := database.Database.Where("id = ?", codeIdInt).First(&types.SignupCode{}).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "No signup code found for the given user ID",
		})
		return
	}

	database.Database.Where("username = ?", code.DesignatedUsername).Delete(&types.User{})

	if err := database.Database.Where("id = ?", codeIdInt).Delete(&types.SignupCode{}).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Error in deleting user",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User deleted",
	})
}
