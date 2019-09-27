package upload

import (
	"com/models/wx"
	"fmt"
)

type Video struct {
	wx.Model
	OpenId string `gorm:"column:openid"json:"openid"`
	Title string  `gorm:"column:title"json:"title"`
	Src string `gorm:"column:src"json:"src"`
}

func (this *Video) CreatedVideo() (video Video,err error) {
	find := wx.Db.Create(this).Scan(&video)
	if err = find.Error; err != nil {
		fmt.Println("创建失败", err)
		return
	}
	return
}