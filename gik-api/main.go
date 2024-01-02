/*
main sets up the API server and connects to a database.

The settings are writen in the .env file.
*/

package main

import (
	"GIK_Web/database"
	"GIK_Web/env"
	"GIK_Web/src/routers"
	"fmt"
	"net/http"
	//"github.com/fvbock/endless"
)

// Main function, when run will initialize the server.
func main() {
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	// router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	// your custom format
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))
	// router.Run(":8080")

	// Set up environmental data from .env file
	fmt.Println("Start to run GIN server")
	env.SetEnv()
	fmt.Println("IsLocalDB = ", env.IsLocalDB)
	// Initalize the router
	routersInit := routers.InitRouter()

	// Connect to the database
	database.ConnectDatabase()

	fmt.Printf("\nServer running at %s:%s\n", env.WebserverHost, env.WebserverPort)

	//routersInit.Run(":" + env.WebserverPort)

	server := &http.Server{
		Handler: routersInit,
		Addr:    ":" + env.WebserverPort,
	}

	fmt.Printf("\n Server running at port %s\n", env.WebserverPort)

	if env.HTTPS {
		server.ListenAndServeTLS(".cert/server.crt", ".cert/server.key")
	} else {
		server.ListenAndServe()
	}

}
