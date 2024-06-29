package handler

import (
	"net/http"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/kesyafebriana/e-wallet-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Token struct {
	TokenUsecase usecase.Token
}

func NewToken(TokenUsecase usecase.Token) *Token {
	return &Token{
		TokenUsecase: TokenUsecase,
	}
}

func (h *Token) Create(ctx *gin.Context) {
	token := &dto.ResetTokenRequest{}
	if err := ctx.ShouldBindJSON(token); err != nil {
		ctx.Error(err)
		return
	}

	tokenRes, err := h.TokenUsecase.Create(ctx, token)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := dto.ConvertFromTokenEntity(tokenRes)

	dto.ResponseSuccessJSON(ctx, http.StatusCreated, constant.TokenSentMsg, response)
}

func (h *Token) ChangePassword(ctx *gin.Context) {
	request := &dto.ResetPasswordRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.Error(err)
		return
	}

	_, err := h.TokenUsecase.UpdatePassword(ctx, request)
	if err != nil {
		ctx.Error(err)
		return
	}

	dto.ResponseSuccessJSON(ctx, http.StatusOK, constant.PasswordChangedMsg, nil)
}
