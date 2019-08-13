package comment

import (
	"com/models/servser_model"
	"com/models/servser_model/users"
	"com/models/servser_model/video"
)

//评论表
type Comment struct {
	servser_model.Model
	VideoId int         `gorm:"column:video_id"json:"video_id"` //视频id
	Video   video.Video       `gorm:"ForeignKey:VideoId:AssociationForeignKey:ID"json:"video"`
	Content string      `gorm:"column:content"json:"content"` //评论内容
	UserId  int         `gorm:"column:user_id"json:"user_id"` //评论人的id
	User    users.Users `gorm:"ForeignKey:UserId:AssociationForeignKey:ID"`
}

//回复表
type Reply struct {
	servser_model.Model
	CommentId int     `gorm:"column:comment_id"json:"comment_id"` //通过评论id可以知道自己属于哪条评论
	Comment   Comment `gorm:"ForeignKey:CommentId:AssociationForeignKey:ID"json:"comment"`
	ReplyId   int     `gorm:"column:reply_id"json:"reply_id"` //回复目标id
	Content   string  `gorm:"column:content"json:"content"`   //回复目标内容
	UserId    int     `gorm:"column:user_id"json:"user_id"`   //回复用户id
}

type TotalComment struct {
	Comments []Comment `json:"comments"`
	Total  int     `json:"total"`
}

func (this *Comment) QueryComment(id string)(totalComment TotalComment,err error){
	query := servser_model.Db.Raw("select * from comment inner join video on video.id = comment.video_id inner join users on comment.user_id = users.id where video.id = ?",&id).Scan(&totalComment.Comments)
	if err = query.Error; err != nil{
		return
	}
	return
}

func (this *Comment) DeleteComment(id string)(bool bool,err error){
	delete := servser_model.Db.Exec("delete from comment where comment.id = ?",&id)
	if err = delete.Error; err != nil{
		return false, err
	}
	return true,nil
}