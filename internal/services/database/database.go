package database

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const ServiceName = "Database"

type Service struct {
	DB *gorm.DB
}

func (s *Service) String() string { return ServiceName }

func (s *Service) Start(context.Context) (err error) {
	s.DB, err = gorm.Open(postgres.Open(fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)))
	if err != nil {
		return err
	}
	return s.DB.AutoMigrate(
	// Insert auto migrated models here
	)
}

func (s *Service) State(ctx context.Context) (string, error) {
	db, err := s.DB.DB()
	if err != nil {
		return "Unhealthy", err
	}
	err = db.Ping()
	if err != nil {
		return "Unhealthy", err
	}
	return "Healthy", nil
}

func (s *Service) Terminate(ctx context.Context) error {
	db, err := s.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
