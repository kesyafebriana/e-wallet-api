package repository

import (
	"context"
	"database/sql"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/entity"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/kesyafebriana/e-wallet-api/internal/query"
)

type Wallet interface {
	Create(ctx context.Context, wallet *dto.WalletRequest) (*entity.Wallet, error)
	GetByWalletNumber(ctx context.Context, walletNumber string) ([]entity.Wallet, error)
	GetByUserId(ctx context.Context, userId int64) ([]entity.Wallet, error)
	IncreaseBalance(ctx context.Context, request *dto.TransactionRequest) (*entity.Wallet, error)
	DecreaseBalance(ctx context.Context, request *dto.TransactionRequest) (*entity.Wallet, error)
}

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{
		db: db,
	}
}

func (r *WalletRepository) Create(ctx context.Context, wallet *dto.WalletRequest) (*entity.Wallet, error) {
	var responseWallet entity.Wallet

	err := r.db.QueryRowContext(ctx, query.CreateWallet, wallet.UserId).Scan(&responseWallet.Id, &responseWallet.WalletNumber, &responseWallet.Balance, &responseWallet.CreatedAt, &responseWallet.UpdatedAt)

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	responseWallet.UserId = wallet.UserId

	return &responseWallet, nil
}

func (r *WalletRepository) GetByWalletNumber(ctx context.Context, walletNumber string) ([]entity.Wallet, error) {
	res := []entity.Wallet{}

	rows, err := r.db.QueryContext(ctx, query.FindByWalletNumber, walletNumber)

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	defer rows.Close()

	for rows.Next() {
		var responseWallet entity.Wallet

		err := rows.Scan(&responseWallet.Id,
			&responseWallet.UserId,
			&responseWallet.WalletNumber,
			&responseWallet.Balance,
			&responseWallet.CreatedAt,
			&responseWallet.UpdatedAt)

		if err != nil {
			return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
		}

		res = append(res, responseWallet)
	}

	err = rows.Err()

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	return res, nil
}

func (r *WalletRepository) GetByUserId(ctx context.Context, userId int64) ([]entity.Wallet, error) {
	res := []entity.Wallet{}

	rows, err := r.db.QueryContext(ctx, query.FindByUserId, userId)

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	defer rows.Close()

	for rows.Next() {
		var responseWallet entity.Wallet

		err := rows.Scan(&responseWallet.Id,
			&responseWallet.UserId,
			&responseWallet.WalletNumber,
			&responseWallet.Balance,
			&responseWallet.CreatedAt,
			&responseWallet.UpdatedAt)

		if err != nil {
			return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
		}

		res = append(res, responseWallet)
	}

	err = rows.Err()

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	return res, nil
}

func (r *WalletRepository) IncreaseBalance(ctx context.Context, request *dto.TransactionRequest) (*entity.Wallet, error) {
	var walletRes entity.Wallet

	err := r.db.QueryRowContext(ctx, query.IncreaseBalance, request.Amount, request.RecipientWalletNumber).Scan(
		&walletRes.Id,
		&walletRes.UserId,
		&walletRes.WalletNumber,
		&walletRes.Balance,
		&walletRes.CreatedAt,
		&walletRes.UpdatedAt,
	)

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}
	
	return &walletRes, nil
}

func (r *WalletRepository) DecreaseBalance(ctx context.Context, request *dto.TransactionRequest) (*entity.Wallet, error) {
	var walletRes entity.Wallet

	err := r.db.QueryRowContext(ctx, query.DecreaseBalance, request.Amount, request.SenderWalletNumber).Scan(
		&walletRes.Id,
		&walletRes.UserId,
		&walletRes.WalletNumber,
		&walletRes.Balance,
		&walletRes.CreatedAt,
		&walletRes.UpdatedAt,
	)

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}
	
	return &walletRes, nil
}