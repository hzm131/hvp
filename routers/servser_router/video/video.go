package videoManagement

import (
	"com/models/servser_model/video"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smokezl/govalidators"
	"io/ioutil"
	"net/http"
)

func CreateVideo(c *gin.Context) {
	validator := govalidators.New()
	video := video.Video{}
	value, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println("value", &value)
	if err != nil {
		fmt.Println("video反序列化是失败", err)
		return
	}
	json.Unmarshal(value, &video)
	errList := validator.Validate(&video)
	if errList != nil {
		for _, err := range errList {
			c.JSON(http.StatusOK, gin.H{
				"status": 400,
				"error":  err,
				"data":   "json数据类型不匹配",
			})
		}
		return
	}
	id, err := video.CreatedVideo()
	if id == 0 || err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"id":     id,
	})
}

func QueryVideo(c *gin.Context) {
	limit := c.Request.URL.Query().Get("limit")
	offset := c.Request.URL.Query().Get("offset")
	condition := c.Request.URL.Query().Get("condition")
	orderBy := c.Request.URL.Query().Get("order_by")
	fmt.Println("condition", condition)
	video := video.Video{}
	value, err := video.QueryVideos(condition,orderBy, limit, offset)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   value,
	})
}

func FindVideo(c *gin.Context) {
	videoId := c.Param("id")
	if videoId == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"data":   "没有更新id",
		})
		return
	}
	video := video.Video{}
	value, err := video.FindVideo(videoId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   value,
	})
}

func UpdateVideo(c *gin.Context) {
	videoId := c.Param("id")
	if videoId == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"data":   "没有更新id",
		})
		return
	}
	video := video.Video{}
	value, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("video反序列化是失败", err)
		return
	}
	json.Unmarshal(value, &video)
	fmt.Println("video:", video)
	err = video.UpdateVideo(videoId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   "更新成功",
	})
}

func DeleteVideo(c *gin.Context) {
	videoId := c.Param("id")
	if videoId == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"data":   "没有传id",
		})
		return
	}
	video := video.Video{}
	fmt.Println("video:", video)
	err := video.DeleteVideo(videoId)
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
