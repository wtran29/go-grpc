package main

import (
	"context"
	"log"

	pb "github.com/wtran29/go-grpc/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg: %v\n", err)
	}

	nums := []int32{1, 2, 3, 4}
	for _, num := range nums {
		log.Printf("Sending number: %d\n", num)
		stream.Send(&pb.AvgRequest{
			Num: num,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg: %v\n", res.Result)
}
