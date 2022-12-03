package main

import "fmt"

type Person2 interface {
	SayHello()
}

type person2 struct {
	name string
	age  int
}

func (p *person2) SayHello() {
	fmt.Printf("Name: %s, age: %d\n", p.name, p.age)
}

func NewPerson2(name string, age int) Person2 {
	return &person2{name: name, age: age}
}

func Factory_InterfaceFactory() {
	p := NewPerson2("Bod", 20)
	p.SayHello()
}
