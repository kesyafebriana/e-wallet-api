package middleware

import (
	"strings"

	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	helper "github.com/kesyafebriana/e-wallet-api/internal/pkg/helper"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer_"
		authToken := c.GetHeader("Authorization")

		if authToken == "" {
			err := apperror.StatusUnauthorized(constant.ErrorUnauthorized, constant.UnauthorizedMsg)
			c.Error(err)
			return
		}

		if !strings.HasPrefix(authToken, "Bearer ") {
			err := apperror.StatusUnauthorized(constant.ErrorUnauthorized, constant.UnauthorizedMsg)
			c.Error(err)
			return
		}

		clientToken := authToken[len(BearerSchema):]

		helper := &helper.TokenImplementation{}

		jwt, err := helper.ParseAndVerify(clientToken)
		if err != nil {
			c.Error(err)
			return
		}

		userId, ok := jwt["user_id"].(float64)
		if !ok {
			err = apperror.StatusUnauthorized(constant.ErrorUnauthorized, constant.UnauthorizedMsg)
			c.Error(err)
			return
		}

		walletNumber, ok := jwt["wallet_number"].(string)
		if !ok {
			err = apperror.StatusUnauthorized(constant.ErrorUnauthorized, constant.UnauthorizedMsg)
			c.Error(err)
			return
		}

		c.Set(constant.MyUserId, userId)
		c.Set(constant.MyWalletNumber, walletNumber)
		c.Next()
	}
}
