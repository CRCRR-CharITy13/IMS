/*
router sets up the HTTP routing.

Contains the routing for all the GET, SET, DELETE, etc. commands and their
associated function.
*/

package routers

// import "github.com/gin-gonic/gin"

import (
	"GIK_Web/src/routers/classification"
	"GIK_Web/src/routers/status"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Create a new Gin Engine
	r := gin.New()

	r.GET("/ping", status.Ping)

	categoryApis := r.Group("/classification")
	{
		categoryApis.GET("/list-category", classification.ListCategory)
		categoryApis.GET("/list-subcategory1", classification.ListSubCategory1)
		categoryApis.GET("/list-subcategory2", classification.ListSubCategory2)
		categoryApis.GET("/list-itemDescription", classification.ListItemDescription)
		categoryApis.GET("/list-itemSize", classification.ListSize)
	}

	// createItemApis := r.Group("/items")
	// {
	// 	createItemApis.POST("/create-item", classification.CreateItem)
	// }

	return r
}
