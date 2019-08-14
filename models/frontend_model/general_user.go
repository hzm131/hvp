package frontend_model

import "com/models/servser_model"

type GeneralUser struct {
	servser_model.Model
	UserName string `gorm:"column:username"json:"username"validate:"required||string"`
	PassWord string `gorm:"column:password"json:"password"validate:"required||string"`
	Birthday *string `gorm:"column:birthday"json:"birthday"` //生日
	AvatarId int `gorm:"column:avatar_id"json:"avatar_id"` //头像
	Avatar Avatar `gorm:"ForeignKey:AvatarId;AssociationForeignKey:ID"json:"avatar"`
	Address *string `gorm:"column:address"json:"address"` //地址
	Email *string `gorm:"column:email"json:"email"`
}

type Avatar struct {
	servser_model.Model
	Src string `gorm:"column:src"json:"src"`
}