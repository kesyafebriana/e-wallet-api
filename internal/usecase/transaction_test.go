package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/entity"
	"github.com/kesyafebriana/e-wallet-api/internal/mocks/repomock"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/kesyafebriana/e-wallet-api/internal/repository"
	"github.com/kesyafebriana/e-wallet-api/internal/usecase"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransactionImplementation_GetAllByWalletNumber(t *testing.T) {
	t.Run("should return error when repository returned error", func(t *testing.T) {
		// given
		var ctx context.Context
		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockTransactionRepo.On("GetAllByWalletNumber", ctx, mock.Anything, mock.Anything).Return(nil, nil, errors.New("test error"))
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transaction, pagination, err := u.GetAllByWalletNumber(ctx, "777001", &dto.PaginationInfo{})

		// then
		assert.Error(t, errors.New("test error"), err)
		assert.Empty(t, transaction)
		assert.Nil(t, pagination)
	})

	t.Run("should return transaction response and page info when get all transaction success", func(t *testing.T) {
		// given
		var ctx context.Context

		page := &entity.Pagination{
			TotalData: 1,
			TotalPage: 1,
		}

		transaction := []entity.Transaction{
			{Id: 1},
		}

		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)

		mockTransactionRepo.On("GetAllByWalletNumber", ctx, mock.Anything, mock.Anything).Return(transaction, page, nil)
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, paginationRes, err := u.GetAllByWalletNumber(ctx, "777001", &dto.PaginationInfo{})

		// then
		assert.NoError(t, err)
		assert.Equal(t, page, paginationRes)
		assert.Equal(t, transaction, transactionRes)
	})
}

func TestTransactionImplementation_CreateTopUp(t *testing.T) {
	t.Run("should return error when create error repository returned error", func(t *testing.T) {
		// given
		var ctx context.Context

		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)

		mockTransactionRepo.On("Create", ctx, mock.Anything).Return(nil, errors.New("test error"))
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, err := u.CreateTopUp(ctx, &dto.TransactionRequest{SenderWalletNumber: "Test"})

		// then
		assert.Error(t, errors.New("test error"), err)
		assert.Nil(t, transactionRes)
	})

	t.Run("should return error when increase balance repository returned error", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &entity.Transaction{
			Id: 1,
		}

		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockTransactionRepo.On("Create", ctx, mock.Anything).Return(transaction, nil)
		mockWalletRepo.On("IncreaseBalance", ctx, mock.Anything).Return(nil, errors.New("test error"))
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, err := u.CreateTopUp(ctx, &dto.TransactionRequest{SenderWalletNumber: "Test"})

		// then
		assert.Error(t, errors.New("test error"), err)
		assert.Nil(t, transactionRes)
	})

	t.Run("should add gacha attempt when amount equal to 10 million", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &entity.Transaction{
			Id: 1,
		}

		wallet := &entity.Wallet{
			Id: 1,
		}

		gacha := &entity.GachaAttempt{
			Id: 1,
		}

		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockTransactionRepo.On("Create", ctx, mock.Anything).Return(transaction, nil)
		mockWalletRepo.On("IncreaseBalance", ctx, mock.Anything).Return(wallet, nil)
		mockGachaRepo.On("AddGachaAttempt", ctx, mock.Anything).Return(gacha, nil)
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, err := u.CreateTopUp(ctx, &dto.TransactionRequest{SenderWalletNumber: "Test"})

		// then
		assert.Nil(t, err)
		assert.Equal(t, transaction, transactionRes)
	})

	t.Run("should return error when add gacha attempt repository returned error", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &entity.Transaction{
			Id: 1,
		}

		wallet := &entity.Wallet{
			Id: 1,
		}

		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockTransactionRepo.On("Create", ctx, mock.Anything).Return(transaction, nil)
		mockWalletRepo.On("IncreaseBalance", ctx, mock.Anything).Return(wallet, nil)
		mockGachaRepo.On("AddGachaAttempt", ctx, mock.Anything).Return(nil, errors.New("test error"))
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, err := u.CreateTopUp(ctx, &dto.TransactionRequest{SenderWalletNumber: "Test"})

		// then
		assert.Error(t, errors.New("test error"), err)
		assert.Nil(t, transactionRes)
	})
}

func TestTransactionImplementation_CreateTransfer(t *testing.T) {
	t.Run("should return error when get recipient wallet repository returned error", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &dto.TransactionRequest{
			SenderWalletNumber: "1",
		}
		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockWalletRepo.On("GetByWalletNumber", ctx, mock.Anything).Return(nil, errors.New("test error"))
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, userRes, err := u.CreateTransfer(ctx, transaction)

		// then
		assert.Error(t, errors.New("test error"), err)
		assert.Nil(t, transactionRes)
		assert.Nil(t, userRes)
	})

	t.Run("should return error when no recipient with wallet number input", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &dto.TransactionRequest{
			SenderWalletNumber: "1",
		}
		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockWalletRepo.On("GetByWalletNumber", ctx, mock.Anything).Return([]entity.Wallet{}, nil)
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, userRes, err := u.CreateTransfer(ctx, transaction)

		// then
		assert.Error(t, constant.ErrorWalletNumberNotFound, err)
		assert.Nil(t, transactionRes)
		assert.Nil(t, userRes)
	})

	t.Run("should return error when get recipient user repository returned error", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &dto.TransactionRequest{
			SenderWalletNumber: "1",
		}
		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockWalletRepo.On("GetByWalletNumber", ctx, mock.Anything).Return([]entity.Wallet{{Id: 1}}, nil)
		mockUserRepo.On("FindById", ctx, mock.Anything).Return(nil, errors.New("test error"))
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, userRes, err := u.CreateTransfer(ctx, transaction)

		// then
		assert.Error(t, errors.New("test error"), err)
		assert.Nil(t, transactionRes)
		assert.Nil(t, userRes)
	})

	t.Run("should return error when sender wallet number same with recipient wallet number", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &dto.TransactionRequest{
			SenderWalletNumber: "1",
		}
		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockWalletRepo.On("GetByWalletNumber", ctx, mock.Anything).Return([]entity.Wallet{{WalletNumber: "1"}}, nil)
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, userRes, err := u.CreateTransfer(ctx, transaction)

		// then
		assert.Error(t, constant.ErrorTransferToAccount, err)
		assert.Nil(t, transactionRes)
		assert.Nil(t, userRes)
	})

	t.Run("should return error when find id repository returned error", func(t *testing.T) {
		var ctx context.Context

		transaction := &dto.TransactionRequest{
			SenderWalletNumber:    "1",
			RecipientWalletNumber: "2",
			Amount:                decimal.NewFromInt(10000),
		}

		recipient := []entity.Wallet{{WalletNumber: "2", UserId: 1}}
		sender := []entity.Wallet{{WalletNumber: "1", UserId: 2, Balance: decimal.NewFromInt(50000)}}

		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.RecipientWalletNumber).Return(recipient, nil)
		mockUserRepo.On("FindById", ctx, recipient[0].UserId).Return([]entity.User{{Id: 1}}, nil)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.SenderWalletNumber).Return([]entity.Wallet{{WalletNumber: "1", UserId: 2, Balance: decimal.NewFromInt(50000)}}, nil)
		mockUserRepo.On("FindById", ctx, sender[0].UserId).Return(nil, errors.New("test error"))
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, userRes, err := u.CreateTransfer(ctx, transaction)

		// then
		assert.Error(t, errors.New("test error"), err)
		assert.Nil(t, transactionRes)
		assert.Nil(t, userRes)
	})

	t.Run("should return error when get sender wallet repository returned error", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &dto.TransactionRequest{
			SenderWalletNumber:    "1",
			RecipientWalletNumber: "2",
		}
		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.RecipientWalletNumber).Return([]entity.Wallet{{WalletNumber: "2"}}, nil)
		mockUserRepo.On("FindById", ctx, mock.Anything).Return([]entity.User{{Id: 1}}, nil)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.SenderWalletNumber).Return(nil, errors.New("test error"))
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, userRes, err := u.CreateTransfer(ctx, transaction)

		// then
		assert.Error(t, errors.New("test error"), err)
		assert.Nil(t, transactionRes)
		assert.Nil(t, userRes)
	})

	t.Run("should return error when amount transfer bigger than sender balance", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &dto.TransactionRequest{
			SenderWalletNumber:    "1",
			RecipientWalletNumber: "2",
			Amount:                decimal.NewFromInt(10000),
		}
		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.RecipientWalletNumber).Return([]entity.Wallet{{WalletNumber: "2"}}, nil)
		mockUserRepo.On("FindById", ctx, mock.Anything).Return([]entity.User{{Id: 1}}, nil)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.SenderWalletNumber).Return([]entity.Wallet{{WalletNumber: "1", Balance: decimal.NewFromInt(500)}}, nil)
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, userRes, err := u.CreateTransfer(ctx, transaction)

		// then
		assert.Error(t, constant.ErrorInsufficientBalance, err)
		assert.Nil(t, transactionRes)
		assert.Nil(t, userRes)
	})

	t.Run("should return error when create transaction repository returned error", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &dto.TransactionRequest{
			SenderWalletNumber:    "1",
			RecipientWalletNumber: "2",
			Amount:                decimal.NewFromInt(10000),
		}
		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.RecipientWalletNumber).Return([]entity.Wallet{{WalletNumber: "2"}}, nil)
		mockUserRepo.On("FindById", ctx, mock.Anything).Return([]entity.User{{Id: 1}}, nil)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.SenderWalletNumber).Return([]entity.Wallet{{WalletNumber: "1", Balance: decimal.NewFromInt(100000)}}, nil)
		mockTransactionRepo.On("Create", ctx, mock.Anything).Return(nil, errors.New("test error"))
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, userRes, err := u.CreateTransfer(ctx, transaction)

		// then
		assert.Error(t, errors.New("test error"), err)
		assert.Nil(t, transactionRes)
		assert.Nil(t, userRes)
	})

	t.Run("should return error when increase recipient balance returned error", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &dto.TransactionRequest{
			SenderWalletNumber:    "1",
			RecipientWalletNumber: "2",
			Amount:                decimal.NewFromInt(10000),
		}
		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.RecipientWalletNumber).Return([]entity.Wallet{{WalletNumber: "2"}}, nil)
		mockUserRepo.On("FindById", ctx, mock.Anything).Return([]entity.User{{Id: 1}}, nil)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.SenderWalletNumber).Return([]entity.Wallet{{WalletNumber: "1", Balance: decimal.NewFromInt(100000)}}, nil)
		mockTransactionRepo.On("Create", ctx, transaction).Return(&entity.Transaction{Id: 1}, nil)
		mockWalletRepo.On("IncreaseBalance", ctx, mock.Anything).Return(nil, errors.New("test error"))
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, userRes, err := u.CreateTransfer(ctx, transaction)

		// then
		assert.Error(t, errors.New("test error"), err)
		assert.Nil(t, transactionRes)
		assert.Nil(t, userRes)
	})

	t.Run("should return error when decrease sender balance returned error", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &dto.TransactionRequest{
			SenderWalletNumber:    "1",
			RecipientWalletNumber: "2",
			Amount:                decimal.NewFromInt(10000),
		}
		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.RecipientWalletNumber).Return([]entity.Wallet{{WalletNumber: "2"}}, nil)
		mockUserRepo.On("FindById", ctx, mock.Anything).Return([]entity.User{{Id: 1}}, nil)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.SenderWalletNumber).Return([]entity.Wallet{{WalletNumber: "1", Balance: decimal.NewFromInt(100000)}}, nil)
		mockTransactionRepo.On("Create", ctx, transaction).Return(&entity.Transaction{Id: 1}, nil)
		mockWalletRepo.On("IncreaseBalance", ctx, mock.Anything).Return(&entity.Wallet{UserId: 1}, nil)
		mockWalletRepo.On("DecreaseBalance", ctx, mock.Anything).Return(nil, errors.New("test error"))
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, userRes, err := u.CreateTransfer(ctx, transaction)

		// then
		assert.Error(t, errors.New("test error"), err)
		assert.Nil(t, transactionRes)
		assert.Nil(t, userRes)
	})

	t.Run("should return transaction and user response when transfer success", func(t *testing.T) {
		// given
		var ctx context.Context

		transaction := &dto.TransactionRequest{
			SenderWalletNumber:    "1",
			RecipientWalletNumber: "2",
			Amount:                decimal.NewFromInt(10000),
		}
		expected := &entity.Transaction{Id: 1}
		expectedUser := &dto.UserTransferResponse{
			Recipient: "Boba",
			Sender:    "Boba",
		}
		mockWalletRepo := new(repomock.Wallet)
		mockUserRepo := new(repomock.User)
		mockGachaRepo := new(repomock.Gacha)
		mockTransactionRepo := new(repomock.Transaction)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.RecipientWalletNumber).Return([]entity.Wallet{{WalletNumber: "2"}}, nil)
		mockUserRepo.On("FindById", ctx, mock.Anything).Return([]entity.User{{Id: 1, Name: "Boba"}}, nil)
		mockWalletRepo.On("GetByWalletNumber", ctx, transaction.SenderWalletNumber).Return([]entity.Wallet{{WalletNumber: "1", Balance: decimal.NewFromInt(100000)}}, nil)
		mockTransactionRepo.On("Create", ctx, transaction).Return(expected, nil)
		mockWalletRepo.On("IncreaseBalance", ctx, mock.Anything).Return(&entity.Wallet{UserId: 1}, nil)
		mockWalletRepo.On("DecreaseBalance", ctx, mock.Anything).Return(&entity.Wallet{UserId: 1}, nil)
		u := usecase.NewTransactionImplementation(mockTransactionRepo, mockWalletRepo, mockUserRepo, mockGachaRepo)

		// when
		transactionRes, userRes, err := u.CreateTransfer(ctx, transaction)

		// then
		assert.Nil(t, err)
		assert.Equal(t, expected, transactionRes)
		assert.Equal(t, expectedUser, userRes)
	})
}

func TestTransactionImplementation_WithTrx(t *testing.T) {
	t.Run("should return transaction implementation when withTrx success", func(t *testing.T) {
		// Mocks setup
		txHandle := &sql.Tx{}
		repo := &repository.TransactionRepository{}
		mockTransactionRepo := new(repomock.Transaction)
		mockTransactionRepo.On("WithTrx", mock.Anything).Return(*repo)
		u := usecase.NewTransactionImplementation(mockTransactionRepo, nil, nil, nil)

		// Call the method under test
		modifiedU := u.WithTrx(txHandle)

		// Assertions
		if &modifiedU == nil {
			t.Error("WithTrx() should not return nil")
		}
		if modifiedU.TransactionRepository == nil {
			t.Error("WithTrx() did not set the transaction handle correctly")
		}
	})
}
