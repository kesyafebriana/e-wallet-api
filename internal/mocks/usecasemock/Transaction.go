// Code generated by mockery v2.10.4. DO NOT EDIT.

package usecasemock

import (
	context "context"

	dto "github.com/kesyafebriana/e-wallet-api/internal/dto"
	entity "github.com/kesyafebriana/e-wallet-api/internal/entity"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"

	usecase "github.com/kesyafebriana/e-wallet-api/internal/usecase"
)

// Transaction is an autogenerated mock type for the Transaction type
type Transaction struct {
	mock.Mock
}

// CreateTopUp provides a mock function with given fields: ctx, request
func (_m *Transaction) CreateTopUp(ctx context.Context, request *dto.TransactionRequest) (*entity.Transaction, error) {
	ret := _m.Called(ctx, request)

	var r0 *entity.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, *dto.TransactionRequest) *entity.Transaction); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.TransactionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTransfer provides a mock function with given fields: ctx, request
func (_m *Transaction) CreateTransfer(ctx context.Context, request *dto.TransactionRequest) (*entity.Transaction, *dto.UserTransferResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *entity.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, *dto.TransactionRequest) *entity.Transaction); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transaction)
		}
	}

	var r1 *dto.UserTransferResponse
	if rf, ok := ret.Get(1).(func(context.Context, *dto.TransactionRequest) *dto.UserTransferResponse); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*dto.UserTransferResponse)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *dto.TransactionRequest) error); ok {
		r2 = rf(ctx, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllByWalletNumber provides a mock function with given fields: ctx, walletNumber, pagination
func (_m *Transaction) GetAllByWalletNumber(ctx context.Context, walletNumber string, pagination *dto.PaginationInfo) ([]entity.Transaction, *entity.Pagination, error) {
	ret := _m.Called(ctx, walletNumber, pagination)

	var r0 []entity.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, string, *dto.PaginationInfo) []entity.Transaction); ok {
		r0 = rf(ctx, walletNumber, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Transaction)
		}
	}

	var r1 *entity.Pagination
	if rf, ok := ret.Get(1).(func(context.Context, string, *dto.PaginationInfo) *entity.Pagination); ok {
		r1 = rf(ctx, walletNumber, pagination)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*entity.Pagination)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, *dto.PaginationInfo) error); ok {
		r2 = rf(ctx, walletNumber, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// WithTrx provides a mock function with given fields: _a0
func (_m *Transaction) WithTrx(_a0 *sql.Tx) usecase.TransactionImplementation {
	ret := _m.Called(_a0)

	var r0 usecase.TransactionImplementation
	if rf, ok := ret.Get(0).(func(*sql.Tx) usecase.TransactionImplementation); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(usecase.TransactionImplementation)
	}

	return r0
}