// package main starting the program and links structure of the project.
package main

import (
	"CheckingErrorsHW2/internal/database"
	"CheckingErrorsHW2/internal/handlers"
	"CheckingErrorsHW2/internal/taskservice"
	"CheckingErrorsHW2/internal/userservice"
	"CheckingErrorsHW2/internal/web/tasks"
	"CheckingErrorsHW2/internal/web/users"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// isTableTasksNotExistError проверяет наличие ошибки создания таблицы 'tasks' в БД.
func isTableTasksNotExistError(err error) bool {
	return strings.Contains(err.Error(), "таблица 'tasks' не найдена (код: 42P01)")
}

// isTableUsersNotExistError проверяет наличие ошибки создания таблицы 'users' в БД.
func isTableUsersNotExistError(err error) bool {
	return strings.Contains(err.Error(), "таблица 'users' не найдена (код: 42P01)")
}

func main() {
	dbConnect, dbErr := database.InitDB()
	if dbErr != nil {
		isTasksTableErr := isTableTasksNotExistError(dbErr)
		if isTasksTableErr {
			log.Fatalf("FATAL ERROR: %v\nВыполните: make migrate для Tasks", dbErr)
		}

		isUsersTableError := isTableUsersNotExistError(dbErr)
		if isUsersTableError {
			log.Fatalf("FATAL ERROR: %v\nВыполните: make migrate для Users", dbErr)
		}

		log.Fatal("Ошибки инициализации БД:", dbErr)
	}

	taskRepo := taskservice.NewTaskRepository(dbConnect)
	taskService := taskservice.NewService(taskRepo)
	taskHandler := handlers.NewHandler(taskService)

	userRepo := userservice.NewUserRepository(dbConnect)
	userService := userservice.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTaskHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	err := e.Start(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	log.Println("Server started at localhost:8080")
}
