package main

import "fmt"

type Employee struct {
	Name     string
	Position string
	Income   int
}

func NewEmployeeFactory(position string, income int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:     name,
			Position: position,
			Income:   income,
		}
	}
}

type EmployeeFactory struct {
	Position string
	Income   int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{
		Name:     name,
		Position: f.Position,
		Income:   f.Income,
	}
}

func NewEmployeeFactory2(position string, income int) *EmployeeFactory {
	return &EmployeeFactory{
		Position: position,
		Income:   income,
	}
}

func Factory_FactoryGenerator() {
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 80000)
	developer := developerFactory("Adam")
	manager := managerFactory("Jane")
	fmt.Println(developer)
	fmt.Println(manager)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}
