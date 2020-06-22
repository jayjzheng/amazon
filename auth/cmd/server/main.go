package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jayjzheng/amazon/auth/internal/server"
	"github.com/jayjzheng/amazon/auth/pb"
	"google.golang.org/grpc"
)

func main() {
	ff := parseFlags()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", ff.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterAuthServiceServer(s, &server.Auth{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

type flags struct {
	port int
}

func parseFlags() flags {
	var ff flags

	flag.IntVar(&ff.port, "port", 3000, "port")
	flag.Parse()

	return ff
}
