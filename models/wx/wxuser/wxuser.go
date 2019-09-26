package wxuser

import (
	"com/models/wx"
	"fmt"
)

type WxUser struct {
	wx.Model
	Code string `gorm:"column:code"json:"code"`
	NickName string `gorm:"column:nickName"json:"nickName"`
	AvatarUrl string `gorm:"column:avatarUrl"json:"avatarUrl"`
	SessionKey string `gorm:"column:session_key"json:"session_key"`
	OpenId string `gorm:"column:openid"json:"openid"`
	Language string `gorm:"column:language"json:"language"`
	City string `gorm:"column:city"json:"city"` //市
	Province string `gorm:"column:province"json:"province"` //省
	Country string `gorm:"column:country"json:"country"` //国家
	Gender *int `gorm:"column:gender"json:"gender"` //性别
}

func (this *WxUser) CreateData() (wxUser WxUser, err error) {
	openId := WxUser{}
	wx.Db.Raw("select id from wx_user where openid = ?", this.OpenId).Scan(&openId)
	if Id := int(openId.ID); Id > 0 {
		fmt.Println("openid已存在")
		db := wx.Db.Exec("UPDATE wx_user set code = ?, nickName = ?, avatarUrl = ?, session_key = ?,openid = ?, language = ?, city = ?, province = ?, country = ?, gender = ?  where id = ?",this.Code,this.NickName,this.AvatarUrl,this.SessionKey,this.OpenId,this.Language,this.City,this.Province,this.Country,this.Gender,Id)
		if err = db.Error;err != nil{
			fmt.Println("更新")
			return
		}
		db2 := wx.Db.Raw("select * from wx_user where id = ?", Id).Scan(&wxUser)
		if err = db2.Error; err != nil{
			fmt.Println("查询")
			return
		}
	}else{
		db := wx.Db.Create(this).Scan(&wxUser)
		if err = db.Error; err != nil {
			fmt.Println("创建失败")
			return
		}
	}
	return
}

func (this *WxUser) FindOpenId() (wxUser WxUser, err error) {
	query := wx.Db.Raw("select * from wx_user where openid = ?", this.OpenId).Scan(&wxUser)
	if err = query.Error;err != nil{
		return
	}
	return
}