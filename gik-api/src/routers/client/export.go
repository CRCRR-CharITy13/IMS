package client

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
)

type export struct {
	ID      uint    `gorm:"primaryKey;autoIncrement:false" csv:"-"`
	Name    string  `json:"name"`
	Contact string  `json:"contact"`
	Phone   string  `json:"phone"`
	Email   string  `json:"email"`
	Address string  `json:"address"`
	Balance float32 `json:"balance"`
}

func ExportClients(c *gin.Context) {
	client := []export{}

	baseQuery := database.Database.Model(&types.Client{})

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

	fileName := uuid.New().String() + ".csv"
	csvFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = gocsv.MarshalFile(&client, csvFile)
	c.File(fileName)

	os.Remove(fileName)
}
