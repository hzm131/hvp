package video_src

import (
	"com/models/servser_model/video"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteImageSrc(c *gin.Context) {
	VideoSrcId := c.Param("id")
	if VideoSrcId == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"data":   "没有传id",
		})
		return
	}
	video_src := video.VideoSrc{}
	err := video_src.DeleteVideoSrc(VideoSrcId)
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
