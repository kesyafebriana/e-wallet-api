package entity

import "time"

type GachaAttempt struct {
	Id        int64
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
