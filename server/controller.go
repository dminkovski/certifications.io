package server

import (
	"net/http"
	"log"
	"fmt"
    "encoding/json"
)

func Login(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	u := req.FormValue("username")
	p := req.FormValue("password")

	fmt.Printf("Login with %v %v", u, p)
	data := make(map[string]string)
	data["username"] = u
	data["password"] = p
	response, err := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}