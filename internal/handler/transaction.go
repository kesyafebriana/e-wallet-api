package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/kesyafebriana/e-wallet-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	TransactionUsecase usecase.Transaction
}

func NewTransaction(TransactionUsecase usecase.Transaction) *Transaction {
	return &Transaction{
		TransactionUsecase: TransactionUsecase,
	}
}

func (h *Transaction) GetAll(ctx *gin.Context) {
	search := ctx.Query("s")
	sort := ctx.Query("sort")
	sortBy := ctx.Query("sortBy")
	startAt := ctx.Query("startAt")
	endAt := ctx.Query("endAt")
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	paginationReq := dto.PaginationInfo{
		Search:    &search,
		Sort:      &sort,
		SortBy:    &sortBy,
		StartDate: &startAt,
		EndDate:   &endAt,
		Page:      &page,
		Limit:     &limit,
	}

	walletNumber, ok := ctx.Get(constant.MyWalletNumber)
	if !ok {
		err := apperror.StatusInternalServerError(constant.ErrorInternalServer, constant.InternalServerErrorMsg)
		ctx.Error(err)
		return
	}

	transactionRes, paginationInfo, err := h.TransactionUsecase.GetAllByWalletNumber(ctx, walletNumber.(string), &paginationReq)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := dto.ConvertFromTransactionEntity(transactionRes, paginationInfo.TotalPage, paginationInfo.TotalData)

	dto.ResponseSuccessJSON(ctx, http.StatusOK, constant.OKMsg, response)
}

func (h *Transaction) TopUp(ctx *gin.Context) {
	request := &dto.TopUpRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.Error(err)
		return
	}

	walletNumber, ok := ctx.Get(constant.MyWalletNumber)
	if !ok {
		err := apperror.StatusInternalServerError(constant.ErrorInternalServer, constant.InternalServerErrorMsg)
		ctx.Error(err)
		return
	}

	transactionReq := dto.ConvertTopUpToTransactionRequest(request, walletNumber.(string))

	transactionRes, err := h.TransactionUsecase.CreateTopUp(ctx, transactionReq)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := dto.ConvertFromTopUpRequest(transactionRes)

	dto.ResponseSuccessJSON(ctx, http.StatusOK, constant.TopUpSuccessMsg, response)
}

func (h *Transaction) Transfer(ctx *gin.Context) {
	log.Println(ctx.Get("db_trx"))
	dbTxInterface, _ := ctx.Get("db_trx")
	if dbTxInterface == nil {
		err := apperror.StatusInternalServerError(constant.ErrorNoDbTrx, constant.InternalServerErrorMsg)
		ctx.Error(err)
		return
	}
	txHandle, ok := dbTxInterface.(*sql.Tx)
	log.Println(txHandle)
	log.Println(txHandle)
	if !ok {
		err := apperror.StatusInternalServerError(constant.ErrorInvalidDbTrx, constant.InternalServerErrorMsg)
		ctx.Error(err)
		return
	}

	request := &dto.TransferRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		txHandle.Rollback()
		ctx.Error(err)
		return
	}

	recipient := ctx.Query("to")

	if recipient == "" {
		err := apperror.StatusBadRequest(constant.ErrorWalletNumberNotFound, constant.RecipientWalletNotFoundMsg)
		txHandle.Rollback()
		ctx.Error(err)
		return
	}

	request.RecipientWalletId = recipient

	walletNumber, ok := ctx.Get(constant.MyWalletNumber)
	if !ok {
		err := apperror.StatusInternalServerError(constant.ErrorInternalServer, constant.InternalServerErrorMsg)
		txHandle.Rollback()
		ctx.Error(err)
		return
	}

	transactionReq := dto.ConvertTransferToTransactionRequest(request, walletNumber.(string))

	transactionRes, userRes, err := h.TransactionUsecase.WithTrx(txHandle).CreateTransfer(ctx, transactionReq)

	if err != nil {
		txHandle.Rollback()
		ctx.Error(err)
		return
	}

	response := dto.ConvertFromTransferRequest(transactionRes, userRes)

	dto.ResponseSuccessJSON(ctx, http.StatusOK, constant.TransferSuccessMsg, response)
}
