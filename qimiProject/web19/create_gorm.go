package main

import "fmt"

func createTest() {
	dbres := GolbalDb.Create(&[]TestUser{
		{Name: "第四", Age: 29},
		{Name: "第五", Age: 15},
		{Name: "第六", Age: 20},
	})
	fmt.Println(dbres.Error, dbres.RowsAffected)
	if dbres.Error != nil {
		fmt.Println("创建失败")
	} else {
		fmt.Println("创建成功")
	}
}
