package main

import (
	"context"
	"log"

	pb "github.com/idfumg/go-grpc-course-2/blog/proto"
)

func ReadBlog(c pb.BlogServiceClient, id string) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
