package main

import (
	"context"
	"log"

	pb "github.com/wtran29/go-grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doCalculate was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 3,
		B: 10,
	})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("Calculating: %v\n", res.Result)
}
