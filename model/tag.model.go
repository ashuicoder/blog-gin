package model

import "server/global"

type Tag struct {
	global.BaseModel
	Name   string `gorm:"type:varchar(200);not null;unique;" json:"name" binding:"required"`
	Icon   string `gorm:"type:varchar(400);default:null" json:"icon"`
	Status uint8  `gorm:"default:0; not null" json:"status"`

	Articles []*Article `gorm:"many2many:article_tags;"`
}
