package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	dsn = "root:15926653820@Hu@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
)

var DB *gorm.DB

func MysqlInit() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql open err : ", err)
	}

	DB = db
}
