package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/idfumg/go-grpc-course-2/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := CreateBlog(c)
	log.Printf("Created a blog with id: %s\n", id)

	blog, _ := ReadBlog(c, id)
	log.Printf("Read a blog: %v\n", blog)

	blog, _ = ReadBlog(c, "invalid")
	log.Printf("Read a blog: %v\n", blog)

	UpdateBlog(c, id)
	ListBlogs(c)
	DeleteBlog(c, id)
}