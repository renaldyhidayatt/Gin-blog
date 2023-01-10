package main

import (
	"ginBlog/config"
	"ginBlog/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	db := config.SetupDatabase()

	routes.NewRoute(db, r)

	r.Run(":5000")
}
