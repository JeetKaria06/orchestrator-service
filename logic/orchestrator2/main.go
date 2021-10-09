package main

import (
	"context"
	"log"
	"net"
	"time"

	pbDataMock "github.com/JeetKaria06/orchestrator-service/datamock/proto"
	pb "github.com/JeetKaria06/orchestrator-service/logic/proto"
	"google.golang.org/grpc"
)

const (
	port    = ":9001"
	address = "localhost:10000"
)

type server struct {
	pb.UnimplementedRealOrchestratorServer
}

func (s *server) GetUser(ctx context.Context, in *pb.NameRequest) (*pb.User, error) {
	name := in.GetName()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbDataMock.NewDatamockClient(conn)

	ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMockUserData(ctx1, &pbDataMock.NameRequest{Name: name})

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
