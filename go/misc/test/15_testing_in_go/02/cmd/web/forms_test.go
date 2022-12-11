package main

import (
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	form := NewForm(nil)

	has := form.Has("whatever")
	if has {
		t.Error("new form should not have any fields")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = NewForm(postedData)

	has = form.Has("a")
	if !has {
		t.Error("this form should have a field")
	}
}

func TestForm_Required(t *testing.T) {
	form := NewForm(nil)
	form.Required("a")
	if form.Valid() {
		t.Error("Should have exactly one error")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = NewForm(postedData)
	form.Required("a")
	if !form.Valid() {
		t.Error("Should have all fields validated successfully")
	}
}

func TestForm_Check(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")
	if form.Valid() {
		t.Error("Form is valid but there should be an error in there")
	}
}

func TestForm_ErrorGet(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")
	s := form.Errors.Get("password")
	if s == "" {
		t.Error("Form is valid but there should have generated an error")
	}

	s = form.Errors.Get("whatever")
	if s != "" {
		t.Error("Should not have an error but got one")
	}
}