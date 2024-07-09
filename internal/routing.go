package imagePal

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetImage(url string) (buffer []byte, contentType string, err error) {
	imageRequest, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer imageRequest.Body.Close()

	buffer, err = io.ReadAll(imageRequest.Body)
	if err != nil {
		return nil, "", err
	}

	contentType = imageRequest.Header.Get("content-type")

	return buffer, contentType, nil
}

type ResizeImageParams struct {
	ImageUrl string `form:"image_url"`
	Width    string `form:"width"`
	Height   string `form:"height"`
}

func ResizeRoute(c *gin.Context) {
	var params ResizeImageParams

	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// TODO: Defaults for width and height
	width, err := strconv.Atoi(params.Width)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	height, err := strconv.Atoi(params.Height)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	buffer, contentType, err := GetImage(params.ImageUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	resizedBuffer, err := Resize(
		&buffer, width, height,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// TODO: Should we cache the result?
	// err = bimg.Write(filename, resizedBuffer)

	c.Data(http.StatusOK, contentType, resizedBuffer)
}
