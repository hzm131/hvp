package users

import (
	"com/models/servser_model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	UserName string `gorm:"column:username"validate:"required||string"`
	PassWord string `gorm:"column:password"validate:"required||string"`
}

func (this *Users) FindId() (int, error) {
	//根据用户名 密码查询用户 将查询到的结果封装在user结构中
	var user Users
	query := servser_model.Db.Raw("select id from users where username=? and password=? limit 1", this.UserName, this.PassWord).Scan(&user)
	if err := query.Error; err != nil {
		fmt.Println("用户名或密码有问题", err)
		return 0, err
	}
	id := int(user.ID)
	return id, nil
}

func (this *Users) CreateData() (int, error) {
	user := Users{}
	servser_model.Db.Raw("select id from users where username=?", this.UserName).Scan(&user)
	if user.ID > 0 {
		fmt.Println("用户名已存在")
		return -1, nil
	}
	idUser := Users{}
	db := servser_model.Db.Create(this).Scan(&idUser)
	if err := db.Error; err != nil {
		fmt.Println("创建失败")
		return 0, err
	}
	id := int(idUser.ID)
	fmt.Println("id", idUser.ID)
	return id, nil
}
