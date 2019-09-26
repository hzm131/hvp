package routers

import (
	"com/middleware/cors"
	_ "com/models"
	"com/routers/auth"
	"com/routers/servser_router/image_src"
	"com/routers/servser_router/upload"
	videoManagement "com/routers/servser_router/video"
	"com/routers/servser_router/video_src"
	commentRoute "com/routers/wx/comment"
	"com/routers/wx/uploadImg"
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
	uploadApi.Use(auth.GetAuth)
	{
		uploadApi.POST("/video", upload.UploadVideo) //上传视频
		uploadApi.POST("/image", uploadImg.UploadImage) //上传图片
	}

	//评论
	commentApi := r.Group("/comment")
	commentApi.Use(auth.GetAuth)
	{
		commentApi.GET("/query", commentRoute.QueryComment)
		commentApi.DELETE("/:id", commentRoute.DeleteComment)
	}
	//回复
	replyApi := r.Group("/reply")
	replyApi.Use(auth.GetAuth)
	{
		replyApi.GET("/query", commentRoute.QueryReply)
		replyApi.DELETE("/:id", commentRoute.DeleteReply)
	}




	videoApi := r.Group("/video")
	videoApi.Use(auth.GetAuth)
	{
		videoApi.POST("/create", videoManagement.CreateVideo)
		videoApi.GET("/query", videoManagement.QueryVideo)
		videoApi.GET("/query/:id", videoManagement.FindVideo)
		videoApi.PUT("/update/:id", videoManagement.UpdateVideo)
		videoApi.DELETE("/delete/:id", videoManagement.DeleteVideo)

		videoApi.DELETE("/image/delete/:id", image_src.DeleteImageSrc)
		videoApi.DELETE("/video/delete/:id", video_src.DeleteImageSrc)
	}

	return r
}
