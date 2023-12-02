package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"time"
)

var db = NewDB()

func HandleTime(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/time.html", "templates/base.html")
	t.Execute(w, time.Now())
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	var err error = nil

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		if password, ok := db.Load(username); ok && password == r.FormValue("password") {
			http.Redirect(w, r, "/time", http.StatusSeeOther)
			return
		}
		err = errors.New("Неправильный логин или пароль")
	}

	t, _ := template.ParseFiles("templates/index.html", "templates/base.html")
	t.Execute(w, err)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleIndex)
	mux.HandleFunc("/time", HandleTime)

	db.Store("user1", "password1")
	db.Store("user2", "password2")

	log.Println("server started on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
