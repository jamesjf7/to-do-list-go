package main

import (
	"github.com/gin-gonic/gin"

	"to-do-list-go/src/config"
	"to-do-list-go/src/routes"
)

func main() {
	config.DB = config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(config.DB) // defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns (artinya jalan terus :D

	app := gin.Default()
	app.Use(func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT,DELETE")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	routes.Router(app)

	app.Run(":8080")
}
