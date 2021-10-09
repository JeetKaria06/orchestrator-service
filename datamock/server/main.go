package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	pb "github.com/JeetKaria06/orchestrator-service/datamock/proto"
	"google.golang.org/grpc"
)

const (
	port = ":10000"
)

type server struct {
	pb.UnimplementedDatamockServer
}

func (s *server) GetMockUserData(ctx context.Context, in *pb.NameRequest) (*pb.User, error) {
	name := in.GetName()
	if len(name) < 6 {
		return nil, errors.New("length of name is less than 6")
	}
	return &pb.User{Name: name, Class: strconv.Itoa(len(name)), Roll: int64(len(name) * 10)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDatamockServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
