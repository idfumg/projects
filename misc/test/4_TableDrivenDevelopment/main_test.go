// Table Driven Testing

package main

import "testing"

type TestCase struct {
	name     string
	value    int
	expected bool
}

func TestIsArmstrong(t *testing.T) {
	tests := []TestCase{
		{name: "Testing value for: 153", value: 153, expected: true},
		{name: "Testing value for: 370", value: 370, expected: true},
		{name: "Testing value for: 371", value: 371, expected: true},
		{name: "Testing value for: 407", value: 407, expected: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsArmstrong(test.value)
			if test.expected != actual {
				t.Fail()
			}
		})
	}
}

func TestIsNotArmstrong(t *testing.T) {
	tests := []TestCase{
		{name: "Testing value for: 350", value: 350, expected: false},
		{name: "Testing value for: 300", value: 300, expected: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsArmstrong(test.value)
			if test.expected != actual {
				t.Fail()
			}
		})
	}
}
