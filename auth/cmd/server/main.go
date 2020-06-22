package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jayjzheng/amazon/auth/internal/domain"
	"github.com/jayjzheng/amazon/auth/internal/jwt"
	"github.com/jayjzheng/amazon/auth/internal/memory"
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
	auth := &server.Auth{
		Service: domain.AuthService{
			UserStore: memory.NewClient(),
			TokenGenerator: jwt.NewJWT(
				jwt.WithSecret([]byte(ff.jwtSecret)),
			),
		},
	}

	pb.RegisterAuthServiceServer(s, auth)

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

type flags struct {
	port      int
	jwtSecret string
}

func parseFlags() flags {
	var ff flags

	flag.IntVar(&ff.port, "port", 3000, "port")
	flag.StringVar(&ff.jwtSecret, "jwt-secret", "", "jwt secret")
	flag.Parse()

	if ff.jwtSecret == "" {
		ff.jwtSecret = os.Getenv("JWT_SECRET")
	}

	return ff
}
