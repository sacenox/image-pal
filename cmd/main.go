package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	imageTools "github.com/sacenox/image-pal/internal"
)

type ImageInput struct {
	ImageUrl string `form:"image_url"`
	Width    string `form:"width"`
	Height   string `form:"height"`
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"health": "ok",
		})
	})

	r.GET("/resize", func(c *gin.Context) {
		var action ImageInput

		if err := c.ShouldBind(&action); err != nil {
			c.JSON(400, gin.H{
				"msg": err.Error(),
			})
			return
		}

		log.Print(action)

		// Get the image:
		imageRequest, err := http.Get(action.ImageUrl)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": err.Error(),
			})
			return
		}
		defer imageRequest.Body.Close()

		buffer, err := io.ReadAll(imageRequest.Body)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": err.Error(),
			})
			return
		}

		// TODO: Defaults for width and height
		width, err := strconv.Atoi(action.Width)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": err.Error(),
			})
			return
		}

		height, err := strconv.Atoi(action.Height)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": err.Error(),
			})
			return
		}

		resizedBuffer, err := imageTools.Resize(
			buffer, width, height,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": err.Error(),
			})
			return
		}

		// TODO: Should we cache the result?
		// err = bimg.Write(filename, resizedBuffer)

		contentType := imageRequest.Header.Get("content-type")
		c.Data(200, contentType, resizedBuffer)
	})

	r.Run()
}
