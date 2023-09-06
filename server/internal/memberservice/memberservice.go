package memberservice

import (
	db "github.com/gilwong00/go-discord/db/sqlc"
	"github.com/gilwong00/go-discord/gen/proto/v1/member/memberv1connect"
)

type memberService struct {
	store db.Store
}

func NewMemberService(store db.Store) memberv1connect.MemberServiceHandler {
	return &memberService{
		store: store,
	}
}
