package main

import "testing"

var tests = []struct {
	name     string
	dividend float64
	divisor  float64
	expected float64
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr && err == nil {
			t.Error("Expected an error but did not get one")
		} 
		if !tt.isErr && err != nil {
			t.Error("Did not expect an error but got one")
		}
		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}
