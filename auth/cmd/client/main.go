package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/jayjzheng/amazon/auth/pb"
	"google.golang.org/grpc"
)

func main() {
	ff := parseFlags()

	conn, err := grpc.Dial(ff.server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)
	ctx := context.Background()

	user := pb.User{
		Login:    ff.login,
		Password: ff.password,
	}

	switch ff.action {
	case "create":
		_, err = client.CreateUser(ctx, &user)
	case "change":
		_, err = client.ChangePassword(ctx, &pb.ChangePasswordRequest{
			User:        &user,
			NewPassword: ff.newPassword,
		})
	case "token":
		var resp *pb.CreateTokenResponse
		resp, err = client.CreateToken(ctx, &user)
		fmt.Printf("%+v\n", resp)
	default:
		log.Fatalf("unknown action: %s, 'create', 'change' or 'token'\n", ff.action)
	}

	fmt.Println("error", err)
}

type flags struct {
	server      string
	login       string
	password    string
	newPassword string
	action      string
}

func parseFlags() flags {
	var ff flags

	flag.StringVar(&ff.server, "server", "localhost:3000", "server address")
	flag.StringVar(&ff.action, "action", "", "grpc actions, 'create', 'change', or 'token'")
	flag.StringVar(&ff.login, "login", "", "login")
	flag.StringVar(&ff.password, "password", "", "password")
	flag.StringVar(&ff.newPassword, "new-pass", "", "new password")
	flag.Parse()

	return ff
}
