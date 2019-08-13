package users

import (
	"com/models/servser_model"
	"fmt"
)

type Users struct {
	servser_model.Model
	UserName string `gorm:"column:username"json:"username"validate:"required||string"`
	PassWord string `gorm:"column:password"json:"password"validate:"required||string"`
}

func (this *Users) FindId() (Users, error) {
	//根据用户名 密码查询用户 将查询到的结果封装在user结构中
	user := Users{}
	query := servser_model.Db.Raw("select * from users where username=? and password=? limit 1", this.UserName, this.PassWord).Scan(&user)
	if err := query.Error; err != nil {
		fmt.Println("用户名或密码有问题", err)
		return user, err
	}
	user.PassWord = "你猜猜"
	fmt.Println("user:", user)
	return user, nil
}

func (this *Users) CreateData() (Id int, user Users, err error) {
	userId := Users{}
	servser_model.Db.Raw("select id from users where username=?", this.UserName).Scan(&userId)
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
