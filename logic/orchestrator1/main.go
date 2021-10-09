package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/JeetKaria06/orchestrator-service/logic/proto"
	"google.golang.org/grpc"
)

const (
	port    = ":9000"
	address = "localhost:9001"
)

type server struct {
	pb.UnimplementedRealOrchestratorServer
}

func (s *server) GetUserByName(ctx context.Context, in *pb.NameRequest) (*pb.User, error) {
	name := in.GetName()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRealOrchestratorClient(conn)

	ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx1, &pb.NameRequest{Name: name})

	return &pb.User{Name: r.GetName(), Class: r.GetClass(), Roll: r.GetRoll()}, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRealOrchestratorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
