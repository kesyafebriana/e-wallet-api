package usecase

import (
	"context"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/entity"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	helper "github.com/kesyafebriana/e-wallet-api/internal/pkg/helper"
	"github.com/kesyafebriana/e-wallet-api/internal/repository"
)

type User interface {
	GetProfile(ctx context.Context, userId int64) (*entity.User, *entity.Wallet, error)
	Create(ctx context.Context, user *dto.UserRequest) (*entity.User, *entity.Wallet, error)
	Login(ctx context.Context, user *dto.LoginRequest) (*dto.LoginResponse, error)
}

type UserImplementation struct {
	UserRepository   repository.User
	WalletRepository repository.Wallet
}

func NewUserImplementation(UserRepository repository.User, WalletRepository repository.Wallet) *UserImplementation {
	return &UserImplementation{
		UserRepository:   UserRepository,
		WalletRepository: WalletRepository,
	}
}

func (u *UserImplementation) GetProfile(ctx context.Context, userId int64) (*entity.User, *entity.Wallet, error) {
	userRes, err := u.UserRepository.FindById(ctx, userId)

	if err != nil {
		return nil, nil, err
	}

	walletRes, err := u.WalletRepository.GetByUserId(ctx, userId)

	if err != nil {
		return nil, nil, err
	}

	if len(walletRes) == 0 {
		return nil, nil, apperror.StatusInternalServerError(constant.ErrorNoWallet, constant.InternalServerErrorMsg)
	}

	return &userRes[0], &walletRes[0], nil
}

func (u *UserImplementation) Create(ctx context.Context, user *dto.UserRequest) (*entity.User, *entity.Wallet, error) {
	var userRes *entity.User
	var walletRes *entity.Wallet
	walletReq := &dto.WalletRequest{}
	hashUtil := &helper.HashImplementation{}

	exists, err := u.UserRepository.FindByEmail(ctx, user.Email)

	if err != nil {
		return nil, nil, err
	}

	if len(exists) > 0 {
		return nil, nil, apperror.StatusBadRequest(constant.ErrorDuplicateRecord, constant.DuplicateEmailMsg)
	}

	user.Password, err = hashUtil.HashPassword(user.Password)

	if err != nil {
		return nil, nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	userRes, err = u.UserRepository.Create(ctx, user)

	if err != nil {
		return nil, nil, err
	}

	walletReq.UserId = userRes.Id

	walletRes, err = u.WalletRepository.Create(ctx, walletReq)

	if err != nil {
		return nil, nil, err
	}

	return userRes, walletRes, nil
}

func (u *UserImplementation) Login(ctx context.Context, user *dto.LoginRequest) (*dto.LoginResponse, error) {
	var token string
	hashUtil := &helper.HashImplementation{}
	tokenUtil := &helper.TokenImplementation{}

	exists, err := u.UserRepository.FindByEmail(ctx, user.Email)

	if err != nil {
		return nil, err
	}

	if len(exists) == 0 {
		return nil, apperror.StatusBadRequest(constant.ErrorAccountNotFound, constant.InvalidLoginInputMsg)
	}

	isCorrect, err := hashUtil.CheckPassword(user.Password, []byte(exists[0].Password))

	if err != nil {
		return nil, apperror.StatusBadRequest(constant.ErrorWrongPassword, constant.InvalidLoginInputMsg)
	}

	walletRes, err := u.WalletRepository.GetByUserId(ctx, exists[0].Id)

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	if len(walletRes) == 0 {
		return nil, apperror.StatusInternalServerError(constant.ErrorNoWallet, constant.InternalServerErrorMsg)
	}

	if isCorrect {
		token, err = tokenUtil.CreateAndSign(exists[0].Id, walletRes[0].WalletNumber)

		if err != nil {
			return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
		}
	}

	res := &dto.LoginResponse{
		Token: token,
	}

	return res, nil
}
