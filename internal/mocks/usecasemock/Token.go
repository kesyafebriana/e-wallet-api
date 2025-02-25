// Code generated by mockery v2.10.4. DO NOT EDIT.

package usecasemock

import (
	context "context"

	dto "github.com/kesyafebriana/e-wallet-api/internal/dto"
	entity "github.com/kesyafebriana/e-wallet-api/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// Token is an autogenerated mock type for the Token type
type Token struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, request
func (_m *Token) Create(ctx context.Context, request *dto.ResetTokenRequest) (*entity.PasswordTokens, error) {
	ret := _m.Called(ctx, request)

	var r0 *entity.PasswordTokens
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ResetTokenRequest) *entity.PasswordTokens); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.PasswordTokens)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.ResetTokenRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePassword provides a mock function with given fields: ctx, request
func (_m *Token) UpdatePassword(ctx context.Context, request *dto.ResetPasswordRequest) (*entity.PasswordTokens, error) {
	ret := _m.Called(ctx, request)

	var r0 *entity.PasswordTokens
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ResetPasswordRequest) *entity.PasswordTokens); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.PasswordTokens)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.ResetPasswordRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
