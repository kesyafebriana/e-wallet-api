package constant

import "errors"

var (
	ErrorEmailNotFound        = errors.New("email not found")
	ErrorUnknownClaims        = errors.New("unknown claims")
	ErrorDuplicateRecord      = errors.New("duplicate record")
	ErrorAccountNotFound      = errors.New("email not found")
	ErrorWalletNumberNotFound = errors.New("wallet number not found")
	ErrorWrongPassword        = errors.New("wrong password")
	ErrorInvalidToken         = errors.New("invalid token")
	ErrorUnauthorized         = errors.New("unauthorized")
	ErrorAmountValidation     = errors.New("amount validation error")
	ErrorInsufficientBalance  = errors.New("insufficient balance")
	ErrorInternalServer       = errors.New("internal server error")
	ErrorNoWallet             = errors.New("this account doesn't have a wallet")
	ErrorTransferToAccount    = errors.New("transfer to own account")
	ErrorNoDbTrx              = errors.New("database transaction not found in context")
	ErrorInvalidDbTrx         = errors.New("invalid type for database transaction in context")
	ErrorNoAttempt            = errors.New("account have no gacha attempt")
)
