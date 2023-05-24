package server

import (
	"context"

	pb "github.com/krishnakantha1/expenseTrackerAuth/pb"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (a *AuthServer) Auth(ctx context.Context, rq *pb.AuthRequestNormal) (*pb.AuthResponse, error) {
	return &pb.AuthResponse{}, nil
}

func (a *AuthServer) AuthJWT(ctx context.Context, rq *pb.AuthRequestJWT) (*pb.AuthResponse, error) {
	return &pb.AuthResponse{}, nil
}

func (a *AuthServer) SignUp(ctx context.Context, rq *pb.SignUpRequest) (*pb.AuthResponse, error) {
	return &pb.AuthResponse{}, nil
}
