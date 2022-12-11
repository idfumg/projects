package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/idfumg/go-grpc-course-1/greet/proto"
)

var addr string = "localhost:50051"
var useTLS bool = true

func GenClientOpts(tls bool) []grpc.DialOption {
	opts := []grpc.DialOption{}
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Failed loading CA Turst certificates: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	return opts
}

func main() {
	conn, err := grpc.Dial(addr, GenClientOpts(!useTLS)...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	log.Printf("Connected to: %s\n", addr)

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	doSum(c)
	doGreetManyTimes(c)
	doPrimes(c)
	doLongGreet(c)
	doGreetEveryone(c)
	doSqrt(c, -16)
	doGreetWithDeadline(c, 4*time.Second)
	doGreetWithDeadline(c, 1*time.Second)
}
