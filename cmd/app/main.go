package main

import (
	"GOproject/internal/database"
	"GOproject/internal/handlers"
	"GOproject/internal/messagesService"
	"GOproject/internal/userService"
	"GOproject/internal/web/messages"
	"GOproject/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&messagesService.Message{})
	if err != nil {
		log.Fatal(err)
	}

	messagesRepo := messagesService.NewMessageRepository(database.DB)
	messagesService := messagesService.NewService(messagesRepo)
	messagesHandler := handlers.NewMessagesHandler(messagesService)

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewService(userRepo)
	userHandler := handlers.NewUserHandler(*userService)
	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := messages.NewStrictHandler(messagesHandler, nil) // тут будет ошибка
	strictHandler2 := users.NewStrictHandler(userHandler, nil)
	messages.RegisterHandlers(e, strictHandler)
	users.RegisterHandlers(e, strictHandler2)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
