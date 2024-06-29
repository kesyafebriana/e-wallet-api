// Code generated by mockery v2.10.4. DO NOT EDIT.

package usecasemock

import (
	context "context"

	dto "github.com/kesyafebriana/e-wallet-api/internal/dto"
	entity "github.com/kesyafebriana/e-wallet-api/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *User) Create(ctx context.Context, user *dto.UserRequest) (*entity.User, *entity.Wallet, error) {
	ret := _m.Called(ctx, user)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, *dto.UserRequest) *entity.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 *entity.Wallet
	if rf, ok := ret.Get(1).(func(context.Context, *dto.UserRequest) *entity.Wallet); ok {
		r1 = rf(ctx, user)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*entity.Wallet)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *dto.UserRequest) error); ok {
		r2 = rf(ctx, user)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetProfile provides a mock function with given fields: ctx, userId
func (_m *User) GetProfile(ctx context.Context, userId int64) (*entity.User, *entity.Wallet, error) {
	ret := _m.Called(ctx, userId)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.User); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 *entity.Wallet
	if rf, ok := ret.Get(1).(func(context.Context, int64) *entity.Wallet); ok {
		r1 = rf(ctx, userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*entity.Wallet)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, int64) error); ok {
		r2 = rf(ctx, userId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Login provides a mock function with given fields: ctx, user
func (_m *User) Login(ctx context.Context, user *dto.LoginRequest) (*dto.LoginResponse, error) {
	ret := _m.Called(ctx, user)

	var r0 *dto.LoginResponse
	if rf, ok := ret.Get(0).(func(context.Context, *dto.LoginRequest) *dto.LoginResponse); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.LoginResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dto.LoginRequest) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
