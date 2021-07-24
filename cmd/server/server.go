package main

import (
	"log"
	"net"

	"github.com/LuizEduardoCardozo/fullcycle-go-grpc-example/pb"
	"github.com/LuizEduardoCardozo/fullcycle-go-grpc-example/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:5051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &service.UserService{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
