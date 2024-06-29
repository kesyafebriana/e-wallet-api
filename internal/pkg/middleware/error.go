package middleware

import (
	"errors"
	"net/http"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			firstError := c.Errors[0].Err
			errResponse := checkError(firstError)
			c.AbortWithStatusJSON(errResponse.Code, errResponse)
			return
		}

	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "max":
		return "Should be less than " + fe.Param()
	case "min":
		return "Should be greater than " + fe.Param()
	case "email":
		return "Should input valid email"
	}
	return "Unknown error"
}

func checkError(err error) dto.ErrorResponse {
	var ve validator.ValidationErrors
	var appErr *apperror.Error

	if errors.As(err, &ve) {
		details := GenerateValidtionErrs(ve)
		return dto.ErrorResponse{Code: http.StatusBadRequest, Message: constant.ValidationErrorMsg, Details: details}
	}

	if errors.As(err, &appErr) {
		return dto.ErrorResponse{Code: appErr.GetStatusCode(), Message: appErr.GetErrorMessage()}
	}

	return dto.ErrorResponse{Code: http.StatusInternalServerError, Message: constant.InternalServerErrorMsg}
}

func GenerateValidtionErrs(ve validator.ValidationErrors) []dto.ValidationErrorResponse {
	details := make([]dto.ValidationErrorResponse, len(ve))

	for i, fe := range ve {
		details[i] = dto.ValidationErrorResponse{Field: fe.Field(), Message: getErrorMsg(fe)}
	}

	return details
}
