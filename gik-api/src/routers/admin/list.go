/*
list lists the users/admin.
*/

package admin

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	// Get list of users sorted by time regestered
	users := []types.User{}
	if err := database.Database.Order("registered_at desc").Find(&users).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Error listing users",
		})
		return
	}

	// Return the list
	c.JSON(200, gin.H{
		"success": true,
		"message": "Retrieved users",
		"data":    users,
	})
}

type formattedUser struct {
	ID   string `json:"value"`
	Name string `json:"label"`
}

func ListAdminsAndNonAdmins(c *gin.Context) {
	admins := []types.User{}
	nonAdmins := []types.User{}

	// Get list of admins
	if err := database.Database.Model(&types.User{}).Where("admin = ?", true).Find(&admins).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error getting admins",
		})
		return
	}

	// Get list of non-admins
	if err := database.Database.Model(&types.User{}).Where("admin = ?", false).Find(&nonAdmins).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error getting non-admins",
		})
		return
	}

	// Format the users into a more sparse version
	adminsFormatted := []formattedUser{}
	nonAdminsFormatted := []formattedUser{}

	for _, admin := range admins {
		adminsFormatted = append(adminsFormatted, formattedUser{
			ID:   fmt.Sprintf("%d", admin.ID),
			Name: admin.Username,
		})
	}

	for _, nonAdmin := range nonAdmins {
		nonAdminsFormatted = append(nonAdminsFormatted, formattedUser{
			ID:   fmt.Sprintf("%d", nonAdmin.ID),
			Name: nonAdmin.Username,
		})
	}

	// Return two lists
	c.JSON(200, gin.H{
		"success": true,
		"data": [][]formattedUser{
			nonAdminsFormatted,
			adminsFormatted,
		},
	})
}
