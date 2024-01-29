package client

import (
	"GIK_Web/database"
	"GIK_Web/types"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func ImportClients(c *gin.Context) {
	fileParent, err := c.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid file",
		})
		return
	}

	file, err := fileParent.Open()

	defer file.Close()

	clients := []company{}

	gocsv.Unmarshal(file, &clients)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid file",
		})
		return
	}

	for _, client := range clients {

		var count int64

		database.Database.Model(&types.Client{}).Where(types.Client{OrgName: client.Name}).Count(&count)

		if count == 0 {

			newClient := types.Client{
				OrgName:     client.Name,
				Contact:     client.Contact,
				PhoneNumber: client.Phone,
				Email:       client.Email,
				Address:     client.Address,
				Balance:     float32(client.Balance),
			}

			err := database.Database.Model(&types.Client{}).Create(&newClient).Error
			if err != nil {
				c.JSON(500, gin.H{
					"success": false,
					"message": "Unable to create item",
					"error":   err.Error(),
				})
				return
			}
		}
	}

	c.JSON(400, gin.H{
		"success": true,
		"message": "Items Added",
	})
}
