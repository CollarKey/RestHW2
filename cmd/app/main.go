package main

import (
	"CheckingErrorsHW2/internal/database"
	"CheckingErrorsHW2/internal/handlers"
	"CheckingErrorsHW2/internal/taskService"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var task string

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.RequestBody{})

	repo := taskService.NewTaskRepository(database.DB)
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
