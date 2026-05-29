package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/prachotx/real-time-chat/api/config"
	"github.com/prachotx/real-time-chat/api/internal/handler"
	"github.com/prachotx/real-time-chat/api/internal/middleware"
	"github.com/prachotx/real-time-chat/api/internal/model"
	"github.com/prachotx/real-time-chat/api/internal/repository"
	"github.com/prachotx/real-time-chat/api/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func main() {
	cfg := config.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})

	app := fiber.New(fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
	})

	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	api := app.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.Post("/login", authHandler.Login)
			auth.Post("/register", authHandler.Register)
			auth.Get("/me", middleware.AuthMiddleware, authHandler.Profile)
		}
	}

	app.Listen(":" + cfg.Port)
}
