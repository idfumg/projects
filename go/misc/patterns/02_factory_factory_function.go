package main

type Person3 struct {
	Name string
	Age  int
}

func NewPerson3(name string, age int) *Person3 {
	return &Person3{Name: name, Age: age}
}

func Factory_FactoryFunction() {
	_ = NewPerson3("Bob", 20)
}
