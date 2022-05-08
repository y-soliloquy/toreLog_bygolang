package main

import (
	"fmt"
	"torelog_bygolang/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// ログ
	// log.Println("test")

	// ユーザーのテーブルが作れているか
	// fmt.Println(models.Db)

	// ユーザーを作成する
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// 作成されたユーザー情報を取得する
	u, _ := models.GetUser(1)
	fmt.Println(u)

	// ユーザー情報を更新する
	u.Name = "test2"
	u.Email = "test2@example.com"
	u.UpdateUser()

	u, _ = models.GetUser(1)
	fmt.Println(u)
}
