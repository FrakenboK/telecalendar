package database

import (
	"fmt"
	"telecalendar/internal/config"
	"telecalendar/internal/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow"

func Init(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				dsn,
				cfg.Postgres.Host,
				cfg.Postgres.Username,
				cfg.Postgres.Password,
				cfg.Postgres.Database,
				cfg.Postgres.Port,
			),
		),
		&gorm.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf("postgres connection failure: %w", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Calendar{}, &models.Event{})
	if err != nil {
		return nil, fmt.Errorf("automigrate failure: %w", err)
	}
	return db, nil
}
