package auth

import (
	"GIK_Web/database"
	"GIK_Web/types"

	"github.com/gin-gonic/gin"
)

func GetSignupCodeInfo(c *gin.Context) {
	code := c.Query("code")

	if code == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid sign up code",
		})
		return
	}

	// try to find signup code
	signupCode := types.SignupCode{}
	if err := database.Database.Where(&types.SignupCode{
		Code: code,
	}).First(&signupCode).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid sign up code",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Signup code found",
		"data":    signupCode.DesignatedUsername,
	})

}
