package main

import (
	"context"
	"log"

	pb "github.com/wtran29/go-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Calculate func was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.A + in.B,
	}, nil
}
