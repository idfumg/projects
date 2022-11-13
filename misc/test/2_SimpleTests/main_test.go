package main

import "testing"

type TestCase struct {
	value    int
	expected bool
}

func TestIsArmstrong(t *testing.T) {
	testCase := TestCase{
		value: 371,
		expected: true,
	}
	actual := IsArmstrong(testCase.value)
	if testCase.expected != actual {
		t.Fail()
	}
}

func TestIsNotArmstrong(t *testing.T) {
	testCase := TestCase{
		value: 350,
		expected: false,
	}
	actual := IsArmstrong(testCase.value)
	if testCase.expected != actual {
		t.Fail()
	}
}