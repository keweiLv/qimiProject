package main

import "gorm.io/gorm"

type Dog struct {
	gorm.Model
	Name      string
	GirlGodID uint
	GirlGod   GirlGod
}

type GirlGod struct {
	gorm.Model
	Name string
}

func one2one() {
	g := GirlGod{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "美美",
	}
	d := Dog{
		Model: gorm.Model{
			ID: 1,
		},
		Name:    "皮皮",
		GirlGod: g,
	}
	//GolbalDb.AutoMigrate(&Dog{});
	GolbalDb.Create(&d)
}
