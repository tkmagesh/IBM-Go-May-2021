package main

import (
	"context"
	"grpc-demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Add(ctx context.Context, req *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	res := &proto.CalculatorResponse{Result: result}
	return res, nil
}

func (s *server) Subtract(ctx context.Context, req *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x - y
	res := &proto.CalculatorResponse{Result: result}
	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(grpcServer, &server{})

	e := grpcServer.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
}
