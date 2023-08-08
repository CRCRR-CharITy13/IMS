package auth

import (
	"GIK_Web/database"
	"GIK_Web/env"
	"GIK_Web/types"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckTfaStatusBeforeLogin(c *gin.Context) {
	verification := c.Query("verification")
	username := c.Query("username")

	if verification == "" {
		c.JSON(401, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	// check if JWT is expired or invalid
	token, err := jwt.Parse(verification, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.JWTSigningPassword), nil
	})
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid verification",
		})
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	// check if JWT is valid
	if !token.Valid {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid verification",
		})
		return
	}

	// check if JWT is expired
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Verification is expired",
		})
		return
	}

	// check if username matches up
	username = strings.ToLower(username)

	if claims["username"].(string) != username {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Verification does not match",
		})
		return
	}

	// check if user exists
	user := types.User{}
	if err := database.Database.Model(&types.User{}).Where(&types.User{
		Username: username,
	}).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid username or password",
			"code":    1,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Status retrieved",
		"data":    user.TwoFactorSecret != "",
	})

}
