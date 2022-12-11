package main

import (
	"context"
	"log"
	"net"

	pb "github.com/idfumg/go-grpc-course-2/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	collection *mongo.Collection
	addr       string = "0.0.0.0:50051"
	useTLS     bool   = true
)

// Use if we want to enable SSL feature.
func GenServerOpts(tls bool) []grpc.ServerOption {
	opts := []grpc.ServerOption{}
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}
	return opts
}

func CreateMongo() *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("Error creating the mongo client: %v\n", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to the mongo: %v\n", err)
	}

	return (*mongo.Collection)(client.Database("blogdb").Collection("blog"))
}

func main() {
	collection = CreateMongo()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer(GenServerOpts(!useTLS)...)
	pb.RegisterBlogServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
