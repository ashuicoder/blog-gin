package model

import "server/global"

type Ad struct {
	global.BaseModel
	Title  string `gorm:"type:varchar(200);not null" json:"title" binding:"required"`
	Cover  string `gorm:"type:varchar(400);default:null" json:"cover"`
	Link   string `json:"link" binding:"required"`
	Status uint8  `gorm:"default:0; not null" json:"status"`
	View   uint   `gorm:"default:0; not null" json:"view"`
}
