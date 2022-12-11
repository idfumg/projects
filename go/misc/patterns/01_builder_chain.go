package main

import "fmt"

type Person struct {
	StreetAddress string
	PostCode      string
	City          string
	CompanyName   string
	Position      string
	AnnualIncome  int
}

type PersonBuilder struct {
	person *Person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &Person{},
	}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonAddressBuilder) WithStreetAddress(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) WithCity(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	b.person.PostCode = postcode
	return b
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (b *PersonJobBuilder) WithCompany(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

func (b *PersonJobBuilder) WithPosition(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) WithAnnualIncome(annualIncome int) *PersonJobBuilder {
	b.person.AnnualIncome = annualIncome
	return b
}

func Builder_Chain() {
	person :=
		NewPersonBuilder().
			Lives().
			WithStreetAddress("123 London Road").
			WithCity("London").
			WithPostcode("SW12BC").
			Works().
			WithCompany("Fabrikam").
			WithPosition("Programmer").
			WithAnnualIncome(123000).
			Build()

	fmt.Println(person)
}
