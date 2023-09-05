package userservice

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	userv1 "github.com/gilwong00/go-discord/gen/proto/v1/user"
)

func (s *userService) Login(
	ctx context.Context,
	req *connect.Request[userv1.LoginRequest],
) (*connect.Response[userv1.LoginResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}
