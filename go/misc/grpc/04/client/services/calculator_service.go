package services

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type CalculatorService interface {
	Hello(name string) error
	Fibonacci(n uint32) error
	Average(numbers ...float64) error
	Sum(numbers ...int32) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) CalculatorService {
	return calculatorService{
		calculatorClient: calculatorClient,
	}
}

func (base calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name:        name,
		CreatedDate: timestamppb.Now(),
	}
	res, err := base.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}
	fmt.Printf("Service: Hello\n")
	fmt.Printf("Request: %v\n", req.Name)
	fmt.Printf("Response: %v\n", res.Result)
	return nil
}

func (base calculatorService) Fibonacci(n uint32) error {
	req := FibonacciRequest{
		N: n,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := base.calculatorClient.Fibonacci(ctx, &req)
	if err != nil {
		return err
	}

	fmt.Printf("Service: Fibonacci\n")
	fmt.Printf("Request: %v\n", req.N)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Response: %v\n", res.Result)
	}
	return nil
}

func (base calculatorService) Average(numbers ...float64) error {
	stream, err := base.calculatorClient.Average(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Service: Average\n")
	for _, number := range numbers {
		req := AverageRequest{
			Number: number,
		}
		stream.Send(&req)
		fmt.Printf("Request: %v\n", req.Number)
		time.Sleep(time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Printf("Response: %v\n", res.Result)
	return nil
}

func (base calculatorService) Sum(numbers ...int32) error {
	stream, err := base.calculatorClient.Sum(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Service: Sum\n")
	done := make(chan struct{})
	errs := make(chan error)
	wg := sync.WaitGroup{}
	wg.Add(2)
	wasError := false
	go func() {
		defer wg.Done()
		for _, number := range numbers {
			req := SumRequest{
				Number: number,
			}
			err := stream.Send(&req)
			if err != nil {
				errs <- err
				wasError = true
				return
			}
			fmt.Printf("Request: %v\n", req.Number)
		}
		err := stream.CloseSend()
		if err != nil {
			errs <- err
			wasError = true
			return
		}

	}()
	go func() {
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				errs <- err
				wasError = true
				return
			}
			fmt.Printf("Response: %v\n", res.Result)
		}
	}()
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()
	select {
	case <-done:
		if !wasError {
			return nil
		}
	case err := <-errs:
		return err
	}
	return nil
}
