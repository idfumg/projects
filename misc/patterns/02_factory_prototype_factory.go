package main

import "fmt"

type Employee2 struct {
	Name     string
	Position string
	Income   int
}

type EmployeeType int

const (
	Developer = iota
	Manager
)

func NewEmployee2(role int) *Employee2 {
	switch role {
	case Developer:
		return &Employee2{"", "developer", 60000}
	case Manager:
		return &Employee2{"", "manager", 80000}
	default:
		panic("unsupported role")
	}
}

func Factory_PrototypeFactory() {
	m := NewEmployee2(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
