package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	Id           int64
	UserId       int64
	WalletNumber string
	Balance      decimal.Decimal
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
