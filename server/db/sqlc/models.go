// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	uuid "github.com/gofrs/uuid/v5"
)

type MemberRole string

const (
	MemberRoleADMIN     MemberRole = "ADMIN"
	MemberRoleGUEST     MemberRole = "GUEST"
	MemberRoleMODERATOR MemberRole = "MODERATOR"
)

func (e *MemberRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = MemberRole(s)
	case string:
		*e = MemberRole(s)
	default:
		return fmt.Errorf("unsupported scan type for MemberRole: %T", src)
	}
	return nil
}

type NullMemberRole struct {
	MemberRole MemberRole
	Valid      bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullMemberRole) Scan(value interface{}) error {
	if value == nil {
		ns.MemberRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.MemberRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullMemberRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.MemberRole, nil
}

type Member struct {
	ID        uuid.UUID
	Role      MemberRole
	UserID    uuid.UUID
	ServerID  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Server struct {
	ID              uuid.UUID
	Name            string
	ImageUrl        sql.NullString
	InviteCode      sql.NullString
	CreatedByUserID uuid.NullUUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	AvatarUrl sql.NullString
	CreatedAt time.Time
	DeletedAt time.Time
}
