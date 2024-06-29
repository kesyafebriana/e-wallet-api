package dto

import (
	"time"

	"github.com/kesyafebriana/e-wallet-api/internal/entity"
)

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,max=50,email"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type UserResponse struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Wallet    any        `json:"wallet"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type UserTransferResponse struct {
	Sender    string `json:"sender_name"`
	Recipient string `json:"recipient_name"`
}

func ConvertFromUserEntity(userEntity *entity.User, walletEntity *entity.Wallet) *UserResponse {
	walletRes := ConvertFromWalletEntity(walletEntity)

	response := &UserResponse{
		Id:        userEntity.Id,
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		Wallet:    walletRes,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
		DeletedAt: nil,
	}

	if userEntity.DeletedAt != nil {
		response.DeletedAt = userEntity.DeletedAt
	}

	return response
}
