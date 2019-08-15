package comment

import (
	"com/models/frontend_model"
	"com/models/servser_model"
	"com/models/servser_model/video"
	"fmt"
)

//评论表
type Comment struct {
	servser_model.Model
	VideoId       int                        `gorm:"column:video_id"json:"video_id"validate:"required||integer"` //视频id
	Video         video.Video                `gorm:"ForeignKey:VideoId:AssociationForeignKey:ID"json:"video"`
	Content       string                     `gorm:"column:content"json:"content"validate:"required||integer"`                 //评论内容
	GeneralUserId int                        `gorm:"column:general_user_id"json:"general_user_id"validate:"required||integer"` //评论人的id
	GeneralUser   frontend_model.GeneralUser `gorm:"ForeignKey:GeneralUserId:AssociationForeignKey:ID"json:"general_user"`
	Awesome       *int                       `gorm:"column:awesome"json:"awesome"` //点赞
}

//回复表
type Reply struct {
	servser_model.Model
	CommentId     int                        `gorm:"column:comment_id"json:"comment_id"` //通过评论id可以知道自己属于哪条评论
	Comment       Comment                    `gorm:"ForeignKey:CommentId:AssociationForeignKey:ID"json:"comment"`
	ReplyId       int                        `gorm:"column:reply_id"json:"reply_id"` //回复目标id
	ReplylUser    frontend_model.GeneralUser `gorm:"ForeignKey:ReplyId:AssociationForeignKey:ID"json:"reply_user"`
	Content       string                     `gorm:"column:content"json:"content"`                 //回复目标内容
	GeneralUserId int                        `gorm:"column:general_user_id"json:"general_user_id"` //回复用户id
	GeneralUser   frontend_model.GeneralUser `gorm:"ForeignKey:GeneralUserId:AssociationForeignKey:ID"json:"general_user"`
	Awesome       *int                       `gorm:"column:awesome"json:"awesome"` //点赞
}

type TotalReply struct {
	Replys []Reply `json:"reply"`
	Total  int     `json:"total"`
}

func (this *Reply) QueryReply(commentId string, limit string, offset string) (totalReply TotalReply, err error) {
	if limit == "" {
		limit = "10"
	}
	if offset == "" {
		offset = "0"
	}
	var str = `select reply.id,reply.content,reply.reply_id,reply.comment_id,reply.general_user_id,reply_user.id,reply_user.username,ar.src,general_user.id,general_user.username,avatar.src
from comment inner join reply on
    comment.id = reply.comment_id inner join general_user on
        general_user.id = reply.general_user_id inner join general_user as reply_user on
            general_user.id = reply_user.id inner join avatar on
            avatar.id = general_user.avatar_id inner join avatar as ar on ar.id = reply_user.avatar_id
where comment.id = 1;`
	fmt.Println(str)
	count := servser_model.Db.Raw("select reply.id,reply.content,reply.reply_id,reply.comment_id,reply.general_user_id,general_user.id,general_user.username,avatar.id,avatar.src from comment inner join reply on comment.general_user_id = reply.comment_id inner join general_user on general_user.id = reply.general_user_id inner join avatar on avatar.id = general_user.avatar_id where comment.id = ?", &commentId).Scan(&totalReply.Replys).RowsAffected
	totalReply.Total = int(count)
	totalReply.Replys = nil

	rows, err := servser_model.Db.Raw("select reply.id,reply.content,reply.reply_id,reply.comment_id,reply.general_user_id,general_user.id,general_user.username,avatar.id,avatar.src from comment inner join reply on comment.general_user_id = reply.comment_id inner join general_user on general_user.id = reply.general_user_id inner join avatar on avatar.id = general_user.avatar_id where comment.id = ? limit ? offset ?", &commentId, &limit, &offset).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		reply := Reply{}
		err = rows.Scan(&reply.ID, &reply.Content, &reply.Content, &reply.ReplyId, &reply.CommentId, &reply.GeneralUserId, &reply.GeneralUser.ID, &reply.GeneralUser.UserName, &reply.GeneralUser.Avatar.ID, &reply.GeneralUser.Avatar.Src)
		if err != nil {
			return
		}
		totalReply.Replys = append(totalReply.Replys, reply)
	}
	return
}

type TotalComment struct {
	Comments []Comment `json:"comments"`
	Total    int       `json:"total"`
}

func (this *Comment) QueryComment(condition string, orderBy string, limit string, offset string) (totalComment TotalComment, err error) {
	cond := "%" + condition + "%"
	fmt.Println("cond", cond)
	if limit == "" {
		limit = "10"
	}
	if offset == "" {
		offset = "0"
	}
	count := servser_model.Db.Raw("select comment.id,comment.video_id,video.name,comment.content,comment.general_user_id,general_user.username,avatar.src from video inner join comment on comment.video_id = video.id inner join general_user on comment.general_user_id = general_user.id inner join avatar on general_user.avatar_id = avatar.id where concat(video.name,general_user.username) like ? order by ? Desc", &cond, &orderBy).Scan(&totalComment.Comments).RowsAffected
	totalComment.Total = int(count)
	totalComment.Comments = nil

	rows, err := servser_model.Db.Raw("select comment.id,comment.video_id,video.name,comment.content,comment.general_user_id,general_user.username,avatar.src from video inner join comment on comment.video_id = video.id inner join general_user on comment.general_user_id = general_user.id inner join avatar on general_user.avatar_id = avatar.id where concat(video.name,general_user.username) like ? order by ? Desc limit ? offset ?", &cond, &orderBy, &limit, &offset).Rows()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		comment := Comment{}
		rows.Scan(&comment.ID, &comment.VideoId, &comment.Video.Name, &comment.Content, &comment.GeneralUserId, &comment.GeneralUser.UserName, &comment.GeneralUser.Avatar.Src)
		totalComment.Comments = append(totalComment.Comments, comment)
	}
	return
}

func (this *Comment) DeleteComment(id string) (err error) {
	delete := servser_model.Db.Exec("delete from comment where comment.id = ?", &id)
	if err = delete.Error; err != nil {
		return
	}
	return
}
