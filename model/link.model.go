package model

import "server/global"

type Link struct {
	global.BaseModel
	Title  string `gorm:"type:varchar(200);not null" json:"title" binding:"required"`
	Icon   string `gorm:"type:varchar(400);default:null" json:"icon"`
	Link   string `json:"link" binding:"required"`
	Desc   string `gorm:"text;default:null" json:"desc"`
	Status uint8  `gorm:"default:0; not null" json:"status"`
	View   uint   `gorm:"default:0; not null" json:"view"`
}
