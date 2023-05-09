package model

import "server/global"

type Music struct {
	global.BaseModel
	Title  string `gorm:"type:varchar(200);not null" json:"title" binding:"required"`
	Desc   string `gorm:"text;default:null" json:"desc"`
	Cover  string `gorm:"type:varchar(400);default:null" json:"cover"`
	Status uint8  `gorm:"default:0; not null" json:"status"`
	Link   string `json:"link" binding:"required"`
}
