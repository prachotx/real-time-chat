package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/prachotx/real-time-chat/api/internal/dto"
	"github.com/prachotx/real-time-chat/api/internal/service"
	"github.com/prachotx/real-time-chat/api/pkg/response"
)

type MessageHandler struct {
	messageService service.MessageService
}

func NewMessageHandler(messageService service.MessageService) *MessageHandler {
	return &MessageHandler{messageService}
}

func (h *MessageHandler) Create(c fiber.Ctx) error {
	roomID, _ := strconv.Atoi(c.Params("id"))
	userID := c.Locals("user_id").(uint)

	var req dto.CreateMessageDto
	if err := c.Bind().Body(&req); err != nil {
		return response.Send(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	err := h.messageService.Create(req, uint(roomID), userID)
	if err != nil {
		return response.Send(c, fiber.StatusInternalServerError, "Failed to create message", nil)
	}

	return response.Send(c, fiber.StatusCreated, "Message created successfully", nil)
}
