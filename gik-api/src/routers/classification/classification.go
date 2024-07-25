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

func ListSubCategory1(c *gin.Context) {
	row, err := database.Database.Query("SELECT * FROM SubCategories1")

	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	SubCategories1 := []types.SubCategory1{}
	for row.Next() {
		thisSubCategory1 := types.SubCategory1{}
		err = row.Scan(&thisSubCategory1.SubCategory1Code, &thisSubCategory1.SubCategory1Name)
		fmt.Printf("Category_code: %d, Category_name: %s", thisSubCategory1.SubCategory1Code, thisSubCategory1.SubCategory1Name)
		if err != nil {
			log.Fatal(err)
		}
		SubCategories1 = append(SubCategories1, thisSubCategory1)
	}
	totalSubCategory1 := len(SubCategories1)
	c.JSON(200, gin.H{"success": true, "data": gin.H{
		"data":  SubCategories1,
		"total": totalSubCategory1,
	}})
}
