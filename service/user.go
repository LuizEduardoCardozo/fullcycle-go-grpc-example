package service

import (
	"context"
	"fmt"

	"github.com/LuizEduardoCardozo/fullcycle-go-grpc-example/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (*UserService) NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	fmt.Printf("User name: %s", req.GetName())

	return &pb.User{
		Id:    "1",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil

}
