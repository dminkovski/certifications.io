package main

import (
	_ "github.com/dminkovski/certifications.io/controller"
	_ "github.com/dminkovski/certifications.io/model"
	"github.com/dminkovski/certifications.io/server"
)

func main() {
	server := server.Server{":8080"}
	server.Start()
}
