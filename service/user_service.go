package service

import (
	"context"

	"github.com/lavquik/totality/api/pb"
	"github.com/lavquik/totality/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserSservice struct {
	pb.UnimplementedUserServiceServer
}

func (u *UserSservice) GetUserDetails(ctx context.Context, req *pb.GetUserRequest) (*pb.UserDetailsResponse, error) {
	if len(req.Id) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid userId")
	}

	for _, user := range db.Users {
		if user.Id == req.Id {
			userProto := getUserProto(user)
			return &pb.UserDetailsResponse{User: userProto}, nil
		}
	}

	return nil, status.Error(codes.NotFound, "user not found")
}

func (u *UserSservice) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.UserListResponse, error) {
	var users []*pb.UserDetails
	for _, id := range req.Ids {
		for _, user := range db.Users {
			if user.Id == id {
				users = append(users, getUserProto(user))
			}
		}
	}

	var skip = req.PageNumber * req.PageSize
	var limit = skip + req.PageSize
	if limit > int32(len(users)) {
		limit = int32(len(users))
	}

	return &pb.UserListResponse{Users: users[skip:limit]}, nil
}

func (u *UserSservice) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.UserListResponse, error) {
	var users []db.UserDetails

	for _, user := range db.Users {
		if (req.Filters.Id == "" || req.Filters.Id == user.Id) &&
			(req.Filters.Name == "" || req.Filters.Name == user.Name) &&
			(req.Filters.City == "" || req.Filters.City == user.City) &&
			(req.Filters.Phone == 0 || req.Filters.Phone == user.Phone) &&
			(req.Filters.Height == 0 || req.Filters.Height == user.Height) &&
			(req.Filters.Married == pb.Married_UNSPECIFIED || req.Filters.Married == pb.Married(pb.Married_value[user.Married])) {
			users = append(users, user)
		}
	}

	var userProto []*pb.UserDetails
	for _, user := range users {
		userProto = append(userProto, getUserProto(user))
	}

	var skip = req.PageNumber * req.PageSize
	var limit = skip + req.PageSize
	if limit > int32(len(users)) {
		limit = int32(len(users))
	}

	return &pb.UserListResponse{Users: userProto[skip:limit]}, nil

}

func getUserProto(user db.UserDetails) *pb.UserDetails {
	return &pb.UserDetails{
		Id:      user.Id,
		Name:    user.Name,
		City:    user.City,
		Phone:   user.Phone,
		Height:  user.Height,
		Married: pb.Married(pb.Married_value[user.Married]),
	}
}
