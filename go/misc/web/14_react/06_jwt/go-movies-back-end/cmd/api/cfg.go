package main

import (
	"flag"
	"time"
)

type Cfg struct {
	DSN           string
	JWTSecret     string
	JWTIssuer     string
	JWTAudience   string
	CookieDomain  string
	CookiePath    string
	CookieName    string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	Domain        string
}

func NewCfg() *Cfg {
	var (
		dsn           string
		jwtSecret     string
		jwtIssuer     string
		jwtAudience   string
		cookieDomain  string
		domain        string
	)

	flag.StringVar(&dsn, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=utc connect_timeout=5", "Postgres connection string")
	flag.StringVar(&jwtSecret, "jwt-secret", "verysecret", "signing secret")
	flag.StringVar(&jwtIssuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&jwtAudience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&cookieDomain, "cookie-domain", "127.0.0.1", "cookie domain")
	flag.StringVar(&domain, "domain", "example.com", "domain")
	flag.Parse()

	return &Cfg{
		DSN:           dsn,
		JWTSecret:     jwtSecret,
		JWTIssuer:     jwtIssuer,
		JWTAudience:   jwtAudience,
		CookieDomain:  cookieDomain,
		CookiePath:    "/",
		CookieName:    "refresh_token",
		TokenExpiry:   15 * time.Minute,
		RefreshExpiry: 24 * time.Hour,
		Domain:        domain,
	}
}
