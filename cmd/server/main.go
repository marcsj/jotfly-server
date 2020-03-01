package main

import (
	"github.com/marcsj/jotfly-server/server"
	"log"
)

func main() {
	server.NewNotesServer()
	server.NewUsersServer()
	log.Println("started joyfly server")
}
