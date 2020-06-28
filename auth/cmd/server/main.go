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
			Publisher: mockPublisher{},
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
	flag.StringVar(&ff.jwtSecret, "jwt-secret", "", "jwt secret, env var JWT_SECRET")
	flag.Parse()

	if ff.jwtSecret == "" {
		ff.jwtSecret = os.Getenv("JWT_SECRET")
	}

	return ff
}

type mockPublisher struct{}

func (m mockPublisher) PublishUserCreated(u *domain.User) error {
	log.Printf("UserCreated: %+v\n", u)
	return nil
}

func (m mockPublisher) PublishTokenCreated(u *domain.User) error {
	log.Printf("TokenCreated: %+v\n", u)
	return nil
}

func (m mockPublisher) PublishPasswordChanged(u *domain.User) error {
	log.Printf("PasswordChanged: %+v\n", u)
	return nil
}
