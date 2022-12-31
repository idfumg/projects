package services

import (
	context "context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}

func (c calculatorServer) mustEmbedUnimplementedCalculatorServer() {

}

func (c calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if req.Name == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"name is required",
		)
	}
	return &HelloResponse{
		Result: fmt.Sprintf("Hello %v at %v", req.Name, req.CreatedDate.AsTime().Local()),
	}, nil
}

func (c calculatorServer) Fibonacci(req *FibonacciRequest, stream Calculator_FibonacciServer) error {
	for n := uint32(0); n <= req.N; n++ {
		result := fib(n)
		res := FibonacciResponse{
			Result: result,
		}
		stream.Send(&res)
		time.Sleep(time.Second)
	}
	return nil
}

func fib(n uint32) uint32 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

func (c calculatorServer) Average(stream Calculator_AverageServer) error {
	sum := 0.0
	count := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count += 1
		fmt.Printf("req.Number: %f\n", req.Number)
		fmt.Printf("sum: %f\n", sum)
	}
	res := AverageResponse{
		Result: sum / count,
	}
	return stream.SendAndClose(&res)
}

func (c calculatorServer) Sum(stream Calculator_SumServer) error {
	sum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		res := SumResponse{
			Result: sum,
		}
		err = stream.Send(&res)
		if err != nil {
			return err
		}
	}
	return nil
}
