package database

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type Config interface {
	GetDSN() string
	GetMaxOpenConns() int
	GetMaxIdleConns() int
	GetConnMaxIdleTime() time.Duration
	GetConnMaxLifetime() time.Duration
	GetAttempts() int
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

	maxAttempts := cfg.GetAttempts()
	for i := 0; i < maxAttempts; i++ {

		err = db.Ping()
		if err == nil {
			break
		}

		log.Debug().Err(err).Msg("database ping attempt failed...")
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Error().Err(err).Msgf("failed ping the database")
		return nil, err
	}

	log.Debug().Err(err).Msg("database ping attempt succeeded")

	return db, nil
}
