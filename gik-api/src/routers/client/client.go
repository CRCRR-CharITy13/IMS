package client

import (
	"GIK_Web/database"
	"GIK_Web/type_news"

	"GIK_Web/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type company struct {
	Name    string  `json:"name" binding:"required"`
	Contact string  `json:"contact" binding:"required"`
	Phone   string  `json:"phone" binding:"required"`
	Email   string  `json:"email" binding:"required"`
	Address string  `json:"address" binding:"required"`
	Balance float32 `json:"balance" binding:"required"`
}

type ListClientResponse struct {
	Id      int     `json:"ID" binding:"required"`
	Name    string  `json:"name" binding:"required"`
	Contact string  `json:"contact" binding:"required"`
	Phone   string  `json:"phone" binding:"required"`
	Email   string  `json:"email" binding:"required"`
	Address string  `json:"address" binding:"required"`
	Balance float32 `json:"balance" binding:"required"`
}

// Takes in a few queries of name, contact, address, phone, email and returns a list of clients who match these requests.
func ListClient(c *gin.Context) {
	// Create an empty array to store the list of clients.
	listClient := []type_news.Client{}

	// Creating the initial query for the model client
	baseQuery := database.Database.Model(&listClient)

	// Filter the query based on the field.
	name := c.Query("name")
	if name != "" {
		baseQuery = baseQuery.Where("org_name LIKE ?", name)
	}

	contact := c.Query("contact")
	if contact != "" {
		baseQuery = baseQuery.Where("contact LIKE ?", contact)
	}

	phone := c.Query("phone")
	if phone != "" {
		baseQuery = baseQuery.Where("phone_number LIKE ?", phone)
	}

	email := c.Query("email")
	if email != "" {
		baseQuery = baseQuery.Where("email LIKE ?", email)
	}

	address := c.Query("address")
	if address != "" {
		baseQuery = baseQuery.Where("address LIKE ?", address)
	}

	// Get and store the clients into the array.
	err := baseQuery.Find(&listClient).Error
	if err != nil {
		// If failed, return message and quit
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to query clients",
		})

		return
	}

	clientResponse := make([]ListClientResponse, len(listClient))
	for i, client := range listClient {
		clientResponse[i] = ListClientResponse{
			Id:      int(client.ID),
			Name:    client.OrgName,
			Contact: client.Contact,
			Phone:   client.PhoneNumber,
			Email:   client.Email,
			Address: client.Address,
			Balance: client.Balance,
		}
	}
	// Send message with the results.
	c.JSON(200, gin.H{
		"success": true,
		"message": "Client data retrieved",
		"data":    clientResponse,
	})

}

// Take an ID query and a JSON body of values and update the item.
func UpdateClient(c *gin.Context) {
	// Get the ID
	id := c.Query("id")
	if id == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	// Check that the ID is an integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	// Get the JSON body.
	json := company{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	client := type_news.Client{}
	if err := database.Database.Where("id = ?", idInt).First(&client).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid donor",
		})
		return
	}

	client.OrgName = json.Name
	client.Contact = json.Contact
	client.PhoneNumber = json.Phone
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

	newClient := type_news.Client{
		OrgName:     json.Name,
		Contact:     json.Contact,
		PhoneNumber: json.Phone,
		Email:       json.Email,
		Address:     json.Address,
		Balance:     float32(json.Balance),
	}

	if err := database.Database.Model(&type_news.Client{}).Create(&newClient).Error; err != nil {
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

	utils.CreateSimpleLog(c, "Client created")

}

// Take an ID query and delete that client.
func DeleteClient(c *gin.Context) {
	// Gets the ID
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	client := type_news.Client{}
	if err := database.Database.Model(&type_news.Client{}).Where("id = ?", idInt).First(&client).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid client",
		})
		return
	}

	if err := database.Database.Model(&type_news.Client{}).Delete(&client).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to delete client",
			"error":   err.Error(),
		})
		return
	}

	utils.CreateSimpleLog(c, "Client deleted")

	c.JSON(200, gin.H{
		"success": true,
		"message": "Client successfully deleted.",
	})
}
