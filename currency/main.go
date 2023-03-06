package main

import (
	"net"
	"os"
	"projects/go-microservices/currency/server"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protos "projects/go-microservices/currency/protos/currency"
)

func main() {
	log := hclog.Default()

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// create an instance of the Currency server
	c := server.NewCurrency(log)

	// register the currency server
	protos.RegisterCurrencyServer(gs, c)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC server
	// disable in production
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	// listen for requests
	gs.Serve(l)
}
