package controllers

import (
	"log"
	"net/http"
	"torelog_bygolang/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "navbar", "top")
	} else {
		http.Redirect(w, r, "/trainingLogs", 302)
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

func trainingLogNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "trainingLog_new")
	}
}

func trainingLogSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		content := r.PostFormValue("content")
		satisfaction := r.PostFormValue("satisfaction")
		weather := r.PostFormValue("weather")
		if err := user.CreateTrainingLog(content, satisfaction, weather); err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, "/trainingLogs", 302)
	}
}

func trainingLogEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		t, err := models.GetTrainingLog(id)
		if err != nil {
			log.Fatalln(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "trainingLog_edit")
	}
}

func trainingLogUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		content := r.PostFormValue("content")
		satisfaction := r.PostFormValue("satisfaction")
		weather := r.PostFormValue("weather")
		t := &models.TrainingLog{ID: id, Content: content, Satisfaction: satisfaction, Weather: weather, UserID: user.ID}
		if err := t.UpdateTrainingLog(); err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, "/trainingLogs", 302)
	}
}

func trainingLogDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		l, err := models.GetTrainingLog(id)
		if err := l.DeleteTrainingLog(); err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, "/trainingLogs", 302)
	}
}
