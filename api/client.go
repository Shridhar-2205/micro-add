package main

import (
	"context"
	"fmt"
	"log"

	pb "api/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewAddServiceClient(conn)

	req := &pb.AddRequest{A: 10, B: 2}
	if res, err := client.Compute(context.Background(), req); err == nil {
		fmt.Println(res.Result)
	}
}
