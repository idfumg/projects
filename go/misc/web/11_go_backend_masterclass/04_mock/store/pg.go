package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Logger interface {
	Printf(format string, v ...any)
}

type Config interface {
	GetDBHost() string
	GetDBPort() string
	GetDBUsername() string
	GetDBTable() string
	GetDBPassword() string
	GetDBSslMode() string
	GetDBUrl() string
}

type DbTx interface {
	ExecContext(ctx context.Context, query string, params ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, params ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, params ...interface{}) *sql.Row
}

type Pg struct {
	db     DbTx
	logger Logger
	config Config
}

func NewPg(logger Logger, config Config) (*Pg, error) {
	connectString, err := GetConnectString(config)
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return nil, errors.New("could not connect to the database: " + err.Error())
	}

	if err := db.PingContext(context.Background()); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(30 * time.Second)
	db.SetConnMaxIdleTime(30 * time.Second)

	return &Pg{
		db:     db,
		logger: logger,
		config: config,
	}, nil
}

func GetConnectString(config Config) (string, error) {
	if len(config.GetDBHost()) != 0 {
		return GetDBUrl(
			config.GetDBHost(),
			config.GetDBPort(),
			config.GetDBUsername(),
			config.GetDBTable(),
			config.GetDBPassword(),
			config.GetDBSslMode()), nil
	}
	return pq.ParseURL(config.GetDBUrl()) // postgres://postgres:postgres@localhost:5432/dbname?sslmode=disable
}

func GetDBUrl(host, port, username, table, password, sslmode string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, username, table, password, sslmode,
	)
}
