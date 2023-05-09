package model

import "server/global"

type Article struct {
	global.BaseModel
	Title     string `gorm:"type:varchar(200);not null" json:"title" binding:"required"`
	Cover     string `gorm:"type:varchar(400);default:null" json:"cover"`
	Content   string `gorm:"text;not null" json:"content" binding:"required"`
	Status    uint8  `gorm:"default:0; not null" json:"status"`
	Recommond uint8  `gorm:"default:0; not null" json:"recommond"`
	View      uint   `gorm:"default:0; not null" json:"view"`

	Collections []Collection `gorm:"many2many:article_collections;"`
	Tags        []Tag        `gorm:"many2many:article_tags;"`
}

type ArticleMany struct {
	TagIds        []uint
	CollectionIds []uint
}
