package general_user

import (
	"com/models/servser_model"
	"fmt"
)

type GeneralUser struct {
	servser_model.Model
	UserName string `gorm:"column:username"json:"username"validate:"required||string"`
	PassWord string `gorm:"column:password"json:"password"validate:"required||string"`
	Birthday *string `gorm:"column:birthday"json:"birthday"` //生日
	AvatarId int `gorm:"column:avatar_id"json:"avatar_id"` //头像
	Avatar Avatar `gorm:"ForeignKey:AvatarId;AssociationForeignKey:ID"json:"avatar"`
	Address *string `gorm:"column:address"json:"address"` //地址
	Email *string `gorm:"column:email"json:"email"`
	Phone *string `gorm:"column:phone"json:"phone"`
	Vip int  `gorm:"column:vip;default:0"json:"vip"`
}

type Avatar struct {
	servser_model.Model
	Src string `gorm:"column:src"json:"src"`
}


func (this *GeneralUser) FindId() (GeneralUser, error) {
	//根据用户名 密码查询用户 将查询到的结果封装在user结构中
	user := GeneralUser{}
	query := servser_model.Db.Raw("select * from general_user where username=? and password=? limit 1", this.UserName, this.PassWord).Scan(&user)
	if err := query.Error; err != nil {
		fmt.Println("用户名或密码有问题", err)
		return user, err
	}
	user.PassWord = "你猜猜"
	fmt.Println("user:", user)
	return user, nil
}

func (this *GeneralUser) CreateData() (Id int, user GeneralUser, err error) {
	userId := GeneralUser{}
	servser_model.Db.Raw("select id from general_user where username=?", this.UserName).Scan(&userId)
	if Id = int(userId.ID); Id > 0 {
		fmt.Println("用户名已存在")
		return
	}
	db := servser_model.Db.Create(this).Scan(&user)
	if err = db.Error; err != nil {
		fmt.Println("创建失败")
		return
	}
	return
}