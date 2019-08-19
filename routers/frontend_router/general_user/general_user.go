package general_user

import (
	"com/middleware/jwt"
	"com/models/frontend_model/general_user"
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
	general_user := general_user.GeneralUser{}
	value, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println("value", value)
	if err != nil {
		return
	}
	json.Unmarshal(value, &general_user)
	errList := validator.Validate(&general_user)
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

	user, err := general_user.FindId()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "用户名或密码有问题",
		})
		return
	}

	str, err := jwt.CreateuUserJWT(user)
	if err != nil {
		fmt.Errorf("生成jwt失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   str,
		"user":   user,
	})
}

//注册
func Registered(c *gin.Context) {
	general_user := general_user.GeneralUser{}
	value, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	json.Unmarshal(value, &general_user)
	errList := govalidators.New().Validate(&general_user)
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
	id, user, err := general_user.CreateData()
	if id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  "用户名已经存在",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  "创建失败",
		})
		return
	}
	str, err := jwt.CreateuUserJWT(user) //返回完整token
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
		"user":   user,
	})
}
