package service

import (
	"github.com/prachotx/real-time-chat/api/internal/dto"
	"github.com/prachotx/real-time-chat/api/internal/model"
	"github.com/prachotx/real-time-chat/api/internal/repository"
)

type MessageService interface {
	Create(input dto.CreateMessageDto, roomID uint, userID uint) error
	FindByRoomID(roomID uint) ([]dto.MessageResponse, error)
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

func (s *messageService) FindByRoomID(roomID uint) ([]dto.MessageResponse, error) {
	messages, err := s.messageRepo.FindByRoomID(roomID)
	if err != nil {
		return nil, err
	}

	var result []dto.MessageResponse
	for _, m := range messages {
		result = append(result, dto.MessageResponse{
			ID:        m.ID,
			CreatedAt: m.CreatedAt.String(),
			UpdatedAt: m.UpdatedAt.String(),
			Content:   m.Content,
			RoomID:    m.RoomID,
			SenderID:  m.SenderID,
			Sender: dto.UserResponse{
				ID:       m.Sender.ID,
				Username: m.Sender.Username,
				Email:    m.Sender.Email,
			},
		})
	}
	return result, nil
}
