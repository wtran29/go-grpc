package main

import (
	"log"

	pb "github.com/wtran29/go-grpc/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes func was invoked with: %v\n", in)

	k := int64(2)
	n := in.X
	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			n /= k
		} else {
			k++
		}
	}
	return nil
}
