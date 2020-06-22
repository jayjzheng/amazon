package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jayjzheng/amazon/auth/pb"
	"google.golang.org/grpc"
)

var (
	port int
)

func main() {
	flag.IntVar(&port, "port", 3000, "port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})

	s.Serve(lis)
}

type server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *server) CreateUser(context.Context, *pb.User) (*pb.CreateUserResponse, error) {
	return nil, nil
}

func (s *server) ChangePassword(context.Context, *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	return nil, nil
}

func (s *server) CreateToken(context.Context, *pb.User) (*pb.CreateTokenResponse, error) {
	return nil, nil
}
