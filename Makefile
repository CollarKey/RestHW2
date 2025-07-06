# DB_DSN и MIGRATE Переменные которые будут использоваться в наших командах (Таргетах).
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -source file://migrations -database $(DB_DSN)

# Исполнение миграции, даже если файл миграции создан.
.PHONY: migrate migrate-down migrate-down-all

# migrate-new Таргет для создания новой миграции.
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# migrate Применение миграций.
migrate:
	$(MIGRATE) up

# migrate-down Откат миграции.
migrate-down:
	$(MIGRATE) down

# migrate-down-all Откат всех миграций.
migrate-down-all:
	$(MIGRATE) down -all

# run запускает приложение.
run:
	go run cmd/app/main.go # При вызове make run мы запустим наш сервер

# проверка версии migrate.
migrate-version:
	migrate -version

# lint проверка линтерами.
lint:
	golangci-lint run --color=always

# genUsers исполнение кодогенерации для Users.
genUsers:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

# genTasks исполнение кодогенерации для Tasks.
genTasks:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go