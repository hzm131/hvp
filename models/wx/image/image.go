package image

import (
	"com/models/wx"
	"fmt"
)

type Image struct {
	wx.Model
	OpenId string `gorm:"column:openid"json:"openid"`
	Title string  `gorm:"column:title"json:"title"`
	Src string `gorm:"column:src"json:"src"`
}

func (this *Image) CreatedImage() (image Image,err error) {
	find := wx.Db.Create(this).Scan(&image)
	if err = find.Error; err != nil {
		fmt.Println("创建失败", err)
		return
	}
	return
}