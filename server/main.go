package main

import (
	"fmt"
	"log"
	"net"
	"server/services"
	"os"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	port := os.Getenv("PORT")
	listener,err := net.Listen("tcp",":"+port)
	if err != nil{
		log.Fatal(err)
	}

	services.RegisterActivityServer(s, services.NewActivityServer())

	fmt.Println("gRPC server listening on port "+port)
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}