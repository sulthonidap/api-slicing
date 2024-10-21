package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Init() {
	// Use the public URL for connecting to the database
	dsn := "root:usAqtgbnkocGEpuLvBMjBcBDnbzztfFs@tcp(autorack.proxy.rlwy.net:24422)/railway?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err)
	}

	log.Println("Connected to the database successfully:", dsn)

	DBConn = db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Posting{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Like{})
}
