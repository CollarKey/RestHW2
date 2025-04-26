package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	dsn := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableAutomaticPing:                     true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed connect to database, %v ", err)
	}
	// обработка ошибки отсутствия созданной таблицы tasks в БД
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
		return nil, fmt.Errorf("таблица 'tasks' не найдена (код: 42P01)")
	}

	return DB, nil
}
