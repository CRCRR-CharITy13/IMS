package settings

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
)

func GetTfaStatus(c *gin.Context) {
	user := types.User{}
	if err := database.Database.Model(&types.User{}).Where("id = ?", c.MustGet("userId").(uint)).First(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to find user",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    user.TwoFactorSecret != "",
	})
}

func GenerateTwoFactorSecret(c *gin.Context) {

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "GiftsInKind_Dash",
		AccountName: fmt.Sprintf("User %d", c.MustGet("userId").(uint)),
	})
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to generate secret",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"url":    key.URL(),
			"secret": key.Secret(),
		},
	})
}

func ValidateAndSetupTwoFactor(c *gin.Context) {
	secret := c.Query("secret")
	code := c.Query("code")

	if secret == "" || code == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	if !totp.Validate(code, secret) {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid code",
		})
		return
	}

	// edit user
	user := types.User{}
	if err := database.Database.Model(&types.User{}).Where("id = ?", c.MustGet("userId").(uint)).First(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to find user",
		})
		return
	}

	user.TwoFactorSecret = secret

	if err := database.Database.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to save user",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "2fa setup. You're awesome!",
	})

}
