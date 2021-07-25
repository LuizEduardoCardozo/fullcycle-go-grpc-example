package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

func AddUserVerbose(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "3",
		Name:  "Eduardo",
		Email: "eduard.cardoz@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(stream)

	}

}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "Eduardo",
			Email: "eduard.cardoz@gmail.com",
		},
		{
			Id:    "2",
			Name:  "Wesley Willians",
			Email: "wesley@fullcycle.com.br",
		},
		{
			Id:    "3",
			Name:  "Rodrigo",
			Email: "rod.abreu@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error while creating request: %s\n", err)
	}

	for counter, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
		fmt.Printf("Request number %d sended\n", counter+1)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving the response: %s\n", err)
	}

	fmt.Println(res)
}

func AddUsersStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUsersStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error while receiving the stream: %s\n", err)
	}

	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "Eduardo",
			Email: "eduard.cardoz@gmail.com",
		},
		{
			Id:    "2",
			Name:  "Wesley Willians",
			Email: "wesley@fullcycle.com.br",
		},
		{
			Id:    "3",
			Name:  "Rodrigo",
			Email: "rod.abreu@gmail.com",
		},
	}

	wait := make(chan int)

	go func() {

		for _, req := range reqs {
			fmt.Printf("sending user: %s", req.GetEmail())

			stream.Send(req)
			time.Sleep(time.Second * 2)

		}
		stream.CloseSend()

	}()

	go func() {

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving the data: %s\n", err)
				break
			}

			fmt.Printf("received user %s with the status %s",
				res.GetUser().GetEmail(),
				res.GetResult(),
			)
		}
		close(wait)
	}()

	<-wait

}

func main() {

	connection, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	fmt.Println("Adding users")
	AddUsersStreamBoth(client)
	fmt.Println("Users added!")
}
