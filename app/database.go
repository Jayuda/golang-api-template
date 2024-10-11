package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	// err = database.AutoMigrate(
	// 	// Bank
	// 	&domain.Bank{},
	// )
	// if err != nil {
	// 	panic("failed to auto migrate schema")
	// }

	// RUN after_auto_migrate.sql
	// helper.RunSQLFromFile(database, "app/database/after_auto_migrate.sql")

	return database
}
