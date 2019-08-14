package image_src

import (
	"com/models/servser_model/video"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteImageSrc(c *gin.Context) {
	imageId := c.Param("id")
	if imageId == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"data":   "没有传id",
		})
		return
	}
	image_src := video.ImageSrc{}
	fmt.Println("video:", image_src)
	err := image_src.DeleteImageSrc(imageId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   "删除成功",
	})
}
