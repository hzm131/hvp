package upload

import (
	"com/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smokezl/govalidators"
	"io/ioutil"
	"net/http"
	"strconv"
)

//创建视频
func PostVideo(c *gin.Context) {
	validator := govalidators.New()
	video := models.Videos{}
	value, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
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

	videoId, err := models.CreatedVideo(video)
	if err != nil || videoId == 0 {
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
		"data":   videoId,
	})
}

//更新视频
func UpdateVideo(c *gin.Context) {
	key := c.Param("id")
	id, error := strconv.Atoi(key)
	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  error,
			"data":   "url中id获取失败",
		})
		return
	}
	fmt.Println("id:", id)
	validator := govalidators.New()
	video := models.Videos{}
	value, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
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
	bool := models.UpdatedVideo(id, video)
	if bool == false {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  "数据库更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   "更新成功",
	})
}

//查询视频
func QueryVideo(c *gin.Context) {
	limit := c.Request.URL.Query().Get("limit")
	limitint, error := strconv.Atoi(limit)
	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  error,
			"data":   "url中id获取失败",
		})
		return
	}
	offset := c.Request.URL.Query().Get("offset")
	offsetint, error := strconv.Atoi(offset)
	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  error,
			"data":   "url中id获取失败",
		})
		return
	}

	value, err := models.QueryVideo(limitint, offsetint)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  error,
			"data":   "查询有错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   value,
	})
}
