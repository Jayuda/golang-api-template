package app

import (
	"log"
	"os"
	"time"
	"voltunes-chick-api-master-product/model/domain"

	"gorm.io/driver/postgres"

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
	dsn := "host=192.168.0.1 user=postgres password=postgres dbname=dbtest port=6655 sslmode=disable TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// RUN before_auto_migrate.sql
	// helper.RunSQLFromFile(database, "app/database/before_auto_migrate.sql")

	err = database.AutoMigrate(
		// Bank
		&domain.Bank{},
	)
	if err != nil {
		panic("failed to auto migrate schema")
	}

	// RUN after_auto_migrate.sql
	// helper.RunSQLFromFile(database, "app/database/after_auto_migrate.sql")

	return database
}
