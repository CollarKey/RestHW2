// Package handlers содержит обработчики HTTP-запросов, прием запросов, первичную валидацию данных,
// обработку ошибок и передачу данных на низлежащие слои бизнес-логики.
//
//nolint:exhaustruct
package handlers

import (
	"CheckingErrorsHW2/internal/taskservice"
	"CheckingErrorsHW2/internal/web/tasks"
	"context"
	"errors"
	"fmt"
	"log"
)

// ErrReqBodyNil проверяет, что тело запроса равно nil.
var ErrReqBodyNil = errors.New("request body cannot be nil")

// ErrNotFound указывает, что задача не найдена.
var ErrNotFound = errors.New("cannot find the Task")

// Handler является HTTP-обработчиком, содержащим ссылку на сервис бизнес-логики
// используется для обработки входящих HTTP-запросов и взаимодействия с сервисами.
type Handler struct {
	Service *taskservice.TaskService
}

// NewHandler создает новый экземпляр Handler,
// является точкой входа для вызова слоя бизнес-логики taskService и возвращает результат клиенту.
func NewHandler(service *taskservice.TaskService) *Handler {
	return &Handler{Service: service}
}

// DeleteTasksId обрабатывает HTTP-запрос на удаление Task по ID
// проводит первичную валидацию входных данных, перенаправляет запрос в нижний бизнес-слой для удаления Task
// возвращает соответствующий ответ в зависимости от результата операции.
//
//nolint:revive
func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	if id == 0 {
		return tasks.DeleteTasksId400Response{}, nil
	}

	err := h.Service.DeleteTask(id)
	if err != nil {
		return tasks.DeleteTasksId404Response{}, ErrNotFound
	}

	return tasks.DeleteTasksId204Response{}, nil
}

// PatchTasksId обрабатывает HTTP-запрос на обновление Task по ID
// проводит первичную валидацию входных данных, перенаправляет запрос в нижний бизнес-слой для обновления Task
// возвращает соответствующий ответ в зависимости от результата операции.
//
//nolint:revive
func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id
	taskRequest := request.Body

	errResp400 := tasks.PatchTasksId400Response{}
	if id == 0 {
		return errResp400, nil
	}

	if taskRequest.Task == nil && taskRequest.IsDone == nil {
		return errResp400, nil
	}

	taskToUpdate := taskservice.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(id, taskToUpdate)
	if err != nil {
		return nil, fmt.Errorf("failed to updated the task by ID: %w", err)
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		IsDone: &updatedTask.IsDone,
		Task:   &updatedTask.Task,
	}

	return response, nil
}

// GetTasks обрабатывает HTTP-запрос на получение всех Task
// возвращает соответствующий ответ в зависимости от результата операции.
func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %w", err)
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

// GetTaskByID обрабатывает HTTP-запрос на получение Task по ID.
func (h *Handler) GetTaskByID(_ context.Context, request tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	id := request.Id

	if id == 0 {
		return tasks.GetTasksId400Response{}, nil
	}

	task, err := h.Service.GetTaskByID(id)
	if err != nil {
		return tasks.GetTasksId404Response{}, ErrNotFound
	}

	response := tasks.GetTasksId200JSONResponse{
		Id:     &task.ID,
		Task:   &task.Task,
		IsDone: &task.IsDone,
	}

	return response, nil
}

// PostTasks обрабатывает HTTP-запрос на создание Task
// проводит первичную валидацию на nil тело запроса
// возвращает соответствующий ответ в зависимости от результата операции.
func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		log.Printf("Invalid request: request body is nil")

		return nil, ErrReqBodyNil
	}

	taskToCreate := taskservice.Task{
		Task:   "",
		IsDone: false,
	}

	if taskReq := request.Body.Task; taskReq != nil {
		taskToCreate.Task = *taskReq
	}

	if isDoneReq := request.Body.IsDone; isDoneReq != nil {
		taskToCreate.IsDone = *isDoneReq
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, fmt.Errorf("failed to create the task: %w", err)
	}

	response := &tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}