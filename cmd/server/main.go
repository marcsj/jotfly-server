package main

import (
	"github.com/marcsj/jotfly-server"
	"log"
)

func main() {
	jotfly.NewNotesServer()
	jotfly.NewUsersServer()
	log.Println("started joyfly server")
}
