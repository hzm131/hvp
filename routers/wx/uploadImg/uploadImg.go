package uploadImg

import (
	image2 "com/models/wx/image"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
)





func UploadImage(c *gin.Context) {
	value,_ := c.Get("user")

	//获取用户id
	id,err := MapInter(value)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "",
		})
		return
	}
	intId := int(id)
	strId := strconv.Itoa(intId)

	//判断文件夹是否已经存在
	bool,err := PathExists("public/upload/images/" + strId)
	if err != nil || bool == false{
		//不存在就创建目录
		err = os.Mkdir("public/upload/images/" + strId, os.ModePerm)
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

	//判断文件是否已经存在
	bool,err = PathExists("public/upload/images/"+strId+"/"+filename)
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
	out, err := os.Create("public/upload/images/" + strId+"/"+filename)
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
	str := "http://192.168.2.166:3000/images/"+ strId+"/" + filename
	src := image2.Image{
		Src: str,
		Title:filename,
	}
	image, err := src.CreatedImage()
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

func MapInter (value interface{}) (id float64,err error){
	v,ok := value.(map[string]interface{})
	if !ok {
		err = errors.New("user数据不是map[string]interface{}")
		return
	}
	for key,val := range v {
		if key == "id"{
			fmt.Println("v1 type:", reflect.TypeOf(val))
			int := val.(float64)
			id = int
		}
	}
	return
}