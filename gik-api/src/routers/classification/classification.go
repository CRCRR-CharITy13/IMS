package classification

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	row, err := database.Database.Query("SELECT * FROM categories")

	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	categories := []types.Category{}
	for row.Next() {
		thisCategory := types.Category{}
		err = row.Scan(&thisCategory.CategoryCode, &thisCategory.CategoryName)
		fmt.Printf("Category_code: %s, Category_name: %s", thisCategory.CategoryCode, thisCategory.CategoryName)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, thisCategory)
	}
	totalCategory := len(categories)
	c.JSON(200, gin.H{"success": true, "data": gin.H{
		"data":  categories,
		"total": totalCategory,
	}})
}