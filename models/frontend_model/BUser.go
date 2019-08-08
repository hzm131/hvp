package frontend_model

import "github.com/jinzhu/gorm"

type BUsers struct {
	gorm.Model
	UserName string `gorm:"column:username"validate:"required||string"`
	PassWord string `gorm:"column:password"validate:"required||string"`
}
