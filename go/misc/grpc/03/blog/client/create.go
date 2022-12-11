package main

import (
	"context"
	"log"

	pb "github.com/idfumg/go-grpc-course-2/blog/proto"
)

func CreateBlog(c pb.BlogServiceClient) string {
	log.Printf("CreateBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Clement",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.GetId())
	return res.GetId()
}
