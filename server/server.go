package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/dminkovski/goblog/model"
	"regexp"
	"os"
)

type Server struct {
	Port string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/**.html"))
}

func Images(w http.ResponseWriter, req *http.Request){
	re := regexp.MustCompile(`[a-zA-Z]+\.(png|jpg)\b`)
	tr := regexp.MustCompile(`(png|jpg)\b`)
	img := re.FindString(req.URL.String())
	ty := tr.FindString(img)
	path := fmt.Sprintf("./assets/%v",string(img))
	file, err := os.ReadFile(string(path))
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", fmt.Sprintf("image/%s",ty))
	w.Write(file)
}

func index(w http.ResponseWriter, req *http.Request) {
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

func (server Server) Start() {
	fmt.Println("Starting server at port:", server.Port)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/", index)
	http.HandleFunc("/assets/img/", Images)
	log.Panic(http.ListenAndServe(server.Port, nil))
}
