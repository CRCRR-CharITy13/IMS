package items

import (
	"github.com/gin-gonic/gin"
)

func LookupItem(c *gin.Context) {
	// product id
	// productId := c.Query("id")

	// if productId == "" {
	// 	c.JSON(400, gin.H{
	// 		"success": false,
	// 		"message": "Product ID not provided",
	// 	})
	// 	return
	// }

	// // convert to int
	// productIdInt, err := strconv.Atoi(productId)
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"success": false,
	// 		"message": "Invalid product ID",
	// 	})
	// 	return
	// }

	// var postData item
	// database.Database.Model(&types.Item{}).Where("id = ?", productIdInt).Scan(&postData)

	// c.JSON(200, gin.H{
	// 	"success": true,
	// 	"data":    postData,
	// })

}
