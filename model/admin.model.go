package model

import (
	"server/global"
)

type Admin struct {
	global.BaseModel
	Username  string `gorm:"type:varchar(200);not null;unique" json:"username" binding:"required"`
	Passworld string `gorm:"not null;" json:"password" binding:"required"`
}

type AdminLogin struct {
	Username  string `json:"username" binding:"required"`
	Passworld string `json:"password" binding:"required"`
}
