package routers

import (
	_ "com/models"
	"com/routers/auth"
	"com/routers/servser_router/upload"
	"com/routers/servser_router/user"
	videoManagement "com/routers/servser_router/video"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//静态文件服务 提供图片路径
	r.Static("/images", "./public/upload/images")
	//r.Static("/static/images", "./public/images")
	r.Static("/videos", "./public/upload/videos")
	r.Static("/index.html", "./public/dist")

	//登录模块
	userApi := r.Group("/user")
	{
		userApi.POST("/login", user.Login)
		userApi.POST("/registered", user.Registered)
	}

	//上传模块
	uploadApi := r.Group("/upload")
	uploadApi.Use(auth.GetAuth)
	{
		uploadApi.POST("/video", upload.UploadVideo) //上传视频
		uploadApi.POST("/image", upload.UploadImage) //上传视频
	}

	videoApi := r.Group("/video")
	videoApi.Use(auth.GetAuth)
	{
		videoApi.POST("/create", videoManagement.CreateVideo)
		videoApi.GET("/query", videoManagement.QueryVideo)
		videoApi.GET("/query/:id", videoManagement.FindVideo)
		videoApi.PUT("/update/:id", videoManagement.UpdateVideo)
		videoApi.DELETE("/delete/:id", videoManagement.DeleteVideo)
	}

	testApi := r.Group("/test")
	testApi.Use(auth.GetAuth)
	{
		testApi.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			value, _ := c.Get("user")
			fmt.Println("value", value)
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"error":  nil,
				"data":   "徐倩雯。。。",
				"id":     id,
			})
		})
	}
	return r
}
