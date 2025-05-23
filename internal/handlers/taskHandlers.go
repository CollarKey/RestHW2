package handlers

import (
	"CheckingErrorsHW2/internal/taskService"
	"CheckingErrorsHW2/internal/web/tasks"
	"context"
	//"encoding/json"
	//"fmt"
	//"github.com/gorilla/mux"
	"log"
	//"net/http"
	//"strconv"
	//"golang.org/x/crypto/ssh"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	if id == 0 {
		return tasks.DeleteTasksId400Response{}, nil
	}

	err := h.Service.DeleteTask(id)
	if err != nil {
		return tasks.DeleteTasksId404Response{}, nil
	}

	return tasks.DeleteTasksId204Response{}, nil
}

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

	taskToUpdate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(uint(id), taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		IsDone: &updatedTask.IsDone,
		Task:   &updatedTask.Task,
	}

	return response, nil
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
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

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	if request.Body == nil {
		log.Printf("Invalid request: request body is nil")
		return nil, nil
	}

	taskToCreate := taskService.Task{
		Task:   "",
		IsDone: false,
	}

	if request.Body.Task != nil {
		taskToCreate.Task = *taskRequest.Task
	}
	if request.Body.IsDone != nil {
		taskToCreate.IsDone = *taskRequest.IsDone
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

//func (h *Handler) PatchHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	idStr := vars["id"]
//	if idStr == "" {
//		http.Error(w, "Incorrect ID, no ID specified", http.StatusBadRequest)
//		return
//	}
//
//	var task taskService.Task
//	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	id, err := strconv.ParseUint(idStr, 10, 32)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	updatedTask, err := h.Service.UpdateTaskByID(uint(id), task)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	if err := json.NewEncoder(w).Encode(updatedTask); err != nil {
//		log.Printf("Could not encode response %v", err)
//		return
//	}
//}
//
//func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	StrId := vars["id"]
//	if StrId == "" {
//		http.Error(w, "No ID specified", http.StatusBadRequest)
//		return
//	}
//
//	id, err := strconv.ParseUint(StrId, 10, 32)
//	if err != nil {
//		http.Error(w, "Invalid ID format", http.StatusBadRequest)
//		return
//	}
//
//	if err := h.Service.DeleteTask(uint(id)); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//}
