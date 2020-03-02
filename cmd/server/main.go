package main

import (
	"fmt"
	"github.com/marcsj/jotfly-server/controller"
	"github.com/marcsj/jotfly-server/notes"
	"github.com/marcsj/jotfly-server/repo"
	"github.com/marcsj/jotfly-server/server"
	"github.com/marcsj/jotfly-server/users"
	"github.com/marcsj/jotfly-server/util"
	"github.com/namsral/flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"os"
	"path/filepath"
)

func init() {
	grpcLog := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(grpcLog)
}

var (
	grpcPort = flag.Int(
		"grpc_port",
		50051,
		"port for gRPC calls")
	mainPath = flag.String(
		"main_path",
		"/var/joyfly",
		"main path for storing jotfly files",
		)
	usersPath = flag.String(
		"users_path",
		"/users",
		"path for storing users",
	)
	notesPath = flag.String(
		"notes_path",
		"/notes",
		"path for storing notes",
	)
)

func main() {
	newKey, err := util.GenerateRandomBytes(util.SaltLength)
	if err != nil {
		log.Fatal(err)
	}
	errChannel := make(chan error)

	usersRepo := repo.NewFileUsersRepo(filepath.Join(*mainPath, *usersPath))
	notesRepo := repo.NewFileNotesRepo(filepath.Join(*mainPath, *usersPath))

	usersController := controller.NewUsersController(usersRepo, newKey)
	notesController := controller.NewNotesController(usersRepo, notesRepo)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *grpcPort))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	usersServer := server.NewUsersServer(usersController)
	notesServer := server.NewNotesServer(notesController)
	users.RegisterUsersServer(grpcServer, usersServer)
	notes.RegisterNotesServer(grpcServer, notesServer)

	log.Println("started joyfly server")

	// running gRPC server
	go func () {
		grpclog.Infof("Starting gRPC server. tcp port: %v", *grpcPort)
		errChannel <- grpcServer.Serve(lis)
	}()

	for {
		select {
		case err := <-errChannel:
			if err != nil {
				log.Fatal(err)
			}
		}
	}


}
