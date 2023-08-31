package userservice

import (
	db "github.com/gilwong00/go-discord/db/sqlc"
	"github.com/gilwong00/go-discord/gen/proto/v1/user/userv1connect"
)

type userService struct {
	store db.Store
}

func NewUserService(store db.Store) userv1connect.UserServiceHandler {
	return &userService{
		store: store,
	}
}
