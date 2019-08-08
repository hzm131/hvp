package video

import "github.com/jinzhu/gorm"

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
