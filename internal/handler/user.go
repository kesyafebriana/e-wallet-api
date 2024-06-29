package handler

import (
	"net/http"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/kesyafebriana/e-wallet-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserUsecase usecase.User
}

func NewUser(UserUsecase usecase.User) *User {
	return &User{
		UserUsecase: UserUsecase,
	}
}

func (h *User) GetProfile(ctx *gin.Context) {
	userId, ok := ctx.Get(constant.MyUserId)
	if !ok {
		err := apperror.StatusInternalServerError(constant.ErrorInternalServer, constant.InternalServerErrorMsg)
		ctx.Error(err)
		return
	}

	userRes, walletRes, err := h.UserUsecase.GetProfile(ctx, int64(userId.(float64)))

	if err != nil {
		ctx.Error(err)
		return
	}

	response := dto.ConvertFromUserEntity(userRes, walletRes)

	dto.ResponseSuccessJSON(ctx, http.StatusOK, constant.OKMsg, response)
}

func (h *User) Register(ctx *gin.Context) {
	user := &dto.UserRequest{}
	if err := ctx.ShouldBindJSON(user); err != nil {
		ctx.Error(err)
		return
	}

	userRes, walletRes, err := h.UserUsecase.Create(ctx, user)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := dto.ConvertFromUserEntity(userRes, walletRes)

	dto.ResponseSuccessJSON(ctx, http.StatusCreated, constant.UserCreatedMsg, response)
}

func (h *User) Login(ctx *gin.Context) {
	loginData := &dto.LoginRequest{}
	if err := ctx.ShouldBindJSON(loginData); err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.UserUsecase.Login(ctx, loginData)
	if err != nil {
		ctx.Error(err)
		return
	}

	dto.ResponseSuccessJSON(ctx, http.StatusOK, constant.UserLoggedMsg, res)
}
