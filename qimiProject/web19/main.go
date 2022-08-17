package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var GolbalDb *gorm.DB

func main() {
	//dsn := "root:Kezi520@tcp(1.14.161.232:3306)/goFunc?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:Kezi520@tcp(1.14.161.232:3306)/goFunc?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 191,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: false, // 逻辑外键
	})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	GolbalDb = db
	TestUserCreate()
	//createTest()
	//TestFind()
	one2one()
}
