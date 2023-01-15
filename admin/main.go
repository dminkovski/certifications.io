package main

import (
	_ "github.com/dminkovski/certifications.io/admin/controller"
	"log"
	"net/http"
)

func main() {
	log.Panic(http.ListenAndServe(":8080", nil))
}
