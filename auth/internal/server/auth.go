package server

import (
	"context"

	"github.com/jayjzheng/amazon/auth/internal/domain"
	"github.com/jayjzheng/amazon/auth/pb"
)

type Auth struct {
	pb.UnimplementedAuthServiceServer

	Service domain.AuthService
}

func (s *Auth) CreateUser(_ context.Context, req *pb.User) (*pb.CreateUserResponse, error) {
	if err := s.Service.CreateUser(&domain.User{
		Login:    req.GetLogin(),
		Password: req.GetPassword(),
	}); err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{}, nil
}

func (s *Auth) ChangePassword(_ context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	if err := s.Service.ChangePassword(
		req.User.GetLogin(),
		req.User.GetPassword(),
		req.GetNewPassword(),
	); err != nil {
		return nil, err
	}

	return &pb.ChangePasswordResponse{}, nil
}

func (s *Auth) CreateToken(_ context.Context, req *pb.User) (*pb.CreateTokenResponse, error) {
	token, err := s.Service.CreateToken(&domain.User{
		Login:    req.GetLogin(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateTokenResponse{
		Token: token,
	}, nil
}
