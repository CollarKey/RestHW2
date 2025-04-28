package main

import (
	"CheckingErrorsHW2/internal/database"
	"CheckingErrorsHW2/internal/handlers"
	"CheckingErrorsHW2/internal/taskService"
	"CheckingErrorsHW2/internal/web/tasks"
	//"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	//"net/http"
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

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	log.Println("Server started at localhost:8080")
	err := e.Start(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
