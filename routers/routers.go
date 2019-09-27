package routers

import (
	"com/middleware/cors"
	_ "com/models"
	"com/routers/auth"
	commentRoute "com/routers/wx/comment"
	"com/routers/wx/upload"
	"com/routers/wx/user"
	"com/routers/wx/wxLogin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(cors.Cors())

	//静态文件服务 提供图片路径
	r.Static("/images", "./public/upload/images")
	//r.Static("/static/images", "./public/images")
	r.Static("/videos", "./public/upload/videos")
	r.Static("/index.html", "./public/dist")


	//微信登录模块
	wx := r.Group("/wx")
	{
		wx.POST("/login",wxLogin.Login)
	}


	test := r.Group("/ttt")
	test.Use(auth.ParseAES)
	{
		test.GET("/a", func(c *gin.Context) {
			value,_ := c.Get("openId")
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"error":  nil,
				"data":   value,
			})
		})
	}

	//登录模块
	userApi := r.Group("/user")
	{
		userApi.POST("/login", user.Login)
		userApi.POST("/registered", user.Registered)
	}


	//上传模块
	uploadApi := r.Group("/upload")
	uploadApi.Use(auth.ParseAuth)
	{
		uploadApi.POST("/video", upload.UploadVideo) //上传视频
		uploadApi.POST("/image", upload.UploadImage) //上传图片
		uploadApi.POST("/audio", upload.UploadAudio) //上传音频
	}


	//评论
	commentApi := r.Group("/comment")
	//commentApi.Use(auth.ParseAuth)
	{
		commentApi.GET("/query", commentRoute.QueryComment)
		commentApi.DELETE("/:id", commentRoute.DeleteComment)
	}
	//回复
	replyApi := r.Group("/reply")
	//replyApi.Use(auth.ParseAuth)
	{
		replyApi.GET("/query", commentRoute.QueryReply)
		replyApi.DELETE("/:id", commentRoute.DeleteReply)
	}

	return r
}
