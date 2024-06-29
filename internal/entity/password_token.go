package entity

import (
	"time"
)

type PasswordTokens struct {
	Id        int64
	UserId    int64
	Token     string
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
