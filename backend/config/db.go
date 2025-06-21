package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=Phong16052004@@ dbname=telecom_db port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Kết nối DB thất bại")
	}
	DB = database
	fmt.Println("✅ Đã kết nối PostgreSQL")
}
