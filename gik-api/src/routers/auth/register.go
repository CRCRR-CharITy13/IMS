package auth

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

type registerJSON struct {
	Password     string `json:"password" binding:"required"`
	PasswordConf string `json:"passwordConf" binding:"required"`
	SignupCode   string `json:"signupCode" binding:"required"`
	Eula         bool   `json:"eula" binding:"required"`
}

func Register(c *gin.Context) {
	var json registerJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Missing values",
		})
		return
	}

	// check signup code
	signupCode := types.SignupCode{}
	if err := database.Database.Where(types.SignupCode{
		Code:    json.SignupCode,
		Expired: false,
	}).First(&signupCode).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid sign up code",
		})
		return
	}

	// hash password
	// hash password with argon2id

	// !! - Remember to change params in final version
	hash, err := argon2id.CreateHash(json.Password, argon2id.DefaultParams)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to hash password",
		})
		return
	}

	// create new user
	newUser := types.User{
		Username:     signupCode.DesignatedUsername,
		Password:     hash,
		RegisteredAt: time.Now().Unix(),
	}

	// save the user
	if err := database.Database.Create(&newUser).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to create user",
		})
		return
	}

	// update the sign up code details
	//signupCode.CreatedByUserID = newUser.ID
	signupCode.Expired = true

	// save the sign up code
	if err := database.Database.Save(&signupCode).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to update sign up code",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User created",
		"data":    newUser.ID,
	})

}
