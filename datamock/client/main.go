package main

import (
	"context"
	"log"
	"time"

	pb "github.com/JeetKaria06/orchestrator-service/datamock/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:10000"
)

func main() {
	// Setting up a connection to server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDatamockClient(conn)

	// Contacting server and priting its response
	name := "Jeet Karia"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMockUserData(ctx, &pb.NameRequest{Name: name})
	if err != nil {
		log.Fatalf("could not get response because: %v", err)
	} else {
		log.Printf("%s", r)
	}
}
