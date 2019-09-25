package uploadImg

import (
	"com/models/wx/image"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
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

	bool,err := PathExists("public/upload/images/"+filename)
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
	out, err := os.Create("public/upload/images/" + filename)
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
	//封面表添加
	str := "http://127.0.0.1:3000/images/" + filename
	src := image.Image{
		Src: str,
		Title:filename,
	}
	image, err := src.CreatedImage()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "封面添加数据库失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   image,
	})
}

func PathExists(path string) (bool,error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false,nil
	}
	return false, err
}