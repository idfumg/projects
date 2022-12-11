package api

import (
	"myapp/pkg/config"
	"net/http"
	"testing"
)

func TestLoadCSRF(t *testing.T) {
	conf, err := config.New(config.Test)
	if err != nil {
		t.Error(err)
	}

	h := LoadCSRF(conf)(&myHandler{})

	switch v := h.(type) {
	case http.Handler:
		break
	default:
		t.Errorf("Expected type http.Handler, but got %T", v)
	}
}

func TestLoadSession(t *testing.T) {
	conf, err := config.New(config.Test)
	if err != nil {
		t.Error(err)
	}

	h := LoadSession(conf)(&myHandler{})

	switch v := h.(type) {
	case http.Handler:
		break
	default:
		t.Errorf("Expected type http.Handler, but got %T", v)
	}
}
