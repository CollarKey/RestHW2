package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var task string

func getHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []requestBody

	answer := DB.Find(&tasks)
	if answer.Error != nil {
		http.Error(w, "Failed to get the tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Println(w, "Error encoding JSON", err)
		return
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if body.Task == "" {
		http.Error(w, "Task could not be empty", http.StatusBadRequest)
		return
	}

	process := DB.Create(&body)
	if process.Error != nil {
		http.Error(w, "Could not create the task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("Error encoding JSON: %s", err)
		return
	}
}

// Хочу обновлять статус IsDone по ID
func patchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Incorrect ID, no ID specified", http.StatusBadRequest)
		return
	}

	// Создаем структуру для защиты от перезатирвания полей,
	// Используем указатель, чтобы проверить, что поле передано - совет ИИ DS
	var updatedData struct {
		Task   *string `json:"task"`
		IsDone *bool   `json:"is_done"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if updatedData.Task == nil && updatedData.IsDone == nil {
		http.Error(w, "On of the fields 'Task' or 'Is_done' must be set", http.StatusBadRequest)
		return
	}

	var task requestBody
	if result := DB.First(&task, id); result.Error != nil {
		http.Error(w, "Could not find the task", http.StatusNotFound)
		return
	}

	if updatedData.Task != nil {
		task.Task = *updatedData.Task
	}

	if updatedData.IsDone != nil {
		task.IsDone = *updatedData.IsDone
	}

	if result := DB.Save(&task); result.Error != nil {
		http.Error(w, "Could not update task's status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Printf("Erorr encoding JSON in patchHandler %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error encoding JSON in patchHandler")
		return
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Incorrect ID, no ID specified", http.StatusBadRequest)
		return
	}

	var task requestBody
	if result := DB.First(&task, id); result.Error != nil {
		http.Error(w, "Could not find the task", http.StatusNotFound)
		return
	}

	if result := DB.Delete(&task); result.Error != nil {
		log.Printf("Error deleting task %v", result.Error)
		http.Error(w, "Could not delete task", http.StatusInternalServerError)
		return
	}
}

func main() {
	InitDB()

	// Здесь используйте &Task{}
	DB.AutoMigrate(&requestBody{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", getHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/tasks", postHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/tasks/{id}", patchHandler).Methods(http.MethodPatch)
	router.HandleFunc("/api/tasks/{id}", deleteHandler).Methods(http.MethodDelete)

	log.Println("Server started at localhost:8080")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
