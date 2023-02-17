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

	"github.com/dminkovski/certifications.io/api/database"
	"github.com/dminkovski/certifications.io/api/model"
	"github.com/dminkovski/certifications.io/api/utils"
)

var tpl *template.Template

func Index(w http.ResponseWriter, req *http.Request) {
	certifications := make([]model.Certification, 0)
	err := tpl.ExecuteTemplate(w, "certifications.html", struct {
		Certifications []model.Certification
	}{
		certifications,
	})
	if err != nil {
		log.Panic(err)
	}
}

// Gets Certifications from DB or local JSON
func GetCertifications(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		certifications := database.GetCertifications()
		response, err := json.Marshal(certifications)
		if err != nil {
			http.Error(w, "JSON representation of the object was not possible.", http.StatusInternalServerError)
		}
		utils.PrepareResponse(w, response)
	} else {
		http.Error(w, "Only GET allowed", http.StatusBadRequest)
	}
}

// Handle Certification Route GET and POST
func GetAndPostCertificationById(w http.ResponseWriter, req *http.Request) {
	utils.PrepareResponse(w, nil)

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
			utils.PrepareResponse(w, response)
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
			fmt.Println(mr.Error())
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
		err = database.SaveCertification(c)
		if err != nil {
			http.Error(w, "Saving to local JSON file was not possible.", http.StatusInternalServerError)
		}
		utils.PrepareResponse(w, response)
	}

}

func AddCourse(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var c model.CourseDTO
		err := DecodeJsonBody(w, req, &c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		course := c.Course
		certId := c.CertificationId
		err = database.SaveCourse(course, certId)
		if err != nil {
			msg := fmt.Sprintf("Adding course to %d not possible.", certId)
			http.Error(w, msg, http.StatusInternalServerError)
		}
		response, err := json.Marshal(c)
		if err != nil {
			http.Error(w, "JSON representation of the object was not possible.", http.StatusInternalServerError)
			return
		}
		utils.PrepareResponse(w, response)
	}
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/**.html"))
	http.HandleFunc("/", Index)
	http.HandleFunc("/api/certification", GetAndPostCertificationById)
	http.HandleFunc("/api/course", AddCourse)
	http.HandleFunc("/api/certifications", GetCertifications)
}
