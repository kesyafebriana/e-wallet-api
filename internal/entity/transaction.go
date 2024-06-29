package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Id                    int64
	SenderWalletNumber    string
	RecipientWalletNumber string
	Amount                decimal.Decimal
	SourceOfFunds         string
	Description           *string
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             *time.Time
}
