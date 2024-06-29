package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/entity"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/helper"
	"github.com/kesyafebriana/e-wallet-api/internal/query"
)

type Transaction interface {
	GetAllByWalletNumber(ctx context.Context, walletNumber string, pagination *dto.PaginationInfo) ([]entity.Transaction, *entity.Pagination, error)
	Create(ctx context.Context, transaction *dto.TransactionRequest) (*entity.Transaction, error)
	WithTrx(trxHandle *sql.Tx) TransactionRepository
}

type TransactionRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func NewTransactionRepository(db *sql.DB) Transaction {
	return &TransactionRepository{
		db: db,
	}
}

func (r TransactionRepository) GetAllByWalletNumber(ctx context.Context, walletNumber string, pagination *dto.PaginationInfo) ([]entity.Transaction, *entity.Pagination, error) {
	res := []entity.Transaction{}
	pageInfo := &entity.Pagination{}
	var queryBuilder string

	err := r.db.QueryRowContext(ctx, query.CalculateTransaction, walletNumber).Scan(&pageInfo.TotalData)

	if err != nil {
		return nil, nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	queryBuilder, pageInfo.TotalPage = helper.QueryGetAllTransaction(pagination, pageInfo.TotalData)

	rows, err := r.db.QueryContext(ctx, queryBuilder, walletNumber)

	if err != nil {
		return nil, nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}
	defer rows.Close()

	for rows.Next() {
		var transactionRes entity.Transaction

		err := rows.Scan(&transactionRes.Id,
			&transactionRes.SenderWalletNumber,
			&transactionRes.RecipientWalletNumber,
			&transactionRes.Amount,
			&transactionRes.SourceOfFunds,
			&transactionRes.Description,
			&transactionRes.CreatedAt,
			&transactionRes.UpdatedAt)
		if err != nil {
			return nil, nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
		}
		res = append(res, transactionRes)
	}

	err = rows.Err()
	if err != nil {
		return nil, nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	return res, pageInfo, nil
}

func (r TransactionRepository) Create(ctx context.Context, transaction *dto.TransactionRequest) (*entity.Transaction, error) {
	var responseTransaction entity.Transaction

	err := r.db.QueryRowContext(ctx,
		query.CreateTransaction,
		transaction.SenderWalletNumber,
		transaction.RecipientWalletNumber,
		transaction.Amount,
		transaction.SourceOfFund,
		transaction.Description).Scan(&responseTransaction.Id,
		&responseTransaction.CreatedAt,
		&responseTransaction.UpdatedAt)

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	res := entity.Transaction{Id: responseTransaction.Id,
		SenderWalletNumber:    transaction.SenderWalletNumber,
		RecipientWalletNumber: transaction.RecipientWalletNumber,
		Amount:                transaction.Amount,
		SourceOfFunds:         transaction.SourceOfFund,
		Description:           transaction.Description,
		CreatedAt:             responseTransaction.CreatedAt,
		UpdatedAt:             responseTransaction.UpdatedAt}

	return &res, nil
}

func (r TransactionRepository) WithTrx(trxHandle *sql.Tx) TransactionRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return r
	}

	r.tx = trxHandle
	return r
}
