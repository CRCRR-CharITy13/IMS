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
	env.SetEnv()

	// Initalize the router
	routersInit := routers.InitRouter()

	// Connect to the database
	database.ConnectDatabase()

	fmt.Printf("\nServer running at %s:%s\n", env.WebserverHost, env.WebserverPort)

	// Try to run the server

	// if env.HTTPS {
	// 	endless.ListenAndServeTLS(":"+env.WebserverPort, ".cert/server.crt", ".cert/server.key", routersInit)
	// } else {
	// 	endless.ListenAndServe(env.WebserverHost+":"+env.WebserverPort, routersInit)
	// }

	server := &http.Server{
		Handler: routersInit,
		Addr:    env.WebserverHost + ":" + env.WebserverPort,
	}

	fmt.Printf("\n Server running at %s:%s\n", env.WebserverHost, env.WebserverPort)

	if env.HTTPS {
		server.ListenAndServeTLS(".cert/server.crt", ".cert/server.key")
	} else {
		server.ListenAndServe()
	}

}
