package main

import (
	"github.com/marcsj/jotfly-server/server"
	"log"
)

func main() {
	//TODO: create notes controller, validation checking, server integration
	//TODO: more testing
	//TODO: make flags setup an admin account, with password
	//TODO: make integration tests and do everything important
	//TODO: add things to the readme on how to run, what the purpose of jotfly is

	server.NewNotesServer()
	server.NewUsersServer()
	log.Println("started joyfly server")
}
