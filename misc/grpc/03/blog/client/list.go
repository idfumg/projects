package main

import (
	"context"
	"io"
	"log"

	pb "github.com/idfumg/go-grpc-course-2/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ListBlogs(c pb.BlogServiceClient) ([]*pb.Blog, error) {
	log.Printf("ListBlogs was invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		return nil, err
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}

	return nil, nil
}
