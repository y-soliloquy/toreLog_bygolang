package controllers

import (
	"log"
	"net/http"
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
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		trainingLogs, _ := user.GetTrainingLogsByUser()
		user.TrainingLogs = trainingLogs
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}
