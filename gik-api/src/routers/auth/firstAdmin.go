package auth

import (
	"GIK_Web/database"
	"GIK_Web/type_news"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

func CreateFirstAdmin(c *gin.Context) {

	// verify that there is no other user
	var count int64
	database.Database.Model(&type_news.User{}).Count(&count)
	if count > 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "There is already a user",
		})
		return
	}

	newPassword := c.Query("password")

	if newPassword == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid password",
		})
		return
	}

	// create new admin user
	hash, err := argon2id.CreateHash(newPassword, argon2id.DefaultParams)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to hash password",
		})
		return
	}

	newUser := type_news.User{
		Username: "admin",
		Password: hash,
		//TwoFactorSecret: "",
		RegisteredAt: time.Now().Unix(),
		Admin:        true,
		Disabled:     false,
	}

	if err := database.Database.Create(&newUser).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to create user",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User created",
	})
}
