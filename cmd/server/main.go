package main

import (
	"context"
	"log"
	"myapp/pb"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedSendMessageServer
}

func (service *Server) RequestMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Print("mensage recive : ", req.GetMessage())

	response := &pb.Response{
		Status: 1,
	}

	return response, nil
}
func (service *Server) mustEmbedUnimplementedSendMessageServer() {}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterSendMessageServer(grpcServer, &Server{})

	port := ":5000"

	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	grpcError := grpcServer.Serve(listener)

	if grpcError != nil {
		log.Fatal(grpcError)
	}

}
