package main

import (
	"flag"
	"log"
	"myapp/client/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	host = flag.String("host", "localhost:50051", "host to connect to")
	tls  = flag.Bool("tls", false, "use a secure TLS connection")
)

func GetCredentials(tls bool) credentials.TransportCredentials {
	if tls {
		certFile := "./tls/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatal(err)
		}
		return creds
	} else {
		return insecure.NewCredentials()
	}
}

func GetClientConn(tls bool, host string) *grpc.ClientConn {
	creds := GetCredentials(tls)
	cc, err := grpc.Dial(host, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	return cc
}

func main() {
	flag.Parse()

	cc := GetClientConn(*tls, *host)
	defer cc.Close()

	client := services.NewCalculatorClient(cc)
	service := services.NewCalculatorService(client)

	err := service.Hello("Bob")
	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			log.Fatalf("[%v] %v\n", grpcErr.Code(), grpcErr.Message())
		} else {
			log.Fatal(err)
		}
	}

	// err = service.Fibonacci(3)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = service.Average(1, 2, 3, 4, 5)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = service.Sum(1, 2, 3, 4, 5)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
