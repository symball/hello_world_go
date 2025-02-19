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
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"secret": "AIzaSyDaGmWKa4JsXZ-HjGw7ISLn_3namBGewQe",
		})
	})
	port := ":" + os.Getenv("GOHW_PORT")

	r.Run(port) // listen and serve on 0.0.0.0:8080
}


