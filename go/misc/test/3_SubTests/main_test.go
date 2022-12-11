// Subtests

package main

import "testing"

type TestCase struct {
	value    int
	expected bool
}

func TestIsArmstrong(t *testing.T) {
	t.Run("should return true for 371", func(t *testing.T) {
		testCase := TestCase{
			value: 371,
			expected: true,
		}
		actual := IsArmstrong(testCase.value)
		if testCase.expected != actual {
			t.Fail()
		}
	})

	t.Run("should return true for 370", func(t *testing.T) {
		testCase := TestCase{
			value: 370,
			expected: true,
		}
		actual := IsArmstrong(testCase.value)
		if testCase.expected != actual {
			t.Fail()
		}
	})
}

func TestIsNotArmstrong(t *testing.T) {
	t.Run("should return false for 350", func(t *testing.T) {
		testCase := TestCase{
			value: 350,
			expected: false,
		}
		actual := IsArmstrong(testCase.value)
		if testCase.expected != actual {
			t.Fail()
		}
	})
	t.Run("should return false for 300", func(t *testing.T) {
		testCase := TestCase{
			value: 300,
			expected: false,
		}
		actual := IsArmstrong(testCase.value)
		if testCase.expected != actual {
			t.Fail()
		}
	})
}