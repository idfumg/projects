package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"myapp/server/services"
)

var (
	port = flag.Int("port", 50051, "port number to listen to")
	tls = flag.Bool("tls", false, "use a secure TLS connection")
)

func NewServer(tls bool) *grpc.Server {
	if tls {
		certFile := "./tls/server.crt"
		keyFile := "./tls/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatal(err)
		}
		return grpc.NewServer(grpc.Creds(creds))
	} else {
		return grpc.NewServer()
	}
}

func main() {
	flag.Parse()

	port := fmt.Sprintf(":%d", *port)

	s := NewServer(*tls)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s, services.NewCalculatorServer())
	reflection.Register(s)

	fmt.Print("gRPC server listening on port ", port)
	if *tls {
		fmt.Println(" with TLS")
	} else {
		fmt.Println(" without TLS")
	}
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
