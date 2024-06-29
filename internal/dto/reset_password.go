package dto

type ResetPasswordRequest struct {
	Id          int64  `json:"-"`
	UserId      int64  `json:"-"`
	NewPassword string `json:"new_password" binding:"required,min=8,max=50"`
	Token       string `json:"token" binding:"required"`
}
