package main

import (
	_ "github.com/dminkovski/goblog/model"
	"github.com/dminkovski/goblog/server"
)

func main() {
	server := server.Server{":8080"}
	server.Start()
}
