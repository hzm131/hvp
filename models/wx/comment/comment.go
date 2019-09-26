package comment

import (
	"com/models/wx"
	"com/models/wx/article"
	"com/models/wx/wxuser"
	"fmt"
	"time"
)

//评论表
type Comment struct {
	wx.Model
	ArticleId     int                      `gorm:"column:article_id"json:"article_id"validate:"required||integer"` //视频id
	Article       article.Article          `gorm:"ForeignKey:ArticleId:AssociationForeignKey:ID"json:"article"`
	Content       string                   `gorm:"column:content"json:"content"validate:"required||integer"`                 //评论内容
	GeneralOpenId int                      `gorm:"column:general_openid"json:"general_openid"validate:"required||integer"` //评论人的id
	GeneralWxUser wxuser.WxUser  		   `gorm:"ForeignKey:GeneralOpenId:AssociationForeignKey:OpenId"json:"OpenId"`
	Awesome       *int                     `gorm:"column:awesome"json:"awesome"` //点赞
}





type QueryComment struct {
	ID        int       `json:"id"`
	CreatedAt *time.Time  `gorm:"created_at"json:"created_at"`
	Content       string   `json:"content"`
	GeneralOpenId int 	`json:"general_openid"`
	Awesome       *int     `json:"awesome"`
	ArticleId 	int    `json:"article_id"`
	Title 	string    `json:"title"`
	NickName 	string    `json:"nickName"`
}

type TotalComment struct {
	Comments []QueryComment `json:"comments"`
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
	count := wx.Db.Raw("select comment.id,comment.created_at,comment.content,comment.general_openid,comment.awesome,comment.article_id,article.title,wx_user.nickName from wx_user inner join comment on comment.general_openid = wx_user.openid inner join article on comment.article_id = article.id  where concat(article.title,wx_user.nickName) like ? order by ? Desc", &cond, &orderBy).Scan(&totalComment.Comments).RowsAffected
	totalComment.Total = int(count)
	totalComment.Comments = nil

	rows, err := wx.Db.Raw("select comment.id,comment.created_at,comment.content,comment.general_openid,comment.awesome,comment.article_id,article.title,wx_user.nickName from wx_user inner join comment on comment.general_openid = wx_user.openid inner join article on comment.article_id = article.id  where concat(article.title,wx_user.nickName) like ? order by ? Desc limit ? offset ?", &cond, &orderBy, &limit, &offset).Rows()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		comment := QueryComment{}
		rows.Scan(&comment.ID, &comment.CreatedAt, &comment.Content, &comment.GeneralOpenId, &comment.Awesome, &comment.ArticleId, &comment.Title, &comment.NickName)
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



//回复表
type Reply struct {
	wx.Model
	CommentId     int                      `gorm:"column:comment_id"json:"comment_id"` //通过评论id可以知道自己属于哪条评论
	Comment       Comment                  `gorm:"ForeignKey:CommentId:AssociationForeignKey:ID"json:"comment"`
	ReplyOpenId      int                   `gorm:"column:reply_openid"json:"reply_openid"` //回复目标id
	ReplylWxUser  wxuser.WxUser   		   `gorm:"ForeignKey:ReplyId:AssociationForeignKey:OpenId"json:"reply_wxuser"`
	Content       string                   `gorm:"column:content"json:"content"`                 //回复目标内容
	GeneralOpenId int                      `gorm:"column:general_openid"json:"general_openid"` //回复用户id
	GeneralWxUser wxuser.WxUser  		   `gorm:"ForeignKey:GeneralOpenId:AssociationForeignKey:OpenId"json:"general_wxuser"`
	Awesome       *int                     `gorm:"column:awesome"json:"awesome"` //点赞
}



type QueryReply struct {
	ID        int       	`json:"id"`
	CommentId     int     	`son:"comment_id"` //通过评论id可以知道自己属于哪条评论
	CreatedAt *time.Time  	`json:"created_at"`
	Content       string   	`json:"content"`
	ReplyOpenId   int  		`json:"reply_openid"`
	RyNickName string 		`json:"ryNickName"`
	GeneralOpenId  int   	`json:"general_openid"`
	GlNickName string 		`json:"glNickName"`
	Awesome       *int     	`json:"awesome"`
}

type TotalReply struct {
	Replys []QueryReply `json:"replys"`
	Total  int     `json:"total"`
}

func (this *Reply) QueryReply(commentId string, limit string, offset string) (totalReply TotalReply, err error) {
	if limit == "" {
		limit = "10"
	}
	if offset == "" {
		offset = "0"
	}
	count := wx.Db.Raw(`select reply.id,
       reply.content,
       reply.reply_openid,
       reply.comment_id,
       reply.awesome,
       reply.created_at,
       ry.nickName,
       reply.general_openid,
       gl.nickName
from reply inner join comment on
        reply.comment_id = comment.id inner join wx_user as ry on
        ry.openid = reply.reply_openid inner join wx_user as gl on
        reply.general_openid = gl.openid
where comment.id = ?`, &commentId).Scan(&totalReply.Replys).RowsAffected
	totalReply.Total = int(count)
	totalReply.Replys = nil

	rows, err := wx.Db.Raw(`select reply.id,
       reply.content,
       reply.reply_openid,
       reply.comment_id,
       reply.awesome,
       reply.created_at,
       ry.nickName,
       reply.general_openid,
       gl.nickName
from reply inner join comment on
        reply.comment_id = comment.id inner join wx_user as ry on
        ry.openid = reply.reply_openid inner join wx_user as gl on
        reply.general_openid = gl.openid
where comment.id = ? limit ? offset ?`, &commentId, &limit, &offset).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		reply := QueryReply{}
		err = rows.Scan(&reply.ID, &reply.Content, &reply.ReplyOpenId, &reply.CommentId, &reply.Awesome, &reply.CreatedAt, &reply.RyNickName, &reply.GeneralOpenId, &reply.GlNickName)
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


