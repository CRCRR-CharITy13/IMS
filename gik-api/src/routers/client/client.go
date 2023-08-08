package client

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"GIK_Web/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type company struct {
	Name    string `json:"name" binding:"required"`
	Contact string `json:"contact" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Address string `json:"address" binding:"required"`
	Balance int    `json:"balance" binding:"required"`
}

func ListClient(c *gin.Context) {

	client := []types.Client{}

	baseQuery := database.Database.Model(&client)

	name := c.Query("name")
	if name != "" {
		baseQuery = baseQuery.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}

	contact := c.Query("contact")
	if contact != "" {
		baseQuery = baseQuery.Where("contact LIKE ?", fmt.Sprintf("%%%s%%", contact))
	}

	phone := c.Query("phone")
	if phone != "" {
		baseQuery = baseQuery.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", phone))
	}

	email := c.Query("email")
	if email != "" {
		baseQuery = baseQuery.Where("email LIKE ?", fmt.Sprintf("%%%s%%", email))
	}

	address := c.Query("address")
	if address != "" {
		baseQuery = baseQuery.Where("address LIKE ?", fmt.Sprintf("%%%s%%", address))
	}

	err := baseQuery.Find(&client).Error
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to query clients",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Client data retrieved",
		"data":    client,
	})

}

func UpdateClient(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	json := company{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	client := types.Client{}
	if err := database.Database.Where("id = ?", idInt).First(&client).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid donor",
		})
		return
	}

	client.Name = json.Name
	client.Contact = json.Contact
	client.Phone = json.Phone
	client.Email = json.Email
	client.Address = json.Address
	client.Balance = float32(json.Balance)

	if err := database.Database.Save(client).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to update donor",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    client,
	})

	utils.CreateSimpleLog(c, "Updated client")

}

func AddClient(c *gin.Context) {
	json := company{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	newClient := types.Client{
		Name:    json.Name,
		Contact: json.Contact,
		Phone:   json.Phone,
		Email:   json.Email,
		Address: json.Address,
		Balance: float32(json.Balance),
	}

	if err := database.Database.Model(&types.Client{}).Create(&newClient).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to create donor",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Client successfully created.",
		"data":    json,
	})

	utils.CreateSimpleLog(c, "Created client")

}

/*
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Contact string  `json:"contact"`
	Phone   string  `json:"phone"`
	Email   string  `json:"email"`
	Address string  `json:"address"`
	Balance float32 `json:"balance"`
*/

func DeleteClient(c *gin.Context) {

	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	client := types.Client{}
	if err := database.Database.Model(&types.Client{}).Where("id = ?", idInt).First(&client).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid donor",
		})
		return
	}

	if err := database.Database.Model(&types.Client{}).Delete(&client).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to delete donor",
			"error":   err.Error(),
		})
		return
	}

	utils.CreateSimpleLog(c, "Deleted client")

	c.JSON(200, gin.H{
		"success": true,
		"message": "Client successfully deleted.",
	})
}
