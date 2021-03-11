package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username"`
	Password string `gorm:"type:varchar(20);not null " json:"password"`
	//权限 `gorm:"type: json:""
	Role int `gorm:"type:int " json:"role"`
	//头像
	//Avatar string
}
