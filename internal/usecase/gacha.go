package usecase

import (
	"context"
	"math/rand"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/entity"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/kesyafebriana/e-wallet-api/internal/repository"
)

type Gacha interface {
	GetAll(ctx context.Context) ([]entity.Gacha, error)
	SelectGacha(ctx context.Context, request *dto.GachaRequest) (*entity.Gacha, error)
}

type GachaImplementation struct {
	GachaRepository repository.Gacha
}

func NewGachaImplementation(GachaRepository repository.Gacha) *GachaImplementation {
	return &GachaImplementation{
		GachaRepository: GachaRepository,
	}
}

func (u *GachaImplementation) GetAll(ctx context.Context) ([]entity.Gacha, error) {
	res, err := u.GachaRepository.GetAllGacha(ctx)

	if err != nil {
		return nil, err
	}

	newGacha := []entity.Gacha{}

	perm := rand.Perm(9)
	for i, v := range perm {
		gacha := &entity.Gacha{
			Id:     int64(i) + 1,
			Amount: res[v].Amount,
		}

		newGacha = append(newGacha, *gacha)
	}

	return newGacha, nil
}

func (u *GachaImplementation) SelectGacha(ctx context.Context, request *dto.GachaRequest) (*entity.Gacha, error) {
	exists, err := u.GachaRepository.FindGachaAttemptByUserId(ctx, request.UserId)

	if err != nil {
		return nil, err
	}

	if len(exists) == 0 {
		return nil, apperror.StatusBadRequest(constant.ErrorNoAttempt, constant.NoGachaMsg)
	}
	currentGacha, err := u.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	selectedGacha := currentGacha[request.Id]

	err = u.GachaRepository.DeleteGachaAttempt(ctx, exists[0].Id)

	if err != nil {
		return nil, err
	}

	
	return &selectedGacha, nil
}
