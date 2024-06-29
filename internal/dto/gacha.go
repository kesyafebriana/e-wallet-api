package dto

type GachaRequest struct {
	Id     int64 `json:"id" binding:"required,min=1,max=9"`
	UserId int64 `json:"user_id"`
}
