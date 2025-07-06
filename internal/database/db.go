// Package database provides database connection management.
//
//nolint:exhaustruct
package database

import (
	"CheckingErrorsHW2/internal/projecterrors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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
	// tableUserExists обрабатывает ошибку отсутствия созданной таблицы users в БД.
	var tableUserExists bool
	err = DB.Raw(`
        SELECT EXISTS (
            SELECT 1 
            FROM pg_catalog.pg_tables 
            WHERE schemaname = 'public' 
            AND tablename = 'users'
        )
    `).Scan(&tableUserExists).Error

	if err != nil {
		return nil, fmt.Errorf("ошибка проверки таблицы: %w", err)
	}

	if !tableUserExists {
		return nil, projecterrors.ErrNoUserTable
	}

	// tableTaskExists обрабатывает ошибку отсутствия созданной таблицы tasks в БД.
	var tableTaskExists bool
	err = DB.Raw(`
        SELECT EXISTS (
            SELECT 1 
            FROM pg_catalog.pg_tables 
            WHERE schemaname = 'public' 
            AND tablename = 'tasks'
        )
    `).Scan(&tableTaskExists).Error

	if err != nil {
		return nil, fmt.Errorf("ошибка проверки таблицы: %w", err)
	}

	if !tableTaskExists {
		return nil, projecterrors.ErrNoTaskTable
	}

	return DB, nil
}
