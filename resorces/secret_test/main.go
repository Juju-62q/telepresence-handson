package main

import (
  "os"

  "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/secret", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"secret": "secret is " + os.Getenv("MY_SECRET"),
		})
	})
	r.GET("/config", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"config": "configmap is " + os.Getenv("MY_CONFIG"),
		})
	})
  r.Run(":80") // listen and serve on 0.0.0.0:8080
}
