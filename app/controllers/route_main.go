package controllers

import (
	"log"
	"net/http"
	"torelog_bygolang/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "navbar", "top")
	} else {
		http.Redirect(w, r, "/trainingLog", 302)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	println("ログアウトした１")
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		println("ログアウトした２")
		log.Fatalln(err)
	}
	if err != http.ErrNoCookie {
		println("ログアウトした３")
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	println("ログアウトした４")

	http.Redirect(w, r, "/login", 302)
}
