package handlers

import (
	"CheckingErrorsHW2/internal/taskService"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *taskService.TaskService
}

//

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Tasks
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "aaplication/json")
	if err := json.NewEncoder(w).Encode(createdTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) PatchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Incorrect ID, no ID specified", http.StatusBadRequest)
		return
	}

	var task taskService.Tasks
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTask, err := h.Service.UpdateTaskByID(uint(id), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedTask); err != nil {
		log.Printf("Could not encode response %v", err)
		return
	}
}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	StrId := vars["id"]
	if StrId == "" {
		http.Error(w, "No ID specified", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(StrId, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteTask(uint(id)); err != nil {

		//TODO: Добавить кастомные ошибки?
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
