package repository

import (
	"context"
	"database/sql"

	"github.com/kesyafebriana/e-wallet-api/internal/entity"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/kesyafebriana/e-wallet-api/internal/query"
)

type Gacha interface {
	GetAllGacha(ctx context.Context) ([]entity.Gacha, error)
	FindGachaAttemptByUserId(ctx context.Context, userId int64) ([]entity.Gacha, error)
	AddGachaAttempt(ctx context.Context, userId int64) (*entity.GachaAttempt, error)
	DeleteGachaAttempt(ctx context.Context, id int64) error
}

type GachaRepository struct {
	db *sql.DB
}

func NewGachaRepository(db *sql.DB) *GachaRepository {
	return &GachaRepository{
		db: db,
	}
}

func (r *GachaRepository) GetAllGacha(ctx context.Context) ([]entity.Gacha, error) {
	res := []entity.Gacha{}

	rows, err := r.db.QueryContext(ctx, query.FindAllGacha)

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	defer rows.Close()

	for rows.Next() {
		var responseGacha entity.Gacha

		err := rows.Scan(&responseGacha.Id,
			&responseGacha.Amount)

		if err != nil {
			return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
		}

		res = append(res, responseGacha)
	}

	err = rows.Err()

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	return res, nil
}

func (r *GachaRepository) FindGachaAttemptByUserId(ctx context.Context, userId int64) ([]entity.Gacha, error) {
	res := []entity.Gacha{}

	rows, err := r.db.QueryContext(ctx, query.FindOneGacha, userId)
	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}
	defer rows.Close()

	for rows.Next() {
		var gacha entity.Gacha

		err := rows.Scan(&gacha.Id)
		if err != nil {
			return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
		}
		res = append(res, gacha)
	}

	err = rows.Err()
	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	return res, nil
}

func (r *GachaRepository) AddGachaAttempt(ctx context.Context, userId int64) (*entity.GachaAttempt, error) {
	res := &entity.GachaAttempt{}

	err := r.db.QueryRowContext(ctx, query.CreateGachaAttempt, userId).Scan(&res.Id, &res.UserId, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return nil, apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	return res, nil
}

func (r *GachaRepository) DeleteGachaAttempt(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, query.DeleteGachaAttempt, id)

	if err != nil {
		return apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
	}

	return nil
}
