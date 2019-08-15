package models

import (
	"com/models/frontend_model/general_user"
	"com/models/servser_model"
	"com/models/servser_model/comment"
	"com/models/servser_model/users"
	"com/models/servser_model/video"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
	此包的其他文件 去定义表 以及表关系 操作表的方法封装
*/
var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/hzm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	Db.SingularTable(true)
	//创建表 自动迁移
	Db.AutoMigrate(&users.Users{},
	&video.Video{},
	&comment.Comment{},
	&comment.Reply{},
	&video.ImageSrc{},
	&video.VideoSrc{},
	&general_user.GeneralUser{},
	&general_user.Avatar{})

	servser_model.ModelInit(Db)
}

func CloseDB() {
	defer Db.Close()
}
