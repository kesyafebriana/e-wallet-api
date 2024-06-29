package dto

import "github.com/kesyafebriana/e-wallet-api/internal/entity"

type ResetTokenRequest struct {
	Email  string `json:"email" binding:"required,email"`
	UserId int64  `json:"-"`
	Token  string `json:"-"`
}

type ResetTokenResponse struct {
	Token string `json:"token"`
}

func ConvertFromTokenEntity(passwordToken *entity.PasswordTokens) *ResetTokenResponse {
	return &ResetTokenResponse{
		Token: passwordToken.Token,
	}
}
