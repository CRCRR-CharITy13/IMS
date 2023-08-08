package tags

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"strings"

	"github.com/gin-gonic/gin"
)

type tag struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func ListTags(c *gin.Context) {
	name := c.Query("name")
	tags := []types.Tag{}
	database.Database.Model(&types.Tag{}).Where("name LIKE ?", "%"+name+"%").Order("name").Find(&tags)

	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    tagNames,
	})
}

func UpdateTags(c *gin.Context) {
	json := tag{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	database.Database.Model(&tag{}).Where("ID = ?", json.ID).Update("name", json.Name)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Updated tag",
	})

}

func AddTags(c *gin.Context) {
	name := strings.ToLower(c.Query("name"))
	tags := []types.Tag{}

	database.Database.Model(&types.Tag{}).Where("name = ?", name).Find(&tags)

	if len(tags) > 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Tag Already Exists",
		})
		return
	}

	database.Database.Create(&types.Tag{Name: name})

	c.JSON(200, gin.H{
		"success": true,
		"message": "Tag Added",
	})

}

func DeleteTags(c *gin.Context) {
	name := c.Query("name")

	database.Database.Where("name = ?", name).Delete(&types.Tag{})

	c.JSON(200, gin.H{
		"success": true,
		"message": "Tag Deleted",
	})

}
