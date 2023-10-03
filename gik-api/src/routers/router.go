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
	"GIK_Web/src/routers/tags"
	"GIK_Web/src/routers/transaction"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

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

	// r.Use(middleware.AuthMiddleware())

	r.GET("/ping", status.Ping)

	// Set up Analytics router
	analyticsApi := r.Group("/analytics")
	{
		analyticsApi.Use(middleware.AuthMiddleware())
		analyticsApi.Use(middleware.AdvancedLoggingMiddleware())
		analyticsApi.GET("/transaction", analytics.GraphTransactions)
		analyticsApi.GET("/transaction/total", analytics.GraphTotalTransactions)
		analyticsApi.GET("/attention-item", analytics.AttentionRequiredItem)
		analyticsApi.GET("/attention-location", analytics.AttentionRequiredLocation)
		analyticsApi.GET("/activity", analytics.GetRecentActivity)
		analyticsApi.GET("/trending", analytics.GetTrendingItems)
	}

	// Set up authentication router
	authApi := r.Group("/auth")
	{
		authApi.POST("/login", auth.Login)
		authApi.POST("/prelogin", auth.CheckPasswordForLogin)
		authApi.GET("/tfa", auth.CheckTfaStatusBeforeLogin)
		authApi.POST("/register", auth.Register)
		authApi.GET("/first_admin", auth.CreateFirstAdmin)
		authApi.GET("/scode", auth.GetSignupCodeInfo)
		authApi.GET("/status", middleware.AuthMiddleware(), auth.CheckAuthStatus)
		authApi.GET("/logout", auth.Logout)
	}

	// Set up items router
	itemsApi := r.Group("/items")
	{
		itemsApi.Use(middleware.AuthMiddleware())
		itemsApi.Use(middleware.AdvancedLoggingMiddleware())
		itemsApi.GET("/list", items.ListItem)
		itemsApi.GET("/lookup", items.LookupItem)
		itemsApi.GET("/export", items.ExportItems)
		itemsApi.POST("/import", items.ImportItems)
		itemsApi.GET("/suggest", items.GetAutoSuggest)
		itemsApi.PUT("/add", items.AddItem)
		itemsApi.PATCH("/update", items.UpdateItem)
		itemsApi.DELETE("/delete", items.DeleteItem)
		itemsApi.GET("/list-location-for-item", items.ListLocationForItem)
		itemsApi.GET("/get-unstored-quantity", items.GetUnstoredQuantity)
	}

	// Set up tags router
	tagsApi := r.Group("/tags")
	{
		tagsApi.Use(middleware.AuthMiddleware())
		tagsApi.Use(middleware.AdvancedLoggingMiddleware())
		tagsApi.GET("/list", tags.ListTags)
		tagsApi.PUT("/add", tags.AddTags)
		tagsApi.PATCH("/update", tags.UpdateTags)
		tagsApi.DELETE("/delete", tags.DeleteTags)

	}

	// Set up client router
	clientsApi := r.Group("/client")
	{
		clientsApi.Use(middleware.AuthMiddleware())
		clientsApi.Use(middleware.AdvancedLoggingMiddleware())
		clientsApi.GET("/list", client.ListClient)
		clientsApi.GET("/export", client.ExportClients)
		clientsApi.POST("/import", client.ImportClients)
		// clientsApi.GET("/lookup", client.LookupLocation)
		clientsApi.PUT("/add", client.AddClient)
		clientsApi.DELETE("/delete", client.DeleteClient)
		clientsApi.PATCH("/update", client.UpdateClient)
	}

	// Set up location router
	locationsApi := r.Group("/location")
	{
		locationsApi.Use(middleware.AuthMiddleware())
		locationsApi.Use(middleware.AdvancedLoggingMiddleware())
		locationsApi.GET("/list", location.ListLocation)
		// locationsApi.GET("/lookup", location.LookupLocation)
		locationsApi.PUT("/add", location.AddLocation)
		locationsApi.DELETE("/delete", location.DeleteLocation)
		locationsApi.PATCH("/update", location.UpdateLocation)
		//locationsApi.GET("/scan", location.GetScannedData)
		locationsApi.PUT("/add-item-to-location", location.AddItemToLocation)
		locationsApi.GET("/list-item-in-location", location.ListItemInLocation)
	}

	// Set up transactions router
	transactionApi := r.Group("/transaction")
	{
		transactionApi.Use(middleware.AuthMiddleware())
		transactionApi.Use(middleware.AdvancedLoggingMiddleware())
		transactionApi.GET("/list", transaction.ListTransactions)
		// transactionApi.GET("/listItem", transaction.ListTransactionItem)
		// transactionApi.PATCH("/updateItem", transaction.UpdateTransactionItem)
		transactionApi.PUT("/add", transaction.AddTransaction)
		// transactionApi.PUT("/addItem", transaction.AddTransactionItem)
		transactionApi.DELETE("/delete", transaction.DeleteTransaction)
		// transactionApi.DELETE("/deleteItem", transaction.DeleteTransactionItem)
		transactionApi.GET("/items", transaction.GetTransactionItems)
	}

	// Set up logs
	logsApi := r.Group("/logs")
	{
		logsApi.Use(middleware.AuthMiddleware())
		// Temporary commented by tuan on Sept 23 to avoid creating too many logs
		// during the audit logs display
		// logsApi.Use(middleware.AdvancedLoggingMiddleware())
		logsApi.GET("/advanced", logs.GetAdvancedLogs)
		logsApi.GET("/simple", logs.GetSimpleLogs)
	}

	// Set up admin
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
		adminApi.GET("/scodes", admin.GetSignupCodes)
		adminApi.GET("/scode", admin.CreateSignupCode)
		adminApi.PATCH("/scode/toggle", admin.ToggleSignupCode)
		adminApi.DELETE("/scode/delete", admin.DeleteSignupCodes)
		adminApi.GET("/lists", admin.ListAdminsAndNonAdmins)
		adminApi.GET("/users", admin.ListUsers)
		adminApi.PATCH("/user/toggle", admin.ToggleUser)
		adminApi.DELETE("/user/delete", admin.DeleteUser)
		adminApi.PATCH("/admins", admin.EditAdmins)
	}

	// Set up settings
	settingsApi := r.Group("/settings")
	{
		settingsApi.Use(middleware.AuthMiddleware())
		settingsApi.Use(middleware.AdvancedLoggingMiddleware())
		settingsApi.PATCH("/password", settings.ChangePassword)
		tfaApi := settingsApi.Group("/tfa")
		{
			tfaApi.GET("/status", settings.GetTfaStatus)
			tfaApi.GET("/generate", settings.GenerateTwoFactorSecret)
			tfaApi.PATCH("/setup", settings.ValidateAndSetupTwoFactor)
		}
	}

	// Set up QR code
	qrApi := r.Group("/qr")
	{
		qrApi.Use(middleware.AuthMiddleware())
		qrApi.Use(middleware.AdvancedLoggingMiddleware())
		qrApi.GET("/codes", qr.GetQRCodes)
	}

	// Set up invoice
	invoiceApi := r.Group("/invoice")
	{
		invoiceApi.Use(middleware.AuthMiddleware())
		invoiceApi.Use(middleware.AdvancedLoggingMiddleware())
		invoiceApi.POST("/generate", invoice.GetInvoice)
	}

	// Set up basic info
	infoApi := r.Group("/info")
	{
		infoApi.Use(middleware.AuthMiddleware())

		infoApi.GET("/username", info.GetUsername)
		infoApi.GET("/currentusername", info.GetCurrentUsername)
		infoApi.GET("/client", info.GetClientInfo)
	}

	return r
}
