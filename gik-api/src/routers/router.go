/*
router sets up the HTTP routing.

Contains the routing for all the GET, SET, DELETE, etc. commands and their
associated function.
*/

package routers

// import "github.com/gin-gonic/gin"

import (
	"GIK_Web/src/middleware"
	"GIK_Web/src/routers/auth"
	"GIK_Web/src/routers/classification"
	"GIK_Web/src/routers/status"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Create a new Gin Engine
	r := gin.New()

	r.GET("/ping", status.Ping)

	// Enable the cross-origin resource sharing middleware
	r.Use(middleware.CORSMiddleware())

	// Set up authentication router
	authApi := r.Group("/auth")
	{
		authApi.GET("/first_admin", auth.CreateFirstAdmin)
		authApi.GET("/status", middleware.AuthMiddleware(), auth.CheckAuthStatus)
	}

	categoryApis := r.Group("/classification")
	{
		categoryApis.GET("/list-category", classification.ListCategory)
		categoryApis.GET("/list-subcategory1", classification.ListSubCategory1)
	}

	// createItemApis := r.Group("/items")
	// {
	// 	createItemApis.POST("/create-item", classification.CreateItem)
	// }

	return r
}
