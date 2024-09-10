package auth

import (
	"GIK_Web/database"
	"log"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

func CreateFirstAdmin(c *gin.Context) {
	// verify that there is no other user
	var count int64
	err := database.Database.QueryRow("SELECT COUNT(*) FROM users WHERE is_admin == 1").Scan(&count)
	if err != nil {
		errMsg := "Cannot query the database"
		log.Println(errMsg)
		c.JSON(400, gin.H{
			"success": false,
			"message": errMsg,
		})
		return
	}
	log.Println(count)
	if count > 0 {
		errMsg := "There is already admin in the database"
		log.Println(errMsg)
		c.JSON(400, gin.H{
			"success": false,
			"message": errMsg,
		})
		return
	}

	// check if password is valid
	password := c.Query("password")
	if password == "" {
		errMsg := "Invalid passowrd"
		log.Println(errMsg)
		c.JSON(400, gin.H{
			"success": false,
			"message": errMsg,
		})
		return
	}

	// hash password
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		errMsg := "Unable to hash password"
		log.Println(errMsg)
		c.JSON(500, gin.H{
			"success": false,
			"message": errMsg,
		})
		return
	}

	// insert to the database
	log.Println("Insert the first admin to the users table")
	registeredTime := time.Now().Format(time.RFC3339)
	createdTime := time.Now().Format(time.RFC3339)
	updatedTime := time.Now().Format(time.RFC3339)
	deletedTime := time.Time{}.Format(time.RFC3339)
	_, err = database.Database.Exec("INSERT INTO users (user_id, username, password, registered_at, is_admin, is_disabled,created_at,updated_at,deleted_at) VALUES (?, ?, ?, ?, ?, ?,?, ?, ?)", 1, "admin", hash, registeredTime, true, false, createdTime, updatedTime, deletedTime)
	//
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": "Failed to create the first admin"})
		log.Printf("Error creating the first admin: %v", err)
		return
	}
	//
	c.JSON(200, gin.H{
		"success": true,
		"message": "The first admin has been successfully created",
	})
}
