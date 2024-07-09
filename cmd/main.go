package main

import (
	"github.com/gin-gonic/gin"
	imagePal "github.com/sacenox/image-pal/internal"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"health": "ok",
		})
	})

	r.GET("/resize", imagePal.ResizeRoute)

	r.Run()
}
