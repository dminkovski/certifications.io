package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("./templates/**.html"))
	http.HandleFunc("/", GetDashboard)
	fmt.Println("Initialized Admin Dashboard")
}

func GetDashboard(w http.ResponseWriter, r *http.Request) {
	data := "test"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err := templates.ExecuteTemplate(w, "dashboard.html", data)
	if err != nil {
		log.Panic("Problem with Templating")
	}
}
