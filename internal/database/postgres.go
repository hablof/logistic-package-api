package database

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/jmoiron/sqlx"
)

type Config interface {
	GetDSN() string
	GetMaxOpenConns() int
	GetMaxIdleConns() int
	GetConnMaxIdleTime() time.Duration
	GetConnMaxLifetime() time.Duration
}

// NewPostgres returns DB
func NewPostgres(cfg Config, driver string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, cfg.GetDSN())
	if err != nil {
		log.Error().Err(err).Msgf("failed to create database connection")

		return nil, err
	}
	db.SetMaxOpenConns(cfg.GetMaxOpenConns())
	db.SetMaxIdleConns(cfg.GetMaxIdleConns())
	db.SetConnMaxIdleTime(cfg.GetConnMaxIdleTime())
	db.SetConnMaxLifetime(cfg.GetConnMaxLifetime())

	// need to uncomment for homework-4
	if err = db.Ping(); err != nil {
		log.Error().Err(err).Msgf("failed ping the database")

		return nil, err
	}

	return db, nil
}
