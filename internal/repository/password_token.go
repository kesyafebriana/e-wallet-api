package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/entity"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/kesyafebriana/e-wallet-api/internal/query"
)

type Token interface {
	Create(ctx context.Context, request *dto.ResetTokenRequest) (*entity.PasswordTokens, error)
	FindToken(ctx context.Context, request *dto.ResetPasswordRequest) ([]entity.PasswordTokens, error)
	DeleteToken(ctx context.Context, request *dto.ResetPasswordRequest) (*entity.PasswordTokens, error)
}

type TokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (r *TokenRepository) Create(ctx context.Context, request *dto.ResetTokenRequest) (*entity.PasswordTokens, error) {
	var responseToken entity.PasswordTokens

	err := r.db.QueryRowContext(ctx, query.CreateToken, request.UserId, request.Token).Scan(
		&responseToken.Id,
		&responseToken.ExpiredAt,
		&responseToken.CreatedAt,
		&responseToken.UpdatedAt)

		log.Println(responseToken.ExpiredAt)
	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	res := entity.PasswordTokens{Id: responseToken.Id, UserId: responseToken.UserId, Token: request.Token, ExpiredAt: responseToken.ExpiredAt, CreatedAt: responseToken.CreatedAt, UpdatedAt: responseToken.UpdatedAt}
	return &res, nil
}

func (r *TokenRepository) FindToken(ctx context.Context, request *dto.ResetPasswordRequest) ([]entity.PasswordTokens, error) {
	res := []entity.PasswordTokens{}

	rows, err := r.db.QueryContext(ctx, query.FindToken, request.Token)
	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}
	defer rows.Close()

	for rows.Next() {
		var token entity.PasswordTokens

		err := rows.Scan(&token.Id, &token.UserId, &token.Token, &token.ExpiredAt, &token.DeletedAt)
		if err != nil {
			return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
		}
		res = append(res, token)
	}

	err = rows.Err()
	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	return res, nil
}

func (r *TokenRepository) DeleteToken(ctx context.Context, request *dto.ResetPasswordRequest) (*entity.PasswordTokens, error) {
	var responseToken entity.PasswordTokens

	err := r.db.QueryRowContext(ctx, query.DeleteResetToken, request.Id).Scan(&responseToken.Id, &responseToken.Token, &responseToken.UpdatedAt, &responseToken.DeletedAt)

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	return &responseToken, nil
}
