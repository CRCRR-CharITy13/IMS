package category

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
		err = row.Scan(&thisCategory.Category_code, &thisCategory.Category_name)
		fmt.Printf("Category_code: %s, Category_name: %s", thisCategory.Category_code, thisCategory.Category_name)
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
