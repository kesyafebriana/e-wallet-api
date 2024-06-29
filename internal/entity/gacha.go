package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Gacha struct {
	Id        int64           `json:"id"`
	Amount    decimal.Decimal `json:"amount"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt *time.Time      `json:"-"`
}
