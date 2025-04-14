package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableAutomaticPing:                     true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database, %v ", err)
	}
}
