package handler

import (
	"net/http"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/kesyafebriana/e-wallet-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Gacha struct {
	GachaUsecase       usecase.Gacha
	TransactionUsecase usecase.Transaction
}

func NewGacha(GachaUsecase usecase.Gacha, TransactionUsecase usecase.Transaction) *Gacha {
	return &Gacha{
		GachaUsecase:       GachaUsecase,
		TransactionUsecase: TransactionUsecase,
	}
}

func (h *Gacha) GetGacha(ctx *gin.Context) {
	gacha, err := h.GachaUsecase.GetAll(ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

	dto.ResponseSuccessJSON(ctx, http.StatusOK, constant.OKMsg, gacha)
}

func (h *Gacha) SelectGacha(ctx *gin.Context) {
	request := &dto.GachaRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.Error(err)
		return
	}

	userId, ok := ctx.Get(constant.MyUserId)
	if !ok {
		err := apperror.StatusInternalServerError(constant.ErrorInternalServer, constant.InternalServerErrorMsg)
		ctx.Error(err)
		return
	}

	request.UserId = int64(userId.(float64))

	selectedGacha, err := h.GachaUsecase.SelectGacha(ctx, request)

	if err != nil {
		ctx.Error(err)
		return
	}

	walletNumber, ok := ctx.Get(constant.MyWalletNumber)
	if !ok {
		err := apperror.StatusInternalServerError(constant.ErrorInternalServer, constant.InternalServerErrorMsg)
		ctx.Error(err)
		return
	}

	description := "Reward from Gacha"

	topUp := &dto.TransactionRequest{
		RecipientWalletNumber: walletNumber.(string),
		SenderWalletNumber:    walletNumber.(string),
		Amount:                selectedGacha.Amount,
		SourceOfFund:          "Reward",
		Description:           &description,
	}

	_, err = h.TransactionUsecase.CreateTopUp(ctx, topUp)

	if err != nil {
		ctx.Error(err)
		return
	}

	dto.ResponseSuccessJSON(ctx, http.StatusOK, constant.OKMsg, topUp)
}
