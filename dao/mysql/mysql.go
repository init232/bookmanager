package mysql

import (
	"bookmanager/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	dsn := "root:root@1234@tcp(172.19.145.44:3310)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	//创建表结构
	if err := DB.AutoMigrate(model.Book{}, model.Book{}); err != nil {
		panic(err)
	}

}
