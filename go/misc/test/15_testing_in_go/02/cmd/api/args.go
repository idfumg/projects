package main

import "flag"

type args struct {
	DSN       string
	Domain    string
	JWTSecret string
}

func NewArgs() *args {
	ans := &args{}
	flag.StringVar(&ans.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection dsn")
	flag.StringVar(&ans.Domain, "domain", "example.com", "Domain for application")
	flag.StringVar(&ans.JWTSecret, "jwt-secret", "827ccb0eea8a706c4c34a16891f84e7b", "signing secret")
	flag.Parse()
	return ans
}
