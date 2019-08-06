package upload

import (
	"com/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func UploadImages(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "错误请求",
		})
		return
	}
	//文件的名称
	filename := header.Filename
	//创建空文件
	out, err := os.Create("public/upload/images/" + filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "文件创建失败",
		})
		return
	}
	defer out.Close()
	// 将文件复制进空文件
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "复制文件失败",
		})
		return
	}
	//封面表添加
	id, err := models.CreatedImage("http://192.168.2.219:3000/images/" + filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "封面添加数据库失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   id,
	})
}

func UploadVideos(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "错误请求",
		})
		return
	}
	//文件的名称
	filename := header.Filename
	//创建文件
	out, err := os.Create("public/upload/videos/" + filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "文件创建失败",
		})
		return
	}

	defer out.Close()
	// 将文件复制进空文件
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "复制文件失败",
		})
		return
	}
	// 返回视频路径id
	id, err := models.CreatedVideoSrc("http://192.168.2.219:3000/videos/" + filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "视频路径添加数据库失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   id,
	})
}
