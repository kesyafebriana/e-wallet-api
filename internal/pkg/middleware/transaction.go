package middleware

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/kesyafebriana/e-wallet-api/internal/pkg/apperror"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/constant"
	"github.com/gin-gonic/gin"
)

func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

func DBTransactionMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle, err := db.BeginTx(c, &sql.TxOptions{Isolation: sql.LevelSerializable})
		if err != nil {
			err := apperror.StatusInternalServerError(err, constant.InternalServerErrorMsg)
			c.Error(err)
			return
		}
		log.Print("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				log.Print("rolling back transaction due to status code: ", http.StatusInternalServerError)
				txHandle.Rollback()
			}
		}()

		c.Set("db_trx", txHandle)

		c.Next()
		if tx, ok := c.Get("db_trx"); ok {
			txHandle := tx.(*sql.Tx)
			if StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
				if err := txHandle.Commit(); err != nil {
					log.Print("trx commit error: ", err)
					return
				}
				log.Print("committed!")
			} else {
				log.Print("rolling back transaction due to status code: ", c.Writer.Status())
				txHandle.Rollback()
				return
			}
		} else {
			log.Println("transaction object not found in context")
		}
	}
}
