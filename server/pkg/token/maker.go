package token

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Maker interface {
	CreateToken(userId uuid.UUID, duration time.Duration) (string, *Payload, error)
	ValidateToken(token string) (*Payload, error)
}
