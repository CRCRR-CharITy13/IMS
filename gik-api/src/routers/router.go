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
	}

	return r
}
