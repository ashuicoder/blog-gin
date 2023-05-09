package start

import (
	"fmt"
	"server/global"
	"server/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysql() {
	username := global.Config.Db.Username
	fmt.Printf("username: %v\n", username)
	password := global.Config.Db.Password
	dbname := global.Config.Db.Dbname
	dsn := username + ":" + password + "@tcp(127.0.0.1:3306)/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database ==============")
	}
	fmt.Println("connect database success ====================")

	global.Db = db

	db.AutoMigrate(&model.Article{}, &model.Admin{}, &model.Collection{}, &model.Ad{}, &model.Link{}, &model.Tag{}, &model.Music{})

}
