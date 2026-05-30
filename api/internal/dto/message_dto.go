package dto

type CreateMessageDto struct {
	Content string `json:"content" validate:"required"`
}

type MessageResponse struct {
	ID        uint         `json:"id"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
	Content   string       `json:"content"`
	RoomID    uint         `json:"room_id"`
	SenderID  uint         `json:"sender_id"`
	Sender    UserResponse `json:"sender"`
}
