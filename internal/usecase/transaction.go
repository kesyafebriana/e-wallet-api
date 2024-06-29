package usecase

import (
	"context"
	"database/sql"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/entity"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/kesyafebriana/e-wallet-api/internal/repository"
	"github.com/shopspring/decimal"
)

type Transaction interface {
	GetAllByWalletNumber(ctx context.Context, walletNumber string, pagination *dto.PaginationInfo) ([]entity.Transaction, *entity.Pagination, error)
	CreateTopUp(ctx context.Context, request *dto.TransactionRequest) (*entity.Transaction, error)
	CreateTransfer(ctx context.Context, request *dto.TransactionRequest) (*entity.Transaction, *dto.UserTransferResponse, error)
	WithTrx(*sql.Tx) TransactionImplementation
}

type TransactionImplementation struct {
	TransactionRepository repository.Transaction
	WalletRepository      repository.Wallet
	UserRepository        repository.User
	GachaRepository       repository.Gacha
}

func NewTransactionImplementation(TransactionRepository repository.Transaction, WalletRepository repository.Wallet, UserRepository repository.User, GachaRepository repository.Gacha) *TransactionImplementation {
	return &TransactionImplementation{
		TransactionRepository: TransactionRepository,
		WalletRepository:      WalletRepository,
		UserRepository:        UserRepository,
		GachaRepository:       GachaRepository,
	}
}

func (u *TransactionImplementation) GetAllByWalletNumber(ctx context.Context, walletNumber string, pagination *dto.PaginationInfo) ([]entity.Transaction, *entity.Pagination, error) {
	transactionRes, pageInfo, err := u.TransactionRepository.GetAllByWalletNumber(ctx, walletNumber, pagination)

	if err != nil {
		return nil, nil, err
	}

	return transactionRes, pageInfo, nil
}

func (u TransactionImplementation) CreateTopUp(ctx context.Context, request *dto.TransactionRequest) (*entity.Transaction, error) {
	transactionRes, err := u.TransactionRepository.Create(ctx, request)

	if err != nil {
		return nil, err
	}

	walletRes, err := u.WalletRepository.IncreaseBalance(ctx, request)

	if err != nil {
		return nil, err
	}

	if request.Amount.LessThanOrEqual(decimal.NewFromInt(10000000)) && request.SourceOfFund != "Reward" {
		_, err := u.GachaRepository.AddGachaAttempt(ctx, walletRes.UserId)

		if err != nil {
			return nil, err
		}
	}

	return transactionRes, nil
}

func (u TransactionImplementation) CreateTransfer(ctx context.Context, request *dto.TransactionRequest) (*entity.Transaction, *dto.UserTransferResponse, error) {
	recipientWalletRes, err := u.WalletRepository.GetByWalletNumber(ctx, request.RecipientWalletNumber)

	if err != nil {
		return nil, nil, err
	}

	if len(recipientWalletRes) == 0 {
		return nil, nil, apperror.StatusBadRequest(constant.ErrorWalletNumberNotFound, constant.RecipientWalletNotFoundMsg)
	}

	if recipientWalletRes[0].WalletNumber == request.SenderWalletNumber {
		return nil, nil, apperror.StatusBadRequest(constant.ErrorTransferToAccount, constant.TransferToAccountMsg)
	}

	recipientUserRes, err := u.UserRepository.FindById(ctx, recipientWalletRes[0].UserId)

	if err != nil {
		return nil, nil, err
	}

	senderWalletRes, err := u.WalletRepository.GetByWalletNumber(ctx, request.SenderWalletNumber)

	if err != nil {
		return nil, nil, err
	}

	if request.Amount.GreaterThan(senderWalletRes[0].Balance) {
		return nil, nil, apperror.StatusBadRequest(constant.ErrorInsufficientBalance, constant.InsufficientBalanceMsg)
	}

	senderUserRes, err := u.UserRepository.FindById(ctx, senderWalletRes[0].UserId)

	if err != nil {
		return nil, nil, err
	}

	transactionRes, err := u.TransactionRepository.Create(ctx, request)

	if err != nil {
		return nil, nil, err
	}

	_, err = u.WalletRepository.IncreaseBalance(ctx, request)

	if err != nil {
		return nil, nil, err
	}

	_, err = u.WalletRepository.DecreaseBalance(ctx, request)

	if err != nil {
		return nil, nil, err
	}

	userRes := &dto.UserTransferResponse{
		Sender:    senderUserRes[0].Name,
		Recipient: recipientUserRes[0].Name,
	}

	return transactionRes, userRes, nil
}

func (u TransactionImplementation) WithTrx(trxHandle *sql.Tx) TransactionImplementation {
	u.TransactionRepository = u.TransactionRepository.WithTrx(trxHandle)
	return u
}
