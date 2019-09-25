package comment

import (
	"com/models/frontend_model/general_user"
	"com/models/servser_model/video"
	"com/models/wx"
	"fmt"
)

//评论表
type Comment struct {
	wx.Model
	VideoId       int                      `gorm:"column:video_id"json:"video_id"validate:"required||integer"` //视频id
	Video         video.Video              `gorm:"ForeignKey:VideoId:AssociationForeignKey:ID"json:"video"`
	Content       string                   `gorm:"column:content"json:"content"validate:"required||integer"`                 //评论内容
	GeneralUserId int                      `gorm:"column:general_user_id"json:"general_user_id"validate:"required||integer"` //评论人的id
	GeneralUser   general_user.GeneralUser `gorm:"ForeignKey:GeneralUserId:AssociationForeignKey:ID"json:"general_user"`
	Awesome       *int                     `gorm:"column:awesome"json:"awesome"` //点赞
}

//回复表
type Reply struct {
	wx.Model
	CommentId     int                      `gorm:"column:comment_id"json:"comment_id"` //通过评论id可以知道自己属于哪条评论
	Comment       Comment                  `gorm:"ForeignKey:CommentId:AssociationForeignKey:ID"json:"comment"`
	ReplyId       int                      `gorm:"column:reply_id"json:"reply_id"` //回复目标id
	ReplylUser    general_user.GeneralUser `gorm:"ForeignKey:ReplyId:AssociationForeignKey:ID"json:"reply_user"`
	Content       string                   `gorm:"column:content"json:"content"`                 //回复目标内容
	GeneralUserId int                      `gorm:"column:general_user_id"json:"general_user_id"` //回复用户id
	GeneralUser   general_user.GeneralUser `gorm:"ForeignKey:GeneralUserId:AssociationForeignKey:ID"json:"general_user"`
	Awesome       *int                     `gorm:"column:awesome"json:"awesome"` //点赞
}

type TotalReply struct {
	Replys []Reply `json:"replys"`
	Total  int     `json:"total"`
}

func (this *Reply) QueryReply(commentId string, limit string, offset string) (totalReply TotalReply, err error) {
	if limit == "" {
		limit = "10"
	}
	if offset == "" {
		offset = "0"
	}
	count := wx.Db.Raw(`select reply.id,reply.content,reply.reply_id,reply.comment_id,reply.general_user_id,reply.awesome,reply.created_at,reply_user.id,reply_user.username,ar.id,ar.src,general_user.id,general_user.username,avatar.id,avatar.src
                     from comment inner join reply on
                             comment.id = reply.comment_id inner join general_user on
                             general_user.id = reply.general_user_id inner join general_user as reply_user on
                             reply_user.id = reply.reply_id inner join avatar on
                             avatar.id = general_user.avatar_id inner join avatar as ar on ar.id = reply_user.avatar_id
                     where comment.id = ? `, &commentId).Scan(&totalReply.Replys).RowsAffected
	totalReply.Total = int(count)
	totalReply.Replys = nil

	rows, err := wx.Db.Raw(`select reply.id,reply.content,reply.reply_id,reply.comment_id,reply.general_user_id,reply.awesome,reply.created_at,reply_user.id,reply_user.username,ar.id,ar.src,general_user.id,general_user.username,avatar.id,avatar.src
                     from comment inner join reply on
                             comment.id = reply.comment_id inner join general_user on
                             general_user.id = reply.general_user_id inner join general_user as reply_user on
                             reply_user.id = reply.reply_id inner join avatar on
                             avatar.id = general_user.avatar_id inner join avatar as ar on ar.id = reply_user.avatar_id
                     where comment.id = ? limit ? offset ?`, &commentId, &limit, &offset).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		reply := Reply{}
		err = rows.Scan(&reply.ID, &reply.Content, &reply.ReplyId, &reply.CommentId, &reply.GeneralUserId, &reply.Awesome, &reply.CreatedAt, &reply.ReplylUser.ID, &reply.ReplylUser.UserName, &reply.ReplylUser.Avatar.ID, &reply.ReplylUser.Avatar.Src, &reply.GeneralUser.ID, &reply.GeneralUser.UserName, &reply.GeneralUser.Avatar.ID, &reply.GeneralUser.Avatar.Src)
		if err != nil {
			return
		}
		totalReply.Replys = append(totalReply.Replys, reply)
	}
	return
}

func (this *Reply) DeleteReply(id string) (err error) {
	delete := wx.Db.Exec("delete from reply where id = ?", &id)
	if err = delete.Error; err != nil {
		return
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
	count := wx.Db.Raw("select comment.id,comment.created_at,comment.awesome,comment.video_id,video.name,comment.content,comment.general_user_id,general_user.username,avatar.src from video inner join comment on comment.video_id = video.id inner join general_user on comment.general_user_id = general_user.id inner join avatar on general_user.avatar_id = avatar.id where concat(video.name,general_user.username) like ? order by ? Desc", &cond, &orderBy).Scan(&totalComment.Comments).RowsAffected
	totalComment.Total = int(count)
	totalComment.Comments = nil

	rows, err := wx.Db.Raw("select comment.id,comment.created_at,comment.awesome,comment.video_id,video.name,comment.content,comment.general_user_id,general_user.username,avatar.src from video inner join comment on comment.video_id = video.id inner join general_user on comment.general_user_id = general_user.id inner join avatar on general_user.avatar_id = avatar.id where concat(video.name,general_user.username) like ? order by ? Desc limit ? offset ?", &cond, &orderBy, &limit, &offset).Rows()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		comment := Comment{}
		rows.Scan(&comment.ID, &comment.CreatedAt, &comment.Awesome, &comment.VideoId, &comment.Video.Name, &comment.Content, &comment.GeneralUserId, &comment.GeneralUser.UserName, &comment.GeneralUser.Avatar.Src)
		totalComment.Comments = append(totalComment.Comments, comment)
	}
	return
}

func (this *Comment) DeleteComment(id string) (err error) {
	delete := wx.Db.Exec("delete comment,reply from comment inner join reply on comment.id = reply.comment_id where comment.id = ?", &id)
	if err = delete.Error; err != nil {
		return
	}
	return
}
