package main

import (
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	listener,err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatal(err)
	}

	services.RegisterActivityServer(s, services.NewActivityServer())

	fmt.Println("gRPC server listening on port 8080")
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}