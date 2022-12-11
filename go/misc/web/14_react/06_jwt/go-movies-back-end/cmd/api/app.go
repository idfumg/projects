package main

import (
	"database/sql"
	"myapp/internal/repository"
	"myapp/internal/repository/dbrepo"
)

type application struct {
	DB   repository.DatabaseRepo
	cfg  *Cfg
	auth *Auth
}

func NewApp(conn *sql.DB, cfg *Cfg) *application {
	db := &dbrepo.PostgresDBRepo{
		DB: conn,
	}

	auth := &Auth{
		Issuer:        cfg.JWTIssuer,
		Audience:      cfg.JWTAudience,
		Secret:        cfg.JWTSecret,
		TokenExpiry:   cfg.TokenExpiry,
		RefreshExpiry: cfg.RefreshExpiry,
		CookiePath:    cfg.CookiePath,
		CookieName:    cfg.CookieName,
		CookieDomain:  cfg.CookieDomain,
	}

	return &application{
		DB:   db,
		cfg:  cfg,
		auth: auth,
	}
}
