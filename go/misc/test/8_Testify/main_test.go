package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterUnique(t *testing.T) {
	input := []Developer{
		{Name: "Elliot"},
		{Name: "Elliot"},
		{Name: "David"},
		{Name: "Alexander"},
		{Name: "Eva"},
		{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"David",
		"Alexander",
		"Eva",
		"Alan",
	}

	assert.Equal(t, expected, FilterUnique(input))
}

func TestNegativeFilterUnique(t *testing.T) {
	input := []Developer{
		{Name: "Elliot"},
		{Name: "Elliot"},
		{Name: "David"},
		{Name: "Alexander"},
		{Name: "Eva"},
		{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"Eva",
		"Alan",
	}

	assert.NotEqual(t, expected, FilterUnique(input))
}