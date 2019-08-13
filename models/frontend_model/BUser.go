package frontend_model

import "com/models/servser_model"

type BUsers struct {
	servser_model.Model
	UserName string `gorm:"column:username"validate:"required||string"`
	PassWord string `gorm:"column:password"validate:"required||string"`
}
