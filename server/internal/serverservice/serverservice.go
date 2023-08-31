package serverservice

import (
	db "github.com/gilwong00/go-discord/db/sqlc"
	"github.com/gilwong00/go-discord/gen/proto/v1/server/serverv1connect"
)

type serverService struct {
	store db.Store
}

func NewServerService(store db.Store) serverv1connect.ServerServiceHandler {
	return &serverService{
		store: store,
	}
}
