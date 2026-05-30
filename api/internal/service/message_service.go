package service

import (
	"github.com/prachotx/real-time-chat/api/internal/dto"
	"github.com/prachotx/real-time-chat/api/internal/model"
	"github.com/prachotx/real-time-chat/api/internal/repository"
)

type MessageService interface {
	Create(input dto.CreateMessageDto, roomID uint, userID uint) error
}

type messageService struct {
	messageRepo repository.MessageRepository
}

func NewMessageService(messageRepo repository.MessageRepository) MessageService {
	return &messageService{messageRepo}
}

func (s *messageService) Create(input dto.CreateMessageDto, roomID uint, userID uint) error {
	message := model.Message{
		Content:  input.Content,
		RoomID:   roomID,
		SenderID: userID,
	}

	return s.messageRepo.Create(&message)
}
