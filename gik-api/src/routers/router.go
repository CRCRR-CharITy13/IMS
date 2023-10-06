/*
router sets up the HTTP routing.

Contains the routing for all the GET, SET, DELETE, etc. commands and their
associated function.
*/
package routers

// import "github.com/gin-gonic/gin"

import (
	"GIK_Web/src/middleware"
	"GIK_Web/src/routers/admin"
	"GIK_Web/src/routers/analytics"
	"GIK_Web/src/routers/auth"
	"GIK_Web/src/routers/client"
	"GIK_Web/src/routers/info"
	"GIK_Web/src/routers/invoice"
	"GIK_Web/src/routers/items"
	"GIK_Web/src/routers/location"
	"GIK_Web/src/routers/logs"
	"GIK_Web/src/routers/qr"
	"GIK_Web/src/routers/settings"
	"GIK_Web/src/routers/status"
	"GIK_Web/src/routers/transaction"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// InitRouter initializes the router and the connections before returning it
func InitRouter() *gin.Engine {
	// Create a new Gin Engine
	r := gin.New()

	// Set up simple logger middleware
	r.Use(gin.Logger())

	// Set up middleware to send 500 when panic
	r.Use(gin.Recovery())

	// Enable the cross-origin resource sharing middleware
	r.Use(middleware.CORSMiddleware())

	// Set up cookies
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))

	// Set up a basic ping
	r.GET("/ping", status.Ping)

	// Set up Analytics router
	// Requires Authentication + Logging
	analyticsApi := r.Group("/analytics")
	{
		analyticsApi.Use(middleware.AuthMiddleware())
		analyticsApi.Use(middleware.AdvancedLoggingMiddleware())

		analyticsApi.GET("/transaction", analytics.GraphTransactions)
		analyticsApi.GET("/transaction/total", analytics.GraphTotalTransactions)
		analyticsApi.GET("/attention", analytics.AttentionRequired)
		analyticsApi.GET("/activity", analytics.GetRecentActivity)
		analyticsApi.GET("/trending", analytics.GetTrendingItems)
	}

	// Set up authentication router
	authApi := r.Group("/auth")
	{
		authApi.POST("/login", auth.Login)
		authApi.POST("/prelogin", auth.CheckPasswordForLogin)
		authApi.POST("/register", auth.Register)
		authApi.GET("/first_admin", auth.CreateFirstAdmin)
		authApi.GET("/scode", auth.GetSignupCodeInfo)
		authApi.GET("/status", middleware.AuthMiddleware(), auth.CheckAuthStatus)
		authApi.GET("/logout", auth.Logout)
	}

	// Set up items router
	// Requires Authentication + Logging
	itemsApi := r.Group("/item")
	{
		itemsApi.Use(middleware.AuthMiddleware())
		itemsApi.Use(middleware.AdvancedLoggingMiddleware())

		itemsApi.GET("/", items.ListItem)
		itemsApi.POST("/", items.AddItem)
		itemsApi.PUT("/", items.UpdateItem)
		itemsApi.PATCH("/", items.EditItem)
		itemsApi.DELETE("/", items.DeleteItem)

		itemsApi.GET("/suggest", items.GetAutoSuggest)
		itemsApi.PUT("/import", items.ImportItems)
		itemsApi.GET("/lookup", items.LookupItem)
		itemsApi.GET("/export", items.ExportItems)
		itemsApi.GET("/location", items.ListLocationForItem)
	}

	// Set up client router
	// Requires Authentication + Logging
	clientsApi := r.Group("/client")
	{
		clientsApi.Use(middleware.AuthMiddleware())
		clientsApi.Use(middleware.AdvancedLoggingMiddleware())

		clientsApi.GET("/", client.ListClient)
		clientsApi.POST("/", client.AddClient)
		clientsApi.PUT("/", client.UpdateClient)
		clientsApi.DELETE("/", client.DeleteClient)

		clientsApi.GET("/export", client.ExportClients)
		clientsApi.PUT("/import", client.ImportClients)
	}

	// Set up location router
	// Requires Authentication + Logging
	locationsApi := r.Group("/location")
	{
		locationsApi.Use(middleware.AuthMiddleware())
		locationsApi.Use(middleware.AdvancedLoggingMiddleware())

		locationsApi.GET("/", location.ListLocation)
		locationsApi.POST("/", location.AddLocation)
		locationsApi.DELETE("/", location.DeleteLocation)
		locationsApi.PUT("/", location.UpdateLocation)

		locationsApi.POST("/item", location.AddItemToLocation)
		locationsApi.GET("/item", location.ListItemInLocation)
	}

	// Set up transactions router
	// Requires Authentication + Logging
	transactionApi := r.Group("/transaction")
	{
		transactionApi.Use(middleware.AuthMiddleware())
		transactionApi.Use(middleware.AdvancedLoggingMiddleware())

		transactionApi.GET("/", transaction.ListTransactions)
		transactionApi.POST("/", transaction.AddTransaction)
		transactionApi.DELETE("/", transaction.DeleteTransaction)
		transactionApi.GET("/item", transaction.GetTransactionItems)
	}

	// Set up logs
	// Requires Authentication + Logging
	logsApi := r.Group("/logs")
	{
		logsApi.Use(middleware.AuthMiddleware())
		logsApi.Use(middleware.AdvancedLoggingMiddleware())

		logsApi.GET("/advanced", logs.GetAdvancedLogs)
		logsApi.GET("/simple", logs.GetSimpleLogs)
	}

	// Set up admin
	// Requires Authentication + Logging + Admin
	adminApi := r.Group("/admin")
	{
		adminApi.Use(middleware.AuthMiddleware())
		adminApi.Use(middleware.AdvancedLoggingMiddleware())
		adminApi.Use(middleware.AdminMiddleware())

		adminApi.GET("/status", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"success": true,
				"message": "Admin verified",
			})
		})

		adminApi.GET("/scode", admin.GetSignupCodes)
		adminApi.POST("/scode", admin.CreateSignupCode)
		adminApi.PATCH("/scode", admin.ToggleSignupCode)
		adminApi.DELETE("/scode", admin.DeleteSignupCodes)

		adminApi.GET("/", admin.ListAdminsAndNonAdmins)
		adminApi.GET("/users", admin.ListUsers)
		adminApi.PATCH("/users", admin.ToggleUser)
		adminApi.DELETE("/users", admin.DeleteUser)

		adminApi.PATCH("/admins", admin.EditAdmins)
	}

	// Set up settings
	// Requires Authentication + Logging
	settingsApi := r.Group("/settings")
	{
		settingsApi.Use(middleware.AuthMiddleware())
		settingsApi.Use(middleware.AdvancedLoggingMiddleware())

		settingsApi.PATCH("/", settings.ChangePassword)
	}

	// Set up QR code
	// Requires Authentication + Logging
	qrApi := r.Group("/qr")
	{
		qrApi.Use(middleware.AuthMiddleware())
		qrApi.Use(middleware.AdvancedLoggingMiddleware())

		qrApi.GET("/codes", qr.GetQRCodes)
	}

	// Set up invoice
	// Requires Authentication + Logging
	invoiceApi := r.Group("/invoice")
	{
		invoiceApi.Use(middleware.AuthMiddleware())
		invoiceApi.Use(middleware.AdvancedLoggingMiddleware())

		invoiceApi.GET("/", invoice.GetInvoice)
	}

	// Set up basic info
	// Requires Authentication + Logging
	infoApi := r.Group("/info")
	{
		infoApi.Use(middleware.AuthMiddleware())

		infoApi.GET("/username", info.GetUsername)
		infoApi.GET("/currentusername", info.GetCurrentUsername)
		infoApi.GET("/client", info.GetClientInfo)
	}

	return r
}
