package messageservice

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	messagev1 "github.com/gilwong00/go-discord/gen/proto/v1/message"
)

func (s *messageService) SendDirectMessage(
	ctx context.Context,
	req *connect.Request[messagev1.SendDirectMessageRequest],
	stream *connect.ServerStream[messagev1.SendDirectMessageResponse],
) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}
