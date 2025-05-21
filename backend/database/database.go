package database

import (
	"fmt"
	"log"
	"sistem-manajemen-armada/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var (
	db_host      = config.Env.DB_HOST
	db_user      = config.Env.DB_USER
	db_password  = config.Env.DB_PASSWORD
	db_name      = config.Env.DB_NAME
	db_port      = config.Env.DB_PORT
	db_ssl_mode  = config.Env.DB_SSL_MODE
	db_time_zone = config.Env.DB_TIME_ZONE
)

func InitializeDB() {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", db_host, db_user, db_password, db_name, db_port, db_ssl_mode, db_time_zone)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
}
