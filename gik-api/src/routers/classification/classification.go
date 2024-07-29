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
		fmt.Printf("SubCategory1_code: %d, SubCategory1_name: %s; ", thisSubCategory1.SubCategory1Code, thisSubCategory1.SubCategory1Name)
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

func ListSubCategory2(c *gin.Context) {
	row, err := database.Database.Query("SELECT * FROM SubCategories2")

	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	SubCategories2 := []types.SubCategory2{}
	for row.Next() {
		thisSubCategory2 := types.SubCategory2{}
		err = row.Scan(&thisSubCategory2.SubCategory2Code, &thisSubCategory2.SubCategory2Name)
		fmt.Printf("SubCategory2_code: %s, SubCategory2_name: %s; ", thisSubCategory2.SubCategory2Code, thisSubCategory2.SubCategory2Name)
		if err != nil {
			log.Fatal(err)
		}
		SubCategories2 = append(SubCategories2, thisSubCategory2)
	}
	totalSubCategory2 := len(SubCategories2)
	c.JSON(200, gin.H{"success": true, "data": gin.H{
		"data":  SubCategories2,
		"total": totalSubCategory2,
	}})
}

func ListItemDescription(c *gin.Context) {
	row, err := database.Database.Query("SELECT * FROM Items_Descriptions")

	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	ItemsDescriptions := []types.ItemsDescription{}
	for row.Next() {
		thisItemsDescription := types.ItemsDescription{}
		err = row.Scan(&thisItemsDescription.DescriptionCode, &thisItemsDescription.DescriptionName)
		// fmt.Printf("Description_code: %d, Descriptions_name: %s; ", thisItemsDescription.DescriptionCode, thisItemsDescription.DescriptionName)
		if err != nil {
			log.Fatal(err)
		}
		ItemsDescriptions = append(ItemsDescriptions, thisItemsDescription)
	}
	totalItemsDescription := len(ItemsDescriptions)
	c.JSON(200, gin.H{"success": true, "data": gin.H{
		"data":  ItemsDescriptions,
		"total": totalItemsDescription,
	}})
}

func ListSize(c *gin.Context) {
	row, err := database.Database.Query("SELECT * FROM Sizes")

	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	ItemsSizes := []types.Size{}
	for row.Next() {
		thisItemSize := types.Size{}
		err = row.Scan(&thisItemSize.ItemSize, &thisItemSize.Description)
		fmt.Printf("Size: %s, Description: %s; ", thisItemSize.ItemSize, thisItemSize.Description)
		if err != nil {
			log.Fatal(err)
		}
		ItemsSizes = append(ItemsSizes, thisItemSize)
	}
	totalItemsSize := len(ItemsSizes)
	c.JSON(200, gin.H{"success": true, "data": gin.H{
		"data":  ItemsSizes,
		"total": totalItemsSize,
	}})
}
