package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "example.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	PORT = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	user_id := int32(rand.Intn(1000))
	user := &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}
	return user, nil
}

func main() {
	// This net listener here will listen for incoming requests and manage traffic towards server that we will run.
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Error while creating listener: %v", err)
	}
	// This is instance of server that will be listening at above mentioned listener and server request using service we will create.
	s := grpc.NewServer()
	// Here we are registering service that will server.
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	fmt.Printf("Server listening at port no: %v", lis.Addr())
	// Here our server will start listening for traffic using listener we created.
	err2 := s.Serve(lis)
	if err2 != nil {
		log.Fatalf("Server didn't started:%v", err2)
	}
}
