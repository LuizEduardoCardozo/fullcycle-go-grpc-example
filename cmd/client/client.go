package main

import (
	"context"
	"fmt"
	"log"

	"github.com/LuizEduardoCardozo/fullcycle-go-grpc-example/pb"
	"google.golang.org/grpc"
)

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "1",
		Name:  "Eduardo",
		Email: "eduard.cardoz@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(res)
}

func main() {

	connection, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUser(client)

}
