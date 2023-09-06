package memberservice

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	memberv1 "github.com/gilwong00/go-discord/gen/proto/v1/member"
)

func (s *memberService) CreateServerMember(
	ctx context.Context,
	req *connect.Request[memberv1.CreateServerMemberRequest],
) (*connect.Response[memberv1.CreateServerMemberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}
