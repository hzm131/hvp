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
	video := video.Video{}
	value, err := video.QueryVideos(condition,limit, offset)
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
