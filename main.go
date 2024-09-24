package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	port := ":" + os.Getenv("GOHW_PORT")

	r.Run(port) // listen and serve on 0.0.0.0:8080
}
