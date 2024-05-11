package controllers

import (
	models "hackaton/model"
	"html/template"
	"net/http"
)

func FAQHandler(w http.ResponseWriter, r *http.Request) {
	var Connected PackageInfo
	cookie, _ := r.Cookie("User")
	if cookie == nil {
		Connected.Connected = false
		println("Connected = false")
	} else {
		Connected.Connected = true
		println("Connected = true")
	
	}

	if r.URL.Path != "/FAQ" {
		models.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("./view/FAQ.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, Connected)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
