package handlers

import (
	"GOproject/internal/messagesService"
	"GOproject/internal/web/messages"
	"context"
)

type MessagesHandler struct {
	messageService *messagesService.MessageService
}

// Нужна для создания структуры MessagesHandler на этапе инициализации приложения

func NewMessagesHandler(messagesService *messagesService.MessageService) *MessagesHandler {
	return &MessagesHandler{
		messageService: messagesService,
	}
}

func (h *MessagesHandler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	// Получение всех сообщений из сервиса
	allMessages, err := h.messageService.GetAllMessages()
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

func (h *MessagesHandler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body
	messageToCreate := messagesService.Message{Text: *messageRequest.Message}
	createdMessage, err := h.messageService.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}

	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}

	return response, nil
}

func (h *MessagesHandler) DeleteMessages(_ context.Context, request messages.DeleteMessagesRequestObject) (messages.DeleteMessagesResponseObject, error) {
	messageRequest := request.Body
	Id := messageRequest.Id
	err := h.messageService.DeleteMessageByID(int(*Id))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *MessagesHandler) PatchMessages(_ context.Context, request messages.PatchMessagesRequestObject) (messages.PatchMessagesResponseObject, error) {
	messageRequest := request.Body
	NewMessage := messagesService.Message{Text: *messageRequest.Message}
	Id := request.Body.Id
	changedMessage, err := h.messageService.UpdateMessageByID(int(*Id), NewMessage)
	if err != nil {
		return nil, err
	}
	response := messages.PatchMessages201JSONResponse{
		Id:      Id,
		Message: &changedMessage.Text,
	}

	return response, nil
}
