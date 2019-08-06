package user

import (
	"com/middleware/jwt"
	"com/models/servser_model/users"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smokezl/govalidators"
	"io/ioutil"
	"net/http"
)

// 登录
func Login(c *gin.Context) {
	validator := govalidators.New()
	user := users.Users{}
	value, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println("value", value)
	if err != nil {
		return
	}
	json.Unmarshal(value, &user)
	errList := validator.Validate(&user)
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

	id, err := user.FindId()
	if id == 0 || err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "用户不存在",
		})
		return
	}

	str, err := jwt.CreateJWT(id)
	if err != nil {
		fmt.Errorf("生成jwt失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   str,
	})
}

//注册
func Registered(c *gin.Context) {
	user := users.Users{}
	value, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	json.Unmarshal(value, &user)
	errList := govalidators.New().Validate(&user)
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
	id, err := user.CreateData()
	if id == -1 {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  "用户名已经存在",
		})
		return
	}
	if err != nil || id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  "创建失败",
		})
		return
	}
	str, err := jwt.CreateJWT(id) //返回完整token
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "jwt生成失败",
		})
		return
	}
	fmt.Println("打印完整的token:", str) //打印token
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   str,
		"userId": id,
	})
}
