package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address4 struct {
	Street  string
	City    string
	Country string
}

func (a *Address4) DeepCopy() *Address4 {
	return &Address4{
		Street:  a.Street,
		City:    a.City,
		Country: a.Country,
	}
}

type Person4 struct {
	Name    string
	Address *Address4
	Friends []string
}

func (p *Person4) DeepCopy() *Person4 {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func (p *Person4) DeepCopySerialization() *Person4 {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	ans := Person4{}
	_ = d.Decode(&ans)

	return &ans
}

var mainOffice = &Person4{"", &Address4{"123 London Rd", "London", "UK"}, []string{}}

func newPerson(proto *Person4, name string) *Person4 {
	ans := proto.DeepCopySerialization()
	ans.Name = name
	return ans
}

func NewLondonPerson(name string) *Person4 {
	return newPerson(mainOffice, name)
}

func Prototype() {
	john := &Person4{
		"John",
		&Address4{"123 London Rd", "London", "UK"},
		[]string{"Sam", "Jane"},
	}
	fmt.Println(john, john.Address)

	bob := john.DeepCopy()
	bob.Name = "Bob"
	bob.Address.City = "Moscow"
	bob.Friends = append(bob.Friends, "Angela")
	fmt.Println(bob, bob.Address)

	jane := bob.DeepCopySerialization()
	jane.Address.City = "Praha"
	fmt.Println(jane, jane.Address)

	sam := NewLondonPerson("Sam")
	fmt.Println(sam, sam.Address)
}
