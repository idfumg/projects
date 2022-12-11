package test

import (
	"log"
	pb "myapp/internal/proto/rocket/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetClient() pb.RocketServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	rocketClient := pb.NewRocketServiceClient(conn)
	return rocketClient
}
