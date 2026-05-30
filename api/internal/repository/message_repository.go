package repository

import (
	"github.com/prachotx/real-time-chat/api/internal/model"
	"gorm.io/gorm"
)

type MessageRepository interface {
	Create(message *model.Message) error
	FindByRoomID(roomID uint) ([]model.Message, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db}
}

func (r *messageRepository) Create(message *model.Message) error {
	return r.db.Create(message).Error
}

func (r *messageRepository) FindByRoomID(roomID uint) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Preload("Sender").Where("room_id = ?", roomID).Find(&messages).Error
	return messages, err
}
