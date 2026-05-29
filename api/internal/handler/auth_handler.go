package handler

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/prachotx/real-time-chat/api/internal/dto"
	"github.com/prachotx/real-time-chat/api/internal/service"
	"github.com/prachotx/real-time-chat/api/pkg/response"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var input dto.LoginDto
	if err := c.Bind().Body(&input); err != nil {
		return response.Send(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	tokenString, err := h.authService.Login(input)
	if err != nil {
		return response.Send(c, fiber.StatusUnauthorized, "Invalid credentials", nil)
	}

	cookie := fiber.Cookie{
		Name:     "access_token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	}

	c.Cookie(&cookie)

	return response.Send(c, fiber.StatusOK, "Login successful", nil)
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req dto.RegisterDto
	if err := c.Bind().Body(&req); err != nil {
		return response.Send(c, fiber.StatusUnprocessableEntity, err.Error(), nil)
	}

	err := h.authService.Register(req)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return response.Send(c, fiber.StatusConflict, "Email already exists", nil)
		}
		return response.Send(c, fiber.StatusInternalServerError, "Failed to register user", nil)
	}

	return response.Send(c, fiber.StatusOK, "Register successful", nil)
}

func (h *AuthHandler) Profile(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	profile, err := h.authService.FindProfile(userID)
	if err != nil {
		return response.Send(c, fiber.StatusNotFound, "Profile not found", nil)
	}

	return response.Send(c, fiber.StatusOK, "Profile found", profile)
}
