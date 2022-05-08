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
	fmt.Println(models.Db)
}
