package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/wtran29/go-grpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax func is invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		nums := []int32{4, 7, 2, 19, 4, 6, 32}

		for _, num := range nums {
			log.Printf("Sending number: %v\n", num)
			stream.Send(&pb.MaxRequest{
				Num: num,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Println("Problem while reading server stream: %v\n", err)
				break
			}
			log.Printf("Received a new maximum: %d\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
