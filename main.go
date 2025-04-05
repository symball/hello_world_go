package main

import (
	"log"
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
	r.GET("/test-duplicate", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"secret": "BIzaSyDaGmWKa4JsXZ-HjGw7ISLn_3namBGfuQe",
		})
	})
	port := ":" + os.Getenv("GOHW_PORT")

	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
