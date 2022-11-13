package main

import (
	"context"
	"log"

	pb "github.com/idfumg/go-grpc-course-2/blog/proto"
)

func UpdateBlog(c pb.BlogServiceClient, in string) {
	log.Printf("UpdateBlog was invoked")

	req := &pb.Blog{
		Id:       in,
		AuthorId: "Clement",
		Title:    "A new title",
		Content:  "Content of the first blog with some useful additions!",
	}

	_, err := c.UpdateBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed while updating a blog: %v\n", err)
	}

	log.Printf("Blog was updated")
}
