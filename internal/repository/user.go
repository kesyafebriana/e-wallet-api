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

type User interface {
	FindByEmail(ctx context.Context, email string) ([]entity.User, error)
	FindById(ctx context.Context, id int64) ([]entity.User, error)
	Create(ctx context.Context, user *dto.UserRequest) (*entity.User, error)
	UpdatePassword(ctx context.Context, user *dto.ResetPasswordRequest) (*entity.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) ([]entity.User, error) {
	res := []entity.User{}

	rows, err := r.db.QueryContext(ctx, query.FindUserByEmail, email)
	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}
	defer rows.Close()

	for rows.Next() {
		var User entity.User

		err := rows.Scan(&User.Id, &User.Name, &User.Email, &User.Password, &User.CreatedAt, &User.UpdatedAt)
		if err != nil {
			return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
		}
		res = append(res, User)
	}

	err = rows.Err()
	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	return res, nil
}

func (r *UserRepository) FindById(ctx context.Context, id int64) ([]entity.User, error) {
	res := []entity.User{}

	rows, err := r.db.QueryContext(ctx, query.FindUserById, id)
	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}
	defer rows.Close()

	for rows.Next() {
		var User entity.User

		err := rows.Scan(&User.Id, &User.Name, &User.Email, &User.Password, &User.CreatedAt, &User.UpdatedAt)
		if err != nil {
			return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
		}
		res = append(res, User)
	}

	err = rows.Err()
	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	return res, nil
}

func (r *UserRepository) Create(ctx context.Context, user *dto.UserRequest) (*entity.User, error) {
	var responseUser entity.User

	err := r.db.QueryRowContext(ctx, query.CreateUser, user.Name, user.Email, user.Password).Scan(&responseUser.Id, &responseUser.CreatedAt, &responseUser.UpdatedAt)

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	res := entity.User{Id: responseUser.Id, Name: user.Name, Email: user.Email, CreatedAt: responseUser.CreatedAt, UpdatedAt: responseUser.UpdatedAt}
	return &res, nil
}

func (r *UserRepository) UpdatePassword(ctx context.Context, user *dto.ResetPasswordRequest) (*entity.User, error) {
	var responseUser entity.User

	err := r.db.QueryRowContext(ctx, query.UpdatePassword, user.NewPassword, user.UserId).Scan(&responseUser.Id, &responseUser.UpdatedAt)

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	return &responseUser, nil
}
