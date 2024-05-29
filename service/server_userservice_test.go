package service

import (
	"context"
	"testing"

	"github.com/lavquik/totality/api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUserDetailsService(t *testing.T) {
	ctx := context.Background()
	client, closer := server()
	defer closer()

	type expectation struct {
		out *pb.UserDetailsResponse
		err error
	}

	tests := map[string]struct {
		req *pb.GetUserRequest
		exp expectation
	}{
		"valid user": {
			req: &pb.GetUserRequest{Id: "1"},
			exp: expectation{
				out: &pb.UserDetailsResponse{
					User: &pb.UserDetails{
						Id:      "1",
						Name:    "Steve",
						City:    "LA",
						Phone:   1234567890,
						Height:  5.8,
						Married: pb.Married_TRUE,
					},
				},
				err: nil,
			},
		},
		"invalid user": {
			req: &pb.GetUserRequest{Id: "4"},
			exp: expectation{
				out: nil,
				err: status.Error(codes.NotFound, "user not found"),
			},
		},
		"empty user": {
			req: &pb.GetUserRequest{Id: ""},
			exp: expectation{
				out: nil,
				err: status.Error(codes.InvalidArgument, "invalid userId"),
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.GetUserDetails(ctx, tt.req)
			if err != nil {
				if tt.exp.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.exp.err, err)
				}
			} else {
				if tt.exp.out.User.Id != out.User.Id ||
					tt.exp.out.User.Name != out.User.Name ||
					tt.exp.out.User.City != out.User.City ||
					tt.exp.out.User.Phone != out.User.Phone ||
					tt.exp.out.User.Height != out.User.Height ||
					tt.exp.out.User.Married != out.User.Married {
					t.Errorf("Out -> \nWant: %v\nGot: %v\n", tt.exp.out, out)
				}
			}

		})
	}
}

func TestListUserService(t *testing.T) {
	ctx := context.Background()
	client, closer := server()
	defer closer()
	type expectation struct {
		out *pb.UserListResponse
		err error
	}
	tests := map[string]struct {
		req *pb.GetUserListRequest
		exp expectation
	}{
		"valid user list": {
			req: &pb.GetUserListRequest{Ids: []string{"1", "2"}, PageNumber: 0, PageSize: 2},
			exp: expectation{
				out: &pb.UserListResponse{
					Users: []*pb.UserDetails{
						{
							Id:      "1",
							Name:    "Steve",
							City:    "LA",
							Phone:   1234567890,
							Height:  5.8,
							Married: pb.Married_TRUE,
						},
						{
							Id:      "2",
							Name:    "John",
							City:    "NY",
							Phone:   2345678901,
							Height:  5.9,
							Married: pb.Married_TRUE,
						},
					},
				},
				err: nil,
			},
		},
		"invalid user list": {
			req: &pb.GetUserListRequest{Ids: []string{"5", "4"}, PageNumber: 0, PageSize: 2},
			exp: expectation{
				out: &pb.UserListResponse{
					Users: []*pb.UserDetails{},
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.GetUserList(ctx, tt.req)
			if err != nil {
				if tt.exp.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.exp.err, err)
				}
			} else {
				if len(tt.exp.out.Users) == 0 && len(out.Users) != 0 {
					t.Errorf("Out -> \nWant: %v\nGot: %v\n", tt.exp.out, out)
				} else {
					for i, user := range out.Users {
						if tt.exp.out.Users[i].Id != user.Id ||
							tt.exp.out.Users[i].Name != user.Name ||
							tt.exp.out.Users[i].City != user.City ||
							tt.exp.out.Users[i].Phone != user.Phone ||
							tt.exp.out.Users[i].Height != user.Height ||
							tt.exp.out.Users[i].Married != user.Married {
							t.Errorf("Out -> \nWant: %v\nGot: %v\n", tt.exp.out, out)
						}
					}
				}
			}
		})
	}

}

func TestSearchUserService(t *testing.T) {
	ctx := context.Background()
	client, closer := server()
	defer closer()
	type expectation struct {
		out *pb.UserListResponse
		err error
	}
	tests := map[string]struct {
		req *pb.SearchUsersRequest
		exp expectation
	}{
		"valid user list": {
			req: &pb.SearchUsersRequest{
				Filters: &pb.UserDetails{
					Phone:   1234567890,
					Married: pb.Married_TRUE,
				},
				PageNumber: 0,
				PageSize:   2,
			},
			exp: expectation{
				out: &pb.UserListResponse{
					Users: []*pb.UserDetails{
						{
							Id:      "1",
							Name:    "Steve",
							City:    "LA",
							Phone:   1234567890,
							Height:  5.8,
							Married: pb.Married_TRUE,
						},
					},
				},
				err: nil,
			},
		},
		"NonExistent users": {
			req: &pb.SearchUsersRequest{
				Filters: &pb.UserDetails{
					Phone:   999999,
					Married: pb.Married_FALSE,
				},
				PageNumber: 0,
				PageSize:   2,
			},
			exp: expectation{
				out: &pb.UserListResponse{
					Users: []*pb.UserDetails{},
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.SearchUsers(ctx, tt.req)
			if err != nil {
				if tt.exp.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.exp.err, err)
				}
			} else {
				if len(tt.exp.out.Users) == 0 && len(out.Users) != 0 {
					t.Errorf("Out -> \nWant: %v\nGot: %v\n", tt.exp.out, out)
				} else {
					for i, user := range out.Users {
						if tt.exp.out.Users[i].Id != user.Id ||
							tt.exp.out.Users[i].Name != user.Name ||
							tt.exp.out.Users[i].City != user.City ||
							tt.exp.out.Users[i].Phone != user.Phone ||
							tt.exp.out.Users[i].Height != user.Height ||
							tt.exp.out.Users[i].Married != user.Married {
							t.Errorf("Out -> \nWant: %v\nGot: %v\n", tt.exp.out, out)
						}
					}
				}
			}
		})
	}

}
