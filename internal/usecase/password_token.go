package usecase

import (
	"context"
	"time"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/entity"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	helper "github.com/kesyafebriana/e-wallet-api/internal/pkg/helper"
	"github.com/kesyafebriana/e-wallet-api/internal/repository"
)

type Token interface {
	Create(ctx context.Context, request *dto.ResetTokenRequest) (*entity.PasswordTokens, error)
	UpdatePassword(ctx context.Context, request *dto.ResetPasswordRequest) (*entity.PasswordTokens, error)
}

type TokenImplementation struct {
	TokenRepository repository.Token
	UserRepository  repository.User
}

func NewTokenImplementation(TokenRepository repository.Token, UserRepository repository.User) *TokenImplementation {
	return &TokenImplementation{
		TokenRepository: TokenRepository,
		UserRepository:  UserRepository,
	}
}

func (u *TokenImplementation) Create(ctx context.Context, request *dto.ResetTokenRequest) (*entity.PasswordTokens, error) {
	var responseToken *entity.PasswordTokens


	exists, err := u.UserRepository.FindByEmail(ctx, request.Email)

	if err != nil {
		return nil, err
	}

	if len(exists) == 0 {
		return nil, apperror.StatusBadRequest(constant.ErrorEmailNotFound, constant.UserNotRegisteredMsg)
	}

	token := helper.GetToken(10)

	request.UserId = exists[0].Id
	request.Token = token

	responseToken, err = u.TokenRepository.Create(ctx, request)

	if err != nil {
		return nil, err
	}

	return responseToken, nil
}

func (u *TokenImplementation) UpdatePassword(ctx context.Context, request *dto.ResetPasswordRequest) (*entity.PasswordTokens, error) {
	var responseToken *entity.PasswordTokens
	hashUtil := &helper.HashImplementation{}

	exists, err := u.TokenRepository.FindToken(ctx, request)

	if err != nil {
		return nil, err
	}

	if len(exists) == 0 {
		return nil, apperror.StatusBadRequest(constant.ErrorInvalidToken, constant.InvalidResetTokenMsg)
	}

	if !time.Now().Before(*helper.Timezone(exists[0].ExpiredAt)) {
		return nil, apperror.StatusBadRequest(constant.ErrorInvalidToken, constant.ExpiredResetTokenMsg)
	}

	if exists[0].DeletedAt != nil {
		return nil, apperror.StatusBadRequest(constant.ErrorInvalidToken, constant.UsedTokenMsg)
	}

	newPassword, err := hashUtil.HashPassword(request.NewPassword)

	if err != nil {
		return nil, err
	}

	request.Id = exists[0].Id
	request.UserId = exists[0].UserId
	request.NewPassword = newPassword

	_, err = u.UserRepository.UpdatePassword(ctx, request)

	if err != nil {
		return nil, err
	}

	responseToken, err = u.TokenRepository.DeleteToken(ctx, request)
	
	if err != nil {
		return nil, err
	}

	return responseToken, nil
}
