package main

import "flag"

type args struct {
	DSN string
}

func NewArgs() *args {
	ans := &args{}
	flag.StringVar(&ans.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection dsn")
	flag.Parse()
	return ans
}
