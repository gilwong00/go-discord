package userservice

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	userv1 "github.com/gilwong00/go-discord/gen/proto/user/v1"
)

func (s *userService) CreateUser(
	ctx context.Context,
	req *connect.Request[userv1.CreateUserRequest],
) (*connect.Response[userv1.CreateUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}
