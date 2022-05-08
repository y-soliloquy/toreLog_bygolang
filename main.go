package main

import (
	"fmt"
	"log"
	"torelog_bygolang/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	// ログ
	log.Println("test")
}
