package grpc

import (
	"context"
	"log"
	"myapp/internal/rocket"
	"net"

	"google.golang.org/grpc"

	pb "myapp/internal/proto/rocket/v1"
)

type RocketService interface {
	GetRocket(ctx context.Context, id string) (rocket.Rocket, error)
	AddRocket(ctx context.Context, rocketRocket rocket.Rocket) (rocket.Rocket, error)
	DelRocket(ctx context.Context, id string) error
}

type Handler struct {
	pb.UnimplementedRocketServiceServer
	RocketService RocketService
}

func NewHandler(rocketService RocketService) Handler {
	return Handler{
		RocketService: rocketService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("could not listen on port 50051")
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("Failed to serve: %s\n", err)
		return err
	}

	return nil
}

func (h Handler) GetRocket(ctx context.Context, req *pb.GetRocketRequest) (*pb.GetRocketResponse, error) {
	rocket, err := h.RocketService.GetRocket(ctx, req.Id)
	if err != nil {
		log.Print("Failed to retrieve a rocket by id")
		return &pb.GetRocketResponse{}, err
	}
	return &pb.GetRocketResponse{
		Rocket: &pb.Rocket{
			Id:   rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

func (h Handler) AddRocket(ctx context.Context, req *pb.AddRocketRequest) (*pb.AddRocketResponse, error) {
	newRkt, err := h.RocketService.AddRocket(ctx, rocket.Rocket{
		ID:   req.Rocket.Id,
		Type: req.Rocket.Type,
		Name: req.Rocket.Name,
	})
	if err != nil {
		return &pb.AddRocketResponse{}, nil
	}
	return &pb.AddRocketResponse{
		Rocket: &pb.Rocket{
			Id:   newRkt.ID,
			Type: newRkt.Type,
			Name: newRkt.Name,
		},
	}, nil
}

func (h Handler) DelRocket(ctx context.Context, req *pb.DelRocketRequest) (*pb.DelRocketResponse, error) {
	err := h.RocketService.DelRocket(ctx, req.Id)
	if err != nil {
		return &pb.DelRocketResponse{}, err
	}
	return &pb.DelRocketResponse{
		Status: "Rocket was deleted",
	}, nil
}
