package main

import (
	"demo/hello_demo/biz/dal"
	"demo/hello_demo/biz/dal/mysql"
	"demo/hello_demo/biz/model"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/pingcap/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	// result := mysql.DB.Create(&model.User{Email: "test@test.com", Password: "123456"})
	result := mysql.DB.Model((&model.User{})).Where("id = ?", 1).Update("password", "new_password")
	if result.Error != nil {
		log.Error(result.Error.Error())
	}
	fmt.Println(result.RowsAffected)

	var row model.User
	result = mysql.DB.First(&row, 1)
	if result.Error != nil {
		log.Error(result.Error.Error())
	}
	fmt.Println(row)
}
