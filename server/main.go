package main

import (
	"fmt"
	"log"

	"github.com/daniilperestoronin/go-grpc-vue-template/server"
)

func main() {
	go func() {
		grpcServer := server.GrpcServer{
			Address:    fmt.Sprintf("%s:%d", "localhost", 7777),
			WebAddress: fmt.Sprintf("%s:%d", "localhost", 7771),
		}
		if err := grpcServer.Serve(); err != nil {
			panic("Failed start grpc service")
		}
	}()

	go func() {
		oauth2Server := server.OAuth2Server{}

		if err := oauth2Server.Serve(); err != nil {
			panic("Failed start oauth service")
		}
	}()

	go func() {
		staticServer := server.StaticServer{}

		if err := staticServer.Serve(); err != nil {
			panic("Failed start oauth service")
		}
	}()

	// infinite loop
	log.Printf("Entering infinite loop")
	select {}
}
