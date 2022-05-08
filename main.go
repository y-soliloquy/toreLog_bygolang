package main

import (
	"fmt"
	"torelog_bygolang/app/controllers"
	"torelog_bygolang/app/models"
)

func main() {
	fmt.Println(models.Db)
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// ログ
	// log.Println("test")

	/**************************************************************************************/
	//	ユーザー周り																		//
	/**************************************************************************************/

	// ユーザーのテーブルが作れているか
	// fmt.Println(models.Db)

	// ユーザーを作成する
	// u := &models.User{}
	// u.Name = "test3"
	// u.Email = "test3@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// 作成されたユーザー情報を取得する
	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// ユーザー情報を更新する
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()

	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	//　ユーザー情報を削除する
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	/**************************************************************************************/
	//	トレーニングログ周り																　//
	/**************************************************************************************/

	// トレーニングログを作成する
	// user, _ := models.GetUser(3)
	// user.CreateTrainingLog("トレログ3-2")

	// トレーニングログを取得する（単数）
	// t, _ := models.GetTrainingLog(1)
	// fmt.Println(t)

	// トレーニングログを取得する（複数）
	// trainingLogs, _ := models.GetTrainingLogs()
	// for _, v := range trainingLogs {
	// 	fmt.Println(v)
	// }

	// ユーザーを絞ってトレーニングログを取得する
	// user2, _ := models.GetUser(3)
	// trainingLogs, _ := user2.GetTrainingLogsByUser()
	// for _, v := range trainingLogs {
	// 	fmt.Println(v)
	// }

	// トレーニングログ情報を更新する
	// t, _ := models.GetTrainingLog(1)
	// t.Content = "更新トレログ"
	// t.UpdateTrainingLog()

	// トレーニングログ情報を削除する
	// t, _ := models.GetTrainingLog(3)
	// t.DeleteTrainingLog()

	/**************************************************************************************/
	//	サーバー周り																　//
	/**************************************************************************************/

	// サーバーの立ち上げ
	controllers.StartMainServer()
}
