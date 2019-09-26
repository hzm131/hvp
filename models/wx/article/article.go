package article

import (
	"com/models/wx"
	"com/models/wx/wxuser"
)

type Article struct {
	wx.Model
	Content  string `gorm:"column:content"json:"content"` //存储富文本的html片段
	Title  string `gorm:"column:title"json:"title"` //标题
	Category  string `gorm:"column:category"json:"category"`  //类别
	OpenId int `gorm:"column:openid"json:"openid"`
	WxUser wxuser.WxUser `gorm:"ForeignKey:OpenId;AssociationForeignKey:OpenId"json:"wxUser"`

}
