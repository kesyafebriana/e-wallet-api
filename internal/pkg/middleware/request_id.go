package middleware

import (
	"net/http"

	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId(c *gin.Context) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Set(constant.RequestId, uuid)
	c.Next()
}
