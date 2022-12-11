package main

import "flag"

type Cfg struct {
	DSN string
}

func NewCfg() *Cfg {
	dsn := ""

	flag.StringVar(&dsn, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=utc connect_timeout=5", "Postgres connection string")
	flag.Parse()

	return &Cfg{
		DSN: dsn,
	}
}
