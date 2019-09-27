package upload

import (
	"com/models/wx"
	"fmt"
)

type Audio struct {
	wx.Model
	OpenId string `gorm:"column:openid"json:"openid"`
	Title string  `gorm:"column:title"json:"title"`
	Src string `gorm:"column:src"json:"src"`
}

func (this *Image) CreatedAudio() (audio Audio,err error) {
	find := wx.Db.Create(this).Scan(&audio)
	if err = find.Error; err != nil {
		fmt.Println("创建失败", err)
		return
	}
	return
}
