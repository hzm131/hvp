package models

import (
	"com/models/servser_model"
	"com/models/servser_model/users"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
	此包的其他文件 去定义表 以及表关系 操作表的方法封装
*/
var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:h891453@tcp(127.0.0.1:3306)/hzm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//创建表 自动迁移
	Db.AutoMigrate(&users.Users{})

	servser_model.ModelInit(Db)
}

func CloseDB() {
	defer Db.Close()
}
