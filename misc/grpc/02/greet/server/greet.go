package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"time"

	pb "github.com/idfumg/go-grpc-course-1/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.GreetServiceServer
}

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %+v\n", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.GetFirstName(),
	}, nil
}

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with: %+v\n", in)

	return &pb.SumResponse{
		Result: in.GetX() + in.GetY(),
	}, nil
}

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.GetFirstName(), i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}

func (s *Server) GetPrimes(in *pb.PrimesRequest, stream pb.GreetService_GetPrimesServer) error {
	log.Printf("GetPrimes was invoked")

	primes := func() []int64 {
		ans := []int64{}
		k := int64(2)
		v := in.GetValue()
		for v > 1 {
			for v%k == 0 {
				ans = append(ans, k)
				v /= k
			}
			k++
		}
		return ans
	}()

	for i := 0; i < len(primes); i++ {
		stream.Send(&pb.PrimesResponse{
			Result: primes[i],
		})
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading from the stream: %v\n", err)
		}

		if len(res) != 0 {
			res += " "
		}
		res += req.GetFirstName()
	}
}

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.GetFirstName()
		err = stream.SendMsg(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to the client: %v\n", err)
		}
	}
}

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt was invoked")

	num := in.GetNumber()

	if num < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative number: %d", num),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(num)),
	}, nil
}

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked")

	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("The client cancelled the request")
			return nil, status.Error(codes.Canceled, "The client cancelled the request")
		} else {
			return nil, status.Error(codes.Unknown, "The unknown context error occuried")
		}
	case <-time.After(3 * time.Second):
		break
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.GetFirstName(),
	}, nil
}
