package utils

import (
	"PromisedLandLab/statics"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	dsn := "BatryCC:BatryCC@tcp(198.211.44.35:3306)/batrycc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("无法连接到数据库")
	}

	db.AutoMigrate(&statics.User{})
	return db
}
