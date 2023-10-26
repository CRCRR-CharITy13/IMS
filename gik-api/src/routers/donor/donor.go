package donor

import (
	"GIK_Web/database"
	"GIK_Web/type_news"

	"github.com/gin-gonic/gin"

	"GIK_Web/utils"
	"strconv"
)

type ListDonorResponse struct {
	Id      int    `json:"ID" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Address string `json:"address" binding:"required"`
}

// Takes in a few queries of name, contact, address, phone, email and returns a list of donors who match these requests.
func ListDonor(c *gin.Context) {
	// Create an empty array to store the list of donors.
	listDonor := []type_news.Donor{}

	// Creating the initial query for the model donor
	baseQuery := database.Database.Model(&listDonor)

	// Filter the query based on the field.
	name := c.Query("name")
	if name != "" {
		baseQuery = baseQuery.Where("org_name LIKE ?", name)
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

	// Get and store the donors into the array.
	err := baseQuery.Find(&listDonor).Error
	if err != nil {
		// If failed, return message and quit
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to query donors",
		})

		return
	}

	donorResponse := make([]ListDonorResponse, len(listDonor))
	for i, donor := range listDonor {
		donorResponse[i] = ListDonorResponse{
			Id:      int(donor.ID),
			Name:    donor.Name,
			Phone:   donor.PhoneNumber,
			Email:   donor.Email,
			Address: donor.Address,
		}
	}
	// Send message with the results.
	c.JSON(200, gin.H{
		"success": true,
		"message": "Donor data retrieved",
		"data":    donorResponse,
	})

}

type donorInfo struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Address string `json:"address" binding:"required"`
}

func AddDonor(c *gin.Context) {
	json := donorInfo{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	newDonor := type_news.Donor{
		Name:        json.Name,
		PhoneNumber: json.Phone,
		Email:       json.Email,
		Address:     json.Address,
	}

	if err := database.Database.Model(&type_news.Donor{}).Create(&newDonor).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to create donor",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Donor successfully created.",
		"data":    json,
	})

	utils.CreateSimpleLog(c, "Donor created")

}

// Take an ID query and delete that donor.
func DeleteDonor(c *gin.Context) {
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

	donor := type_news.Donor{}
	if err := database.Database.Model(&type_news.Donor{}).Where("id = ?", idInt).First(&donor).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid donor",
		})
		return
	}

	if err := database.Database.Model(&type_news.Donor{}).Delete(&donor).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to delete donor",
			"error":   err.Error(),
		})
		return
	}

	utils.CreateSimpleLog(c, "Donor deleted")

	c.JSON(200, gin.H{
		"success": true,
		"message": "Donor successfully deleted.",
	})
}
