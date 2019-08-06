package upload

import (
	"com/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/smokezl/govalidators"
	"io/ioutil"
	"net/http"
)

//第一级分类
func Class(c *gin.Context) {
	validator := govalidators.New()
	class := models.Classes{}
	value, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	json.Unmarshal(value, &class)
	errList := validator.Validate(&class)
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
	ClassId, err := models.CreatedClass(class)
	if err != nil || ClassId == 0 {
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
		"data":   ClassId,
	})

}

//第二级分类
func Secondary(c *gin.Context) {
	validator := govalidators.New()
	secondary := models.Secondary{}
	value, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	json.Unmarshal(value, &secondary)
	errList := validator.Validate(&secondary)
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
	SecondaryId, err := models.CreatedSecondary(secondary)
	if err != nil || SecondaryId == 0 {
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
		"data":   SecondaryId,
	})

}
