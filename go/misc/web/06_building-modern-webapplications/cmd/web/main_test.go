package main

import "testing"

func TestInitAll(t *testing.T) {
	appConfig, err := initAll()
	if err != nil {
		t.Error("Fail initAll()", err)
	}
	if appConfig == nil {
		t.Error("Ivalid appConfig")
	}
}