package configs

import (
	"prakerja3/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:110597@tcp(127.0.0.1:3306)/prakerja_batch_3?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connect database failed")
	}
	migratDatabase()
}

func migratDatabase() {
	DB.AutoMigrate(&models.Karyawan{})
}
