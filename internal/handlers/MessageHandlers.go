package handlers

import (
	"GOproject/internal/messagesService"
	"GOproject/internal/web/messages"
	"context"
	"gorm.io/gorm"
)

type Handler struct {
	Service *messagesService.MessageService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	// Получение всех сообщений из сервиса
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := messages.GetMessages200JSONResponse{}

	// Заполняем слайс response всеми сообщениями из БД
	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body
	messageToCreate := messagesService.Message{Text: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}

	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}

	return response, nil
}

func (h *Handler) DeleteMessages(_ context.Context, request messages.DeleteMessagesRequestObject) (messages.DeleteMessagesResponseObject, error) {
	messageRequest := request.Body
	id := messagesService.Message{
		Model: gorm.Model{ID: *messageRequest.Id},
	}
	err := h.Service.DeleteMessageByID(int(id.ID))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *Handler) PatchMessages(_ context.Context, request messages.PatchMessagesRequestObject) (messages.PatchMessagesResponseObject, error) {
	messageRequest := request.Body
	NewMessage := messagesService.Message{Text: *messageRequest.Message}
	id := messagesService.Message{
		Model: gorm.Model{ID: *messageRequest.Id},
	}
	changedMessage, err := h.Service.UpdateMessageByID(int(id.ID), NewMessage)
	if err != nil {
		return nil, err
	}
	response := messages.PatchMessages201JSONResponse{
		Id:      &id.ID,
		Message: &changedMessage.Text,
	}

	return response, nil
}
