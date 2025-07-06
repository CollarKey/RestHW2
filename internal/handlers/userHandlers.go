// Package handlers userHandlers.go содержит обработчики HTTP-запросов для users, прием запросов,
// первичную валидацию данных, обработку ошибок и передачу данных на низлежащие слои бизнес-логики.
//nolint: exhaustruct, ireturn
package handlers

import (
	"CheckingErrorsHW2/internal/projecterrors"
	"CheckingErrorsHW2/internal/userservice"
	"CheckingErrorsHW2/internal/web/users"
	"context"
	"fmt"
)

// UserHandler является HTTP-обработчиком, содержащим ссылку на сервис бизнес-логики
// используется для обработки входящих HTTP-запросов и взаимодействия с сервисами.
type UserHandler struct {
	Service *userservice.UserService
}

// NewUserHandler создает новый экземпляр UserHandler,
// является точкой входа для вызова слоя бизнес-логики userService и возвращает результат клиенту.
func NewUserHandler(service *userservice.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// PostUsers производит обработку HTTP-запроса на создание User.
func (u UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	if userRequest == nil {
		return nil, projecterrors.ErrReqBodyNilUser
	}

	if userRequest.Email == nil {
		return nil, fmt.Errorf("%w", projecterrors.ErrEmailRequired)
	}

	if userRequest.Password == nil {
		return nil, fmt.Errorf("%w", projecterrors.ErrPasswordRequired)
	}

	userToCreate := userservice.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := u.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, fmt.Errorf("failed to create the user: %w", err)
	}

	response := users.PostUsers201JSONResponse{
		Id: &createdUser.ID,
		Email: &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

// GetUsers производит обработку HTTP-запроса на получение слайса всех Users.
func (u UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetAllUsers()
	if err != nil{
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:     &usr.ID,
			Email:   &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

// GetUsersId производит обработку HTTP-запроса на получение User по его ID.
//nolint:revive
func (u UserHandler) GetUsersId(_ context.Context, request users.GetUsersIdRequestObject) (users.GetUsersIdResponseObject, error) {
	id := request.Id

	if id == 0 {
		return users.GetUsersId400Response{}, nil
	}

	user, err := u.Service.GetUsersByID(id);
	if err != nil {
		return users.GetUsersId404Response{}, projecterrors.ErrNotFoundUser
	}

	response := users.GetUsersId200JSONResponse{
		Id: &user.ID,
		Email: &user.Email,
		Password: &user.Password,
	}

	return response, nil
}

// PatchUsersId производит обработку HTTP-запроса на изменение User по его ID.
//nolint:revive
func (u UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id
	userRequest := request.Body

	if id == 0 {
		return users.PatchUsersId400Response{}, nil
	}

	errResponse400 := users.PatchUsersId400Response{}
	if userRequest.Email == nil && userRequest.Password == nil {
		return errResponse400, nil
	}

	userToUpdate, err := u.Service.GetUsersByID(id)
	if err != nil {
		return users.PatchUsersId404Response{}, projecterrors.ErrNotFoundUser
	}

	if userRequest.Email != nil {
		userToUpdate.Email = *userRequest.Email
	}

	if userRequest.Password != nil {
		userToUpdate.Password = *userRequest.Password
	}

	updatedUser, err := u.Service.UpdateUserByID(id, userToUpdate)
	if err != nil {
		return nil, fmt.Errorf("failed to update user by ID: %w", err)
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return response, nil
}

// DeleteUsersId производит обработку HTTP-запроса на удаление User.
//nolint:revive
func (u UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id

	if id == 0 {
		return users.DeleteUsersId400Response{}, nil
	}

	err := u.Service.DeleteUser(id)
	if err != nil {
		return users.DeleteUsersId404Response{}, projecterrors.ErrNotFoundUser
	}

	return users.DeleteUsersId204Response{}, nil
}