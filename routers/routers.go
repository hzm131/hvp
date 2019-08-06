package routers

import (
	_ "com/models"
	"com/routers/auth"
	"com/routers/servser_router/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//静态文件服务 提供图片路径
	r.Static("/images", "./public/upload/images")
	r.Static("/static/images", "./public/images")
	r.Static("/videos", "./public/upload/videos")
	r.Static("/index.html", "./public/dist")

	//登录模块
	userApi := r.Group("/user")
	{
		userApi.POST("/login", user.Login)
		userApi.POST("/registered", user.Registered)
	}

	testApi := r.Group("/test")
	testApi.Use(auth.GetAuth)
	{
		testApi.POST("/1", func(context *gin.Context) {
			value, _ := context.Get("userId")
			fmt.Println("value", value)
		})
	}
	return r
}