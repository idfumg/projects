package main

import "fmt"

type Product struct {
	value int
}

type Spec interface {
	IsSatisfied(Product) bool
}

type PositiveSpec struct{}

func (s *PositiveSpec) IsSatisfied(p Product) bool { return p.value > 0 }

type NegativeSpec struct{}

func (s *NegativeSpec) IsSatisfied(p Product) bool { return p.value < 0 }

type ValueSpec struct{ value int }

func (s *ValueSpec) IsSatisfied(p Product) bool { return s.value == p.value }

type OrSpec struct{ first, second Spec }

func (s *OrSpec) IsSatisfied(p Product) bool {
	return s.first.IsSatisfied(p) || s.second.IsSatisfied(p)
}

// This function is closed for modification
func Filter(products []Product, fn Spec) []Product {
	ans := []Product{}
	for _, product := range products {
		if fn.IsSatisfied(product) {
			ans = append(ans, product)
		}
	}
	return ans
}

func SOLID_OpenClosed() {
	products := []Product{{-1}, {-2}, {1}, {2}, {3}, {4}}

	fmt.Println()
	for _, p := range Filter(products, &PositiveSpec{}) {
		fmt.Printf("%d ", p.value)
	}
	fmt.Println()

	for _, p := range Filter(products, &NegativeSpec{}) {
		fmt.Printf("%d ", p.value)
	}
	fmt.Println()

	for _, p := range Filter(products, &ValueSpec{value: 3}) {
		fmt.Printf("%d ", p.value)
	}
	fmt.Println()

	for _, p := range Filter(products, &OrSpec{&ValueSpec{3}, &ValueSpec{4}}) {
		fmt.Printf("%d ", p.value)
	}
	fmt.Println()
}
