package model

import "server/global"

type Collection struct {
	global.BaseModel
	Title  string `gorm:"type:varchar(200);not null" json:"title" binding:"required"`
	Cover  string `gorm:"type:varchar(400);default:null" json:"cover"`
	Desc   string `gorm:"default:null" json:"desc"`
	Status uint8  `gorm:"default:0; not null" json:"status"`
	View   uint   `gorm:"default:0; not null" json:"view"`

	Articles []*Article `gorm:"many2many:article_collections;"`
}
