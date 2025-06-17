// Package database provides database connection management.
//
//nolint:exhaustruct
package database

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ErrNoTaskTable указывает отсутствие таблицы 'tasks' в БД (psql код: 42P01).
var ErrNoTaskTable = errors.New("таблица 'tasks' не найдена (код: 42P01)")

// InitDB initialized database connection.
//
// Returns:
// - error if failed connect to database.
// - error if required table 'task' not exist.
func InitDB() (*gorm.DB, error) {
	dsn := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableAutomaticPing:                     true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed connect to database, %v ", err)
	}

	// tableExists обрабатывает ошибку отсутствия созданной таблицы tasks в БД.
	var tableExists bool
	err = DB.Raw(`
        SELECT EXISTS (
            SELECT 1 
            FROM pg_catalog.pg_tables 
            WHERE schemaname = 'public' 
            AND tablename = 'tasks'
        )
    `).Scan(&tableExists).Error

	if err != nil {
		return nil, fmt.Errorf("ошибка проверки таблицы: %w", err)
	}

	if !tableExists {
		return nil, ErrNoTaskTable
	}

	return DB, nil
}
