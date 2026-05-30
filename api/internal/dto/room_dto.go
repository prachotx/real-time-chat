package dto

type CreateRoomDto struct {
	Name string `json:"name" validate:"required"`
}

type RoomUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type RoomResponse struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	CreatedByID uint             `json:"created_by_id"`
	CreatedBy   RoomUserResponse `json:"created_by"`
}
