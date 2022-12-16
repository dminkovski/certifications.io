package server

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"os"
)

type Server struct {
	Port string
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



func (server Server) Start() {
	fmt.Println("Starting server at port:", server.Port)
	http.HandleFunc("/assets/img/", Images)
	log.Panic(http.ListenAndServe(server.Port, nil))
}
