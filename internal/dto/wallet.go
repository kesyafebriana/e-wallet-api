package dto

import (
	"github.com/kesyafebriana/e-wallet-api/internal/entity"
	"github.com/shopspring/decimal"
)

type WalletRequest struct {
	UserId       int64  `json:"user_id" binding:"required"`
	WalletNumber string `json:"wallet_number"`
}

type WalletResponse struct {
	WalletNumber string          `json:"wallet_number"`
	Balance      decimal.Decimal `json:"balance"`
}

func ConvertFromWalletEntity(walletEntity *entity.Wallet) *WalletResponse {

	return &WalletResponse{
		WalletNumber: walletEntity.WalletNumber,
		Balance:      walletEntity.Balance,
	}
}
