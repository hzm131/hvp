package video

import (
	"com/models/servser_model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Video struct {
	gorm.Model
	Path         string  `gorm:"column:path"`
	Name         string  `gorm:"column:name"`
	Pid          int     `gorm:"column:pid"`
	Origin       string  `gorm:"column:origin"`       //产地
	Duration     string  `gorm:"column:duration"`     //时长
	Language     string  `gorm:"column:language"`     //语种
	Years        string  `gorm:"column:years"`        //年份
	Score        float32 `gorm:"column:score"`        //评分
	Introduction string  `gorm:"column:introduction"` //简介

	VideoSrcId int `gorm:"column:video_src_id;"` //视频路径
	ImageSrcId int `gorm:"column:image_src_id;"` //封面路径
}

//上传视频封面
type ImageSrc struct {
	gorm.Model
	SrcPath string `gorm:"column:src_path"`
}


// 上传视频路径
type VideoSrc struct {
	gorm.Model
	SrcPath string `gorm:"column:src_path"`
}


//评论表
type Comment struct {
	gorm.Model
	VideoId int    `gorm:"column:video_id"` //视频id
	Content string `gorm:"column:content"`  //评论内容
	UserId  int    `gorm:"column:user_id"`  //评论人的id
}

//回复表
type Reply struct {
	gorm.Model
	CommentId int    `gorm:"column:comment_id"` //通过评论id可以知道自己属于哪条评论
	ReplyId   int    `gorm:"column:reply_id"`   //回复目标id
	Content   string `gorm:"column:content"`    //回复目标内容
	UserId    int    `gorm:"column:user_id"`    //回复用户id
}


func (this *VideoSrc) CreatedVideoSrc()(int,error){
	fmt.Printf("this",this)
	path := VideoSrc{}
	find := servser_model.Db.Create(this).Scan(&path)
	if err:=find.Error; err!=nil{
		fmt.Println("创建失败",err)
		return 0,err
	}
	//创建成功后返回id
	id := int(path.ID)  //拿到id
	return id,nil
}

func (this *ImageSrc) CreatedImageSrc()(int,error){
	fmt.Printf("this",this)
	path := ImageSrc{}
	find := servser_model.Db.Create(this).Scan(&path)
	if err:=find.Error; err!=nil{
		fmt.Println("创建失败",err)
		return 0,err
	}
	//创建成功后返回id
	id := int(path.ID)  //拿到id
	return id,nil
}