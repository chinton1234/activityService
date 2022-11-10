package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"server/services"

	"google.golang.org/grpc"

	"fmt"
	"log"
	"net/http"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

func serviceRegistryWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Println(err)
	}

	serviceID := "activity-service1"
	port, _ := strconv.Atoi(getPort()[1:len(getPort())])
	address := getHostname()

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "activity-service",
		Port:    port,
		Address: getHostname(),
		Check: &consulapi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/check", address, port),
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	regiErr := consul.Agent().ServiceRegister(registration)

	if regiErr != nil {
		log.Printf("Failed to register service: %s:%v ", address, port)
	} else {
		log.Printf("successfully register service: %s:%v", address, port)
	}
}

func check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Consul check")
}

func getPort() (port string) {
	port = os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	port = ":" + port
	return
}

func getHostname() (hostname string) {
	hostname, _ = os.Hostname()
	return
}

func main() {
	s := grpc.NewServer()

	port := os.Getenv("PORT")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterActivityServer(s, services.NewActivityServer())

	fmt.Println("gRPC server listening on port " + port)
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
