package serverservice

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	serverv1 "github.com/gilwong00/go-discord/gen/proto/v1/server"
)

func (s *serverService) CreateServer(
	ctx context.Context,
	req *connect.Request[serverv1.CreateServerRequest],
) (*connect.Response[serverv1.CreateServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}
