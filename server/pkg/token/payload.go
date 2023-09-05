package token

import (
	"errors"
	"time"

	"github.com/gofrs/uuid/v5"
)

type Payload struct {
	UserID    uuid.UUID `json:"userId"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

var (
	errExpiredToken = errors.New("token has expired")
	errInvalidToken = errors.New("invalid token")
)

func NewTokenPayload(
	userID uuid.UUID,
	duration time.Duration,
) *Payload {
	return &Payload{
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

func (p *Payload) Validate() error {
	if time.Now().After(p.ExpiredAt) {
		return errExpiredToken
	}
	return nil
}
