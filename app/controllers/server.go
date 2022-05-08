package controllers

import (
	"net/http"
	"torelog_bygolang/config"
)

// サーバーの設定
func StartMainServer() (err error) {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
