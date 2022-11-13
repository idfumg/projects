package main

import "fmt"

type weapon interface {
	attack()
}

type weapons []weapon

func (w weapons) attack() {
	for _, w := range w {
		w.attack()
	}
}

type star struct {
	Name string
}

type sword struct {
	Name string
}

func (s star) attack() {
	fmt.Println(s.Name, "swining a sword")
}

func (s sword) attack() {
	fmt.Println(s.Name, "throwing a start")
}

func main() {
	ws := weapons{star{"Bob"}, sword{"Keit"}}
	ws.attack()
}