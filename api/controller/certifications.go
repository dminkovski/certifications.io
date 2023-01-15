package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/dminkovski/certifications.io/api/model"
)

var tpl *template.Template

func GetCertifications(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		certifications := model.LoadCertifications()
		response, err := json.Marshal(certifications)
		if err != nil {
			http.Error(w, "JSON representation of the object was not possible.", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	} else {
		http.Error(w, "Only GET allowed", http.StatusBadRequest)
	}
}

func Index(w http.ResponseWriter, req *http.Request) {
	certifications := model.LoadCertifications()
	err := tpl.ExecuteTemplate(w, "certifications.html", struct {
		Certifications []model.Certification
	}{
		certifications,
	})
	if err != nil {
		log.Panic(err)
	}
}

func GetCreateCertification(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		err := tpl.ExecuteTemplate(w, "create-certification.html", nil)
		if err != nil {
			log.Panic(err)
		}
	}
}

func GetAndPostCertificationById(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		GetCertificationById(w, req)
	} else if req.Method == "POST" {
		PostCreateCertification(w, req)
	}
}

func GetCertificationById(w http.ResponseWriter, req *http.Request) {

	re := regexp.MustCompile(`\d+$`)
	url := req.URL.String()
	if re.Match([]byte(url)) {
		idString := re.FindString(url)
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Id not of type number", http.StatusBadRequest)
		} else {
			data := map[string]string{
				"response": fmt.Sprintf("Congrats %d", id),
			}
			response, err := json.Marshal(data)
			if err != nil {
				http.Error(w, "JSON representation of the object was not possible.", http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
			w.Write(response)
		}
	} else {
		http.Error(w, "Id not provided", http.StatusBadRequest)
	}

}

func PostCreateCertification(w http.ResponseWriter, req *http.Request) {
	var c model.Certification
	err := DecodeJsonBody(w, req, &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		var mr *MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Error(), mr.Status())
		} else {
			log.Panic(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	} else {
		response, err := json.Marshal(c)
		if err != nil {
			http.Error(w, "JSON representation of the object was not possible.", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

}

func init() {
	tpl = template.Must(template.ParseGlob("templates/**.html"))
	http.HandleFunc("/", Index)
	http.HandleFunc("/certifications/create", GetCreateCertification)
	http.HandleFunc("/api/certification", GetAndPostCertificationById)
	http.HandleFunc("/api/certifications", GetCertifications)
}
