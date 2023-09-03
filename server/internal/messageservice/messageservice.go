package messageservice

import (
	db "github.com/gilwong00/go-discord/db/sqlc"
	"github.com/gilwong00/go-discord/gen/proto/v1/message/messagev1connect"
)

type messageService struct {
	store db.Store
}

func NewMessageService(store db.Store) messagev1connect.ServerServiceHandler {
	return &messageService{
		store: store,
	}
}
