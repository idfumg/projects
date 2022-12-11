package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/idfumg/go-grpc-course-1/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Artem",
	})

	if err != nil {
		log.Fatalf("Could not read: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.GetResult())
}

func doSum(c pb.GreetServiceClient) {
	log.Printf("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		X: 1,
		Y: 2,
	})

	if err != nil {
		log.Fatalf("Could not read: %v\n", err)
	}

	log.Printf("Sum: %d\n", res.GetResult())
}

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Clement",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.GetResult())
	}
}

func doPrimes(c pb.GreetServiceClient) {
	log.Printf("doPrimes was invoked")

	req := &pb.PrimesRequest{
		Value: 120,
	}

	stream, err := c.GetPrimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GetPrimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Receiving prime: %d\n", msg.Result)
	}
}

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongRead was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Marie"},
		{FirstName: "Peter"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Error while sending message to the stream: %v\n", err)
		}
		time.Sleep(500 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response from the stream: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.GetResult())
}

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Printf("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GreetEveryone: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Marie"},
		{FirstName: "Peter"},
	}

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			err := stream.SendMsg(req)
			if err != nil {
				log.Fatalf("Error sending a request: %v\n", err)
			}
			time.Sleep(1 * time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Error closing the stream: %v\n", err)
		}
	}()

	for {
		res := pb.GreetResponse{}
		err := stream.RecvMsg(&res)

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error receiving a response: %v\n", err)
		}

		log.Printf("Received: %v\n", res.GetResult())
	}
}

func doSqrt(c pb.GreetServiceClient, n int64) {
	log.Printf("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from the server: %s\n", e.Message())
			log.Printf("Error code from the server: %s\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Printf("We probably sent a negative number (InvalidArgument)\n")
			}
			log.Printf("Sqrt response: nil\n")
			return
		} else {
			log.Fatalf("Error calling Sqrt rpc function: %v\n", err)
		}
	}

	log.Printf("Sqrt response: %f\n", res.GetResult())
}

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Printf("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Clement",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline exceeded!")
				log.Printf("GreetWithDeadline: nil\n")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error occuried: %v\n", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.GetResult())
}
