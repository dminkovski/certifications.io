package main

import (
	_ "github.com/dminkovski/certifications.io/api/controller"
	_ "github.com/dminkovski/certifications.io/api/database"
	_ "github.com/dminkovski/certifications.io/api/model"
	"github.com/dminkovski/certifications.io/api/server"
)

func main() {
	server := server.Server{Port: ":8080"}
	server.Start()
}
