package main

import (
	"context"
	"log"

	pb "github.com/idfumg/go-grpc-course-2/blog/proto"
)

func DeleteBlog(c pb.BlogServiceClient, id string) {
	log.Printf("DeleteBlog was invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Printf("Blog was deleted!")
}
