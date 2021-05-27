package main

import (
	"context"
	"fmt"
	"grpc-demo/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewCalculatorServiceClient(conn)
	addRequest := &proto.CalculatorRequest{X: 100, Y: 200}
	addResponse, err := client.Add(context.Background(), addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Result = %d\n", addResponse.GetResult())

}
