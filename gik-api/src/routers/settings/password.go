package settings

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"GIK_Web/utils"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

type changeRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

func ChangePassword(c *gin.Context) {
	json := changeRequest{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Missing values",
		})
		return
	}

	userId := c.MustGet("userId").(uint)

	// check if old password is correct
	user := types.User{}
	if err := database.Database.Model(&types.User{}).Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Error finding user",
		})
		return
	}

	// confirm password
	valid, err := utils.ComparePassword(json.OldPassword, user.Password)
	if !valid || err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid old password",
		})
		return
	}

	// hash new password
	hashedPassword, err := argon2id.CreateHash(json.NewPassword, argon2id.DefaultParams)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to hash new password",
		})
		return
	}

	// update password
	user.Password = hashedPassword
	if err := database.Database.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error updating password",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Password updated",
	})

}
