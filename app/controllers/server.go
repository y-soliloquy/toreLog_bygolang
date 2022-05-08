package controllers

import (
	"net/http"
	"torelog_bygolang/config"
)

// サーバーの設定
func StartMainServer() (err error) {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
