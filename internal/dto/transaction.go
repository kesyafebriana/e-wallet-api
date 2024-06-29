package dto

import (
	"fmt"
	"time"

	"github.com/kesyafebriana/e-wallet-api/internal/entity"
	"github.com/shopspring/decimal"
)

type TopUpRequest struct {
	Amount       any    `json:"amount" binding:"required,numeric,min=50000,max=10000000"`
	SourceOfFund string `json:"source_of_fund" binding:"required"`
}

type TransferRequest struct {
	RecipientWalletId string  `json:"recipient_wallet_id"`
	Amount            any     `json:"amount" binding:"required,numeric,min=1000,max=50000000"`
	Description       *string `json:"description" binding:"omitempty,max=35"`
}

type TransactionRequest struct {
	SenderWalletNumber    string          `json:"sender_wallet_id" binding:"required"`
	RecipientWalletNumber string          `json:"recipient_wallet_id" binding:"required"`
	Amount                decimal.Decimal `json:"amount" binding:"required,numeric,min=1000,max=50000000"`
	SourceOfFund          string          `json:"source_of_fund" binding:"required"`
	Description           *string         `json:"description"`
}

type TransactionResponse struct {
	Id           int64           `json:"id"`
	Sender       string          `json:"sender,omitempty"`
	Recipient    string          `json:"recipient"`
	Amount       decimal.Decimal `json:"amount"`
	SourceOfFund string          `json:"source_of_fund"`
	Description  *string         `json:"description,omitempty"`
	CreatedAt    time.Time       `json:"transaction_date"`
}

type ListTransactionResponse struct {
	TotalPage    int                   `json:"total_page"`
	TotalItem    int                   `json:"total_item"`
	Transactions []TransactionResponse `json:"transactions"`
}

func generateDescription(request TopUpRequest) *string {
	description := fmt.Sprintf("Top Up From %s", request.SourceOfFund)
	return &description
}

func ConvertTopUpToTransactionRequest(request *TopUpRequest, walletNumber string) *TransactionRequest {
	return &TransactionRequest{
		SenderWalletNumber:    walletNumber,
		RecipientWalletNumber: walletNumber,
		Amount:                decimal.NewFromFloat(request.Amount.(float64)),
		SourceOfFund:          request.SourceOfFund,
		Description:           generateDescription(*request),
	}
}

func ConvertTransferToTransactionRequest(request *TransferRequest, walletNumber string) *TransactionRequest {
	return &TransactionRequest{
		SenderWalletNumber:    walletNumber,
		RecipientWalletNumber: request.RecipientWalletId,
		Amount:                decimal.NewFromFloat(request.Amount.(float64)),
		SourceOfFund:          "Transfer",
		Description:           request.Description,
	}
}

func ConvertFromTransferRequest(transaction *entity.Transaction, user *UserTransferResponse) *TransactionResponse {
	recipient := fmt.Sprintf("%s (%s)", user.Recipient, transaction.RecipientWalletNumber)

	return &TransactionResponse{
		Id:           transaction.Id,
		Sender:       "You",
		Recipient:    recipient,
		Amount:       transaction.Amount,
		SourceOfFund: transaction.SourceOfFunds,
		Description:  transaction.Description,
		CreatedAt:    transaction.CreatedAt,
	}
}

func ConvertFromTopUpRequest(transaction *entity.Transaction) *TransactionResponse {
	return &TransactionResponse{
		Id:           transaction.Id,
		Recipient:    "You",
		Amount:       transaction.Amount,
		SourceOfFund: transaction.SourceOfFunds,
		Description:  transaction.Description,
		CreatedAt:    transaction.CreatedAt,
	}
}

func ConvertFromTransactionEntity(transaction []entity.Transaction, totalPage int, totalData int) *ListTransactionResponse {
	response := &ListTransactionResponse{}
	transactions := []TransactionResponse{}

	for _, v := range transaction {
		res := TransactionResponse{
			Id:           v.Id,
			Sender:       v.SenderWalletNumber,
			Recipient:    v.RecipientWalletNumber,
			Amount:       v.Amount,
			SourceOfFund: v.SourceOfFunds,
			Description:  v.Description,
			CreatedAt:    v.CreatedAt,
		}

		transactions = append(transactions, res)
	}

	response.Transactions = transactions
	response.TotalItem = totalData
	response.TotalPage = totalPage + 1

	return response
}
