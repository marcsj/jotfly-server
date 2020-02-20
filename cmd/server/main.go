package main

import (
	"github.com/marcsj/jotfly-server"
	"log"
)

func main() {
	jotfly.NewNotesServer()
	log.Println("started joyfly server")
}
