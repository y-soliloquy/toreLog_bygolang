package controllers

import "net/http"

func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "navbar", "signup")
}
