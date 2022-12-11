package main

import (
	pb "github.com/idfumg/go-grpc-course-2/blog/proto"
)

type Server struct {
	pb.BlogServiceServer
}
