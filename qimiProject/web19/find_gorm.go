package main

import "fmt"

func TestFind() {
	//result := map[string]interface{}{}
	//dbRes := GolbalDb.Model(&TestUser{}).First(&result, 11)
	//fmt.Println(errors.Is(dbRes.Error, gorm.ErrRecordNotFound))

	var User TestUser
	GolbalDb.Where("name = ?", "第二").First(&User)
	fmt.Println(User)
}
