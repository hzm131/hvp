package upload

import (
	"com/models/wx/upload"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)


func UploadAudio(c *gin.Context) {
	value,_ := c.Get("user")

	//获取用户id
	openid,err := MapInter(value)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "",
		})
		return
	}

	//判断文件夹是否已经存在
	bool,err := PathExists("public/upload/audios/" + openid)
	if err != nil || bool == false{
		//不存在就创建目录
		err = os.Mkdir("public/upload/audios/" + openid, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 400,
				"error":  err,
				"data":   "创建目录失败",
			})
			return
		}
	}

	//获取上传的信息
	file, header, err := c.Request.FormFile("audio")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "错误请求",
		})
		return
	}
	//文件的名称
	filename := header.Filename

	//判断文件是否已经存在
	bool,err = PathExists("public/upload/audios/"+openid+"/"+filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "未知",
		})
		return
	}
	if bool == true {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  nil,
			"data":   "文件名已存在",
		})
		return
	}

	//创建空文件
	out, err := os.Create("public/upload/audios/" + openid+"/"+filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "文件创建失败",
		})
		return
	}
	defer out.Close()
	// 将文件复制进空文件
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "复制文件失败",
		})
		return
	}
	str := "http://192.168.2.166:3000/audios/"+ openid+"/" + filename
	src := upload.Video{
		Src: str,
		Title:filename,
	}
	video, err := src.CreatedVideo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "添加数据库失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   video,
	})
}
