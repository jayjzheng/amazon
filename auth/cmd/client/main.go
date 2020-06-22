package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jayjzheng/amazon/auth/pb"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := "localhost:3000"
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)
	ctx := context.Background()

	user := pb.User{
		Login:    "foo",
		Password: "bar",
	}

	if _, err := client.CreateUser(ctx, &user); err != nil {
		log.Fatalln(err)
	}

	resp, err := client.CreateToken(ctx, &user)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v\n", resp)

	resp, err = client.CreateToken(ctx, &pb.User{
		Login:    "foo",
		Password: "barbar",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v\n", resp)

	_, err = client.ChangePassword(ctx, &pb.ChangePasswordRequest{
		User: &pb.User{
			Login:    "foo",
			Password: "bar",
		},
		NewPassword: "barbar",
	})
	if err != nil {
		log.Println(err)
	}

	resp, err = client.CreateToken(ctx, &pb.User{
		Login:    "foo",
		Password: "barbar",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("changed: %+v\n", resp)
}
