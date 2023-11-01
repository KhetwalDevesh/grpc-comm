package main

import (
	"context"
	"github.com/KhetwalDevesh/project-totality/data"
	pb "github.com/KhetwalDevesh/project-totality/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

var address = "0.0.0.0:50051"

type UserServiceServer struct {
	pb.UserServiceServer
}

var users = data.GetUsers()

func (s *UserServiceServer) GetUserById(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	// Find the user by ID
	for _, user := range users {
		if user.ID == in.Id {
			return &pb.UserResponse{
				User: &pb.User{
					Id:      user.ID,
					Name:    user.Name,
					City:    user.City,
					Phone:   user.Phone,
					Height:  user.Height,
					Married: user.Married,
				},
			}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "User not found")
}

func (s *UserServiceServer) GetUsersByIds(ctx context.Context, in *pb.UsersRequest) (*pb.UsersResponse, error) {
	var userList []*pb.User
	// Find the users by IDs
	for _, id := range in.Ids {
		for _, user := range users {
			if id == user.ID {
				userList = append(userList, &pb.User{
					Id:      user.ID,
					Name:    user.Name,
					City:    user.City,
					Phone:   user.Phone,
					Height:  user.Height,
					Married: user.Married,
				})
				break
			}
		}
	}
	if len(userList) == 0 {
		return nil, status.Errorf(codes.NotFound, "No User found")
	}
	return &pb.UsersResponse{Users: userList}, nil
}

func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	log.Printf("Listening on %s\n", address)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserServiceServer{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve : %v\n", err)
	}
}
