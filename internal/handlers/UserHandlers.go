package handlers

import (
	"GOproject/internal/userService"
	"GOproject/internal/web/users"
	"context"
)

type UserHandler struct {
	userService userService.UserService
}

func NewUserHandler(userService userService.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	// Получение всех сообщений из сервиса
	allUsers, err := h.userService.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	// Заполняем слайс response всеми сообщениями из БД
	for _, msg := range allUsers {
		user := users.User{
			Id:       &msg.ID,
			Email:    &msg.Email,
			Password: &msg.Password,
		}
		response = append(response, user)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{Password: *userRequest.Password, Email: *userRequest.Email}
	createdUser, err := h.userService.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func (h *UserHandler) DeleteUsers(_ context.Context, request users.DeleteUsersRequestObject) (users.DeleteUsersResponseObject, error) {
	userRequest := request.Body
	Id := userRequest.Id
	err := h.userService.DeleteUserByID(int(*Id))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *UserHandler) PatchUsers(_ context.Context, request users.PatchUsersRequestObject) (users.PatchUsersResponseObject, error) {
	userRequest := request.Body
	NewUser := userService.User{Email: *userRequest.Email, Password: *userRequest.Password}
	Id := request.Body.Id
	changedUser, err := h.userService.UpdateUserByID(int(*Id), NewUser)
	if err != nil {
		return nil, err
	}
	response := users.PatchUsers201JSONResponse{
		Id:       Id,
		Email:    &changedUser.Email,
		Password: &changedUser.Password,
	}

	return response, nil
}
