package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/JeetKaria06/orchestrator-service/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50001"
)

type server struct {
	pb.UnimplementedOrchestratorServer
}

func (s *server) GetUserByName(ctx context.Context, in *pb.NameRequest) (*pb.User, error) {
	return nil, errors.New("not implemented yet. Jeet will implement me")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrchestratorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
