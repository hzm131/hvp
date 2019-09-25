package article

import (
	"com/models/wx"
	"com/models/wx/user"
)

type Article struct {
	wx.Model
	Content  string `gorm:"column:content"json:"content"` //存储富文本的html片段
	Title  string `gorm:"column:title"json:"title"` //标题
	Category  string `gorm:"column:category"json:"category"`  //类别
	OpenId int `gorm:"column:openid"json:"openid"`
	WxUser user.WxUser `gorm:"ForeignKey:OpenId;AssociationForeignKey:OpenId"json:"wxUser"`

}
