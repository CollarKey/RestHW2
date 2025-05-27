// package main starting the program and links structure of the project.
package main

import (
	"CheckingErrorsHW2/internal/database"
	"CheckingErrorsHW2/internal/handlers"
	"CheckingErrorsHW2/internal/taskservice"
	"CheckingErrorsHW2/internal/web/tasks"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// isTableNotExistError проверяет наличие ошибки создания таблицы 'tasks' в БД.
func isTableNotExistError(err error) bool {
	return strings.Contains(err.Error(), "таблица 'tasks' не найдена (код: 42P01)")
}

func main() {
	dbConnect, dbErr := database.InitDB()
	if dbErr != nil {
		if isTableNotExistError(dbErr) {
			log.Fatalf("FATAL ERROR: %v\nВыполните: make migrate", dbErr)
		}

		log.Fatal("Ошибки инициализации БД:", dbErr)
	}

	repo := taskservice.NewTaskRepository(dbConnect)
	service := taskservice.NewService(repo)
	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	err := e.Start(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	log.Println("Server started at localhost:8080")
}