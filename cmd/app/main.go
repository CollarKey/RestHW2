package main

import (
	"CheckingErrorsHW2/internal/database"
	"CheckingErrorsHW2/internal/handlers"
	"CheckingErrorsHW2/internal/taskService"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

// Проверка наличия ошибки создания таблицы 'tasks' в БД
func isTableNotExistError(err error) bool {
	return strings.Contains(err.Error(), "таблица 'tasks' не найдена (код: 42P01)")
}

func main() {
	dbConnect, dbErr := database.InitDB()
	if dbErr != nil {
		if isTableNotExistError(dbErr) {
			log.Fatalf("FATAL ERROR: %v\nВыполните: make migrate", dbErr)
		}
		log.Fatal("Ошибки иницилизации БД:", dbErr)
	}

	repo := taskService.NewTaskRepository(dbConnect)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", handler.GetHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/tasks", handler.PostHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/tasks/{id}", handler.PatchHandler).Methods(http.MethodPatch)
	router.HandleFunc("/api/tasks/{id}", handler.DeleteHandler).Methods(http.MethodDelete)

	log.Println("Server started at localhost:8080")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
