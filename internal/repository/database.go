package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	Dbname          string
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
}

func NewDatabase(config DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.Dbname,
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)

	return db, nil
}
