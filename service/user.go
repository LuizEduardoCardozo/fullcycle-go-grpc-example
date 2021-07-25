package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {

	stream.Send(&pb.UserResultStream{
		Result: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 1)

	stream.Send(&pb.UserResultStream{
		Result: "Found a name",
		User: &pb.User{
			Name: "Eduardo",
		},
	})

	time.Sleep(time.Second * 1)

	stream.Send(&pb.UserResultStream{
		Result: "Found a email",
		User: &pb.User{
			Name:  "Eduardo",
			Email: "eduard.cardoz@gmail.com",
		},
	})

	return nil

}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {

	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err != io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatalf("Error while receiving the stream: %s\n", err)
		}

		fmt.Printf("Adding user %s\n", req.GetEmail())

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})

	}

}
