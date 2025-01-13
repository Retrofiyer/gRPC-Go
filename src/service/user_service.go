package service

import (
	"context"
	"grpc-go/src/model"
	"grpc-go/src/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) AddUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {
	user := model.User{ID: req.Id, Name: req.Name}
	model.Users = append(model.Users, user)
	return &pb.UserResponse{Message: "User added successfully!"}, nil
}

func (s *UserService) GetUsers(ctx context.Context, req *pb.Empty) (*pb.UserList, error) {
	var users []*pb.User
	for _, u := range model.Users {
		users = append(users, &pb.User{Id: u.ID, Name: u.Name})
	}
	return &pb.UserList{Users: users}, nil
}
