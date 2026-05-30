package dto

type CreateMessageDto struct {
	Content string `json:"content" validate:"required"`
}
