package main

import (
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	//CURD
	// mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "ofeifjlsf"}) //增
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "222222222") //改

	var row model.User
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row) //查

	fmt.Printf("row: %+v\n", row)

	//软删
	// mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})
	//强制删除
	//mysql.DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{})

}
