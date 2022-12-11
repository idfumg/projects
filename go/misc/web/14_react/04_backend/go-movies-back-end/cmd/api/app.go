package main

type application struct {
	Domain string
}

func NewApp() *application {
	return &application{}
}