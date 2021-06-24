package main

import (
	"context"
	"log"
	"myapp/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewSendMessageClient(conn)

	req := &pb.Request{
		Message: "Hello Gpc",
	}

	res, errReq := client.RequestMessage(context.Background(), req)

	if errReq != nil {
		log.Fatal(errReq)
	}

	log.Print(res.GetStatus())
}
