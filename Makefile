# Makefile для создания миграций

# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -source file://migrations -database $(DB_DSN)

.PHONY: migrate migrate-down migrate-down-all

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграции
migrate-down:
	$(MIGRATE) down

# Откат всех миграций
migrate-down-all:
	$(MIGRATE) down -all

# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/app/main.go # При вызове make run мы запустим наш сервер


# проврка версии migrate
migrate-version:
	migrate -version