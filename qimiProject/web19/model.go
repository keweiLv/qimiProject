package main

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type TestUser struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
}

func TestUserCreate() {
	err := GolbalDb.AutoMigrate(&TestUser{})
	if err != nil {
		return
	}
}
